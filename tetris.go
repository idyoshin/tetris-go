package main

import (
	"sync"
)

const FIELD_WIDTH = 10

const FIELD_HEIGHT = 20

type Location struct {
	x int
	y int
}

type Figure struct {
	data [4][4]int
	x    int
	y    int
	lock sync.Mutex
}

type Field struct {
	data [FIELD_HEIGHT][FIELD_WIDTH]int
}

func (f *Field) initialize() {
	for y := 0; y < FIELD_HEIGHT; y++ {
		for x := 0; x < FIELD_WIDTH; x++ {
			f.data[y][x] = 0
		}
	}
}

func (f *Figure) center() {
	f.x = FIELD_WIDTH/2 - 2
}

func (f *Figure) left(field *Field) bool {
	f.lock.Lock()

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if f.data[i][j] == 1 {

				if f.x-1+j < 0 {
					f.lock.Unlock()
					return false
				}

				if field.data[f.y+i][f.x-1+j] == 1 {
					f.lock.Unlock()
					return false
				}
			}
		}
	}

	f.x--
	f.lock.Unlock()
	return true
}

func (f *Figure) right(field *Field) bool {
	f.lock.Lock()

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if f.data[i][j] == 1 {

				if f.x+1+j >= FIELD_WIDTH {
					f.lock.Unlock()
					return false
				}

				if field.data[f.y+i][f.x+1+j] == 1 {
					f.lock.Unlock()
					return false
				}
			}
		}
	}

	f.x++
	f.lock.Unlock()
	return true
}

func __rotation_representation(rect [4][4]int) [4][4]int {

	for i := 0; i < 4; i++ {
		for j := i; j < 4; j++ {
			tmp := rect[i][j]
			rect[i][j] = rect[j][i]
			rect[j][i] = tmp
		}
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			k := 3 - j
			tmp := rect[j][i]
			rect[j][i] = rect[k][i]
			rect[k][i] = tmp
		}
	}

	return rect
}

func (f *Figure) rotate(field *Field) bool {
	f.lock.Lock()
	var new_representation [4][4]int = __rotation_representation(f.data)

	for i := 0; i < 4; i++ {

		for j := 0; j < 4; j++ {

			if new_representation[i][j] == 1 {
				// exit in case there is  situation out of range
				if f.x+j < 0 || f.x+j >= FIELD_WIDTH {
					f.lock.Unlock()
					return false
				}

				if f.y+i < 0 || f.y+i >= FIELD_HEIGHT {
					f.lock.Unlock()
					return false
				}

				if field.data[f.y+i][f.x+j] == 1 {
					f.lock.Unlock()
					return false
				}
			}
		}

	}

	f.data = new_representation

	f.lock.Unlock()
	return true
}

func (f *Figure) down(field *Field) bool {
	f.lock.Lock()

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {

			if f.data[i][j] == 1 {

				if f.y+1+i >= FIELD_HEIGHT {
					f.lock.Unlock()
					return false
				}

				if field.data[f.y+1+i][f.x+j] == 1 {
					f.lock.Unlock()
					return false
				}
			}
		}
	}

	f.y++

	f.lock.Unlock()
	return true
}

func (f *Figure) fixate(field *Field) bool {
	f.lock.Lock()
	res := f.can_fixate(field)

	if res {

		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {

				if f.data[i][j] == 1 {
					field.data[f.y+i][f.x+j] = 1
				}
			}
		}
	}

	f.lock.Unlock()
	return res
}

func (f *Figure) can_fixate(field *Field) bool {
	for i := 0; i < 4; i++ {

		for j := 0; j < 4; j++ {

			if f.data[i][j] == 1 {
				if f.y+1+i == FIELD_HEIGHT {
					return true
				}

				if field.data[f.y+1+i][f.x+j] == 1 {
					return true
				}
			}
		}

	}

	return false
}

func pyramid() Figure {
	representation := [4][4]int{
		{0, 0, 0, 0},
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 0, 0, 0},
	}
	return Figure{data: representation}
}

func square() Figure {
	representation := [4][4]int{
		{0, 0, 0, 0},
		{0, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}
	return Figure{data: representation}
}

func palka() Figure {
	representation := [4][4]int{
		{0, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 1, 0, 0},
	}
	return Figure{data: representation}

}

func galka() Figure {
	representation := [4][4]int{
		{0, 0, 1, 0},
		{0, 1, 1, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 0},
	}
	return Figure{data: representation}
}

func galka2() Figure {
	representation := [4][4]int{
		{0, 1, 0, 0},
		{0, 1, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
	}
	return Figure{data: representation}
}

func shit() Figure {
	representation := [4][4]int{
		{0, 1, 1, 0},
		{0, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 0},
	}
	return Figure{data: representation}
}

func shit2() Figure {
	representation := [4][4]int{
		{0, 1, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
	}
	return Figure{data: representation}
}
