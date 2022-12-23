package main

import (
	"fmt"
	tm "github.com/buger/goterm"
	rand "math/rand"
	"strconv"
	"sync"
	"time"
)

const RENDER_SCALE = 2

type Game struct {
	field         Field
	currentFigure Figure
	nextFigure    Figure

	score    int
	gameOver bool
}

func scale_x(f func()) {
	for scale := 0; scale < RENDER_SCALE; scale++ {
		f()
	}
}

func (g *Game) checkSuccess() {
	startY := g.currentFigure.y

	coefficient := 1.0

	for i := 0; i < 4; i++ {

		if i+startY < FIELD_HEIGHT {
			if g.isFullLine(i + startY) {
				g.score = g.score + int(10.0*coefficient)
				coefficient = coefficient * 1.1

				g.removeLine(i + startY)
			}
		}
	}
}

func (g *Game) removeLine(i int) {
	for line := i; line > 0; line-- {
		g.field.data[line] = g.field.data[line-1]
	}

	for col := 0; col < FIELD_WIDTH; col++ {
		g.field.data[0][col] = 0
	}
}

func (g *Game) isFullLine(i int) bool {
	for j := 0; j < FIELD_WIDTH; j++ {

		if g.field.data[i][j] == 0 {
			return false
		}
	}

	return true
}

func (g *Game) render() {
	tm.MoveCursor(1, 1)
	for i := 0; i < FIELD_HEIGHT; i++ {

		for scale := 0; scale < RENDER_SCALE; scale++ {
			tm.Print(" ")
			tm.Print("|") // start of the box

			for j := 0; j < FIELD_WIDTH; j++ {

				var printed = false
				if g.field.data[i][j] == 1 {
					printed = true
					scale_x(func() { tm.Print("*") })
				}

				if !printed {

					if i >= g.currentFigure.y && i < g.currentFigure.y+4 && j >= g.currentFigure.x && j < g.currentFigure.x+4 {

						y := i - g.currentFigure.y
						x := j - g.currentFigure.x

						if g.currentFigure.data[y][x] == 1 {
							printed = true
							scale_x(func() { tm.Print("*") })
						}
					}
				}

				if !printed {
					scale_x(func() { tm.Print(" ") })
				}
			}

			tm.Print("|")
			tm.Println("")
		}

	}

	tm.Print(" ")
	tm.Print("└")
	for j := 0; j < FIELD_WIDTH; j++ {
		scale_x(func() { tm.Print("-") })
	}
	tm.Print("┘")
	tm.Println("")

	score := fmt.Sprintf("SCORE: %i", g.score)
	tm.Println(tm.Bold(score))

	tm.Flush()
}

func (g *Game) onKeyboardAction(symbol rune) {

	var f func(f *Field) bool = nil

	switch symbol {
	case 'a':
		{
			f = g.currentFigure.left
			break
		}
	case 'w':
		{
			f = g.currentFigure.rotate
			break
		}
	case 's':
		{
			f = g.currentFigure.down
			break
		}
	case 'd':
		{
			f = g.currentFigure.right
			break
		}
	}

	if f != nil {
		if f(&g.field) {
			g.render()
		}
	}
}

//func randomFigure() Figure {
//	return galka()
//}

func randomFigure() Figure {

	// pickup random figure
	rand.Seed(time.Now().UnixNano())
	number := rand.Int() % 7

	var f Figure
	switch number {
	case 0:
		f = pyramid()
	case 1:
		f = square()
	case 2:
		f = palka()
	case 3:
		f = galka()
	case 4:
		f = galka2()
	case 5:
		f = shit()
	case 6:
		f = shit2()
	}

	// rotate randomly
	rand.Seed(time.Now().UnixNano())
	number = rand.Int() % 5
	for i := 0; i < number; i++ {
		f.data = __rotation_representation(f.data)
	}

	return Figure{
		data: f.data,
		lock: sync.Mutex{},
	}
}

func (g *Game) validateFigureStart() {

	for i := 0; i < 4; i++ {

		for j := 0; j < 4; j++ {
			if g.currentFigure.data[i][j] == 1 {
				if g.field.data[g.currentFigure.y+i][g.currentFigure.x+j] == 1 {
					g.gameOver = true
					return
				}
			}
		}
	}

	g.score += 4
}

func (g *Game) pickFigure() {
	g.currentFigure = g.nextFigure
	g.nextFigure = randomFigure()
	g.nextFigure.center()

	g.validateFigureStart()
}

func (g *Game) initializeFigure() {
	g.currentFigure = randomFigure()
	g.currentFigure.center()
	g.nextFigure = randomFigure()
	g.nextFigure.center()
}

func (g *Game) run() {
	for {
		time.Sleep(150 * time.Millisecond)

		g.currentFigure.down(&g.field)

		if g.currentFigure.fixate(&g.field) {
			g.checkSuccess()
			g.pickFigure()
		}
		g.render()

		if g.gameOver {
			fmt.Println("GAME OVER! ", strconv.Itoa((g.score)))
			break
		}
	}
}
