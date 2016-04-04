package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/kussj/actuatorbeat/beater"
)

func main() {
	err := beat.Run("actuatorbeat", "", beater.New())
	if err != nil {
		os.Exit(1)
	}
}
