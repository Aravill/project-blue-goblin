package player

type Flag string

const (
	Quitting Flag = "quitting"
)

func (f Flag) String() string {
	return string(f)
}
