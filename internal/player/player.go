package player

import "blue-goblin/internal/location"

type Player struct {
	CurrentLocation string
}

func (p *Player) MoveTo(location *location.Location) {
	p.CurrentLocation = location.Id
}
