package game

import (
	"fmt"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
)

// Setup initializes sdl by setting window properties
func Setup() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}
	defer sdl.Quit()

	window, renderer, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}
	defer window.Destroy()
	defer renderer.Destroy()

	window.SetTitle("T-REX Runner in Go")

	sdl.PumpEvents() // for mac

	if err := loadGame(renderer); err != nil {
		return fmt.Errorf(" Error Loading Game due to: %v", err)
	}

	return nil
}

// loadGame loads game
func loadGame(renderer *sdl.Renderer) error {
	game, err := newGame(renderer)
	if err != nil {
		return err
	}
	defer game.destroy()

	eventCh := make(chan sdl.Event)
	defer close(eventCh)

	errCh := game.Run(eventCh) // run the game

	runtime.LockOSThread() // locks the goroutine to run in the same thread when it began until program exists
	for {
		select {
		case eventCh <- sdl.WaitEvent():
		case err := <-errCh:
			return err
		}
	}
}
