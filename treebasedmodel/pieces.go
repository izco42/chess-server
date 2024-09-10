package treebasedmodel

/*cada funcion recibira una tupla [x,y] de donde esta la pieza, ademas de un string que indique el color de
la pieza , ya que dependiendo el color sus movimientos seran diferentes
y retornara un arreglo con las posibles jugadas

recordar que es un tablero de 14x14 y que los indices invalidos son
--ezquina superior izquierda--
[0,0], [0,1] , [0,2]
[1,0], [1,1] , [1,2]
[2,0], [2,1] , [2,2]

--ezquina superior derecha--

[0,11], [0,12] , [0,13]
[1,11], [1,12] , [1,13]
[2,11], [2,12] , [2,13]


--ezquina inferior izquierda--
[11,0], [11,1] , [11,2]
[12,0], [12,1] , [12,2]
[13,0], [13,1] , [13,2]

--ezquina inferior derecha--
[11,11], [11,12] , [11,13]
[12,11], [12,12] , [12,13]
[13,11], [13,12] , [13,13]
*/

func clearInvalidMoves(moves [][2]int) [][2]int {
	var validMoves [][2]int
	for _, move := range moves {
		if move[0] < 0 || move[0] > 13 || move[1] < 0 || move[1] > 13 {
			continue
		}

		// Esquinas inválidas
		if (move[0] <= 2 && move[1] <= 2) || // Esquina superior izquierda
			(move[0] <= 2 && move[1] >= 11) || // Esquina superior derecha
			(move[0] >= 11 && move[1] <= 2) || // Esquina inferior izquierda
			(move[0] >= 11 && move[1] >= 11) { // Esquina inferior derecha
			continue
		}

		validMoves = append(validMoves, move)
	}
	return validMoves
}

/*la primer posicion indica altura la segunda anchura
ir hacia abajo implica sumar 1 a la altura
ir hacia arriba implica restar 1 a la altura
ir hacia la derecha implica sumar 1 a la anchura
ir hacia la izquierda implica restar 1 a la anchura
*/

func GetMovesPawn(position [2]int, board [][]string, color string) [][2]int {
	var moves [][2]int
	if color == "B" {
		//movimiento hacia adelante
		if board[position[0]-1][position[1]] == "-" {
			moves = append(moves, [2]int{position[0] - 1, position[1]})
		}
		//diagonal derecha
		if string(board[position[0]-1][position[1]+1][0]) != "B" && board[position[0]-1][position[1]+1] != "-" {
			moves = append(moves, [2]int{position[0] - 1, position[1] + 1})
		}
		//diagonal izquierda
		if string(board[position[0]-1][position[1]-1][0]) != "B" && board[position[0]-1][position[1]-1] != "-" {
			moves = append(moves, [2]int{position[0] - 1, position[1] - 1})
		}
	}

	if color == "Y" {
		if board[position[0]+1][position[1]] == "-" {
			moves = append(moves, [2]int{position[0] + 1, position[1]})
		}
		//diagonal derecha
		if string(board[position[0]+1][position[1]+1][0]) != "Y" && board[position[0]+1][position[1]+1] != "-" {
			moves = append(moves, [2]int{position[0] + 1, position[1] + 1})
		}
		//diagonal izquierda
		if string(board[position[0]+1][position[1]-1][0]) != "Y" && board[position[0]+1][position[1]-1] != "-" {
			moves = append(moves, [2]int{position[0] + 1, position[1] - 1})
		}
	}

	if color == "R" {
		//movimiento hacia adelante (esta pieza su adelante es a la misma altura pero moverse a la izquierda)
		if board[position[0]][position[1]-1] == "-" {
			moves = append(moves, [2]int{position[0], position[1] - 1})
		}
		//diagonal derecha
		if string(board[position[0]-1][position[1]-1][0]) != "R" && board[position[0]-1][position[1]-1] != "-" {
			moves = append(moves, [2]int{position[0] - 1, position[1] - 1})
		}
		//diagonal izquierda
		if string(board[position[0]+1][position[1]-1][0]) != "R" && board[position[0]+1][position[1]-1] != "-" {
			moves = append(moves, [2]int{position[0] + 1, position[1] - 1})
		}
	}

	if color == "G" {
		//movimiento hacia adelante (esta pieza su adelante es a la misma altura pero moverse a la derecha)
		if board[position[0]][position[1]+1] == "-" {
			moves = append(moves, [2]int{position[0], position[1] + 1})
		}
		//diagonal derecha
		if string(board[position[0]-1][position[1]+1][0]) != "G" && board[position[0]-1][position[1]+1] != "-" {
			moves = append(moves, [2]int{position[0] - 1, position[1] + 1})
		}
		//diagonal izquierda
		if string(board[position[0]+1][position[1]+1][0]) != "G" && board[position[0]+1][position[1]+1] != "-" {
			moves = append(moves, [2]int{position[0] + 1, position[1] + 1})
		}
	}

	return clearInvalidMoves(moves)
}

