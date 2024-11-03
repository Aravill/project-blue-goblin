package location

import "blue-goblin/internal/item"

type Location struct {
	Id          string
	Description string
	Exits       []*Exit
	Items       []*ItemOnLocation
}

func (location *Location) RemoveItem(itemName string) *item.Item {
	for i, itemOnLoc := range location.Items {
		if itemOnLoc.Item.Name == itemName {
			location.Items = append(location.Items[:i], location.Items[i+1:]...)
			return itemOnLoc.Item
		}
	}
	return nil
}

func (location *Location) GetExit(exitKeyword string) *Exit {
	for _, e := range location.Exits {
		if e.Keyword == exitKeyword {
			return e
		}
	}
	return nil
}

func (location *Location) GetFullDescription() string {
	var full string = location.Description

	for _, e := range location.Items {
		full += " " + e.Description
	}

	for _, e := range location.Exits {
		full += " " + e.Description
	}

	return full
}
