package location

import "blue-goblin/internal/item"

type ItemOnLocation struct {
	Item        *item.Item `json:"item"`
	Description string     `json:"description"`
}
