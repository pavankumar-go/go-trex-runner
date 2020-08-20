package trex

import (
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	gravity   = 0.5
	jumpSpeed = 10
)

// TRex holds texture, position & image options for trex
type TRex struct {
	Mu       sync.RWMutex
	Texture  []*sdl.Texture
	Renderer *sdl.Renderer
	X, Y     int32
	W, H     int32
	time     int
	speed    float32
	Rect     sdl.Rect
	Dead     bool
}

// Update ..
func (t *TRex) Update() {
	t.Mu.Lock()
	defer t.Mu.Unlock()
	t.time++
	t.Y -= int32(t.speed)
	t.speed += gravity
	if t.Y < 0 {
		t.Y = 0
	}
}

// Jump  ..
func (t *TRex) Jump() {
	t.Mu.Lock()
	defer t.Mu.Unlock()
	t.speed = -jumpSpeed
	t.X = 10
	t.W = 40
	t.H = 40
}

// Destroy destroys texture on exit
func (t *TRex) Destroy() {
	t.Mu.Lock()
	defer t.Mu.Unlock()
	for _, texture := range t.Texture {
		texture.Destroy()
	}
}

// IsDead ..
func (t *TRex) IsDead() bool {
	t.Mu.RLock()
	defer t.Mu.RUnlock()
	return t.Dead
}
