// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: ff051dd8a6
// Version Date: Sun 15 Mar 23:19:53 UTC 2020

package main

import (
	"flag"

	// This Service
	"github.com/metaverse/truss/cmd/_integration-tests/server/test-service-definitions/0-basic/test-service/svc/server"
)

func main() {
	// Update addresses if they have been overwritten by flags
	flag.Parse()

	server.Run(server.DefaultConfig
}
