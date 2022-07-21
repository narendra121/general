package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	myJob := gocron.NewScheduler(time.UTC)
	myJob.Every(5).Second().Do(func() {
		hello()
	})
	myJob.StartBlocking()
}
func hello() {
	fmt.Println("------------------------Hello----------------------")
}
