package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const screenWidth = 640
const screenHeight = 480

var gWindow *sdl.Window
var gScreenSurface *sdl.Surface
var gHelloWorld *sdl.Surface
var gKeyPressSurfaces [keyPressSurfaceTotal]*sdl.Surface
var gCurrentSurface *sdl.Surface
var gRenderer *sdl.Renderer
var gTexture *sdl.Texture

const (
	keyPressSurfaceDefault = iota
	keyPressSurfaceUp
	keyPressSurfaceDown
	keyPressSurfaceLeft
	keyPressSurfaceRight
	keyPressSurfaceTotal
)

func init() {
	var err error

	if err = sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}

	if gWindow, err = sdl.CreateWindow(
		"Tutorial!",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		screenWidth,
		screenHeight,
		sdl.WINDOW_SHOWN); err != nil {
		panic(err)
	}

	gRenderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)
	imgFlags := img.INIT_PNG
	err = img.Init(imgFlags)
	if err != nil {
		panic(err)
	}

	if gScreenSurface, err = gWindow.GetSurface(); err != nil {
		panic(err)
	}
}

func main() {
	quit := false
	loadMedia()
	gCurrentSurface := gKeyPressSurfaces[keyPressSurfaceDefault]

	for !quit {
		e := sdl.PollEvent()
		for e != nil {
			if e.GetType() == sdl.QUIT {
				quit = true
			} else if e.GetType() == sdl.KEYDOWN {
				switch e.(*sdl.KeyboardEvent).Keysym.Sym {

				case sdl.K_UP:
					gCurrentSurface = gKeyPressSurfaces[keyPressSurfaceUp]

				case sdl.K_DOWN:
					gCurrentSurface = gKeyPressSurfaces[keyPressSurfaceDown]

				case sdl.K_LEFT:
					gCurrentSurface = gKeyPressSurfaces[keyPressSurfaceLeft]

				case sdl.K_RIGHT:
					gCurrentSurface = gKeyPressSurfaces[keyPressSurfaceRight]

				default:
					gCurrentSurface = gKeyPressSurfaces[keyPressSurfaceDefault]
				}
			}

			gRenderer.Clear()
			gRenderer.Copy(gTexture, nil, nil)
			gRenderer.Present()

			var stretchRect sdl.Rect
			stretchRect.X = 0
			stretchRect.Y = 0
			stretchRect.W = screenWidth
			stretchRect.H = screenHeight
			gCurrentSurface.Blit(nil, gScreenSurface, &stretchRect)
			gWindow.UpdateSurface()
			e = sdl.PollEvent()
		}

	}

	close()
}

func loadMedia() {
	gKeyPressSurfaces[keyPressSurfaceDefault] = loadSurface("media/texture.png")
	gKeyPressSurfaces[keyPressSurfaceUp] = loadSurface("media/up.bmp")
	gKeyPressSurfaces[keyPressSurfaceDown] = loadSurface("media/down.bmp")
	gKeyPressSurfaces[keyPressSurfaceLeft] = loadSurface("media/left.bmp")
	gKeyPressSurfaces[keyPressSurfaceRight] = loadSurface("media/right.bmp")

}

func close() {
	gScreenSurface.Free()
	gCurrentSurface.Free()
	gTexture.Destroy()
	gRenderer.Destroy()
	img.Quit()
	sdl.Quit()
}

func loadSurface(path string) (optimizedSurface *sdl.Surface) {
	var err error
	loadedSurface, err := img.Load(path)
	if err != nil {
		fmt.Printf("%v", err)
	} else {
		optimizedSurface, err = loadedSurface.Convert(gScreenSurface.Format, 0)
		if err != nil {
			fmt.Printf("%v", err)
		}
	}
	loadedSurface.Free()
	return
}

func loadTexture(path string) (newTexture *sdl.Texture) {
	loadedSurface := loadSurface(path)
	gRenderer.CreateTextureFromSurface(loadedSurface)
	loadedSurface.Free()
	return
}
