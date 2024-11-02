package console

import (
	"fmt"
	"time"
)

type Narrator struct{}

func SayLine(text string) {
	Say(text)
	fmt.Println()
}

func Say(text string) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(20 * time.Millisecond)
	}
}
