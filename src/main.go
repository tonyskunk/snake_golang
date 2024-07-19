package main

import "example.com/m/v2/Desktop/Projects/dep/snake/src/snake"

func main() {

	s := snake.NewSnake()
	s.Reset()
	g := snake.NewGame()
	g.SetSnake(s)
	g.Run()

}
