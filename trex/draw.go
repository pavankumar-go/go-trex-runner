package trex

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Draw  draws trex on window
func (t *TRex) Draw(renderer *sdl.Renderer) error {
	t.Mu.RLock()
	defer t.Mu.RUnlock()

	rect := TRex{Rect: sdl.Rect{
		X: t.X,
		Y: 600 - t.Y - t.H - 10, // 550 ++
		W: t.W,
		H: t.H,
	}}

	i := t.time / 7 % len(t.Texture)
	if err := renderer.Copy(t.Texture[i], nil, &rect.Rect); err != nil {
		return err
	}

	return nil
}
