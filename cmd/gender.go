package main

import "github.com/oakroots/persona"

type genderValue struct {
	value *persona.Gender
}

func (g *genderValue) String() string {
	if g.value == nil {
		return "u"
	}

	return g.value.String()
}

func (g *genderValue) Set(s string) error {
	pg, err := persona.ParseGender(s)
	if err != nil {
		return err
	}

	*g.value = pg

	return nil
}

func (g *genderValue) Type() string {
	return "gender"
}
