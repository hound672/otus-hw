package cli

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/build"
)

func Main() {
	root := &cobra.Command{
		Version: build.Version,
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}

	migrateRoot := &cobra.Command{
		Use: "migrate",
	}

	migrateStatus := &cobra.Command{
		Use: "status",
		Run: func(cmd *cobra.Command, args []string) {
			migrateStatus()
		},
	}
	migrateUp := &cobra.Command{
		Use: "up",
		Run: func(cmd *cobra.Command, args []string) {
			migrateUp()
		},
	}
	migrateDown := &cobra.Command{
		Use: "down",
		Run: func(cmd *cobra.Command, args []string) {
			migrateDown()
		},
	}
	migrateRoot.AddCommand(migrateStatus, migrateUp, migrateDown)

	root.AddCommand(migrateRoot)
	if err := root.Execute(); err != nil {
		log.Fatalf("root.Execute: %v", err)
	}
}
