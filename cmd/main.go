package main

import (
	"flag"
	"log"

	"github.com/arybolovlev/orkestrator/client"
	"github.com/arybolovlev/orkestrator/manager"
)

func main() {
	// Global options
	var port int
	flag.IntVar(&port, "port", 8162, "Port to connect or listen to")

	// Server options
	var srv bool
	flag.BoolVar(&srv, "server", false, "Run Orkestrator in the server mode")

	// Client options
	var cli bool
	flag.BoolVar(&cli, "client", false, "Run Orkestrator in the client mode")
	var file string
	flag.StringVar(&file, "file", "", "Specification file")

	flag.Parse()

	if srv {
		log.Println("Running Orkestrator Manager")
		manager.Run(port)
	}

	if cli {
		log.Println("Running Orkestrator Client")
		client.Run(port, file)
	}
}
