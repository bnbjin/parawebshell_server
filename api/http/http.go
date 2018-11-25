package http

import (
	cmap "github.com/orcaman/concurrent-map"
	"log"
	"os"
)

var (
	loggerStd = log.New(os.Stdout, "pws/http: ", log.Lshortfile|log.Ltime|log.Ldate)
	loggerErr = log.New(os.Stderr, "pws/http: ", log.Lshortfile|log.Ltime|log.Ldate)
)

type Server struct {
}

func init() {
	loggerStd.Println("http module initialzes successfully.")
}
