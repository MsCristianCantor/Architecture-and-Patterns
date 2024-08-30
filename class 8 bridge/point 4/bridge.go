package main

import "fmt"

type AbstractionFigure interface {
	Drawing()
}

type CircleFigure struct{}

func (a *CircleFigure) Drawing() {
	fmt.Println("Drawing circle")
}

type SquareFigure struct{}

func (a *SquareFigure) Drawing() {
	fmt.Println("Drawing square")
}

type Figure struct {
	Figure AbstractionFigure
}

func (i *Figure) SetFigure(figure AbstractionFigure) {
	i.Figure = figure
}

type ImplementationDrawAPI1 struct {
	Figure
}

func (i *ImplementationDrawAPI1) Draw() {
	i.Figure.Figure.Drawing()
	fmt.Println("Draw from API 1")
}

type ImplementationDrawAPI2 struct {
	Figure
}

func (i *ImplementationDrawAPI2) Draw() {
	i.Figure.Figure.Drawing()
	fmt.Println("Draw from API 2")
}

func main() {
	var figures []AbstractionFigure
	figures = append(figures, &CircleFigure{})
	figures = append(figures, &SquareFigure{})

	for _, v := range figures {
		figure := Figure{}
		figure.SetFigure(v)
		var drawAPI1 = &ImplementationDrawAPI1{
			Figure: figure,
		}
		drawAPI1.Draw()
		var drawAPI2 = &ImplementationDrawAPI2{
			Figure: figure,
		}
		drawAPI2.Draw()
	}
}
