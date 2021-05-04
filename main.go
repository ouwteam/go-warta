package main

import (
	"context"
	"flag"
	"fmt"
	"go-warta/src/route"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	APPDB, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
	if err != nil {
		panic(err.Error())
	}

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	routeMain := route.NewRouteMain(APPDB)

	r := mux.NewRouter()
	r.HandleFunc("/", routeMain.HandleMain)
	r.HandleFunc("/api/send-message", routeMain.HandleSendMessage)
	r.HandleFunc("/api/bot-info", routeMain.HandleBotInfo)
	r.HandleFunc("/api/get-update", routeMain.HandleGetUpdate)
	r.HandleFunc("/hook", routeMain.HandleMain)

	srv := &http.Server{
		Handler: r,
		Addr:    os.Getenv("SERVER_ADDR") + ":" + os.Getenv("SERVER_PORT"),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		fmt.Println("Listen and serve on : " + srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)

	defer cancel()
	log.Println("shutting down")
	APPDB.Close()
	srv.Shutdown(ctx)
	os.Exit(0)
}
