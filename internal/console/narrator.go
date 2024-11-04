package console

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/eiannone/keyboard"
)

var skipDelay int32 // Atomic flag to indicate whether to skip the delay

func InitInputListener() {
	// Open the keyboard listener
	if err := keyboard.Open(); err != nil {
		fmt.Println("Error opening keyboard:", err)
		return
	}
	// Ensure the keyboard is closed when no longer needed
	defer keyboard.Close()

	// Start a goroutine to listen for key presses
	go func() {
		for {
			_, key, err := keyboard.GetKey()
			if err != nil {
				continue
			}
			if key == keyboard.KeySpace { // Check if the spacebar was pressed
				atomic.StoreInt32(&skipDelay, 1)
				break // Exit the loop after pressing spacebar
			}
		}
	}()
}

func SayLine(text string) {
	Say(text)
	fmt.Println()
}

func Say(text string) {
	PrintInCharacters(text)
}

func PrintInCharacters(text string) {
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
}

func SetColorYellow() {
	fmt.Print("\033[33m")
}

func ResetColor() {
	fmt.Print("\033[0m")
}
