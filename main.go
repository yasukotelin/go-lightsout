package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/yasukotelin/scrlib"
)

var (
	width  = flag.Int("w", 5, "field width")
	height = flag.Int("h", 5, "field height")
)

func goLigthsout() string {
	var sb strings.Builder
	sb.WriteString("                   _ _       _     _                   _\n")
	sb.WriteString("  __ _  ___       | (_) __ _| |__ | |_ ___  ___  _   _| |_\n")
	sb.WriteString(" / _` |/ _ \\ _____| | |/ _` | '_ \\| __/ __|/ _ \\| | | | __|\n")
	sb.WriteString("| (_| | (_) |_____| | | (_| | | | | |_\\__ \\ (_) | |_| | |_\n")
	sb.WriteString(" \\__, |\\___/      |_|_|\\__, |_| |_|\\__|___/\\___/ \\__,_|\\__|\n")
	sb.WriteString(" |___/                 |___/")
	return sb.String()
}

func congratulation() string {
	var sb strings.Builder
	sb.WriteString("  ____                            _         _       _   _             _\n")
	sb.WriteString(" / ___|___  _ __   __ _ _ __ __ _| |_ _   _| | __ _| |_(_) ___  _ __ | |\n")
	sb.WriteString("| |   / _ \\| '_ \\ / _` | '__/ _` | __| | | | |/ _` | __| |/ _ \\| '_ \\| |\n")
	sb.WriteString("| |__| (_) | | | | (_| | | | (_| | |_| |_| | | (_| | |_| | (_) | | | |_|\n")
	sb.WriteString(" \\____\\___/|_| |_|\\__, |_|  \\__,_|\\__|\\__,_|_|\\__,_|\\__|_|\\___/|_| |_(_)\n")
	sb.WriteString("                  |___/")
	return sb.String()
}

func ruleMsg() string {
	var sb strings.Builder
	sb.WriteString("Challenge to turn off all lights.\n")
	sb.WriteString("\n")
	sb.WriteString("light on -> !\n")
	sb.WriteString("light off -> .")
	return sb.String()
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	flag.Parse()

	// 初期表示
	scrlib.Clear()
	fmt.Println(goLigthsout())
	fmt.Println()
	fmt.Println(ruleMsg())
	fmt.Println()
	scanAnyBtn()
	scrlib.Clear()

	board := NewBoard(*height, *width)
	board.BuildField()

	result := play(board)

	// ゲームクリア
	fmt.Println(congratulation())
	fmt.Println()
	fmt.Printf("your count is %d\n", result)
	fmt.Println()
	scanAnyBtn()
}

func play(board *Board) int {
	var cnt int
	for {
		board.Disp()
		fmt.Println()
		fmt.Printf("count: %d\n", cnt)
		fmt.Println()
		x, y := scanPosition(board.Width, board.Height)
		board.Update(x-1, y-1)

		cnt++

		scrlib.Clear()
		if board.IsCorrect() {
			return cnt
		}
	}
}

func scanAnyBtn() {
	fmt.Print("Please press any button. ")
	fmt.Scanln()
}

func scanPosition(height int, width int) (x int, y int) {
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&y, &x); err != nil {
			fmt.Println(err)
			fmt.Println("please input like 2 1")
			continue
		}
		if y > height {
			fmt.Printf("max height is %d\n", height)
			continue
		}
		if x > width {
			fmt.Printf("max width is %d\n", width)
			continue
		}
		break
	}
	return x, y
}
