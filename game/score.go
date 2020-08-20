package game

import (
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func (g *Game) displayScore() error {
	if err := ttf.Init(); err != nil {
		return err
	}
	defer ttf.Quit()

	font, err := ttf.OpenFont("./assets/font/PressStart2P.ttf", 30)
	if err != nil {
		return err
	}
	defer font.Close()

	score := strconv.Itoa(int(g.score))
	white := sdl.Color{R: 255, G: 255, B: 255, A: 255}
	surface, err := font.RenderUTF8Solid(score, white)
	if err != nil {
		return err
	}
	defer surface.Free()

	texture, err := g.renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return err
	}
	defer texture.Destroy()

	dst := &sdl.Rect{X: 700, Y: 450, W: 20, H: 20}
	if g.score > 100 {
		dst.W = 25
	} else if g.score > 1000 {
		dst.W = 30
	}

	g.renderer.Copy(texture, nil, dst)

	return nil
}
