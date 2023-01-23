package main


 import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	// set log output to JSON format
	log.SetFormatter(&log.JSONFormatter{})

	// set log output destination to standard output
	log.SetOutput(os.Stdout)

	// set log level to warning, which will only log messages at warning level or higher
	log.SetLevel(log.WarnLevel)

	// create a new log entry with fields for name and age
	log.WithFields(log.Fields{
		"Name": "James Buns",
		"Age":  90,
	}).Fatalln("James name and age")

}
