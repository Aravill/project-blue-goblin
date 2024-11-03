package act

import (
	"blue-goblin/internal/item"
	"blue-goblin/internal/location"
)

func LoadActOne() Act {
	var actOne = NewAct("act-one")

	actOne.AddLocation(location.Location{
		Id:          "act-one:start",
		Description: "You are in a dark room.",
		Exits: []*location.Exit{
			{
				Id:          "act-one:start:north",
				Description: "There is a door to the north.",
				Keyword:     "north",
				Destination: "act-one:end",
				Lock: &location.Lock{
					Description: "The door is locked with a heavy rusty padlock.",
					KeyId:       "act-one:start:rusty-key",
					Status:      true,
				},
			},
		},
		Items: []*location.ItemOnLocation{
			{
				Item: &item.Item{
					Id:          "act-one:start:rusty-key",
					Name:        "rusty key",
					Description: "A worn rusty key",
				},
				Description: "A rusty key lies on a table in the corner of the room.",
			},
		},
	})

	actOne.AddLocation(location.Location{
		Id:          "act-one:end",
		Description: "You are in another dark room. ",
		Exits: []*location.Exit{
			{
				Id:          "act-one:end:south",
				Description: "There is a door to the south.",
				Keyword:     "south",
				Destination: "act-one:start",
			},
		},
		Items: []*location.ItemOnLocation{
			{
				Item: &item.Item{
					Id:          "act-one:end:coin",
					Name:        "coin",
					Description: "A shiny coin.",
				},
				Description: "A shiny coin is on the floor.",
			},
		},
	})

	return actOne
}
