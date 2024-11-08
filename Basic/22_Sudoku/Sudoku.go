package main

import (
	"errors"
	"fmt"
)

const (
	length = 9
	width  = 9
)

type sudoku [length][width]int8

type verify [length][width]bool

func newSudoku(s [length][width]int8) (result sudoku, v verify) {
	result = s
	for i1 := range s {
		for i2, j := range s[i1] {
			if j == 0 {
				v[i1][i2] = true
			}
		}
	}
	return result, v
}

func change(x, y, elem int8, s1 *sudoku, v *verify) error {
	if !v[x][y] {
		return errors.New("no permission to modify the original elements")
	}
	if s1[x][y] == elem {
		return errors.New("element already modified")
	}
	for i := 0; i < length; i++ {
		if s1[i][y] == elem {
			return errors.New("element already exists in column")
		}
		if s1[x][i] == elem {
			return errors.New("element already exists in row")
		}
	}
	// 检查 3x3 子网格
	startRow := (x / 3) * 3
	startCol := (y / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if s1[i][j] == elem {
				return errors.New("element already exists in 3x3 grid")
			}
		}
	}
	s1[x][y] = elem
	return nil
}

func main() {
	s, v := newSudoku([length][width]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})

	err := change(2, 2, 1, &s, &v)
	if err != nil {
		fmt.Println("修改失败:", err)
	} else {
		fmt.Println("修改成功")
	}
}
