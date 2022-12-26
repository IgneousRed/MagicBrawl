package main

import (
	et "github.com/IgneousRed/EduTen"
	m "github.com/IgneousRed/gomisc"
)

type Micros = m.Micros

var MicrosGet = m.MicrosGet

type Rad = et.Rad
type Vec2 = et.Vec2
type Verts = et.Verts
type Trigs = et.Trigs

var V2 = et.V2

func main() {
	// f, err := et.FontNew("FiraCode-Medium.ttf")
	// m.FatalErr("", err)
	// font = f
	et.WindowTitleSet("MagicBrawl")
	et.WindowMaximize()
	et.WindowResizingSet(et.WREnabled)
	g := Game{
		rng:   m.PCG32Init(),
		mages: []Mage{{}},
	}
	et.Run(&g)
}
