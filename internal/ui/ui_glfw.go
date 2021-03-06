// Copyright 2015 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build !js

package ui

import (
	"fmt"
	glfw "github.com/go-gl/glfw3"
	"runtime"
	"time"
)

var currentUI *userInterface

func Init() {
	runtime.LockOSThread()

	glfw.SetErrorCallback(func(err glfw.ErrorCode, desc string) {
		panic(fmt.Sprintf("%v: %v\n", err, desc))
	})
	if !glfw.Init() {
		panic("glfw.Init() fails")
	}
	glfw.WindowHint(glfw.Visible, glfw.False)
	glfw.WindowHint(glfw.Resizable, glfw.False)

	window, err := glfw.CreateWindow(16, 16, "", nil, nil)
	if err != nil {
		panic(err)
	}

	u := &userInterface{
		window: window,
		funcs:  make(chan func()),
	}
	go func() {
		runtime.LockOSThread()
		u.window.MakeContextCurrent()
		glfw.SwapInterval(1)
		for f := range u.funcs {
			f()
		}
	}()

	currentUI = u
}

func ExecOnUIThread(f func()) {
	ch := make(chan struct{})
	currentUI.funcs <- func() {
		defer close(ch)
		f()
	}
	<-ch
}

func Start(width, height, scale int, title string) (actualScale int, err error) {
	return currentUI.start(width, height, scale, title)
}

func Terminate() {
	currentUI.terminate()
}

func DoEvents() error {
	return currentUI.doEvents()
}

func IsClosed() bool {
	return currentUI.isClosed()
}

func SwapBuffers() {
	currentUI.swapBuffers()
}

type userInterface struct {
	window *glfw.Window
	scale  int
	funcs  chan func()
}

func (u *userInterface) start(width, height, scale int, title string) (actualScale int, err error) {
	monitor, err := glfw.GetPrimaryMonitor()
	if err != nil {
		return 0, err
	}
	videoMode, err := monitor.GetVideoMode()
	if err != nil {
		return 0, err
	}
	x := (videoMode.Width - width*scale) / 2
	y := (videoMode.Height - height*scale) / 3

	ch := make(chan struct{})
	window := u.window
	window.SetFramebufferSizeCallback(func(w *glfw.Window, width, height int) {
		window.SetFramebufferSizeCallback(nil)
		close(ch)
	})
	window.SetSize(width*scale, height*scale)
	window.SetTitle(title)
	window.SetPosition(x, y)
	window.Show()

	for {
		done := false
		glfw.PollEvents()
		select {
		case <-ch:
			done = true
		default:
		}
		if done {
			break
		}
	}

	u.scale = scale

	// For retina displays, recalculate the scale with the framebuffer size.
	windowWidth, _ := window.GetFramebufferSize()
	actualScale = windowWidth / width

	return actualScale, nil
}

func (u *userInterface) pollEvents() error {
	glfw.PollEvents()
	return updateInput(u.window, u.scale)
}

func (u *userInterface) doEvents() error {
	if err := u.pollEvents(); err != nil {
		return err
	}
	for u.window.GetAttribute(glfw.Focused) == 0 {
		time.Sleep(time.Second / 60)
		if err := u.pollEvents(); err != nil {
			return err
		}
	}
	return nil
}

func (u *userInterface) terminate() {
	glfw.Terminate()
}

func (u *userInterface) isClosed() bool {
	return u.window.ShouldClose()
}

func (u *userInterface) swapBuffers() {
	u.window.SwapBuffers()
}
