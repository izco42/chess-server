package treebasedmodel

import (
	"fmt"
)

// Estructura de un movimiento
type Move struct {
	From  [2]int
	To    [2]int
	Piece string
}

// Estructura de un nodo
type Node struct {
	Board  [][]string
	Value  int
	Childs []*Node
	Move   Move
}

func NewNode(board [][]string, move Move) *Node {
	return &Node{
		Board:  board,
		Value:  0,
		Childs: nil,
		Move:   move,
	}
}

func (n *Node) AddChild(child *Node) {
	n.Childs = append(n.Childs, child)
}

func (n *Node) ShowData() {
	println("----------------------------------")
	println("Move: ", fmt.Sprintf("%v", n.Move))
	println("value: ", n.Value)
	println("Board: ")
	/*
		for _, row := range n.Board {
			// Imprimir cada fila como una cadena formateada
			fmt.Printf("%v\n", row)
		}*/
}
