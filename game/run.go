package game

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// Run is go-routine where we draw game objects and handle events by the user
// returns an error through error channel
func (g *Game) Run(events <-chan sdl.Event) <-chan error {
	errCh := make(chan error)

	go func() {
		defer close(errCh)
		tick := time.NewTicker(12 * time.Millisecond)
		for {
			select {
			case e := <-events:
				if ok := g.handleEvents(e); ok {
					return
				}

			case <-tick.C:
				g.update()

				if g.trex.IsDead() {
					g.gameOver()
					time.Sleep(1 * time.Second)
					g.restart()
				}

				if err := g.draw(); err != nil {
					errCh <- err
				}
			}
		}
	}()

	return errCh
}
