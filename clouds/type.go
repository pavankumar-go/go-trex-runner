package clouds

import (
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

// Clouds holds texture, position & image options for clouds
type Clouds struct {
	Mu       sync.RWMutex
	Renderer *sdl.Renderer
	Texture  *sdl.Texture
	clouds   []*Cloud
	Speed    float32
}

// Cloud holds texture, position & image options for cacti
type Cloud struct {
	Mu   sync.RWMutex
	X, Y int32
	W, H int32
}

// Update ..
func (c *Clouds) Update() {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	for _, i := range c.clouds {
		i.Mu.Lock()
		i.X -= int32(c.Speed)
		i.Mu.Unlock()
	}
}

// Destroy destroys texture on exit
func (c *Clouds) Destroy() {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Texture.Destroy()
}
