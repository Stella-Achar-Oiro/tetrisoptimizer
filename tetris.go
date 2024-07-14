package main

func InitLengths(lengths *[COLUMNS]int) {
    for i := 0; i < COLUMNS; i++ {
        (*lengths)[i] = ROWS - 1
    }
}

func InitCompletedLines(completedLines *[ROWS]bool) {
    for i := 0; i < ROWS; i++ {
        (*completedLines)[i] = false
    }
}

func InitCheckList(checkList *[ROWS][COLUMNS]bool) {
    for row := 0; row < ROWS; row++ {
        for col := 0; col < COLUMNS; col++ {
            (*checkList)[row][col] = false
        }
    }
}

func InitTetrisMap(tetrisMap *[ROWS][COLUMNS]rune) {
    for row := 0; row < ROWS; row++ {
        for col := 0; col < COLUMNS; col++ {
            (*tetrisMap)[row][col] = '.'
        }
    }
}

func UpdateTetrisMap(tetrisMap *[ROWS][COLUMNS]rune, checkList *[ROWS][COLUMNS]bool) {
    for row := 0; row < ROWS; row++ {
        for col := 0; col < COLUMNS; col++ {
            if (*checkList)[row][col] {
                (*tetrisMap)[row][col] = '#'
            }
        }
    }
}

func CheckTetrisMap(tetrisMap [ROWS][COLUMNS]rune, checkList *[ROWS][COLUMNS]bool, lengths *[COLUMNS]int) {
    for col := 0; col < COLUMNS; col++ {
        for row := 0; row < ROWS; row++ {
            if tetrisMap[row][col] == '#' {
                (*checkList)[row][col] = true
            }
        }
    }
    for col := 0; col < COLUMNS; col++ {
        for row := 0; row < ROWS; row++ {
            if tetrisMap[row][col] == '#' {
                (*lengths)[col] = row - 1
                break
            }
        }
    }
}

func MarkDoneLines(tetrisBooleanMap *[ROWS][COLUMNS]bool, completedLines *[ROWS]bool, eliminatedLines *int) {
    for row := 0; row < ROWS; row++ {
        if (*tetrisBooleanMap)[row][0] &&
            (*tetrisBooleanMap)[row][1] &&
            (*tetrisBooleanMap)[row][2] &&
            (*tetrisBooleanMap)[row][3] &&
            (*tetrisBooleanMap)[row][4] &&
            (*tetrisBooleanMap)[row][5] &&
            (*tetrisBooleanMap)[row][6] &&
            (*tetrisBooleanMap)[row][7] &&
            (*tetrisBooleanMap)[row][8] &&
            (*tetrisBooleanMap)[row][9] {
            (*completedLines)[row] = true
            (*eliminatedLines)++
        }
    }
}

func EliminateLines(tetrisMap *[ROWS][COLUMNS]rune, tetrisBooleanMap *[ROWS][COLUMNS]bool, completedLines *[ROWS]bool, colsLengths *[COLUMNS]int, gameSpeed *float64) {
    for row := ROWS - 1; row >= 0; row-- {
        if (*completedLines)[row] {
            for col := 0; col < COLUMNS; col++ {
                (*tetrisBooleanMap)[row][col] = false
            }
            for r := row; r >= 1; r-- {
                for col := 0; col < COLUMNS; col++ {
                    (*tetrisBooleanMap)[r][col] = (*tetrisBooleanMap)[r-1][col]
                }
            }
            row++
        }
    }

    for row := ROWS - 1; row >= 0; row-- {
        if (*completedLines)[row] {
            for col := 0; col < COLUMNS; col++ {
                (*tetrisMap)[row][col] = '.'
            }
            (*completedLines)[row] = false
            InitLengths(colsLengths)
            CheckTetrisMap(*tetrisMap, tetrisBooleanMap, colsLengths)
        }
    }

    *gameSpeed *= 0.975
}

func IsGameOver(colsLengths [COLUMNS]int) bool {
    for i := 0; i < COLUMNS; i++ {
        if colsLengths[i] == 0 {
            return true
        }
    }
    return false
}
