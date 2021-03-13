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
	Use:   "server",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		router := chi.NewRouter()

		// Suggested basic middleware stack from chi's docs
		router.Use(middleware.RequestID)
		router.Use(middleware.RealIP)
		router.Use(middleware.Logger)
		router.Use(middleware.Recoverer)

		router.HandleFunc("/health", server.Health)
		router.HandleFunc("/echo", server.Echo)

		host := viper.GetString("server.host")
		port := viper.GetString("server.port")
		readTimeout := viper.GetInt("server.timeout.read")
		writeTimeout := viper.GetInt("server.timeout.write")

		srv := &http.Server{
			Handler:      router,
			Addr:         host + ":" + port,
			ReadTimeout:  time.Duration(readTimeout) * time.Second,
			WriteTimeout: time.Duration(writeTimeout) * time.Second,
		}

		fmt.Printf("running http server on port %s...\n", port)
		log.Fatal(srv.ListenAndServe())

	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.PersistentFlags().String("server.host", "", "")
	err := viper.BindPFlag("server.host", serverCmd.PersistentFlags().Lookup("server.host"))
	if err != nil {
		panic(err)
	}

	serverCmd.PersistentFlags().String("server.port", "", "")
	err = viper.BindPFlag("server.port", serverCmd.PersistentFlags().Lookup("server.port"))
	if err != nil {
		panic(err)
	}
}
