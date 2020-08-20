package clouds

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Draw iterates through all clouds and calls draw for each cloud
func (c *Clouds) Draw(renderer *sdl.Renderer) error {
	c.Mu.RLock()
	defer c.Mu.RUnlock()

	for _, cloud := range c.clouds {
		if err := cloud.Draw(renderer, c.Texture); err != nil {
			return err
		}
	}
	return nil
}

// Draw draws cloud on window
func (c *Cloud) Draw(renderer *sdl.Renderer, texture *sdl.Texture) error {
	rect := &sdl.Rect{X: c.X, Y: c.Y, W: c.W, H: c.H}
	if err := renderer.Copy(texture, nil, rect); err != nil {
		return err
	}

	return nil
}
