package ui

import (
	"log"

	"../nes"
	"errors"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

type View interface {
	Enter()
	Exit()
	Update(t, dt float64, action ...byte)
	GetConsole() nes.Console
	GetTexture() uint32
}

type Director struct {
	window    *glfw.Window
	audio     *Audio
	view      View
	menuView  View
	timestamp float64
	auto      bool
	act       int32
	speed     int32
}

func NewDirector(window *glfw.Window, audio *Audio, auto bool, speed int32) *Director {
	director := Director{}
	director.window = window
	director.audio = audio
	director.auto = auto
	director.act = -1
	director.speed = speed
	return &director
}
func (d *Director) SetAct(act int32) {
	d.act = act
}

func (d *Director) GetWindow() (glfw.Window, error) {
	if d.view != nil {
		return *d.window, nil
	}
	return *new(glfw.Window), errors.New("Get Window return nil pointer")
}

func (d *Director) GetView() (View, error) {
	if d.view != nil {
		return d.view, nil
	}
	return nil, errors.New("Get View return nil pointer")
}
func (d *Director) GetTexture() (uint32, error) {
	if d.view != nil {
		return d.view.GetTexture(), nil
	}
	return 0, errors.New("Get Texture return nil pointer")
}

func (d *Director) SetTitle(title string) {
	d.window.SetTitle(title)
}

func (d *Director) SetView(view View) {
	if d.view != nil {
		d.view.Exit()
	}
	d.view = view
	if d.view != nil {
		d.view.Enter()
	}
	d.timestamp = glfw.GetTime()
}

func (d *Director) Step(action ...byte) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	timestamp := glfw.GetTime()
	dt := timestamp - d.timestamp
	d.timestamp = timestamp
	if d.view != nil {
		if action != nil {
			d.view.Update(timestamp, dt, action[0])
		} else {
			d.view.Update(timestamp, dt)
		}
	}
}

func (d *Director) Start(paths []string) {
	d.menuView = NewMenuView(d, paths)
	if len(paths) == 1 {
		d.PlayGame(paths[0])
	} else {
		d.ShowMenu()
	}
	d.Run()
}

func (d *Director) Run() {
	for !d.window.ShouldClose() {
		d.Step()
		d.window.SwapBuffers()
		glfw.PollEvents()
	}
	d.SetView(nil)
}

func (d *Director) PlayGame(path string) {
	hash, err := hashFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	console, err := nes.NewConsole(path)
	if err != nil {
		log.Fatalln(err)
	}
	view := NewGameView(d, console, path, hash)
	d.SetView(view)
	console.Reset()

}

func (d *Director) ShowMenu() {
	if d.menuView != nil {
		d.SetView(d.menuView)
	}
}
