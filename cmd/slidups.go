package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/acamilleri/slidups/internal/handler"
	"github.com/acamilleri/slidups/pkg/logger"
)

var (
	version string

	listenAddr = kingpin.Flag("listen.addr", "listen address").
		Default("0.0.0.0:8080").
		Envar("SLIDUPS_LISTEN_ADDR").
		String()

	uploadDestination = kingpin.Flag("upload.destination", "destination directory of file upload").
		Default("/slides").
		Envar("SLIDUPS_UPLOAD_DESTINATION").
		ExistingDir()

	logLevel = kingpin.Flag("log.level", "log level").
		Envar("SLIDUPS_LOG_LEVEL").
		Default(logrus.InfoLevel.String()).
		Enum("debug", "info", "error")

	logFormat = kingpin.Flag("log.format", "log format").
		Envar("SLIDUPS_LOG_FORMAT").
		Default("text").
		Enum("text", "json")
)

func main() {
	kingpin.Version(version)
	kingpin.Parse()

	// init logger
	log := logger.New()
	log.SetLevel(*logLevel)
	log.SetFormatter(*logFormat)

	// init handler
	hdl := handler.New(log, *uploadDestination)
	hdl.RegisteredRoutes()

	// start http server
	log.Infof("Listen addr: %s", *listenAddr)
	err := http.ListenAndServe(*listenAddr, nil)
	if err != nil {
		log.WithError(err).Fatal("http server failure")
	}
}
