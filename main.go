package main

import (
	"io"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	"urlshortener/common"
	"urlshortener/utils"
)

func main() {
	time.Local = time.UTC

	// log configuration
	logFileWriter := &lumberjack.Logger{
		Filename: "logs/run.log",
		MaxSize:  3, // megabytes
	}
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(io.MultiWriter(os.Stderr, logFileWriter))

	// init step
	log.Infoln("Server initialized")

	// init postgre db
	// from the common package
	_, err := common.InitDB()
	if err != nil {
		panic(err)
	}
	log.Infoln("Database connection established")

	// Run the cleanup job as a background goroutine
	// CRON job here, use a library for it
	// different from goroutine which is not managed by os but by the go run time
	log.Infoln("Cleanup Job Running")
	go utils.RunCleanupJob()

	// migrate db
	// used to update a live database when it is changed e.g. new table
	// it will require a migration script to be written
	// RECOMMENDED ACTIONS: creating a new column of data and copying old column over
	// NOT RECOMMENDED ACTIONS: dropping a column, rename column, mutate column
	// so usually e.g. i want to change a status column from text to int
	// creat new int column, then map the data from the old text column over, then drop the old column
	// db, err := common.GetDB()
	// if err != nil {
	// 	fmt.Println("Error opening database connection:", err)
	// 	return
	// }
	// defer db.Close()

	// // AutoMigrate will update the table structure based on the model
	// db.AutoMigrate(&url.URL{})

	// start the server
	server := &Server{}
	server.Init("8080")
	log.Infoln("Server running!!")
	if err := server.Serve(); err != nil {
		log.Errorln(err)
	}

	log.Infoln("Server stopped")

	// Keep the main goroutine running to allow the background job to continue
	select {}

}
