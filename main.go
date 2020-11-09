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

//simple function to save Image to png
func saveScreenshot(image image.Image, path string) error {
	//creates image.png
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	//encodes it in to the file
	if err := png.Encode(f, image); err != err {
		f.Close()
		return err
	}
	//and closes the file
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

func main() {
	//piano tiles
	var screensht screenshotWindow
	var r, g, b, a uint32
	//screen variable is used to hold the screenshot of a 4 tiles
	//it is wide, small screenshot (look at image.png) used for detecting
	//individual tiles which the mouse will later click on
	var screen image.Image

	//StartPos.x is x coordinate of the top left corner of the screenshot
	screensht.StartPos.x = 630
	//StartPos.y is y coordinate of the top left corner of the screenshot
	screensht.StartPos.y = 569
	//wight of the screenshot (make sure that it is all over the tiles lenght)
	screensht.Widht = 400
	//height of the screenshot (doesnt really matter that much, it has a little overlap)
	screensht.Height = 10
	//tilesDistance is the distance between individual tiles (in pixels)
	screensht.tilesDistance = 100
	//piano tiles has 4 tiles to click on, right?
	screensht.tilesCount = 4

	sdl.Delay(2000)
	/*
	 * this code til row 77 is for debug only
	 * it is used for tunning the screenshot pos
	 * make sure that the screenshot is on the top of the first row of tiles
	 * example: image.png on the github repo
	 */
	screen, err := screenshot.Capture(screensht.StartPos.x, screensht.StartPos.y, screensht.Widht, screensht.Height)
	if err != nil {
		panic(err)
	}
	//simple function to save Image to png without other bullshittery
	saveScreenshot(screen, "image.png")
	r, g, b, a = screen.At(10+screensht.tilesDistance, screensht.Height/2).RGBA()
	fmt.Println(r, g, b, a)

	//this sets the control colors which are used to decide events (such as click on the tile)
	//if the tile is black (this variable defines our black) then mouse will click there
	var blackcolor uint32 = 4369 + 4369 + 4369 + 65535
	//if the screen or tile is red (this variable defines our red) then the program will end
	var redcolor uint32 = 64507 + 15934 + 14392 + 65535

	var index int
	quit := 0
	//welcome to the main "game" loop
	for quit == 0 {
		//first of all you take a screenshot on the pos we entered earlier (StartPos.x and StartPos.y)
		screen, err = screenshot.Capture(screensht.StartPos.x, screensht.StartPos.y, screensht.Widht, screensht.Height)
		if err != nil {
			panic(err)
		}
		/*
		 * this goes through all tiles places (4) and if the color on the tile pos is same as our black (blackcolor)
		 * then the mouse will go to the poss and click on the tile
		 * if the color is red (redcolor) then the program will end
		 */
		for index = 0; index < screensht.tilesCount; index++ {
			//takes the colors from the right pos of each tile
			r, g, b, a = screen.At(10+(screensht.tilesDistance*index), screensht.Height/2).RGBA()

			//and compares it with predeclared colors
			if r+g+b+a == blackcolor {
				robotgo.MoveMouseSmooth(screensht.StartPos.x+(screensht.tilesDistance*index)+50, screensht.StartPos.y+100, 0.0, 0.0)
				robotgo.MouseClick()
			}
			if r+g+b+a == redcolor {
				quit++
			}
		}
		//my tring to stop memory leak XD
		screen = nil
	}
}
