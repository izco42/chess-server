package treebasedmodel

import (
	"math"
	"fmt"
)
// Estructura de un movimiento
type Move struct {
	From  [2]int
	To    [2]int
	Piece string
}

type Node struct {
	Interval [2]float64
	Board [][]string 
	Value float64
	Childs []*Node
	Visited bool
	Move Move
}

func NewNode(board [][]string, move Move) *Node {
	return &Node{
		Interval: [2]float64{math.Inf(-1), math.Inf(1)},
		Board: board,
		Value: 0,
		Childs: nil,
		Visited: false,
		Move: move,
	}
}

func (n *Node) AddChild(child *Node) {
	n.Childs = append(n.Childs, child)
}


/*funcion para mostrar los datos del nodo*/
func (n *Node) ShowData() {
	println("----------------------------------")
	println("Move: ", fmt.Sprintf("%v", n.Move))
	println("Board: ")
	for _, row := range n.Board {
		// Imprimir cada fila como una cadena formateada
		fmt.Printf("%v\n", row)
	}
}


/*funcion para calcular el valor */




