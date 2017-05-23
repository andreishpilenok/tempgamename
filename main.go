package main

import (
	"io/ioutil"
	"runtime"
	"strings"
	"time"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

const (
	screenWidth  = 800
	screenHeight = 1600
	playerSpeed  = 4
)

type platform struct {
	*sf.Sprite
}

var res *Resources
var platforms []*platform

func NewPlatform(char rune, pos sf.Vector2f) *platform {
	p := new(platform)

	var texture *sf.Texture
	if char == 'P' {
		texture = res.images["testimg.png"]
	} else if char == 'C' {
		texture = res.images["..."]
	}

	p.Sprite = sf.NewSprite(texture)
	p.SetPosition(pos)

	return p

}

func main() {

	// music := sf.NewMusic("music.ogg")
	// music.Play()

	runtime.LockOSThread()

	res = NewResources()

	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Game", sf.StyleDefault, nil)
	window.SetVerticalSyncEnabled(true)
	window.SetFramerateLimit(60)

	player := NewPlayer([3]sf.KeyCode{sf.KeyA, sf.KeyD, sf.KeySpace}, sf.Vector2f{screenWidth / 2, screenHeight / 2})

	data, _ := ioutil.ReadFile("assets/levels/level1.txt")

	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		for j, character := range line {
			if character != ' ' {
				platform := NewPlatform(character, sf.Vector2f{float32(j) * 32, screenHeight - float32(i)*32})
				platforms = append(platforms, platform)
			}
		}
	}
	view := sf.NewView()
	view.SetSize(sf.Vector2f{screenWidth, screenHeight})

	view.SetCenter(sf.Vector2f{screenWidth / 2, screenHeight / 2})

	var dt float32

	for window.IsOpen() {
		start := time.Now()

		if event := window.PollEvent(); event != nil {
			switch event.Type {
			case sf.EventClosed:
				window.Close()
			}
		}
		player.Update(dt)
		view.Move(sf.Vector2f{0, -1 * 10 * dt})

		window.SetView(view)
		for _, p := range platforms {
			didCollide, rect := Intersects(p.Sprite, player.Sprite)
			if didCollide && rect.Width > rect.Height {
				player.onGround = true
				player.SetPosition(sf.Vector2f{player.GetPosition().X, p.GetPosition().Y - player.GetGlobalBounds().Height/2})
			}
			if didCollide && rect.Width < rect.Height {
				player.vx = 0
				if player.GetPosition().X > rect.Left+rect.Width/2 {
					player.SetPosition(sf.Vector2f{p.GetPosition().X + player.GetGlobalBounds().Width, player.GetPosition().Y})
				} else {
					player.SetPosition(sf.Vector2f{p.GetPosition().X - player.GetGlobalBounds().Width/2, player.GetPosition().Y})
				}
			}
		}

		window.Clear(sf.ColorBlack)
		for _, platform := range platforms {
			window.Draw(platform)
		}
		window.Draw(player)

		window.Display()

		elapsed := time.Since(start)
		dt = float32(elapsed) / float32(time.Second)
	}

}
