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
	pos Vec2
}
type Mage struct {
	// hp  f64
	pos Vec2
	vel Vec2
	ang Rad
}
type Rang struct {
	owner int
	pos   Vec2
	vel   Vec2
	ang   Rad
	ccw   bool
}
type Fire struct {
	pos   Vec2
	vel   Vec2
	spawn Micros
}
type Asd struct {
	pos  Vec2
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
