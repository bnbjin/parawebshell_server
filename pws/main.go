package main

import (
	"flag"
	"log"
	"time"
	// "github.com/bnbjin/parawebshell_server/config"
)

const ()

var (
// configPath = flag.String("config", "${HOME}/.paraws/config.json", "configuration path")
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	flag.Parse()
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln("paradox web shell panic: ", r)
			time.Sleep(1 * time.Second)
		} else {
			log.Println("paradox web shell exit normally")
		}
	}()

	log.Println("para web shell startup, version ", CurrentVersionNumber)
}
