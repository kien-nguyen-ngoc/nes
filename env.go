package main

import (
	"./nes"
	"./ui"
	"strconv"
	"sync"
	"time"
)

type Environment struct {
	director ui.Director
	audio    ui.Audio
}

const (
	LIVES      = "32"
	SCORE_LOW  = "7E2"
	SCORE_HIGH = "7E3"
	KEY_PRESS  = "F1"
)

// action
func (env *Environment) act() {

}

// get key press
func (env *Environment) GetKeyPress() int {
	view, err := env.director.GetView()
	if err != nil {
		panic(env.director)
	}
	return int(view.GetConsole().Controller2.Read())
}

// set key press
func (env *Environment) SetKeyPress(value byte) {
	_, err := env.director.GetView()
	if err != nil {
		panic(env.director)
	} else {
		env.director.SetAct(int32(value))
	}
}

// get lives
func (env *Environment) GetLives() int {
	view, err := env.director.GetView()
	if err != nil {
		panic(env.director)
	}
	return int(view.GetConsole().RAM[hex2dec(LIVES)])
}

// get scores
func (env *Environment) GetScores() int {
	view, err := env.director.GetView()
	if err != nil {
		panic(env.director)
	}
	console := view.GetConsole()
	return int(console.RAM[hex2dec(SCORE_LOW)]) +
		int(console.RAM[hex2dec(SCORE_HIGH)])*256
}

// start game
func (env *Environment) Start(path string) {
	env.director, env.audio = ui.Init(true, 2)
	env.director.Start([]string{path})
}

// reset game
func (env *Environment) Reset() {
	view, err := env.director.GetView()
	if err != nil {
		panic(env.director)
	}
	console := view.GetConsole()
	console.Reset()
}

// get pixel of current frame in RGBA array for 1 pixel. length of array equal 4 time number of pixel
func (env *Environment) GetFrame() []byte {
	view, err := env.director.GetView()
	if err != nil {
		panic(env.director)
	}
	console := view.GetConsole()
	return console.Buffer().Pix
}

// get control space
func (env *Environment) GetControlSpace() int {
	return 8
}

// get observation space
func (env *Environment) GetObservationSpace() int {
	return len(env.GetFrame())
}

// convert memory address from hex to dec
func hex2dec(hex string) uint32 {
	n, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		panic(err)
	}
	n2 := uint32(n)
	return n2
}

func main() {
	env := new(Environment)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Second)
		for true {
			time.Sleep(1 * time.Second)
			env.SetKeyPress(nes.ButtonA)
			println("press")
		}
	}()

	env.Start("./roms/contra.nes")

	wg.Wait()

}