func GetMovesRook(position [2]int, board [][]string, color string) [][2]int {
	var moves [][2]int
	//movimiento hacia arrina (restar a la altura)
	for i := position[0] - 1; i >= 0; i-- {
		if board[i][position[1]] == "-" {
			moves = append(moves, [2]int{i, position[1]})
		} else if string(board[i][position[1]][0]) != color {
			moves = append(moves, [2]int{i, position[1]})
			break // Detenemos el bucle si encontramos una pieza del color opuesto
		} else {
			break // Detenemos el bucle si encontramos una pieza del mismo color
		}
	}

	//movimiento hacia abajo (sumar a la altura)
	for i := position[0] + 1; i < 14; i++ {
		if board[i][position[1]] == "-" {
			moves = append(moves, [2]int{i, position[1]})
		} else if string(board[i][position[1]][0]) != color {
			moves = append(moves, [2]int{i, position[1]})
			break // Detenemos el bucle si encontramos una pieza del color opuesto
		} else {
			break // Detenemos el bucle si encontramos una pieza del mismo color
		}
	}

	//movimiento hacia la derecha (sumar a la anchura)
	for i := position[1] + 1; i < 14; i++ {
		if board[position[0]][i] == "-" {
			moves = append(moves, [2]int{position[0], i})
		} else if string(board[position[0]][i][0]) != color {
			moves = append(moves, [2]int{position[0], i})
			break // Detenemos el bucle si encontramos una pieza del color opuesto
		} else {
			break // Detenemos el bucle si encontramos una pieza del mismo color
		}
	}

	//movimiento hacia la izquierda (restar a la anchura)
	for i := position[1] - 1; i >= 0; i-- {
		if board[position[0]][i] == "-" {
			moves = append(moves, [2]int{position[0], i})
		} else if string(board[position[0]][i][0]) != color {
			moves = append(moves, [2]int{position[0], i})
			break // Detenemos el bucle si encontramos una pieza del color opuesto
		} else {
			break // Detenemos el bucle si encontramos una pieza del mismo color
		}
	}
	return clearInvalidMoves(moves)
}

