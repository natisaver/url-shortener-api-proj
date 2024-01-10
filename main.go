package main
import (
	"io"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/natisaver/url-shortener-api-proj/urlshortner/common"
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
	_, err:= common.InitDB() 
	if err != nil {
		panic(err)
	}
	log.Infoln("Database connection established")

	// migrate db
	// used to update a live database when it is changed e.g. new table 
	// it will require a migration script to be written
	// RECOMMENDED ACTIONS: creating a new column of data and copying old column over
	// NOT RECOMMENDED ACTIONS: dropping a column, rename column, mutate column
	// so usually e.g. i want to change a status column from text to int
	// creat new int column, then map the data from the old text column over, then drop the old column

	// start the server
	server := &server{}
	server.Init("8080")
	log.Infoln("Server running!!")
	if err := server.Serve(); err != nil {
		log.Errorln(err)
	}

	log.Infoln("Server stopped")

}


