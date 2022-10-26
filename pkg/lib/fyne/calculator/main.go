package main

import (
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Calculator")

	mainResult := widget.NewLabel("")

	mainS := ""
	mainResult.SetText(mainS)

	one := widget.NewButton("1", func() {
		mainS += "1"
	})
	two := widget.NewButton("2", func() {
		mainS += "2"
	})
	three := widget.NewButton("3", func() {
		mainS += "3"
	})
	four := widget.NewButton("4", func() {
		mainS += "4"
	})
	five := widget.NewButton("5", func() {
		mainS += "5"
	})
	six := widget.NewButton("6", func() {
		mainS += "6"
	})
	seven := widget.NewButton("7", func() {
		mainS += "7"
	})
	eigth := widget.NewButton("8", func() {
		mainS += "8"
	})
	nine := widget.NewButton("9", func() {
		mainS += "9"
	})
	zero := widget.NewButton("0", func() {
		mainS += "0"
	})

	clear := widget.NewButton("C", func() {
		mainS = ""
	})

	plus := widget.NewButton("+", func() {
		mainS += "+"
	})

	minus := widget.NewButton("-", func() {
		mainS += "-"
	})

	multiply := widget.NewButton("*", func() {
		mainS += "*"
	})

	divide := widget.NewButton("/", func() {
		mainS += "/"
	})

	// fourth
	equal := widget.NewButton("=", func() {
		expression, err := govaluate.NewEvaluableExpression(mainS)
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				mainS = strconv.FormatFloat(result.(float64), 'f', -1, 64)
			} else {
				mainS = err.Error()
			}
		} else {
			mainS = err.Error()
		}
	})

	// five updating it each time
	go func() {
		for {
			mainResult.SetText(mainS)
		}
	}()

	// creating layout from the given buttons
	myWindow.SetContent(container.NewVBox(
		mainResult,
		container.NewGridWithColumns(3,
			divide,
			multiply,
			minus,
			// plus,
		),
		container.NewGridWithColumns(2,
			container.NewGridWithRows(3,
				container.NewGridWithColumns(3,
					seven,
					eigth,
					nine,
				),
				container.NewGridWithColumns(3,
					four,
					five,
					six,
				),
				container.NewGridWithColumns(3,
					one,
					two,
					three,
				),
			),
			plus,
		),
		container.NewGridWithColumns(2,
			container.NewGridWithColumns(2,
				clear,
				zero,
			),
			equal,
		),
	),
	)
	myWindow.ShowAndRun()
}