/*movimientos para el caballo*/
func GetMovesKnight(position [2]int, board [][]string, color string) [][2]int {
	var moves [][2]int
	//movimiento hacia arriba 2 arriba 1 a la derecha o 1 a la izquierda validar que ahi no haya una pieza del mismo color
	if string(board[position[0]-2][position[1]+1][0]) != color {
		moves = append(moves, [2]int{position[0] - 2, position[1] + 1})
	}

	//movimimento hacia arriba 2 y 1 a la izquierda
	if string(board[position[0]-2][position[1]-1][0]) != color {
		moves = append(moves, [2]int{position[0] - 2, position[1] - 1})
	}

	//movimiento hacia abajo 2 y 1 a la derecha
	if string(board[position[0]+2][position[1]+1][0]) != color {
		moves = append(moves, [2]int{position[0] + 2, position[1] + 1})
	}

	//movimiento hacia abajo 2 y 1 a la izquierda
	if string(board[position[0]+2][position[1]-1][0]) != color {
		moves = append(moves, [2]int{position[0] + 2, position[1] - 1})
	}

	//movimiento hacia la derecha 2 y 1 arriba
	if string(board[position[0]-1][position[1]+2][0]) != color {
		moves = append(moves, [2]int{position[0] - 1, position[1] + 2})
	}

	//movimiento hacia la derecha 2 y 1 abajo
	if string(board[position[0]+1][position[1]+2][0]) != color {
		moves = append(moves, [2]int{position[0] + 1, position[1] + 2})
	}

	//movimiento hacia la izquierda 2 y 1 arriba
	if string(board[position[0]-1][position[1]-2][0]) != color {
		moves = append(moves, [2]int{position[0] - 1, position[1] - 2})
	}

	//movimiento hacia la izquierda 2 y 1 abajo
	if string(board[position[0]+1][position[1]-2][0]) != color {
		moves = append(moves, [2]int{position[0] + 1, position[1] - 2})
	}

	return clearInvalidMoves(moves)
}

func GetMovesBishop(position [2]int, board [][]string, color string) [][2]int {
	var moves [][2]int

	//diagonal superior izquierda
	for i, j := position[0]-1, position[1]-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "-" {
			moves = append(moves, [2]int{i, j})
		} else if string(board[i][j][0]) != color {
			moves = append(moves, [2]int{i, j})
			break
		} else {
			break
		}
	}

	//diagonal superior derecha
	for i, j := position[0]-1, position[1]+1; i >= 0 && j < 14; i, j = i-1, j+1 {
		if board[i][j] == "-" {
			moves = append(moves, [2]int{i, j})
		} else if string(board[i][j][0]) != color {
			moves = append(moves, [2]int{i, j})
			break
		} else {
			break
		}
	}

	//diagonal inferior derecha
	for i, j := position[0]+1, position[1]+1; i < 14 && j < 14; i, j = i+1, j+1 {
		if board[i][j] == "-" {
			moves = append(moves, [2]int{i, j})
		} else if string(board[i][j][0]) != color {
			moves = append(moves, [2]int{i, j})
			break
		} else {
			break
		}

	}
	//diagonal inferior izquierda
	for i, j := position[0]+1, position[1]-1; i < 14 && j >= 0; i, j = i+1, j-1 {
		if board[i][j] == "-" {
			moves = append(moves, [2]int{i, j})
		} else if string(board[i][j][0]) != color {
			moves = append(moves, [2]int{i, j})
			break
		} else {
			break
		}
	}

	return clearInvalidMoves(moves)
}

