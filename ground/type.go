package ground

import (
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

// Grounds holds texture, position & image options for grounds
type Grounds struct {
	Mu       sync.RWMutex
	Renderer *sdl.Renderer
	Texture  *sdl.Texture
	grounds  []*Ground
	Speed    float32
	// SleepTime int32
}

// Ground holds texture, position & image options for ground
type Ground struct {
	Mu   sync.RWMutex
	X, Y int32
	W, H int32
}

// Update reduces the X axis by certain speed which makes ground move
func (g *Grounds) Update() {
	g.Mu.Lock()
	defer g.Mu.Unlock()
	for _, i := range g.grounds {
		i.Mu.Lock()
		i.X -= int32(g.Speed)
		i.Mu.Unlock()
	}
}

// Destroy destroys texture on exit
func (g *Grounds) Destroy() {
	g.Mu.Lock()
	defer g.Mu.Unlock()
	g.Texture.Destroy()
}
