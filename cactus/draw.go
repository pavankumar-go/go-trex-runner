package cactus

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Draw draws cactus on window
func (c *Cactus) Draw(renderer *sdl.Renderer) error {
	c.Mu.RLock()
	defer c.Mu.RUnlock()
	var i int
	for _, cacti := range c.Cactus {
		if i == 2 {
			i = 0
		}
		if err := cacti.Draw(renderer, c.Texture[i]); err != nil {
			return err
		}
		i++
	}
	return nil
}

// Draw draws cactus on window
func (c *Cacti) Draw(renderer *sdl.Renderer, texture *sdl.Texture) error {
	c.Mu.RLock()
	defer c.Mu.RUnlock()
	rect := &sdl.Rect{X: c.X, Y: c.Y, W: c.W, H: c.H}
	if err := renderer.Copy(texture, nil, rect); err != nil {
		return err
	}

	return nil
}
