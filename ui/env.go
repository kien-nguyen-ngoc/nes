package ui

import (
	"../nes"
	"strconv"
)

type Environment struct {
	console *nes.Console
}

const (
	LIVES      = "32"
	SCORE_LOW  = "7E2"
	SCORE_HIGH = "7E3"
)

// get lives
func (env *Environment) getLives() int {
	return int(env.console.RAM[hex2Dec(LIVES)])
}

// get scores
func (env *Environment) getScores() int {
	return int(env.console.RAM[hex2Dec(SCORE_LOW)]) + int(env.console.RAM[hex2Dec(SCORE_HIGH)])*255
}

// reset game
func (env *Environment) reset() {
	env.console.Reset()
}

// get pixel of current frame in R.G.B.A array for 1 pixel. length of array equal 4 time number of pixel
func (env *Environment) getFrame() []byte {
	return env.console.Buffer().Pix
}

// convert memory address from hex to dec
func hex2Dec(hex string) uint32 {
	n, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		panic(err)
	}
	n2 := uint32(n)
	return n2
}
