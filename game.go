package main

import (
	m "github.com/IgneousRed/gomisc"
)

const (
	MageR = 20
	RangR = 20
)

type Pilar struct {
	hp  f64
	pos v2
}
type Mage struct {
	// hp  f64
	pos v2
	vel v2
	ang rad
}
type Rang struct {
	owner int
	pos   v2
	vel   v2
	ang   rad
	ccw   bool
}
type Fire struct {
	pos   v2
	vel   v2
	spawn Micros
}
type Asd struct {
	pos  v2
	size f64
}
type Game struct {
	// Update
	rng    m.PCG32
	pilars []Pilar
	mages  []Mage
	rangs  []Rang
	fires  []Fire
	// Draw
	asds []Asd
}
