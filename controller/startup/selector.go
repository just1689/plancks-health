package startup

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"syscall"
)

//ModeField is the env field used to tell the this app what to do on startup
const ModeField = "MODE"

//ModeInstall will go about setting up this app as a service.
const ModeInstall = "INSTALL"

//ModeNormal is the day-to-day normal way of operating. This app will run as Normal most of the time.
const ModeNormal = "NORMAL"

func Init() {
	mode := findMode()
	if mode == ModeInstall {
		install()
	} else if mode == ModeNormal {
		normal()
	}
}

func findMode() string {
	mode, found := syscall.Getenv(ModeField)
	if !found {
		log.Error(fmt.Sprintf("Could not find a mode. Shutting down."))
		panic(0)
	}
	return mode

}
