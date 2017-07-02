package main

import (
	"flag"
	"log"
	"net"
	"strings"

	"github.com/nogoegst/fileserver"
)

func main() {
	var zipFlag = flag.Bool("z", false, "serve from zip archive")
	var debugFlag = flag.Bool("debug", false, "debug")
	flag.Parse()
	path := strings.Join(flag.Args(), " ")
	l, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Serving on %s", l.Addr().String())
	log.Fatal(fileserver.Serve(l, path, *zipFlag, *debugFlag))
}
