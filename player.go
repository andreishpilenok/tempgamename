package main

import (
	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type Player struct {
	*sf.Sprite

	keys [3]sf.KeyCode

	vx float32
	vy float32

	onGround bool
	onSide   bool
}

func NewPlayer(keys [3]sf.KeyCode, pos sf.Vector2f) *Player {
	player := new(Player)
	player.Sprite = sf.NewSprite(res.images["testplayer.png"])

	player.SetPosition(pos)

	player.keys = keys

	size := player.GetGlobalBounds()
	player.SetOrigin(sf.Vector2f{size.Width / 2, size.Height / 2})

	return player
}

func (p *Player) Update(dt float32) {
	p.vx = 0
	if sf.KeyboardIsKeyPressed(p.keys[0]) {
		p.vx = -5
	}

	if sf.KeyboardIsKeyPressed(p.keys[1]) {
		p.vx = 5
	}

	if sf.KeyboardIsKeyPressed(p.keys[2]) && p.onGround {
		p.onGround = false
		p.vy = -15
	}

	if !p.onGround {
		p.vy += 1
	}
	if p.onGround {
		p.vy = 0
	}

	if p.onSide {
		p.vx = 0
	}
	p.Move(sf.Vector2f{p.vx, p.vy})
}
