package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	CountdownStart = 3
	FinalWord      = "Go!"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const (
	write = "write"
	sleep = "sleep"
)

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := CountdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, FinalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
