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

	host := viper.GetString(srvHostFlag)
	port := viper.GetString(srvPortFlag)
	readTimeout := viper.GetInt(srvTimeOutReadFlag)
	writeTimeout := viper.GetInt(srvTimeOutWriteFlag)

	srv := &http.Server{
		Handler:      router,
		Addr:         host + ":" + port,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
	}

	fmt.Printf("running http server on port %s...\n", port)
	log.Fatal(srv.ListenAndServe())
}

const (
	srvHostFlag         = "server.host"
	srvPortFlag         = "server.port"
	srvTimeOutReadFlag  = "server.timeout.read"
	srvTimeOutWriteFlag = "server.timeout.write"
)

func init() {
	rootCmd.AddCommand(serverCmd)

	cmdFlags := serverCmd.Flags()

	cmdFlags.String(srvHostFlag, "", "")
	_ = viper.BindPFlag(srvHostFlag, cmdFlags.Lookup(srvHostFlag))

	cmdFlags.String(srvPortFlag, "", "")
	_ = viper.BindPFlag(srvPortFlag, cmdFlags.Lookup(srvPortFlag))

	cmdFlags.String(srvTimeOutReadFlag, "", "")
	_ = viper.BindPFlag(srvTimeOutReadFlag, cmdFlags.Lookup(srvTimeOutReadFlag))

	cmdFlags.String(srvTimeOutWriteFlag, "", "")
	_ = viper.BindPFlag(srvTimeOutWriteFlag, cmdFlags.Lookup(srvTimeOutWriteFlag))
}
