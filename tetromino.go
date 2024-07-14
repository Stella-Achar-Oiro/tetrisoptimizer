package main

import "math/rand"

type Tetromino struct {
    Shape  [][]rune
    Width  int
    Height int
    X      int
    Y      int
}

func (t *Tetromino) Rotate90Degs() {
    newShape := make([][]rune, t.Width)
    for i := range newShape {
        newShape[i] = make([]rune, t.Height)
    }

    for i := 0; i < t.Height; i++ {
        for j := 0; j < t.Width; j++ {
            newShape[j][t.Height-i-1] = t.Shape[i][j]
        }
    }

    t.Shape = newShape
    t.Width, t.Height = t.Height, t.Width
}

func CreateRandomTetromino() Tetromino {
    shapes := [][][]rune{
        {
            {'#', '#', '#'},
            {'.', '#', '.'},
        },
        {
            {'#', '#'},
            {'#', '#'},
        },
        {
            {'.', '#', '#'},
            {'#', '#', '.'},
        },
        {
            {'#', '#', '.'},
            {'.', '#', '#'},
        },
        {
            {'#', '#', '#', '#'},
        },
        {
            {'#', '#', '#'},
            {'#', '.', '.'},
        },
        {
            {'#', '#', '#'},
            {'.', '.', '#'},
        },
    }

    shape := shapes[rand.Intn(len(shapes))]
    return Tetromino{
        Shape:  shape,
        Width:  len(shape[0]),
        Height: len(shape),
    }
}

func DropBlockOneRow(matrix *[ROWS][COLUMNS]rune, t *Tetromino, x int) {
    for row := 0; row < t.Height; row++ {
        for col := 0; col < t.Width; col++ {
            if t.Shape[row][col] == '#' {
                (*matrix)[t.Y+row][x+col] = '#'
            }
        }
    }
}
