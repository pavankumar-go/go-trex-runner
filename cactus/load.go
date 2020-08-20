package cactus

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// NewCactus returns a pointer to Cactus with preset values
func NewCactus(renderer *sdl.Renderer) (*Cactus, error) {
	var path string

	cactus := &Cactus{
		Renderer:     renderer,
		Speed:        3.0,
		RandomNumber: 5.0,
	}

	for i := 1; i <= 2; i++ {
		path = fmt.Sprintf("./assets/img/cacti-%d.png", i)
		texture, err := img.LoadTexture(renderer, path)
		if err != nil {
			return nil, err
		}
		cactus.Texture = append(cactus.Texture, texture)
	}

	go func() {
		for {
			cactus.Mu.Lock()
			cactus.Cactus = append(cactus.Cactus, NewCacti())
			cactus.Mu.Unlock()
			time.Sleep(time.Duration(genRandom(int(cactus.RandomNumber), &cactus.Mu)) * time.Second)
		}
	}()

	return cactus, nil
}

func genRandom(random int, mu *sync.RWMutex) int {
	mu.Lock()
	defer mu.Unlock()
	i := 0
	rand.Seed(time.Now().UnixNano())
	p := rand.Perm(random)
	for _, r := range p[:random] {
		i = r
	}

	return i
}

// NewCacti returns a pointer to Cacti with preset values
func NewCacti() *Cacti {
	cacti := &Cacti{
		X: 800,
		Y: 600 - 40 - 10, // 550
		W: 35,
		H: 40,
	}

	return cacti
}