func GetMovesQueen(position [2]int, board [][]string, color string) [][2]int {
	var moves [][2]int

	/*la reina se mueve en todas las direcciones siempre y se detiene al chocar con una pieza suya o enemiga*/
	//movimiento hacia arrina (restar a la altura)
	for i := position[0] - 1; i >= 0; i-- {
		if board[i][position[1]] == "-" {
			moves = append(moves, [2]int{i, position[1]})
		} else if string(board[i][position[1]][0]) != color {
			moves = append(moves, [2]int{i, position[1]})
			break // Detenemos el bucle si encontramos una pieza del color opuesto
		} else {
			break // Detenemos el bucle si encontramos una pieza del mismo color
		}
	}

	//movimiento hacia abajo (sumar a la altura)
	for i := position[0] + 1; i < 14; i++ {
		if board[i][position[1]] == "-" {
			moves = append(moves, [2]int{i, position[1]})
		} else if string(board[i][position[1]][0]) != color {
			moves = append(moves, [2]int{i, position[1]})
			break // Detenemos el bucle si encontramos una pieza del color opuesto
		} else {
			break // Detenemos el bucle si encontramos una pieza del mismo color
		}
	}

	//movimiento hacia la derecha (sumar a la anchura)
	for i := position[1] + 1; i < 14; i++ {
		if board[position[0]][i] == "-" {
			moves = append(moves, [2]int{position[0], i})
		} else if string(board[position[0]][i][0]) != color {
			moves = append(moves, [2]int{position[0], i})
			break // Detenemos el bucle si encontramos una pieza del color opuesto
		} else {
			break // Detenemos el bucle si encontramos una pieza del mismo color
		}

	}

	//movimiento hacia la izquierda (restar a la anchura)
	for i := position[1] - 1; i >= 0; i-- {
		if board[position[0]][i] == "-" {
			moves = append(moves, [2]int{position[0], i})
		} else if string(board[position[0]][i][0]) != color {
			moves = append(moves, [2]int{position[0], i})
			break // Detenemos el bucle si encontramos una pieza del color opuesto
		} else {
			break // Detenemos el bucle si encontramos una pieza del mismo color
		}
	}

	//diagonal superior izquierda
	for i, j := position[0]-1, position[1]-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "-" {
			moves = append(moves, [2]int{i, j})
		} else if string(board[i][j][0]) != color {
			moves = append(moves, [2]int{i, j})
			break
		} else {
			break
		}

	}

	//diagonal superior derecha
	for i, j := position[0]-1, position[1]+1; i >= 0 && j < 14; i, j = i-1, j+1 {
		if board[i][j] == "-" {
			moves = append(moves, [2]int{i, j})
		} else if string(board[i][j][0]) != color {
			moves = append(moves, [2]int{i, j})
			break
		} else {
			break
		}

	}

	//diagonal inferior derecha
	for i, j := position[0]+1, position[1]+1; i < 14 && j < 14; i, j = i+1, j+1 {
		if board[i][j] == "-" {
			moves = append(moves, [2]int{i, j})
		} else if string(board[i][j][0]) != color {
			moves = append(moves, [2]int{i, j})
			break
		} else {
			break
		}

	}

	//diagonal inferior izquierda
	for i, j := position[0]+1, position[1]-1; i < 14 && j >= 0; i, j = i+1, j-1 {
		if board[i][j] == "-" {
			moves = append(moves, [2]int{i, j})
		} else if string(board[i][j][0]) != color {
			moves = append(moves, [2]int{i, j})
			break
		} else {
			break
		}

	}

	return clearInvalidMoves(moves)
}

func GetMovesKing(position [2]int, board [][]string, color string) [][2]int {
	var moves [][2]int

	/*el rey se mueve en cualquier direccion 1 casilla siempre que no choque con una pieza aliada*/

	//movimiento hacia arriba
	if board[position[0]-1][position[1]] == "-" || string(board[position[0]-1][position[1]][0]) != color {
		moves = append(moves, [2]int{position[0] - 1, position[1]})
	}

	//movimiento hacia abajo
	if board[position[0]+1][position[1]] == "-" || string(board[position[0]+1][position[1]][0]) != color {
		moves = append(moves, [2]int{position[0] + 1, position[1]})
	}

	//movimiento hacia la derecha
	if board[position[0]][position[1]+1] == "-" || string(board[position[0]][position[1]+1][0]) != color {
		moves = append(moves, [2]int{position[0], position[1] + 1})
	}

	//movimiento hacia la izquierda
	if board[position[0]][position[1]-1] == "-" || string(board[position[0]][position[1]-1][0]) != color {
		moves = append(moves, [2]int{position[0], position[1] - 1})
	}

	//diagonal superior izquierda
	if board[position[0]-1][position[1]-1] == "-" || string(board[position[0]-1][position[1]-1][0]) != color {
		moves = append(moves, [2]int{position[0] - 1, position[1] - 1})
	}

	//diagonal superior derecha
	if board[position[0]-1][position[1]+1] == "-" || string(board[position[0]-1][position[1]+1][0]) != color {
		moves = append(moves, [2]int{position[0] - 1, position[1] + 1})
	}

	//diagonal inferior derecha
	if board[position[0]+1][position[1]+1] == "-" || string(board[position[0]+1][position[1]+1][0]) != color {
		moves = append(moves, [2]int{position[0] + 1, position[1] + 1})
	}

	//diagonal inferior izquierda
	if board[position[0]+1][position[1]-1] == "-" || string(board[position[0]+1][position[1]-1][0]) != color {
		moves = append(moves, [2]int{position[0] + 1, position[1] - 1})
	}

	return clearInvalidMoves(moves)
}

