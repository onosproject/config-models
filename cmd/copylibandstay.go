package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// Just a simple program copies a file and stays resident
// Allows the container stay alive
// Add as a container to Helm chart deployment.yaml like
//
//  - name: config-models-devicesim-1-0-0
//    image: onosproject/modelplugin/devicesim-1.0.0:latest
//    imagePullPolicy: {{ .Values.image.pullPolicy }}
//    command: ["/copylibandstay"]
//    args: ["devicesim.so.1.0.0", "/usr/local/lib/devicesim.so.1.0.0", "stayrunning"]
//    volumeMounts:
//    - name: shared-data
//      mountPath: /usr/local/lib
func main() {
	if len(os.Args) <= 2 {
		log.Print("Nothing to do - copylibandstay src tgt [stay]")
	}

	if len(os.Args) > 2 {
		sourceFile := os.Args[1]
		targetfile := os.Args[2]
		log.Printf("Copying %s to %s", sourceFile, targetfile)

		from, err := os.Open(sourceFile)
		if err != nil {
			log.Fatal(err)
		}
		defer from.Close()

		to, err := os.OpenFile(targetfile, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer to.Close()

		_, err = io.Copy(to, from)
		if err != nil {
			log.Fatal(err)
		}
	}

	// If there is a 3rd argument then stay resident
	if len(os.Args) == 4 {
		log.Print("Staying alive")
		fmt.Print("Alive")
		for {
			time.Sleep(time.Second)
			fmt.Print(".")
		}
	}
}
