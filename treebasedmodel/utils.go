package treebasedmodel

import (
	"errors"
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*la idea de la siguiente funcion es tener una lista con los nodos de cada nivel , se recorre esa lista y se piden los posibles movimientos para el jugador indicado
con getmovesposibles y se crea un nodo por cada movimiento posible , despues se actualiza la lista de nodos con los nuevos nodos creados y se eliminan
los nodos que ya no se necesitan , ahora se generan los movimientos posibles para el oponente con la funcion
getPossiblesMovesForOpponent y se repite el proceso anterior , se repite el proceso hasta llegar a la profundidad deseada
*/

func GenerateTree(board [][]string, player string, depth int) (*Node, error) {
	if depth < 1 {
		return nil, errors.New("la profundidad debe ser mayor a 0")
	}

	root := NewNode(board, Move{})
	currentLevelNodes := []*Node{root}

	for d := 0; d < depth; d++ {
		var nextLevelNodes []*Node

		for _, node := range currentLevelNodes {
			var moves []Move
			if d%2 == 0 {
				// Generar movimientos para el jugador max (actual)
				moves = GeneratePossibleMoves(node.Board, player)
				//moves = GeneratePossibleMovesForOpponent(node.Board, player)

			} else {
				moves = GeneratePossibleMovesForOpponent(node.Board, player)
			}

			for _, move := range moves {
				newBoard := ApplyMove(node.Board, move)
				valueNode, err := EvaluateBoard(newBoard, player)
				if err != nil {
					fmt.Println("Error:", err)
					return nil, err

				}
				newNode := NewNode(newBoard, move)
				newNode.Value = valueNode
				node.AddChild(newNode)
				nextLevelNodes = append(nextLevelNodes, newNode)
			}
		}

		// Avanzar al siguiente nivel
		currentLevelNodes = nextLevelNodes
	}

	return root, nil
}

/*funcion para calcular el valor de un nodo en base a las piezas presentes en el tablero para ese jugador, recibe un tablero y un jugador , regresa int value o un error , tomar en cuenta que las piezas en la matriz [][]string se ven de la forma "GP" donde el primer caracter es el color del jugador y el segundo caracter es la pieza*/

func EvaluateBoard(board [][]string, player string) (int, error) {
	if player != "B" && player != "Y" && player != "R" && player != "G" {
		return 0, errors.New("Jugador no válido")
	}
	// Definir el valor de las piezas
	pieceValues := map[string]int{
		"P": 1,   // Peón
		"T": 5,   // Torre
		"C": 3,   // Caballo
		"A": 3,   // Alfil
		"D": 9,   // Reina
		"R": 100, // Rey
	}
	// Calcular el valor del nodo
	nodeValue := 0
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			piece := board[row][col]
			if len(piece) > 0 {
				// Verificar si la pieza es del jugador actual
				if IsPlayerPiece(piece, player) {
					// Obtener el valor de la piezas
					pieceValue := pieceValues[string(piece[1])]
					// Sumar el valor de la pieza al valor del nodo
					nodeValue += pieceValue
				}
			}
		}
	}
	return nodeValue, nil
}
