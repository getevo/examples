package main

///
import (
	"fmt"
	"github.com/getevo/evo"
	"github.com/getevo/evo/lib/schedule"
	"time"
)

type Runner struct {
	Round int
}

func (o *Runner) Run() error {
	o.Round++
	fmt.Println("Runner 1:", o.Round)
	return nil
}

type Runner2 struct {
	Round int
}

func (o *Runner2) Run() error {
	o.Round++
	fmt.Println("Runner 2:", o.Round)
	return nil
}

///
func main() {
	evo.Setup()
	runner1 := &Runner{}
	runner1.Round = 10

	runner2 := &Runner2{}

	scheduler := schedule.New()
	scheduler.RepeatN(12, 1*time.Second, runner1)
	scheduler.Every(5*time.Second, runner2)
	evo.Run()
}

///
