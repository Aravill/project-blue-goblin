package player

import "blue-goblin/internal/location"

type Player struct {
	CurrentLocation string
	Flags           []Flag
}

func (p *Player) MoveTo(location *location.Location) {
	p.CurrentLocation = location.Id
}

func (p *Player) HasFlag(flag Flag) bool {
	for _, f := range p.Flags {
		if f == flag {
			return true
		}
	}
	return false
}

func (p *Player) AddFlag(flag Flag) {
	p.Flags = append(p.Flags, flag)
}

func (p *Player) RemoveFlag(flag Flag) {
	var newFlags []Flag
	for _, f := range p.Flags {
		if f != flag {
			newFlags = append(newFlags, f)
		}
	}
	p.Flags = newFlags
}
