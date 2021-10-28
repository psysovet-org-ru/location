package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iamsalnikov/mymigrate"
	"github.com/iamsalnikov/mymigrate/cobracmd"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"location/connect"
	_ "location/migrations"
	"location/repository"
	service2 "location/service"
	"log"
	"net/http"
)

func main() {
	godotenv.Load()

	var c connect.Connect

	db, err := c.Get()
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	mymigrate.SetDatabase(db)

	// MigrateCmd is a cobra command to work with migrations
	var MigrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "work with migrations",
	}

	InitCmd := &cobra.Command{
		Use:   "init",
		Short: "init",
		RunE: func(cmd *cobra.Command, args []string) error {
			return service2.PushData(db)
		},
	}

	ServiceCmd := &cobra.Command{
		Use:   "listen",
		Short: "start service",
		RunE: func(cmd *cobra.Command, args []string) error {
			r := chi.NewRouter()
			r.Use(middleware.Logger)
			service2.Service(r, repository.NewStorage(db))
			err = http.ListenAndServe(":3000", r)
			if err != nil {
				log.Fatalln(err)
			}
			return nil
		},
	}

	MigrateCmd.AddCommand(cobracmd.CreateCmd, cobracmd.HistoryCmd, cobracmd.NewListCmd, cobracmd.ApplyCmd, cobracmd.DownCmd)

	var rootCmd = &cobra.Command{Use: "app"}

	rootCmd.AddCommand(InitCmd, MigrateCmd, ServiceCmd)

	rootCmd.Execute()
}
