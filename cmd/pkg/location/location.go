package location

type Exit struct {
	Id             string
	Dexcription    string
	Keyword        string
	TargetLocation *Location
}

type Location struct {
	Id          string
	Description string
	Exits       []*Exit
}
