package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	colors "golang.org/x/image/colornames"
	"golang.org/x/image/font"
)

// loadTTF loads a new TTF file with a font.
func loadTTF(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)

	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Joystick test",
		Bounds: pixel.R(0, 0, 1280, 720),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)

	if err != nil {
		panic(err)
	}

	fmt.Println("Gamepad is present:", win.JoystickPresent(pixelgl.Joystick1))
	fmt.Println("Name:", win.JoystickName(pixelgl.Joystick1))
	fmt.Println("Button count:", win.JoystickButtonCount(pixelgl.Joystick1))
	fmt.Println("Axis count:", win.JoystickAxisCount(pixelgl.Joystick1))

	// Load a new font.
	face, err := loadTTF("joystix-monospace.ttf", 40)

	if err != nil {
		panic(err)
	}

	atlas := text.NewAtlas(face, text.ASCII)
	txt := text.New(pixel.V(20, 670), atlas)

	for !win.Closed() {
		win.Clear(colors.Blue)

		// Check for stick input.
		lsX := win.JoystickAxis(pixelgl.Joystick1, pixelgl.AxisLeftX)
		lsY := win.JoystickAxis(pixelgl.Joystick1, pixelgl.AxisLeftY)

		rsX := win.JoystickAxis(pixelgl.Joystick1, pixelgl.AxisRightX)
		rsY := win.JoystickAxis(pixelgl.Joystick1, pixelgl.AxisRightY)

		// Check for trigger input.
		lt := win.JoystickAxis(pixelgl.Joystick1, pixelgl.AxisLeftTrigger)
		rt := win.JoystickAxis(pixelgl.Joystick1, pixelgl.AxisRightTrigger)

		// Output axis values.
		txt.Clear()
		txt.Color = colors.White
		fmt.Fprintln(txt, "Left stick X:", lsX)
		fmt.Fprintln(txt, "Left stick Y:", lsY)

		fmt.Fprintln(txt, "Right stick X:", rsX)
		fmt.Fprintln(txt, "Right stick Y:", rsY)

		fmt.Fprintln(txt, "Left trigger:", lt)
		fmt.Fprintln(txt, "Right trigger:", rt)

		// Check for ABYX input.
		if win.JoystickPressed(pixelgl.Joystick1, pixelgl.ButtonA) {
			fmt.Fprintln(txt, "Button A (Cross) pressed")
		}

		if win.JoystickPressed(pixelgl.Joystick1, pixelgl.ButtonB) {
			fmt.Fprintln(txt, "Button B (Circle) pressed")
		}

		if win.JoystickPressed(pixelgl.Joystick1, pixelgl.ButtonY) {
			fmt.Fprintln(txt, "Button Y (Triangle) pressed")
		}

		if win.JoystickPressed(pixelgl.Joystick1, pixelgl.ButtonX) {
			fmt.Fprintln(txt, "Button X (Square) pressed")
		}

		// Check for DPad input.
		if win.JoystickPressed(pixelgl.Joystick1, pixelgl.ButtonDpadUp) {
			fmt.Fprintln(txt, "Button Up pressed")
		}

		if win.JoystickPressed(pixelgl.Joystick1, pixelgl.ButtonDpadDown) {
			fmt.Fprintln(txt, "Button Down pressed")
		}

		if win.JoystickPressed(pixelgl.Joystick1, pixelgl.ButtonDpadLeft) {
			fmt.Fprintln(txt, "Button Left pressed")
		}

		if win.JoystickPressed(pixelgl.Joystick1, pixelgl.ButtonDpadRight) {
			fmt.Fprintln(txt, "Button Right pressed")
		}

		// Check for sticks input.
		if win.JoystickPressed(pixelgl.Joystick1, pixelgl.ButtonLeftThumb) {
			fmt.Fprintln(txt, "Left trigger pressed")
		}

		if win.JoystickPressed(pixelgl.Joystick1, pixelgl.ButtonRightThumb) {
			fmt.Fprintln(txt, "Right trigger pressed")
		}

		// Check for sticks pressing input.
		if win.JoystickPressed(pixelgl.Joystick1, pixelgl.ButtonLeftBumper) {
			fmt.Fprintln(txt, "Button LB (L1) pressed")
		}

		if win.JoystickPressed(pixelgl.Joystick1, pixelgl.ButtonRightBumper) {
			fmt.Fprintln(txt, "Button RB (R1) pressed")
		}

		// Check for Start and Back input.
		if win.JoystickPressed(pixelgl.Joystick1, pixelgl.ButtonStart) {
			fmt.Fprintln(txt, "Button Start pressed")
		}

		if win.JoystickPressed(pixelgl.Joystick1, pixelgl.ButtonBack) {
			fmt.Fprintln(txt, "Button Back pressed")
		}

		txt.Draw(win, pixel.IM)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
