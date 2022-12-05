package main

import (
	"time"

	et "github.com/IgneousRed/EduTen"
	m "github.com/IgneousRed/gomisc"
	// eb "github.com/hajimehoshi/ebiten/v2"
)

var windowSize = m.Vec2F{800, 600}

// var font et.Font

type Mage struct {
	// hp  float64
	pos m.Vec2F
	vel m.Vec2F
	ang m.Rad64
}
type Rang struct {
	owner int
	pos   m.Vec2F
	vel   m.Vec2F
}
type Game struct {
	mages []Mage
	rangs []Rang
}

func (g *Game) Update() {
	if et.KeysDown(et.KeyLeft) {
		g.mages[0].vel = g.mages[0].vel.Add(m.Vec2F{-20, 0})
	}
	if et.KeysDown(et.KeyRight) {
		g.mages[0].vel = g.mages[0].vel.Add(m.Vec2F{20, 0})
	}
	if et.KeysDown(et.KeyDown) {
		g.mages[0].vel = g.mages[0].vel.Add(m.Vec2F{0, -20})
	}
	if et.KeysDown(et.KeyUp) {
		g.mages[0].vel = g.mages[0].vel.Add(m.Vec2F{0, 20})
	}
	velTrg := m.Vec2F{
		m.BToF(et.KeysPressed(et.KeyD)) - m.BToF(et.KeysPressed(et.KeyA)),
		m.BToF(et.KeysPressed(et.KeyW)) - m.BToF(et.KeysPressed(et.KeyS)),
	}.Norm().Mul1(5.)
	g.mages[0].vel = g.mages[0].vel.MoveTowards(velTrg,
		m.Sqrt(g.mages[0].vel.Dst(velTrg))*.5,
	)
	g.mages[0].pos = g.mages[0].pos.Add(g.mages[0].vel)
	g.mages[0].ang = g.mages[0].pos.AngTo64(et.Cursor())
	if et.KeysDown(et.KeySpace) {
		dir := g.mages[0].ang.Vec2F()
		p := g.mages[0].pos.Add(dir.Mul1(50))
		v := dir.Mul1(10)
		g.rangs = append(g.rangs, Rang{0, p, v})
	}
	for i := 0; i < len(g.rangs); i++ {
		dlt := g.mages[g.rangs[i].owner].pos.Sub(g.rangs[i].pos)
		g.rangs[i].vel = g.rangs[i].vel.Add(dlt.MagSet(.3))
		g.rangs[i].vel = g.rangs[i].vel.Sub(dlt.Rot90().MagSet(.2).Project(g.rangs[i].vel))
		g.rangs[i].pos = g.rangs[i].pos.Add(g.rangs[i].vel)
		if g.mages[g.rangs[i].owner].pos.Dst(g.rangs[i].pos) < 40 {
			bound := len(g.rangs) - 1
			g.rangs[i] = g.rangs[bound]
			g.rangs = g.rangs[:bound]
			i--
		}
	}
}
func (g *Game) Draw() {
	et.DrawCircle(g.mages[0].pos, 20, 24, et.Red)
	p := []m.Vec2F{
		{15, 0},
		{-14, 5},
		{-14, -5},
	}
	et.DrawTriangles(m.TranslateVec2F(m.RotateVec2F(p, g.mages[0].ang), g.mages[0].pos),
		[]uint16{0, 1, 2}, et.White,
	)
	for i := range g.rangs {
		p := []m.Vec2F{
			{20, 0},
			{-10, 17},
			{-10, -17},
		}
		a := m.Rad64(float64(time.Now().UnixMilli()) / 1000 * 6)
		et.DrawTriangles(m.TranslateVec2F(m.RotateVec2F(p, a), g.rangs[i].pos),
			[]uint16{0, 1, 2}, et.White,
		)
	}
}

func main() {
	// f, err := et.FontNew("FiraCode-Medium.ttf")
	// m.FatalErr("", err)
	// font = f
	g := Game{}
	g.mages = []Mage{{}}
	et.InitGame("MagicBrawl", windowSize, &g)
}
