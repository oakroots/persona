package persona

const lfsrMask uint32 = 0x1ffffff

type prng struct {
	s uint32
}

func newPRNG(seed uint32) *prng {
	if seed == 0 {
		seed = 1
	}
	return &prng{s: seed & lfsrMask}
}

func (p *prng) next() uint32 {
	s := p.s

	b := ((s >> 0) ^ (s >> 1) ^ (s >> 2) ^ (s >> 6) ^ (s >> 8) ^ (s >> 10) ^ (s >> 25)) & 1
	s = ((s >> 1) | (b << 24)) & lfsrMask

	if s == 0 {
		s = 1
	}

	p.s = s

	return s
}
