// package main

// import (
// 	"os"
// 	"os/exec"
// 	"time"
// )

// func ListenToKeyboard(inputChan chan rune) {
// 	// Set terminal to raw mode to capture key presses
// 	cmd := exec.Command("stty", "raw")
// 	cmd.Stdin = os.Stdin
// 	err := cmd.Run()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer func() {
// 		// Restore terminal to normal mode
// 		cmd := exec.Command("stty", "-raw")
// 		cmd.Stdin = os.Stdin
// 		cmd.Run()
// 	}()

// 	go func() {
// 		buf := make([]byte, 1)
// 		for {
// 			n, err := os.Stdin.Read(buf)
// 			if err != nil || n == 0 {
// 				continue
// 			}
// 			inputChan <- rune(buf[0])
// 		}
// 	}()

// 	// Keep the main goroutine alive
// 	for {
// 		time.Sleep(10 * time.Second)
// 	}
// }


package main

import (
    "bufio"
    "fmt"
    "os"
)

func ListenToKeyboard(inputChan chan rune) {
    reader := bufio.NewReader(os.Stdin)
    for {
        input, err := reader.ReadByte()
        if err != nil {
            fmt.Println("Error reading input:", err)
            continue // You can choose to retry or exit the loop based on the severity of the error
        }
        inputChan <- rune(input)
    }
}

