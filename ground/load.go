package ground

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// NewGrounds returns a pointer to Grounds with preset values
func NewGrounds(renderer *sdl.Renderer) (*Grounds, error) {
	path := fmt.Sprintf("./assets/img/ground.png")
	texture, err := img.LoadTexture(renderer, path)
	if err != nil {
		return nil, err
	}

	grounds := &Grounds{
		Texture:   texture,
		Renderer:  renderer,
		Speed:     3.0,
		SleepTime: 3.0,
	}

	go func() {
		for {
			grounds.grounds = append(grounds.grounds, NewGround())
			time.Sleep(3 * time.Second)
		}
	}()

	return grounds, nil
}

// NewGround returns a pointer to Ground with preset values
func NewGround() *Ground {
	ground := &Ground{
		X: 0,
		Y: 600 - 20,
		W: 1600,
		H: 20,
	}

	return ground
}
