package main

import (
	"log"
	"net/http"
	"os"

	"github.com/varopxndx/bootcamp-workshop-w4/controller"
	"github.com/varopxndx/bootcamp-workshop-w4/router"
	"github.com/varopxndx/bootcamp-workshop-w4/service"
	"github.com/varopxndx/bootcamp-workshop-w4/usecase"

	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

func main() {
	// logger setup
	logger, err := createLogger()
	if err != nil || logger == nil {
		log.Fatal("creating logger: %w", err)
	}

	// Create client
	shapeClient := service.New()

	// Usecase
	useCase := usecase.New(shapeClient)

	// Controllers
	controller := controller.New(useCase, logger, render.New())

	// Setup application routes
	httpRouter := router.New(controller)

	http.ListenAndServe("localhost:8080", httpRouter)
}

func createLogger() (*logrus.Logger, error) {
	logLevel := "DEBUG"
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"log_level": logLevel,
		}).Error("parsing log_level")

		return nil, err
	}

	logger := logrus.New()
	logger.SetLevel(level)
	logger.Out = os.Stdout
	logger.Formatter = &logrus.JSONFormatter{}
	return logger, nil
}
