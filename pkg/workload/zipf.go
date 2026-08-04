package workload

import (
	"math"
	"math/rand"
	"sync/atomic"
)

type ZipfGenerator struct {
	numAccounts int64
	theta       float64
	zetaN       float64
	alpha       float64
	eta         float64
	hottestHits int64
	totalGen    int64
}

func NewZipfGenerator(numAccounts int64, theta float64) *ZipfGenerator {
	if numAccounts <= 0 {
		numAccounts = 10000
	}

	// Avoid 1.0 / (1.0 - theta) division-by-zero singularity when theta == 1.0
	if math.Abs(theta-1.0) < 1e-5 {
		theta = 0.9999
	}

	g := &ZipfGenerator{
		numAccounts: numAccounts,
		theta:       theta,
	}

	if theta > 0.0 {
		g.zetaN = g.computeZeta(numAccounts, theta)
		zeta2 := g.computeZeta(2, theta)
		g.alpha = 1.0 / (1.0 - theta)
		g.eta = (1.0 - math.Pow(2.0/float64(numAccounts), 1.0-theta)) / (1.0 - zeta2/g.zetaN)
	}

	return g
}

func (g *ZipfGenerator) computeZeta(n int64, theta float64) float64 {
	sum := 0.0
	for i := int64(1); i <= n; i++ {
		sum += 1.0 / math.Pow(float64(i), theta)
	}
	return sum
}

// NextAccountID returns a 1-indexed account ID following the Zipfian distribution
func (g *ZipfGenerator) NextAccountID() int64 {
	atomic.AddInt64(&g.totalGen, 1)

	var id int64
	if g.theta <= 0.0 {
		id = rand.Int63n(g.numAccounts) + 1
	} else {
		u := rand.Float64()
		uz := u * g.zetaN

		if uz < 1.0 {
			id = 1
		} else if uz < 1.0+math.Pow(0.5, g.theta) {
			id = 2
		} else {
			id = 1 + int64(float64(g.numAccounts)*math.Pow(g.eta*u-g.eta+1.0, g.alpha))
			if id < 1 {
				id = 1
			} else if id > g.numAccounts {
				id = g.numAccounts
			}
		}
	}

	if id == 1 {
		atomic.AddInt64(&g.hottestHits, 1)
	}

	return id
}

// RealizedHotRowHitRate returns the percentage of generated accounts that hit account ID 1
func (g *ZipfGenerator) RealizedHotRowHitRate() float64 {
	tot := atomic.LoadInt64(&g.totalGen)
	if tot == 0 {
		return 0.0
	}
	hits := atomic.LoadInt64(&g.hottestHits)
	return (float64(hits) / float64(tot)) * 100.0
}
