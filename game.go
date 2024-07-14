package main

import (
	"fmt"
	"time"
)

type Game struct {
	tetrisMap        [ROWS][COLUMNS]rune
	checkList        [ROWS][COLUMNS]bool
	colsLengths      [COLUMNS]int
	completedLines   [ROWS]bool
	currentTetromino Tetromino
	nCurrX           int
	nCurrY           int
	lines            int
	gameSpeed        float64
	bIsRotUsed       bool
	bGameOn          bool
	inputChan        chan rune
}

func NewGame() *Game {
	game := &Game{
		nCurrX:    4,
		nCurrY:    0,
		gameSpeed: 10.0,
		bGameOn:   true,
		inputChan: make(chan rune),
	}
	game.init()
	return game
}

func (g *Game) init() {
	InitLengths(&g.colsLengths)
	InitCheckList(&g.checkList)
	InitTetrisMap(&g.tetrisMap)
	InitCompletedLines(&g.completedLines)
	g.newTetromino()
}

func (g *Game) newTetromino() {
	g.currentTetromino = CreateRandomTetromino()
	g.nCurrX = 4
	g.nCurrY = 0
}

func (g *Game) Run() {
	go ListenToKeyboard(g.inputChan)

	for g.bGameOn {
		g.update()
		g.handleInput()
		time.Sleep(time.Millisecond * time.Duration(20*g.gameSpeed))
		g.render()
	}

	PrintGameOver()
}

func (g *Game) update() {
	g.currentTetromino.X = g.nCurrX
	g.currentTetromino.Y = g.nCurrY - (g.currentTetromino.Height - 1)

	if g.nCurrY >= g.colsLengths[g.nCurrX] {
		InitLengths(&g.colsLengths)
		CheckTetrisMap(g.tetrisMap, &g.checkList, &g.colsLengths)
	}

	if g.nCurrY >= 0 && g.checkList[g.nCurrY][g.nCurrX] {
		g.newTetromino()
		return
	}

	g.nCurrY++
	if g.currentTetromino.Y >= 0 &&
		g.nCurrY >= 0 &&
		g.nCurrY <= g.colsLengths[g.nCurrX] {

		DropBlockOneRow(&g.tetrisMap, &g.currentTetromino, g.nCurrX)
	}

	if IsGameOver(g.colsLengths) {
		g.bGameOn = false
	}

	MarkDoneLines(&g.checkList, &g.completedLines, &g.lines)
	EliminateLines(&g.tetrisMap, &g.checkList, &g.completedLines, &g.colsLengths, &g.gameSpeed)
	UpdateTetrisMap(&g.tetrisMap, &g.checkList)

	if g.bIsRotUsed {
		for row := 0; row < g.nCurrY+g.currentTetromino.Height+1; row++ {
			for col := 0; col < COLUMNS; col++ {
				g.tetrisMap[row][col] = '.'
			}
		}
		g.bIsRotUsed = false
	}
}

func (g *Game) handleInput() {
	select {
	case chr := <-g.inputChan:
		switch chr {
		case 'A', 'a':
			g.nCurrX--
		case 'D', 'd':
			g.nCurrX++
		case 'W', 'w':
			g.bIsRotUsed = true
			g.currentTetromino.Rotate90Degs()
		case 'q':
			g.bGameOn = false
		}

		if g.nCurrX >= 9-g.currentTetromino.Width+1 {
			g.nCurrX = 9 - g.currentTetromino.Width + 1
		} else if g.nCurrX <= 0 {
			g.nCurrX = 0
		}
	default:
		// No input
	}
}

func (g *Game) render() {
	Clear()
	PrintMatrix(g.tetrisMap)
	fmt.Printf("\n Lines: %d\n", g.lines)
}
