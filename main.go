package main

import (
	"log"
	"os/exec"
)

func main() {
	output, err := exec.Command("vcgencmd", " measure_temp").Output()

	if err != nil {
		log.Printf(err.Error())
	}
	if len(output) > 0 {
		log.Printf("output: %s", output)
	}
}
