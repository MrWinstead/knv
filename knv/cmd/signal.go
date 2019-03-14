package cmd

import (
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
)

func registerSignalHandler() {
	notificationChan := make(chan os.Signal)
	signal.Notify(notificationChan, os.Interrupt, os.Kill)
	go func() {
		select {
		case sig := <-notificationChan:
			logrus.Printf("received signal %v, shutting down", sig)
			rootContextCancel()
		case <-rootContext.Done():
			return
		}
	}()
}
