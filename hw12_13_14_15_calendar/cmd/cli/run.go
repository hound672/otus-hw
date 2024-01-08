package cli

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/cmd/application"
)

func waitQuitSignal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func run() {
	app, stop, err := application.CreateApplication()
	if err != nil {
		log.Fatalf("Fail init app: %v", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("Fail run app: %v", err)
	}
	defer stop()

	waitQuitSignal()
}
