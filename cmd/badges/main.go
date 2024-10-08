//go:build !test

package main

import (
	_ "embed"
	"flag"
	"log"
	"os"
)

const cliVersion = "0.1.2"

var serverPort = flag.String("serverPort", "49002", "http service address")

var buildID string

func init() {
	if buildID == "" {
		buildID = "unknown"
	}
}

func main() {
	log.Printf("-> env:ENV_WEB_AUTO_HOST %s", os.Getenv("ENV_WEB_AUTO_HOST"))
	flag.Parse()
	log.Printf("-> run serverPort %v", *serverPort)
	log.Printf("=> now version %v", cliVersion)
	log.Printf("=> now version %v, build id: %s", cliVersion, buildID)
}
