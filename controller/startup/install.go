package startup

import (
	"fmt"
	"git.amabanana.com/plancks-cloud/pc-go-brutus/controller/docker"
	"github.com/sirupsen/logrus"
	"os"
)

func install() {

	welcomeUser()
	installNetwork()
	installHealth()
	checkHealth()
	installComplete()
	os.Exit(0)

}

//welcomeUser tells the user what is about to happen
func welcomeUser() {
	logrus.Infoln(fmt.Sprintf("Welcome to the Planck's Cloud installer."))
}

//install network tries to setup pc-net and tells the user if it succeeded
func installNetwork() {
	logrus.Infoln(fmt.Sprintf(".. Attempting to create overlay network"))
	exists, err := docker.CheckNetworkExists(docker.NetworkName)
	if err != nil {
		logrus.Fatalln("Could not check if the network exists")
	}

	if !exists {
		success, err := docker.CreateOverlayNetwork(docker.NetworkName)

		if err != nil {
			logrus.Fatalln("Could not check if the network exists")
		}

		if !success {
			logrus.Fatalln("Create network was not successful")
		}
	}

	logrus.Infoln(fmt.Sprintf(".. ✅ Success"))

}

//installHealth starts the Health docker image as service on the pc-net network and tells the user if it worked out
func installHealth() {
	logrus.Infoln(fmt.Sprintf(".. Attempting to install health service"))
	//Install health service
	//if err != nil {
	//	logrus.Error(fmt.Sprintf("Failed to install health service. Install failed. Shutting down."))
	//	panic(0)
	//}
	logrus.Infoln(fmt.Sprintf(".. ✅ Success"))

}

//checks that the health service can be contacted in the browser by calling it's health service
func checkHealth() {
	// Call localhost:6108/api/ping to see if it is there.
	// Loop for up to X seconds where X is few seconds.
	// Tell the user to navigate there in their browser.
	// Freak out if it failed.

}

func installComplete() {
	logrus.Infoln(fmt.Sprintf("The installation completed succesfully."))
}
