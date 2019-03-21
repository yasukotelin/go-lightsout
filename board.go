package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Board struct {
	Height int
	Width  int
	Field  [][]Cell
}

func NewBoard(h int, w int) *Board {
	// h * w の2次元スライス生成
	field := make([][]Cell, h)
	for i := 0; i < h; i++ {
		field[i] = make([]Cell, w)
	}
	return &Board{
		Height: h,
		Width:  w,
		Field:  field,
	}
}

func (b *Board) BuildField() {
	rondTime := 100 + rand.Intn(100)
	b.RondomUpdate(rondTime)

	if b.IsCorrect() {
		b.BuildField()
	}
}

func (b *Board) Disp() {
	var sb strings.Builder
	sb.WriteString("  ")
	for w := 1; w < b.Width+1; w++ {
		sb.WriteString(" ")
		sb.WriteString(fmt.Sprintf("%2d", w))
	}
	sb.WriteString("\n")

	for h := range b.Field {
		sb.WriteString(fmt.Sprintf("%2d", h+1))
		sb.WriteString(" ")
		sb.WriteString(b.Field[h][0].String())
		for w := 1; w < b.Width; w++ {
			sb.WriteString(" ")
			sb.WriteString(b.Field[h][w].String())
		}
		sb.WriteString("\n")
	}
	fmt.Print(sb.String())
}

func (b *Board) Update(x int, y int) {
	if x > 0 {
		// Left
		b.Field[y][x-1].IsLightOn = !b.Field[y][x-1].IsLightOn
	}
	if x < b.Width-1 {
		// Right
		b.Field[y][x+1].IsLightOn = !b.Field[y][x+1].IsLightOn
	}
	// Center
	b.Field[y][x].IsLightOn = !b.Field[y][x].IsLightOn
	if y > 0 {
		// Up
		b.Field[y-1][x].IsLightOn = !b.Field[y-1][x].IsLightOn
	}
	if y < b.Height-1 {
		// Bottom
		b.Field[y+1][x].IsLightOn = !b.Field[y+1][x].IsLightOn
	}
}

func (b *Board) RondomUpdate(times int) {
	var x, y int
	for i := 0; i < times; i++ {
		x = rand.Intn(b.Width)
		y = rand.Intn(b.Height)
		b.Update(x, y)
	}
}

func (b *Board) IsCorrect() bool {
	for h := range b.Field {
		for w := range b.Field[h] {
			if b.Field[h][w].IsLightOn {
				return false
			}
		}
	}
	return true
}
