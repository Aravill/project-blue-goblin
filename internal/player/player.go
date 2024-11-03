package player

import (
	"blue-goblin/internal/item"
	"blue-goblin/internal/location"
)

type Player struct {
	CurrentLocation string
	Flags           []Flag
	Items           []item.Item
}

func (player *Player) RemoveItem(itemId string) *item.Item {
	for i, itm := range player.Items {
		if itm.Id == itemId {
			player.Items = append(player.Items[:i], player.Items[i+1:]...)
			return &itm
		}
	}
	return nil
}

func (player *Player) AddItem(item *item.Item) {
	for _, itm := range player.Items {
		if itm.Id == item.Id {
			return
		}
	}
	player.Items = append(player.Items, *item)
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
