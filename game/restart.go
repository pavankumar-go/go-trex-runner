package game

func (g *Game) restart() {
	g.trex.Mu.Lock()
	defer g.trex.Mu.Unlock()

	g.cactus.Mu.Lock()
	defer g.cactus.Mu.Unlock()

	g.ground.Mu.Lock()
	defer g.ground.Mu.Unlock()

	g.clouds.Mu.Lock()
	defer g.clouds.Mu.Unlock()

	g.score = 0
	g.trex.Dead = false
	g.cactus.RandomNumber = 5.0
	g.cactus.Speed = 3.0
	g.ground.Speed = 3.0
	g.clouds.Speed = 1.0
	for _, c := range g.cactus.Cactus {
		c.X = 800
	}
}
