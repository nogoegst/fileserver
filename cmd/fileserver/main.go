package main

import (
	"flag"
	"log"
	"net"

	"github.com/nogoegst/fileserver"
)

func main() {
	var zipFlag = flag.Bool("z", false, "serve from zip archive")
	var debugFlag = flag.Bool("debug", false, "debug")
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatalf("Please specify exactly one path")
	}
	path := flag.Args()[0]
	l, err := net.Listen("tcp4", "0.0.0.0:0")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Serving on %s", l.Addr().String())
	log.Fatal(fileserver.Serve(l, path, *zipFlag, *debugFlag))
}
