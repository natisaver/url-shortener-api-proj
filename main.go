package main
import (
	"fmt"
    "net/http"
    "github.com/gin-gonic/gin"

	"io"
	"os"
	"time"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	"database/sql"
	"github.com/lib/pq"
	"log"
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
	db, err:= common.InitDB() 
	if err != nil {
		panic(err)
	}
	log.Infoln("Database connection established")

	// migrate db

	// start the server
	// 
	log.Infoln("Server stopped")

}


