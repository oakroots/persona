package persona

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newPRNG(t *testing.T) {
	tests := []struct {
		name string
		seed uint32
		want *prng
	}{
		{
			name: "initial seed eq 0 set as 1",
			want: &prng{s: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, newPRNG(tt.seed), "newPRNG(%v)", tt.seed)
		})
	}
}

func Test_prng_next(t *testing.T) {
	type fields struct {
		s uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   uint32
	}{
		{
			name:   "return 1",
			fields: fields{s: 0},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &prng{
				s: tt.fields.s,
			}
			assert.Equalf(t, tt.want, p.next(), "next()")
		})
	}
}
