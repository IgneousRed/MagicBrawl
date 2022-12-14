package main

import (
	et "github.com/IgneousRed/EduTen"
	m "github.com/IgneousRed/gomisc"
)

type s8 = int8
type s16 = int16
type s32 = int32
type s64 = int64
type u8 = uint8
type u16 = uint16
type u32 = uint32
type u64 = uint64
type f32 = float32
type f64 = float64

type Micros = m.Micros
type rad = m.Rad
type v2 = m.Vector2

var MicrosGet = m.MicrosGet
var V2 = m.Vec2

type verts = et.Verts
type trigs = et.Trigs

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
