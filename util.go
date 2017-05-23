package main

import sf "github.com/zyedidia/sfml/v2.3/sfml"

func Intersects(s1, s2 *sf.Sprite) (bool, sf.Rectf) {
	bounds1 := s1.GetGlobalBounds()
	bounds2 := s2.GetGlobalBounds()

	didCollide, rect := bounds1.Intersects(bounds2)
	return didCollide, rect
}

func Intersects2(s1, s2 *sf.Sprite) bool {
	bounds1 := s1.GetGlobalBounds()
	bounds2 := s2.GetGlobalBounds()

	didCollide2, _ := bounds1.Intersects(bounds2)
	return didCollide2
}
