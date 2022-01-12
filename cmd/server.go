package cmd

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/francoposa/echo-server/application/server"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Run: run,
}

func run(cmd *cobra.Command, args []string) {

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.HandleFunc("/health", server.Health)
	router.HandleFunc("/echo", server.Echo)

	host := viper.GetString(serverHostFlag)
	port := viper.GetString(serverPortFlag)
	readTimeout := viper.GetInt(serverTimeOutReadFlag)
	writeTimeout := viper.GetInt(serverTimeOutWriteFlag)

	srv := &http.Server{
		Handler:      router,
		Addr:         host + ":" + port,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
	}

	fmt.Printf("running http server on port %s...\n", port)
	log.Fatal(srv.ListenAndServe())
}

const serverHostFlag = "server.host"
const serverPortFlag = "server.port"
const serverTimeOutReadFlag = "server.timeout.read"
const serverTimeOutWriteFlag = "server.timetout.write"

func init() {
	rootCmd.AddCommand(serverCmd)

	cmdFlags := serverCmd.Flags()

	cmdFlags.String(serverHostFlag, "", "")
	//nolint:ineffassign,staticcheck
	err := viper.BindPFlag(serverHostFlag, cmdFlags.Lookup(serverHostFlag))

	cmdFlags.String(serverPortFlag, "", "")
	//nolint:ineffassign,staticcheck
	err = viper.BindPFlag(serverPortFlag, cmdFlags.Lookup(serverPortFlag))

	cmdFlags.String(serverTimeOutReadFlag, "", "")
	//nolint:ineffassign,staticcheck
	err = viper.BindPFlag(
		serverTimeOutReadFlag, cmdFlags.Lookup(serverTimeOutReadFlag),
	)

	cmdFlags.String(serverTimeOutWriteFlag, "", "")
	//nolint:ineffassign,staticcheck
	err = viper.BindPFlag(
		serverTimeOutWriteFlag, cmdFlags.Lookup(serverTimeOutWriteFlag),
	)
	if err != nil {
		log.Panic(err)
	}

}
