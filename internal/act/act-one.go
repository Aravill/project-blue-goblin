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
