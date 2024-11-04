package console

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/eiannone/keyboard"
)

var skipDelay int32 // Atomic flag to indicate whether to skip the delay

func InitInputListener() {
	go func() {
		defer keyboard.Close()
		for {
			_, key, err := keyboard.GetKey()
			if err != nil {
				continue
			}
			if key == keyboard.KeySpace {
				atomic.StoreInt32(&skipDelay, 1)
				break
			}
		}
	}()
}

func SayLine(text string) {
	Say(text + "\n")
}

func Say(text string) {
	PrintInCharacters(text)
}

func PrintInCharacters(text string) {

	if err := keyboard.Open(); err != nil {
		fmt.Println("Error opening keyboard:", err)
		return
	}

	atomic.StoreInt32(&skipDelay, 0) // Reset the skipDelay flag
	InitInputListener()              // Start listening for input

	for _, char := range text {
		// Curly braces are used to denote key words in text
		if char == '{' {
			SetColorYellow()
			continue
		}
		if char == '}' {
			ResetColor()
			continue
		}
		fmt.Print(string(char))
		if atomic.LoadInt32(&skipDelay) == 0 {
			time.Sleep(20 * time.Millisecond)
		}
	}

	keyboard.Close()
}

func SetColorYellow() {
	fmt.Print("\033[33m")
}

func ResetColor() {
	fmt.Print("\033[0m")
}
