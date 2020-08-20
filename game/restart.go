package game

func (g *Game) restart() {
	g.trex.Mu.Lock()
	g.cactus.Mu.Lock()
	g.ground.Mu.Lock()
	g.clouds.Mu.Lock()

	g.score = 0
	g.trex.Dead = false
	g.cactus.RandomNumber = 5.0
	g.cactus.Speed = 3.0
	g.ground.Speed = 3.0
	g.clouds.Speed = 1.0

	for _, c := range g.cactus.Cactus {
		c.Mu.Lock()
		c.X = 800
		c.Mu.Unlock()
	}

	g.trex.Mu.Unlock()
	g.cactus.Mu.Unlock()
	g.ground.Mu.Unlock()
	g.clouds.Mu.Unlock()

}
