package main

import (
	"log"

	"github.com/docker/go-plugins-helpers/volume"
)

func main() {

	driver := NewExampleDriver()
	handler := volume.NewHandler(driver)
	if err := handler.ServeUnix("docker-volume-driver.sock", 0); err != nil {
		log.Fatalf("Error %v", err)
	}

	for {

	}
}
