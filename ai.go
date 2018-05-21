package main

import (
	"github.com/MaxHalford/gago"
	"math/rand"
	"fmt"
	"time"
	"math"
)

type GameGenome []float64

func sigmoid(score float64) float64 {
	return 1.0 / (1.0 + math.Exp(-score))
}

func transformScore(score float64) float64 {
	return -score
}

func (g GameGenome) Evaluate() (float64, error) {
	game := NewGame(NewArea(10, 10))
	for _, step := range g {
		switch int(step) {
		case 0:
			if err := game.GoForward(); err != nil {
				return transformScore(game.GetScore()), nil
			}
		case 1:
			if err := game.GoLeft(); err != nil {
				return transformScore(game.GetScore()), nil
			}
		case 2:
			if err := game.GoRight(); err != nil {
				return transformScore(game.GetScore()), nil
			}
		}
	}

	return transformScore(game.GetScore()), nil
}

func (g GameGenome) Mutate(rng *rand.Rand) {
	gago.MutNormalFloat64(g, 0.8, rng)
}

func (g GameGenome) Crossover(Y gago.Genome, rng *rand.Rand) {
	gago.CrossUniformFloat64(g, Y.(GameGenome), rng)
}

func (g GameGenome) Clone() gago.Genome {
	var Y = make(GameGenome, len(g))
	copy(Y, g)

	return Y
}

func GameGenomeFactory(rng *rand.Rand) gago.Genome {
	corpus := []float64{
		0.0,
		1.0,
		2.0,
	}
	floats := make([]float64, 10000)
	for i := range floats {
		floats[i] = corpus[rng.Intn(len(corpus))]
	}
	return GameGenome(floats)
}

func GameGenomeRender(g GameGenome) {
	game := NewGame(NewArea(10, 10))
	for _, step := range g {
		CallClear()
		game.Render()

		time.Sleep(100 * time.Millisecond)
		switch int(step) {
		case 0:
			if err := game.GoForward(); err != nil {
				game.Debug()
				return
			}
		case 1:
			if err := game.GoLeft(); err != nil {
				game.Debug()
				return
			}
		case 2:
			if err := game.GoRight(); err != nil {
				game.Debug()
				return
			}
		}
	}
	game.Debug()
}


func RunAI() {
	var ga = gago.Generational(GameGenomeFactory)
	ga.ParallelEval = true
	var err = ga.Initialize()

	if err != nil {
		fmt.Println("Handle init error!")
	}

	fmt.Printf("Best fitness at generation 0: %f\n", ga.HallOfFame[0].Fitness)
	for i := 1; i < 10; i++ {
		err = ga.Evolve()
		if err != nil {
			fmt.Println("Handle error!")
		}
		if math.Mod(float64(i), 100) == 0 {
			fmt.Printf("Best fitness at generation %d: %f\n", i, ga.HallOfFame[0].Fitness)
		}
	}

	GameGenomeRender(ga.HallOfFame[0].Genome.(GameGenome))

}