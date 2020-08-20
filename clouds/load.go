package clouds

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// NewClouds returns a pointer to Clouds with preset values
func NewClouds(renderer *sdl.Renderer) (*Clouds, error) {
	path := fmt.Sprintf("./assets/img/cloud.png")
	texture, err := img.LoadTexture(renderer, path)
	if err != nil {
		return nil, err
	}

	clouds := &Clouds{
		Texture:  texture,
		Renderer: renderer,
		Speed:    1.0,
	}

	go func() {
		for {
			clouds.clouds = append(clouds.clouds, NewCloud())
			time.Sleep(4 * time.Second) // sleeps for 5 seconds before displaying next cloud on wondow
		}
	}()

	return clouds, nil
}

// NewCloud returns a pointer to Cloud with preset values
func NewCloud() *Cloud {
	cloud := &Cloud{
		X: 800,
		Y: 400,
		W: 46,
		H: 13,
	}

	return cloud
}
