package startup

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"syscall"
)

const ModeField = "MODE"
const ModeInstall = "INSTALL"
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
		logrus.Error(fmt.Sprintf("Could not find a mode. Shutting down."))
		panic(0)
	}
	return mode

}
