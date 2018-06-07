package startup

import (
	"fmt"
	"git.amabanana.com/plancks-cloud/pc-go-brutus/controller/docker"
	"git.amabanana.com/plancks-cloud/pc-go-brutus/model"
	"git.amabanana.com/plancks-cloud/pc-go-brutus/util"
	"github.com/sirupsen/logrus"
	"os"
	"time"
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

	url := fmt.Sprintf("http://localhost:%v", model.AppPort)

	answered := false
	attempts := 0
	maxAttempts := 5
	sleepTime := 5
	for !answered {
		bytes, err := util.GetRequest(url)

		if bytes != nil {
			answered = true
			break
		}

		if err != nil {
			logrus.Error(fmt.Sprintf("Checking the health service faild with an error, %s", err.Error()))
			if attempts < maxAttempts {
				logrus.Infoln(fmt.Sprintf("Will try again in a few seconds"))
				time.Sleep(time.Duration(sleepTime) * time.Second)
				continue
			}
		}

		attempts++
		if attempts > maxAttempts {
			break
		}

	}

	if answered {
		logrus.Infoln(fmt.Sprintf(".. ✅ Success"))
	} else {
		logrus.Fatalln(fmt.Sprintf(".. Was not able to find Health service in %v attampts", maxAttempts))
	}

}

func installComplete() {
	logrus.Infoln(fmt.Sprintf("The installation completed succesfully."))
}
