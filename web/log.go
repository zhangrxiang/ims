package web

import (
	"os"
	"simple-ims/utils"
	"time"
)

// Get a filename based on the date, just for the sugar.
func todayFilename() string {
	utils.Mkdir("logs")
	today := time.Now().Format("2006-01-02")
	return "logs/log-" + today + ".txt"
}

func newLogFile() *os.File {
	filename := todayFilename()
	// Open the file, this will append to the today's file if server restarted.
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}
