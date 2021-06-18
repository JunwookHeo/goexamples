package subpkg

import (
	log "github.com/sirupsen/logrus"
)

func LogTest() {
	log.Debug("Debug log")
	log.Info("Info log")
	log.Print("Print log")

	log.Warn("Warning log")
	log.Error("Error log")
}
