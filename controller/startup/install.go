package startup

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func install() {

	welcomeUser()
	installNetwork()
	installHealth()
	checkHealth()

}

//welcomeUser tells the user what is about to happen
func welcomeUser() {
	logrus.Infoln(fmt.Sprintf("Welcome to the Planck's Cloud installer."))
}

//install network tries to setup pc-net and tells the user if it succeeded
func installNetwork() {
	logrus.Infoln(fmt.Sprintf(".. Attempting to create overlay network"))
	//Create network call
	//if err != nil {
	//	logrus.Error(fmt.Sprintf("Failed to create overlay network. Install failed. Shutting down."))
	//	panic(0)
	//}
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
