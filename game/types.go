package game

import (
	"time"

	"github.com/go-dino/cactus"
	"github.com/go-dino/clouds"
	"github.com/go-dino/ground"
	"github.com/go-dino/trex"
	"github.com/veandco/go-sdl2/sdl"
)

// Game contains all objects of trex-runner
type Game struct {
	renderer *sdl.Renderer
	ground   *ground.Grounds
	trex     *trex.TRex
	cactus   *cactus.Cactus
	clouds   *clouds.Clouds
	score    float64
}

// NewGame initializes all game objects
func newGame(renderer *sdl.Renderer) (*Game, error) {
	ground, err := ground.NewGrounds(renderer)
	if err != nil {
		return nil, err
	}

	trex, err := trex.NewTrex(renderer)
	if err != nil {
		return nil, err
	}

	cactus, err := cactus.NewCactus(renderer)
	if err != nil {
		return nil, err
	}

	clouds, err := clouds.NewClouds(renderer)
	if err != nil {
		return nil, err
	}

	return &Game{
		renderer: renderer,
		ground:   ground,
		trex:     trex,
		cactus:   cactus,
		clouds:   clouds,
	}, nil
}

// Draw renderes all game objects on screen
func (g *Game) draw() error {
	g.renderer.Clear()
	if err := g.ground.Draw(g.renderer); err != nil {
		return err
	}

	if err := g.trex.Draw(g.renderer); err != nil {
		return err
	}

	if err := g.cactus.Draw(g.renderer); err != nil {
		return err
	}

	if err := g.clouds.Draw(g.renderer); err != nil {
		return err
	}
	if err := g.displayScore(); err != nil {
		return err
	}

	g.score += 0.05

	g.renderer.Present()

	return nil
}

// Update updates all game objects
func (g *Game) update() {
	g.trex.Update()
	g.cactus.Update()
	g.clouds.Update()
	g.ground.Update()
	g.cactus.IsCollide(g.trex)
	g.inCreaseSpeed()
}

func (g *Game) inCreaseSpeed() {

	go func() {
		tick := time.NewTicker(15 * time.Second)
		for {
			select {
			case <-tick.C:
				g.cactus.Mu.Lock()
				g.clouds.Mu.Lock()
				g.ground.Mu.Lock()

				g.cactus.Speed += 0.001
				g.clouds.Speed += 0.001
				g.ground.Speed += 0.001
				// g.ground.SleepTime -= 0.1
				g.cactus.Mu.Unlock()
				g.clouds.Mu.Unlock()
				g.ground.Mu.Unlock()
			}
		}
	}()
}

// destroy destroys all game objects
func (g *Game) destroy() {
	g.trex.Destroy()
	g.cactus.Destroy()
	g.clouds.Destroy()
	g.ground.Destroy()
}
