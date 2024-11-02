package location

type Location struct {
	Id          string
	Description string
	Exits       []*Exit
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
	for _, e := range location.Exits {
		full += " " + e.Description
	}
	return full
}
