package game

import (
	"github.com/veandco/go-sdl2/sdl"
)

// handleEvents handles events for a particular window action by user
func (g *Game) handleEvents(event sdl.Event) bool {
	switch e := event.(type) {
	case *sdl.QuitEvent:
		return true
	case *sdl.KeyboardEvent:
		if (e.Keysym.Sym == 32) && (e.State == 1) { // SPACE_BAR KEYCODE 32 &&  KEY_PRESSED 1
			if g.trex.Y > 0 {
				return false // avoiding incremental jumps
			}
			g.trex.Jump()
		}
	default:
		return false
	}
	return false
}
