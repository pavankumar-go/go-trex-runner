package cactus

import (
	"sync"

	"github.com/go-dino/trex"
	"github.com/veandco/go-sdl2/sdl"
)

// Cactus holds texture, position & image options for cactus
type Cactus struct {
	Mu           sync.RWMutex
	Renderer     *sdl.Renderer
	Texture      []*sdl.Texture
	Cactus       []*Cacti
	Speed        float32
	RandomNumber float32
}

// Cacti holds texture, position & image options for cacti
type Cacti struct {
	Mu   sync.RWMutex
	X, Y int32
	W, H int32
}

// Update reduces the X axis by certain speed which makes cactus move towards trex
func (c *Cactus) Update() {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	for _, i := range c.Cactus {
		i.Mu.Lock()
		i.X -= int32(c.Speed)
		i.Mu.Unlock()
	}
}

// Destroy destroys texture on exit
func (c *Cactus) Destroy() {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	for _, texture := range c.Texture {
		texture.Destroy()
	}
}

// IsCollide function iterates all cacti to check if it has collided with trex
func (c *Cactus) IsCollide(t *trex.TRex) {
	c.Mu.RLock()
	defer c.Mu.RUnlock()
	for _, cacti := range c.Cactus {
		cacti.isCollide(t)
	}
}

// isCollide function checks if cacti is in contact with trex
func (c *Cacti) isCollide(t *trex.TRex) {
	if t.X > c.X+c.W { // trex is too far left from cactus
		return
	}

	if (t.X + t.W) < c.X { // trex is too far right from cactus
		return
	}

	if c.Y > 550-t.Y+t.H/2 { // when trex is above cactus
		return
	}

	t.Dead = true
}
