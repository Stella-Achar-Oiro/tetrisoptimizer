package main

import (
    "fmt"
)

const (
    ROWS    = 20
    COLUMNS = 10
)

func Clear() {
    fmt.Printf("\033[H\033[2J")
}

func PrintMatrix(matrix [ROWS][COLUMNS]rune) {
    fmt.Print("\n ")
    for i := 0; i < ROWS+3; i++ {
        fmt.Print("7")
    }
    fmt.Println("")
    for i := 0; i < ROWS; i++ {
        fmt.Print(" 7 ")
        for j := 0; j < COLUMNS; j++ {
            fmt.Printf("%c ", matrix[i][j])
        }
        fmt.Println("7")
    }
    fmt.Print(" ")
    for i := 0; i < ROWS+3; i++ {
        fmt.Print("7")
    }
    fmt.Println("")
}
