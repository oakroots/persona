package persona

import (
	"math/bits"
	"math/rand"
	"strings"

	"github.com/oakroots/persona/data"
)

func (g *Generator) GetFirstName() string {
	switch g.gender {
	case Female:
		return g.pickOne(data.FemaleFirstNames)
	case Male:
		return g.pickOne(data.MaleFirstNames)
	default:
		return g.pickOne(data.FantasyFirstNames)
	}
}

func (g *Generator) GetLastName() string {
	switch g.gender {
	case Female:
		return g.pickOne(data.FemaleLastNames)
	case Male:
		return g.pickOne(data.MaleLastNames)
	default:
		return g.pickOne(data.FantasyLastNames)
	}
}

func (g *Generator) GetFullName() string {
	first := g.GetFirstName()
	last := g.GetLastName()

	return strings.TrimSpace(first + " " + last)
}

func (g *Generator) pickOne(fetch func() data.Names) string {
	list := fetch()
	if len(list) == 0 {
		return ""
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	if g.deterministic {
		pr := newPRNG(g.seed)
		pr.next()
		g.seed = pr.s

		idx := idxFromState(pr.s, len(list))

		return list[idx]
	}

	return list[rand.Intn(len(list))]
}

func idxFromState(s uint32, n int) int {
	r := mix64(uint64(s))
	hi, _ := bits.Mul64(r, uint64(n))

	return int(hi)
}

func mix64(x uint64) uint64 {
	x += 0x9E3779B97F4A7C15
	x = (x ^ (x >> 30)) * 0xBF58476D1CE4E5B9
	x = (x ^ (x >> 27)) * 0x94D049BB133111EB

	return x ^ (x >> 28) ^ (x >> 31)
}
