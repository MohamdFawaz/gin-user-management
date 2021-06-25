package lib

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)


type Cronjob struct {

}

func NewCronjob() Cronjob {
	return Cronjob{}
}

func task() {
	fmt.Println("I am running task in background every 10 seconds.")
}

func (cronjob Cronjob) SetupJobs() {
	gocron.Every(10).Seconds().Do(task)
	_, time := gocron.NextRun()
	fmt.Println(time)

	<- gocron.Start()
}