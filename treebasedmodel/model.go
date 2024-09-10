package treebasedmodel

import (
	"fmt"
	"math"
)

type GameState struct {
	Board  [][]string `json:"board"`
	Player string     `json:"player"`
}

/*
los colores de los jugadores son:
		B: blue
		G: green
		Y: yellow
		R: red
    L: loose (jugadores que perdieron)
 
La primer letra de una pieza aliada indica su color y la segunda su tipo

		BP: peon azul
		BT: torre azul
		BC: caballo azul
		BA: alfil azul
		BD: dama azul
		BR: rey azul

		"-" casilla vacia
*/
func GetBestMove(gameState GameState) (Move,error){  

	println("Jugador actual:", gameState.Player)
	/*primero obtenemos el arbol para el jugador actual n movimientos al futuro*/
	tree, err := GenerateTree(gameState.Board, gameState.Player, 4) 
	if err != nil {
		fmt.Println("Error:", err)
    return Move{},err
	}

	_, bestNode := searchNode(tree)
  return bestNode.Move,nil
}

/*Considerar usar pila de esta forma en lugar de recursion*/
func DepthFirstSearch(node *Node) {
	stack := []*Node{node}

	for len(stack) > 0 {
		currentNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		currentNode.ShowData()

		stack = append(stack, currentNode.Childs...)
	}
}

func searchNode(node *Node) (int, *Node) {
	// Valores iniciales para alfa y beta
	alpha := math.MinInt32
	beta := math.MaxInt32

	var bestNode *Node
	bestValue := math.MinInt32

	// Función interna recursiva para aplicar poda alfa-beta
	var alphaBeta func(n *Node, alpha, beta int, maximizingPlayer bool) (int, *Node)
	alphaBeta = func(n *Node, alpha, beta int, maximizingPlayer bool) (int, *Node) { 
		if len(n.Childs) == 0 { // Nodo hoja
			return n.Value, n
		}

		if maximizingPlayer {
			value := math.MinInt32
			var best *Node
			for _, child := range n.Childs {
				childValue, _ := alphaBeta(child, alpha, beta, false)
				if childValue > value {
					value = childValue
					best = child
				}
				alpha = max(alpha, value)
				if beta <= alpha {
					println("Poda tipo alpha en nodo", n.Value, "con alfa:", alpha, "y beta:", beta)
					break
				}
			}
			return value, best
		} else {
			value := math.MaxInt32
			var best *Node
			for _, child := range n.Childs {
				childValue, _ := alphaBeta(child, alpha, beta, true)
				if childValue < value {
					value = childValue
					best = child
				}
				beta = min(beta, value)
				if beta <= alpha {
					fmt.Println("Poda tipo beta en nodo", n.Value, "con alfa:", alpha, "y beta:", beta)
					break
				}
			}
			return value, best
		}
	}

	bestValue, bestNode = alphaBeta(node, alpha, beta, true) // Empezar con el nodo raíz como el nodo maximizador
	return bestValue, bestNode
}
