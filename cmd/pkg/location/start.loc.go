package location

func init() {
	registerLocation(&Location{
		Id:          "start",
		Description: "You are standing in a dark room. There is a door to the north.",
		Exits:       []*Exit{},
	})
}
