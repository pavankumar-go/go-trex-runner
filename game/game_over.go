package game

import (
	"log"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// GameOver loads and draws game-over
func (g *Game) gameOver() {
	texture, err := img.LoadTexture(g.renderer, "./assets/img/game-over.png")
	if err != nil {
		log.Fatalf("could not load font : %v", err)
		return
	}
	defer texture.Destroy()

	rect := &sdl.Rect{X: 300, Y: 500, W: 190, H: 15}
	if err := g.renderer.Copy(texture, nil, rect); err != nil {
		log.Fatalf("could not copy texture : %v", err)
		return
	}

	g.renderer.Present()
	return
}
