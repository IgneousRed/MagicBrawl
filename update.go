package main

import (
	et "github.com/IgneousRed/EduTen"
	m "github.com/IgneousRed/gomisc"
)

func (g *Game) Update() {
	// Camera
	camMove := V2(
		m.BToF(et.KeysPressed(et.KeyRight))-m.BToF(et.KeysPressed(et.KeyLeft)),
		m.BToF(et.KeysPressed(et.KeyUp))-m.BToF(et.KeysPressed(et.KeyDown)),
	)
	camMove = camMove.Add(V2(
		m.BToF(et.Cursor().X >= et.WindowSize().X-10)-m.BToF(et.Cursor().X < 10),
		m.BToF(et.Cursor().Y >= et.WindowSize().Y-10)-m.BToF(et.Cursor().Y < 10),
	))
	et.CamTrans(camMove, 0, m.Pow(2, et.Wheel()))

	// Walk
	max := 5.
	acc := 1.
	velTrg := V2(
		m.BToF(et.KeysPressed(et.KeyD))-m.BToF(et.KeysPressed(et.KeyA)),
		m.BToF(et.KeysPressed(et.KeyW))-m.BToF(et.KeysPressed(et.KeyS)),
	).MagSet(max)
	g.mages[0].vel = g.mages[0].vel.MoveTowards(velTrg,
		m.Sqrt(g.mages[0].vel.Dst(velTrg)/max)*acc,
	)
	g.mages[0].pos = g.mages[0].pos.Add(g.mages[0].vel)
	g.mages[0].ang = et.CamVec2(g.mages[0].pos).AngTo(et.Cursor())

	// Rang
	if et.ButtonDown(et.ButtonL) {
		dir := g.mages[0].ang.Vec2()
		p, v := dir.Mul1(50).Add(g.mages[0].pos), dir.Mul1(5)
		g.rangs = append(g.rangs, Rang{0, p, v, rad(g.rng.Normal64() * m.Tau), false})
	}

	// Fire
	if et.ButtonDown(et.ButtonR) {
		dir := g.mages[0].ang.Vec2()
		p, v := dir.Mul1(50).Add(g.mages[0].pos), dir.Mul1(5)
		g.fires = append(g.fires, Fire{p, v, MicrosGet()})
	}

	// Rangs
	for i := 0; i < len(g.rangs); i++ {
		r := &g.rangs[i]
		ownerPos := g.mages[r.owner].pos
		dlt := ownerPos.Sub(r.pos)
		g.rangs[i].vel = dlt.MagSet(.1).Add(r.vel).
			Sub(dlt.Rot90().MagSet(.01).Project(r.vel))
		g.rangs[i].pos = r.pos.Add(r.vel)
		g.rangs[i].ang += rad((m.Tau + r.vel.Mag()*2) / 60 * m.BToS(r.ccw))
		if ownerPos.Dst(r.pos) < MageR+RangR {
			bound := len(g.rangs) - 1
			g.rangs[i] = g.rangs[bound]
			g.rangs = g.rangs[:bound]
			i--
		}
	}

	// Fires
	for i, f := range g.fires {
		g.fires[i].pos = f.pos.Add(f.vel)
	}
}
