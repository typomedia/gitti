package cmd

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/typomedia/gitti/app"
	"github.com/typomedia/gitti/app/cfg"
	"github.com/typomedia/gitti/app/hdlr"
	"github.com/typomedia/gitti/app/msg"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const Port int = 4000

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the Gitti server",
	Long:  `Start the Gitti server and listen on the given port.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg.Read()

		port, err := cmd.Flags().GetInt("port")
		msg.Check(err)

		startServer(port)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntP("port", "p", Port, "Port to listen on")
}

func startServer(port int) {
	fmt.Println(app.Logo())
	fmt.Println("Starting server on port", port)

	router := mux.NewRouter()
	router.HandleFunc("/", hdlr.Index)
	router.HandleFunc("/status/{project}", hdlr.Status)
	router.HandleFunc("/log/{project}", hdlr.Log)
	router.HandleFunc("/checkout/{project}", hdlr.Checkout)
	router.HandleFunc("/pull/{project}", hdlr.Pull)
	router.HandleFunc("/prune/{project}", hdlr.Prune)
	router.Use(hdlr.Logger)

	if viper.GetBool("auth.enabled") {
		fmt.Println("Authentication token:", viper.GetString("auth.token"))
		router.Use(hdlr.Auth)
	}

	file, err := os.OpenFile(app.App.Name+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(io.MultiWriter(os.Stdout, file))

	// https://stackoverflow.com/questions/19659600
	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%d", port),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	err = srv.ListenAndServe()
	msg.Check(err)
}
