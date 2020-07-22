package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		speed int
		now   time.Time
	)

	fmt.Println(speed, now)

	speed, now = 100, time.Now()
	// variable swapping
	speed = 100
	prevspeed := 50

	speed, prevspeed = prevspeed, speed
}
