package matrix

import (
    "errors"
    "strings"
    "strconv"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (Matrix, error) {
    if s == "" {
        return nil, errors.New("empty matrix")
    }
    
	rows := strings.Split(s, "\n")
    matrix := [][]int{}
    expectedCols := -1
    for i := 0; i < len(rows); i++ {
        numStrings := strings.Fields(rows[i])

        if len(numStrings) == 0 {
            return nil, errors.New("one of the column is empty")
        }
        if expectedCols == -1 {
            expectedCols = len(numStrings)
        } else if len(numStrings) != expectedCols {
            return nil, errors.New("uneven column")
        }
        row := []int{}
        for j := 0; j < len(numStrings); j++ {
            num, err := strconv.Atoi(numStrings[j])
            if err != nil {
                return nil, err
            }
            row = append(row, num)
        }
        matrix = append(matrix, row)
    }
    return matrix, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
    numCols := len(m[0])
    result := [][]int{}
    for i := 0; i < numCols; i++ {
        column := []int{}
        for j := 0; j < len(m); j++ {
            column = append(column, m[j][i])
        }
        result = append(result, column)
    }
    return result
}

func (m Matrix) Rows() [][]int {
	matrix := [][]int{}
    for i := 0; i < len(m); i++ {
        rows := m[i]
        row := []int{}
        for j := 0; j < len(rows); j++ {
            row = append(row, m[i][j])
        }
        matrix = append(matrix, row)
    }
    return matrix
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m) {
        return false
    }
    if col < 0 || col >= len(m[0]) {
        return false
    }
    m[row][col] = val
    return true
}
