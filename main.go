package main

import (
	tm "github.com/buger/goterm"
	tty "github.com/mattn/go-tty"
	"log"
)

var _, _, _ = tty.Open, tm.Clear, log.Fatal

func main() {

	var t, err = tty.Open()
	if err != nil {
		log.Fatal("Unable to initialize keyboard input")
	}

	tm.Clear() // Clear current screen

	var field = Field{}
	field.initialize()

	var g = Game{
		field: field,
		score: 0,
	}

	g.initializeFigure()

	go keyboardThread(&g, t)

	g.run()

}

//
//func main() {
//	var g = galka()
//	d := __rotation_representation(g.data)
//	d = __rotation_representation(d)
//	d = __rotation_representation(d)
//	d = __rotation_representation(d)
//
//	g = palka()
//	d = __rotation_representation(g.data)
//	d = __rotation_representation(d)
//	d = __rotation_representation(d)
//	d = __rotation_representation(d)
//
//	g = pyramid()
//	d = __rotation_representation(g.data)
//	d = __rotation_representation(d)
//	d = __rotation_representation(d)
//	d = __rotation_representation(d)
//
//	g = square()
//	d = __rotation_representation(g.data)
//	d = __rotation_representation(d)
//	d = __rotation_representation(d)
//	d = __rotation_representation(d)
//}

func keyboardThread(g *Game, t *tty.TTY) {

	for {
		r, err := t.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		if r == 0 {
			continue
		}

		g.onKeyboardAction(r)
	}
}
