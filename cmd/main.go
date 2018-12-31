package main

import (
	"context"
	"flag"
	// "github.com/bnbjin/parawebshell_server/config"
	pws "github.com/bnbjin/parawebshell_server"
	"log"
	"time"
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

	ctx := context.Background()

	/* 性能状态记录 */
	proffCanceler, err := profileIfEnabled()
	if nil != err {
		log.Panic(err)
	}
	defer proffCanceler()

	/* 设置信号处理器 */
	intrh, ctx := setupInterruptHandler(ctx)
	defer intrh.Close()

	log.Println("para web shell startup, version ", pws.CurrentVersionNumber)

	time.Sleep(100 * time.Millisecond)
}
