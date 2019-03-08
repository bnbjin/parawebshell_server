package main

import (
	// "context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	// "github.com/bnbjin/parawebshell_server/config"
	pws "github.com/bnbjin/parawebshell_server"
	profile "github.com/bnbjin/parawebshell_server/profile"
	tiny "github.com/go101/tinyrouter"
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

	// ctxbg := context.Background()

	/* profiling */
	proffCanceler, err := profile.ProfileIfEnabled()
	if nil != err {
		log.Panic(err)
	}
	defer proffCanceler()

	/* api router */
	routes := []tiny.Route{
		{
			Method:  "GET",
			Pattern: "/a/b/:c",
			HandleFunc: func(w http.ResponseWriter, req *http.Request) {
				params := tiny.PathParams(req)
				fmt.Fprintln(w, "/a/b/:c", "c =", params.Value("c"))
			},
		},
		{
			Method:  "GET",
			Pattern: "/a/:b/c",
			HandleFunc: func(w http.ResponseWriter, req *http.Request) {
				params := tiny.PathParams(req)
				fmt.Fprintln(w, "/a/:b/c", "b =", params.Value("b"))
			},
		},
		{
			Method:  "GET",
			Pattern: "/a/:b/:c",
			HandleFunc: func(w http.ResponseWriter, req *http.Request) {
				params := tiny.PathParams(req)
				fmt.Fprintln(w, "/a/:b/:c", "b =", params.Value("b"), "c =", params.Value("c"))
			},
		},
		{
			Method:  "GET",
			Pattern: "/:a/b/c",
			HandleFunc: func(w http.ResponseWriter, req *http.Request) {
				params := tiny.PathParams(req)
				fmt.Fprintln(w, "/:a/b/c", "a =", params.Value("a"))
			},
		},
		{
			Method:  "GET",
			Pattern: "/:a/:b/:c",
			HandleFunc: func(w http.ResponseWriter, req *http.Request) {
				params := tiny.PathParams(req)
				fmt.Fprintln(w, "/:a/:b/:c", "a =", params.Value("a"), "b =", params.Value("b"), "c =", params.Value("c"))
			},
		},
	}

	router := tiny.New(tiny.Config{Routes: routes})

	log.Println("Starting service ...")
	log.Fatal(http.ListenAndServe(":8080", router))

	/*
		$ curl localhost:8080/a/b/c
		/a/b/:c c = c
		$ curl localhost:8080/a/x/c
		/a/:b/c b = x
		$ curl localhost:8080/a/x/y
		/a/:b/:c b = x c = y
		$ curl localhost:8080/x/b/c
		/:a/b/c a = x
		$ curl localhost:8080/x/y/z
		/:a/:b/:c a = x b = y c = z
	*/

	log.Println("para web shell startup, version ", pws.CurrentVersionNumber)

	os.Exit(0)
}
