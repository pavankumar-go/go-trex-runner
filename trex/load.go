package trex

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// NewTrex returns a pointer to TRex with preset values
func NewTrex(renderer *sdl.Renderer) (*TRex, error) {
	trex := &TRex{
		X: 10,
		W: 40,
		H: 40,
	}

	for i := 1; i <= 3; i++ {
		path := fmt.Sprintf("./assets/img/trex-%d.png", i)
		texture, err := img.LoadTexture(renderer, path)
		if err != nil {
			return nil, err
		}

		trex.Texture = append(trex.Texture, texture)
	}

	return trex, nil
}
