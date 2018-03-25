package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"

	"github.com/bluele/slack"
)

const (
	webhook = "webHook"
)

func getTemp() float64 {
	reg, err := regexp.Compile("[^0-9.]+")
	if err != nil {
		log.Fatal(err)
	}
	output, execErr := exec.Command("vcgencmd", " measure_temp").Output()

	if execErr != nil {
		log.Fatal(err.Error())
	}
	tmp := reg.ReplaceAllString(string(output[:]), "")
	f, _ := strconv.ParseFloat(tmp, 64)
	return f
}

func partialAlarm(t float64) {
	host := ""
	host, _ = os.Hostname()
	if (t >= 80.0) && (t <= 85.0) {
		hook := slack.NewWebHook(webhook)
		err := hook.PostMessage(&slack.WebHookPostPayload{
			Text: fmt.Sprint("Medium-High Temp Alarm on %s",host),
		})
		if err != nil {
			panic(err)
		}
	}
}

func fullAlarm(t float64) {
        host := ""
        host, _ = os.Hostname()
	if t > 85.0 {
		hook := slack.NewWebHook(webhook)
		err := hook.PostMessage(&slack.WebHookPostPayload{
			Text: fmt.Sprint("Medium-High Temp Alarm on %s",host),
		})
		if err != nil {
			panic(err)
		}
	}
}

func funcSet() {
	output := getTemp()
	partialAlarm(output)
	fullAlarm(output)
}

func doEvery(d time.Duration, f func()) {
        for range time.Tick(d) {
              f() 
        }
}

func main() {
	doEvery(1000*time.Millisecond, funcSet)
}

