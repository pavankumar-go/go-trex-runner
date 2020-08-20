package ground

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Draw iterates through all grounds and calls draw on each ground
func (g *Grounds) Draw(renderer *sdl.Renderer) error {
	g.Mu.Lock()
	defer g.Mu.Unlock()
	for _, ground := range g.grounds {
		if err := ground.Draw(renderer, g.Texture); err != nil {
			return err
		}
	}
	return nil
}

// Draw draws ground on window
func (g *Ground) Draw(renderer *sdl.Renderer, texture *sdl.Texture) error {
	g.Mu.Lock()
	defer g.Mu.Unlock()
	rect := &sdl.Rect{X: g.X, Y: g.Y, W: g.W, H: g.H}
	if err := renderer.Copy(texture, nil, rect); err != nil {
		return err
	}

	return nil
}
