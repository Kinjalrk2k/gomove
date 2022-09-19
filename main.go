package main

import (
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

const timeout = 3 * time.Second

func main() {
	robotgo.MouseSleep = 100
	rand.Seed(time.Now().UnixNano())

	crazyMoves()
}

func crazyMoves() {
	sx, sy := robotgo.GetScreenSize()

	for {
		x := rand.Intn(sx)
		y := rand.Intn(sy)

		robotgo.KeyTap("shift")
		robotgo.MoveSmooth(x, y)
		time.Sleep(timeout)
	}
}