// GeneratePossibleMoves genera todos los movimientos posibles para el jugador actual
func GeneratePossibleMoves(board [][]string, player string) []Move {
	var moves []Move

	// Recorremos el tablero
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			piece := board[row][col]

			// Verificamos si la pieza pertenece al jugador actual
			if len(piece) > 0 && IsPlayerPiece(piece, player) && piece != "-" {
				// Obtenemos los movimientos válidos para la pieza actual
				pieceMoves := GetValidMovesForPiece(board, piece, [2]int{row, col}, player)

				// Agregamos los movimientos a la lista
				for _, to := range pieceMoves {
					moves = append(moves, Move{
						From:  [2]int{row, col},
						To:    to,
						Piece: piece,
					})
				}
			}
		}
	}

	return moves
}

/*esta funcion retornara todos los posibles movimientos para las 
piezas que no sean el jugador que se indica y que ademas no sean piezas grises*/      
func GeneratePossibleMovesForOpponent(board [][]string, player string) []Move {
	var moves []Move
	// Recorremos el tablero
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			piece := board[row][col]

			// Verificamos si la pieza pertenece al jugador actual y que no sea una pieza de un jugador que perdio 
			if len(piece) > 0 && !IsPlayerPiece(piece, player) && piece != "-" && !IsPieceLoose(piece){
 
				pieceMoves := GetValidMovesForPiece(board, piece, [2]int{row, col}, string(piece[0]))

				// Agregamos los movimientos a la lista
				for _, to := range pieceMoves {
					moves = append(moves, Move{
						From:  [2]int{row, col},
						To:    to,
						Piece: piece,
					})
				}
			}
		}
	}

	return moves
}

func IsPlayerPiece(piece string, player string) bool {
	return len(piece) > 0 && len(player) > 0 && piece[0] == player[0]
}

func IsPieceLoose(piece string) bool {
  return len(piece) > 0 && piece[0] == 'L' 
}

// Función que retorna los movimientos válidos de una pieza en base a su tipo
func GetValidMovesForPiece(board [][]string, piece string, position [2]int, color string) [][2]int {
	switch piece[1] { // Asume que el segundo carácter de 'piece' es el tipo de la pieza
	case 'P': // Peón
		return GetMovesPawn(position, board, color)
	case 'R': // Torre
		return GetMovesRook(position, board, color)
	case 'N': // Caballo
		return GetMovesKnight(position, board, color)
	case 'B': // Alfil
		return GetMovesBishop(position, board, color)
	case 'Q': // Reina
		return GetMovesQueen(position, board, color)
	case 'K': // Rey
		return GetMovesKing(position, board, color)
	default:
		return nil

	}
}

// ApplyMove aplica un movimiento en el tablero y devuelve un nuevo tablero
func ApplyMove(board [][]string, move Move) [][]string {
	// Crear una copia del tablero para no modificar el original
	newBoard := make([][]string, len(board))
	for i := range board {
		newBoard[i] = make([]string, len(board[i]))
		copy(newBoard[i], board[i])
	}

	// Obtener las posiciones del movimiento
	fromRow, fromCol := move.From[0], move.From[1]
	toRow, toCol := move.To[0], move.To[1]

	// Mover la pieza
	newBoard[toRow][toCol] = newBoard[fromRow][fromCol] // Mueve la pieza a la nueva posición

	// Vaciar la celda de origen
	newBoard[fromRow][fromCol] = "-" // Dejar vacía la celda de origen

	return newBoard
}
