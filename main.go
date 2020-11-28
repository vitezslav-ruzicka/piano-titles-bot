package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/kbinani/screenshot"
	"github.com/veandco/go-sdl2/sdl"
	"image"
	"image/png"
	"os"
)

var blackcolor uint32 = 4369 + 4369 + 4369 + 65535
var redcolor uint32 = 65535 + 65535
var quit = 0

type pos struct {
	y, x int
}

type screenshotWindow struct {
	StartPos      pos
	Widht         int
	Height        int
	tilesDistance int
	tilesCount    int
}

func saveScreenshot(image image.Image, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	if err := png.Encode(f, image); err != err {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

func (screensht *screenshotWindow) clickTile(screenChan chan image.Image, routineNumber int) {
	var r, g, b, a uint32
	var screen image.Image

	for r+g+b+a != redcolor {
		screen = <-screenChan
		r, g, b, a = screen.At(10+(screensht.tilesDistance*routineNumber), screensht.Height/2).RGBA()

		if r+g+b+a == blackcolor {
			robotgo.MoveMouseSmooth(screensht.StartPos.x+(screensht.tilesDistance*routineNumber)+50, screensht.StartPos.y+100, 0.0, 0.0)
			robotgo.MouseClick()
		}
	}
	quit = 1
}

func main() {
	//piano tiles
	var screensht screenshotWindow
	screensht.StartPos.x = 630
	screensht.StartPos.y = 569
	screensht.Widht = 400
	screensht.Height = 10
	screensht.tilesDistance = 100
	screensht.tilesCount = 4

	sdl.Delay(2000)

	var r, g, b, a uint32
	var screen image.Image
	screenChan := make(chan image.Image)
	screen, err := screenshot.Capture(screensht.StartPos.x, screensht.StartPos.y, screensht.Widht, screensht.Height)
	if err != nil {
		panic(err)
	}
	saveScreenshot(screen, "image.png")
	r, g, b, a = screen.At(10+screensht.tilesDistance, screensht.Height/2).RGBA()
	fmt.Println(r, g, b, a)

	go screensht.clickTile(screenChan, 0)
	go screensht.clickTile(screenChan, 1)
	go screensht.clickTile(screenChan, 2)
	go screensht.clickTile(screenChan, 3)

	for quit == 0 {
		screen, err = screenshot.Capture(screensht.StartPos.x, screensht.StartPos.y, screensht.Widht, screensht.Height)
		if err != nil {
			panic(err)
		}

		screenChan <- screen
		screen = nil
	}
}
