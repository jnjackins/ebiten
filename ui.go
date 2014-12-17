/*
Copyright 2014 Hajime Hoshi

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ebiten

import (
	"errors"
	"fmt"
	glfw "github.com/go-gl/glfw3"
	"image"
	"runtime"
)

type ui struct {
	window          *glfw.Window
	scale           int
	graphicsContext *graphicsContext
	input           input
	funcs           chan func()
}

func newUI(width, height, scale int, title string) (*ui, error) {
	glfw.SetErrorCallback(func(err glfw.ErrorCode, desc string) {
		panic(fmt.Sprintf("%v: %v\n", err, desc))
	})
	if !glfw.Init() {
		return nil, errors.New("glfw.Init() fails")
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	window, err := glfw.CreateWindow(width*scale, height*scale, title, nil, nil)
	if err != nil {
		return nil, err
	}

	u := &ui{
		window: window,
		scale:  scale,
		funcs:  make(chan func()),
	}

	u.run(width, height, scale)

	// For retina displays, recalculate the scale with the framebuffer size.
	windowWidth, _ := window.GetFramebufferSize()
	realScale := windowWidth / width
	u.use(func() {
		u.graphicsContext, err = newGraphicsContext(width, height, realScale)
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *ui) doEvents() {
	glfw.PollEvents()
	u.update()
}

func (u *ui) terminate() {
	glfw.Terminate()
}

func (u *ui) isClosed() bool {
	return u.window.ShouldClose()
}

func (u *ui) Sync(f func()) {
	u.use(f)
}

func (u *ui) draw(f func(GraphicsContext) error) (err error) {
	u.use(func() {
		u.graphicsContext.preUpdate()
	})
	if err = f(&syncGraphicsContext{
		syncer:               u,
		innerGraphicsContext: u.graphicsContext,
	}); err != nil {
		return
	}
	u.use(func() {
		u.graphicsContext.postUpdate()
		u.window.SwapBuffers()
	})
	return
}

func (u *ui) newTextureID(img image.Image, filter int) (TextureID, error) {
	var id TextureID
	var err error
	u.use(func() {
		id, err = idsInstance.createTexture(img, filter)
	})
	return id, err
}

func (u *ui) newRenderTargetID(width, height int, filter int) (RenderTargetID, error) {
	var id RenderTargetID
	var err error
	u.use(func() {
		id, err = idsInstance.createRenderTarget(width, height, filter)
	})
	return id, err
}

func (u *ui) run(width, height, scale int) {
	go func() {
		runtime.LockOSThread()
		u.window.MakeContextCurrent()
		glfw.SwapInterval(1)
		for f := range u.funcs {
			f()
		}
	}()
}

func (u *ui) use(f func()) {
	ch := make(chan struct{})
	u.funcs <- func() {
		f()
		close(ch)
	}
	<-ch
}

func (u *ui) update() {
	u.input.update(u.window, u.scale)
}