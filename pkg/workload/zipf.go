package workload

import (
	"math"
	"math/rand"
	"sync/atomic"
)

// ZipfGenerator holds the exact cumulative Zipf(pmf ~ 1/i^theta) distribution
// over [1, numAccounts]. The cumulative mass table is computed once (finite
// population, so any theta >= 0 — including exactly 1.0 — is well defined) and
// sampled by exact binary-search inversion. This replaces the approximate
// theta < 1 power-law formula, which silently rewrote theta = 1.0 to 0.9999
// and mislabelled the corresponding benchmark cells.
type ZipfGenerator struct {
	numAccounts int64
	theta       float64
	cumulative  []float64 // cumulative[i] = P(X <= i+1), 0-indexed
	hottestHits int64
	totalGen    int64
}

func NewZipfGenerator(numAccounts int64, theta float64) *ZipfGenerator {
	if numAccounts <= 0 {
		numAccounts = 10000
	}
	if theta < 0 {
		theta = 0
	}

	g := &ZipfGenerator{
		numAccounts: numAccounts,
		theta:       theta,
		cumulative:  make([]float64, numAccounts),
	}

	sum := 0.0
	for i := int64(1); i <= numAccounts; i++ {
		sum += 1.0 / math.Pow(float64(i), theta)
		g.cumulative[i-1] = sum
	}
	// Normalize so cumulative[N-1] == 1 exactly (guards float drift).
	for i := range g.cumulative {
		g.cumulative[i] /= sum
	}

	return g
}

// NewSampler returns a worker-local sampler with its own PRNG, eliminating the
// shared math/rand global mutex as a source of artificial client-side
// contention at high concurrency. Hot-row hit accounting stays on the shared
// generator via atomic counters.
func (g *ZipfGenerator) NewSampler(seed int64) *ZipfSampler {
	return &ZipfSampler{
		gen: g,
		rnd: rand.New(rand.NewSource(seed)),
	}
}

type ZipfSampler struct {
	gen *ZipfGenerator
	rnd *rand.Rand
}

// NextAccountID returns a 1-indexed account ID drawn from the exact Zipf
// distribution via binary-search inversion of the cumulative table.
func (s *ZipfSampler) NextAccountID() int64 {
	atomic.AddInt64(&s.gen.totalGen, 1)

	var id int64
	if s.gen.theta == 0 {
		id = s.rnd.Int63n(s.gen.numAccounts) + 1
	} else {
		u := s.rnd.Float64()
		// Lower bound: smallest i with cumulative[i-1] >= u  (i is 1-indexed).
		lo, hi := int64(0), s.gen.numAccounts-1
		for lo < hi {
			mid := (lo + hi) / 2
			if s.gen.cumulative[mid] < u {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
		id = lo + 1
	}

	if id == 1 {
		atomic.AddInt64(&s.gen.hottestHits, 1)
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
