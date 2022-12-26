package main

import (
	et "github.com/IgneousRed/EduTen"
	m "github.com/IgneousRed/gomisc"
)

// Trig data
var basic, mage, rang trigs

func init() {
	basic.Verts = make(verts, 3)
	for i := range basic.Verts {
		basic.Verts[i] = m.Deg(f64(i) * 120).Vec2()
	}
	basic.Inds = []u16{0, 1, 2}

	mage.Verts = verts{
		m.Deg(0).Vec2().Mul1(15),
		m.Deg(160).Vec2().Mul1(15),
		m.Deg(200).Vec2().Mul1(15),
	}
	mage.Inds = basic.Inds

	rang.Verts = make(verts, 6)
	for i := range rang.Verts {
		rang.Verts[i] = m.Deg(f64(i) * 60).Vec2().Mul1(f64(5 + i%2*15))
	}
	rang.Inds = []u16{0, 2, 4, 0, 1, 2, 2, 3, 4, 4, 5, 0}
}

// Other data
var qwe [4]f64

func init() {
	speed := m.Pi
	for i := range qwe {
		qwe[i] = speed
		speed *= -m.Phi
	}
}

func (g *Game) Draw(scr *et.Image) {
	// for _, f := range g.fires {

	// }
	for _, f := range g.fires {
		dltTime := f64(MicrosGet()) / 1000
		clr := et.Red
		clr.A = 128
		for _, s := range qwe {
			et.CamDrawTriangles(scr, basic.Transform1(f.pos, rad(s*dltTime), 20), clr)
		}
	}
	for _, r := range g.rangs {
		et.CamDrawTriangles(scr, rang.Transform1(r.pos, r.ang, 1), et.White)
	}
	for _, m := range g.mages {
		et.CamDrawCircle(scr, m.pos, 20, 32, et.Red)
		et.CamDrawTriangles(scr, mage.Transform1(m.pos, m.ang, 1), et.White)
	}
}
