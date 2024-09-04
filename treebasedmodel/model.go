package treebasedmodel

import (
	"encoding/json"
	"fmt"
	"errors"
	//"math"
)

/*esta funcion implementara poda alpha beta para regresar la mejor jugada para un board recibido
este sera de un tablero 4x4 donde las piezas seran las siguientes:

los colores de los jugadores son:
		B: blue
		G: green
		Y: yellow
		R: red

La primer letra de una pieza aliada indica su color y la segunda su tipo
	--piezas aliadas--

		BP: peon azul
		BT: torre azul
		BC: caballo azul
		BA: alfil azul
		BD: dama azul
		BR: rey azul

		"-" casilla vacia
*/

func GetBestMove() {

	data := `{
	"board": [
		["-","-","-","YT","YC","YA","YR","YD","YA","YC","YT","-","-","-"],
		["-","-","-","YP","YP","YP","YP","YP","YP","YP","YP","-","-","-"],
		["-","-","-","-","-","-","-","-","-","-","-","-","-","-"],
		["GT","GP","-","-","-","-","-","-","-","-","-","-","RP","RT"],
		["GC","GP","-","-","-","-","-","-","-","-","-","-","RP","RC"],
		["GA","GP","-","-","-","-","-","-","-","-","-","-","RP","RA"],
		["GR","GP","-","-","-","-","-","-","-","-","-","-","RP","RR"],
		["GD","GP","-","-","-","-","-","-","-","-","-","-","RP","RD"],
		["GA","GP","-","-","-","-","-","-","-","-","-","-","RP","RA"],
		["GC","GP","-","-","-","-","-","-","-","-","-","-","RP","RC"],
		["GT","GP","-","-","-","-","-","-","-","-","-","-","RP","RT"],
		["-","-","-","-","-","-","-","-","-","-","-","-","-","-"],
		["-","-","-","BP","BP","BP","BP","BP","BP","BP","BP","-","-","-"],
		["-","-","-","BT","BC","BA","BR","BD","BA","BC","BT","-","-","-"]
	],
	"player": "G"
}`
	type GameState struct {
		Board  [][]string `json:"board"`
		Player string     `json:"player"`
	}

	var gameState GameState

	err := json.Unmarshal([]byte(data), &gameState)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// for _, row := range gameState.Board {
	// 	fmt.Println(row)
	// }

	println("Jugador actual:", gameState.Player)
	


	tree, err := GenerateTree(gameState.Board, gameState.Player,2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	DepthFirstSearch(tree)

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
				newNode := NewNode(newBoard, move)
				node.AddChild(newNode)
				nextLevelNodes = append(nextLevelNodes, newNode)
			}
		}

		// Avanzar al siguiente nivel
		currentLevelNodes = nextLevelNodes
	}

	return root, nil
}


/*funcion que hace recorrido en profundidad a partir del nodo que recibe usando pila*/
func DepthFirstSearch(node *Node) {
	stack := []*Node{node}

	for len(stack) > 0 {
		currentNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		currentNode.ShowData()

		stack = append(stack, currentNode.Childs...)
	}
}




/*ejemplo uso de nil (se regresa el valor y nil si todo ha ido bien , )*/
// // suma toma dos enteros y retorna la suma y un error en caso de que algo falle.
// func suma(a, b int) (int, error) {
//     // En este ejemplo, no hay errores posibles, pero podemos simular uno.
//     if a < 0 || b < 0 {
//         return 0, errors.New("no se permiten números negativos")
//     }
//     return a + b, nil
// }

// func main() {
//     resultado, err := suma(3, 5)
//     if err != nil {
//         fmt.Println("Error:", err)
//     } else {
//         fmt.Println("El resultado es:", resultado)
//     }
// }
