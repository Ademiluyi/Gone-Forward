package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const COLUMN = 9
const ROW = 5

type Game struct {
	players map[string]string
	board   [ROW][COLUMN]string
}

// __init__
func NewGame() *Game {
	players := make(map[string]string)
	new_board := [ROW][COLUMN]string{}

	//buffering
	for i := 0; i < ROW; i++ {
		for j := 0; j < COLUMN; j++ {
			new_board[i][j] = " "
		}
	}
	Game := Game{
		players: players,
		board:   new_board,
	}
	return &Game
}

/*
This function plays the game by calling other functions in a loop.
*/
func (game Game) play_game() {
	intro()
	players_ := game.get_and_set_players()
	win := false
	turn := 0
	spot_count := 0

	if game.players[players_[turn]] != "x" {
		turn = 1
	}

	for !win {
		game.print_game_board()
		spot_count++
		// check for a tie
		if spot_count == ROW*COLUMN {
			fmt.Println("A TIE!!!!! PLAY A NEW GAME CHAMPIONS !")
			break
		}
		play := game.get_and_stamp_piece(players_[turn])

		win = game.check_winner(players_[turn], play)

		switch turn {
		case 0:
			turn = 1
		case 1:
			turn = 0
		}
	}

	switch turn {
	case 0:
		turn = 1
	case 1:
		turn = 0
	}

	game.print_game_board()
	fmt.Printf("%v WINS !!!!! ", players_[turn])
}

/*
This function prints Introduction
*/
func intro() {
	fmt.Println()
	game_header := `**************************************************************************
*     Hey, Welcome! Here's a multiplayer game of TicTacConnectFour!      *
**************************************************************************`
	game_rules := `
➢ There can be exactly just 2 players.
➢ Each player will enter their names after which they will randomly be assigned a game piece.
➢ Whoever gets assigned the 'x' piece plays first.
➢ The first player to place 4 consecutive pieces horizontally, vertically, or diagonally wins.
➢ If there is a tie, play a new game & enjoy !
	`
	fmt.Println(game_header)
	fmt.Println()
	fmt.Println("Below are the rules of the game :")
	fmt.Println(game_rules)
}

/*
This function checks for a winner.
*/
func (game Game) check_winner(player, play string) bool {
	uncast_row, uncast_column := convert_to_coords(play)
	row, column := int(uncast_row), int(uncast_column)
	game_piece := game.players[player]

	for i := row - 1; i <= row+1; i++ {
		for j := column - 1; j <= column+1; j++ {
			if (i >= 0 && i <= ROW-1) && (j >= 0 && j <= COLUMN-1) {
				if i == row && j == column {
					continue
				} else {
					if game.board[i][j] == game_piece {
						offset := 3
						row_bound := 4
						column_bound := 8

						if j == column {
							if i < row {
								// check upward 3 spots
								if row-offset >= 0 {
									count := 0
									for i_ := 1; i_ <= offset; i_++ {
										if game.board[row-i_][column] == game_piece {
											count++
										}
									}
									if count == 3 {
										return true
									}
								}
							} else {
								// check downward 3 spots
								if row+offset <= row_bound {
									count := 0
									for i_ := 1; i_ <= offset; i_++ {
										if game.board[row+i_][column] == game_piece {
											count++
										}
									}
									if count == 3 {
										return true
									}
								}

							}
						} else if i == row {
							if j < column {
								// check leftmost 3 spots
								if column-offset >= 0 {
									count := 0
									for i_ := 1; i_ <= offset; i_++ {
										if game.board[row][column-i_] == game_piece {
											count++
										}
									}
									if count == 3 {
										return true
									}
								}
							} else {
								// check rightmost 3 spots
								if column+offset <= column_bound {
									count := 0
									for i_ := 1; i_ <= offset; i_++ {
										if game.board[row][column+i_] == game_piece {
											count++
										}
									}
									if count == 3 {
										return true
									}
								}
							}
						} else if j == column-1 {
							// check top left diagonal
							if i < row {
								fmt.Println(row - offset)
								if row-offset >= 0 && column-offset >= 0 {
									count := 0
									for i_ := 1; i_ <= offset; i_++ {
										if game.board[row-i_][column-i_] == game_piece {
											count++
										}
									}
									if count == 3 {
										return true
									}
								}
							} else {
								// check bottom left diagonal
								if row+offset <= row_bound && column-offset >= 0 {
									count := 0
									for i_ := 1; i_ <= offset; i_++ {
										if game.board[row+i_][column-i_] == game_piece {
											count++
										}
									}
									if count == 3 {
										return true
									}
								}
							}

						} else if j == column+1 {
							// check top right diagonal
							if i < row {
								if row-offset >= 0 && column+offset <= column_bound {
									count := 0
									for i_ := 1; i_ <= offset; i_++ {
										if game.board[row-i_][column+i_] == game_piece {
											count++
										}
									}
									if count == 3 {
										return true
									}
								}
							} else {
								// check bottom right diagonal
								if row+offset <= row_bound && column+offset <= column_bound {
									count := 0
									for i_ := 1; i_ <= offset; i_++ {
										if game.board[row+i_][column+i_] == game_piece {
											count++
										}
									}
									if count == 3 {
										return true
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return false
}

/*
This function gets and sets information on the player, including player names & assigned game piece.
*/
func (game *Game) get_and_set_players() []string {
	game_piece := []string{"x", "o"}
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	player_names := []string{}
	for i := 0; i < 2; i++ {
		fmt.Printf("Player %d enter a name : ", i+1)
		username, _ := reader.ReadString('\n')
		trimmedUser := strings.TrimSpace(username)
		player_names = append(player_names, trimmedUser)
	}

	// randomnizes who gets what piece, and in turn decides who starts the game.
	rand_piece := (rand.Intn(10-0) + 0) % 2
	game.players[player_names[0]] = game_piece[rand_piece]
	switch rand_piece {
	case 0:
		game.players[player_names[1]] = game_piece[1]
	case 1:
		game.players[player_names[1]] = game_piece[0]
	}

	colorGreen := "\033[32m"
	colorReset := "\033[0m"
	fmt.Println()
	fmt.Printf("%v will play with game piece :", strings.ToUpper(player_names[0]))
	fmt.Println(string(colorGreen), game.players[player_names[0]], string(colorReset))
	fmt.Printf("%v will play with game piece :", strings.ToUpper(player_names[1]))
	fmt.Println(string(colorGreen), game.players[player_names[1]], string(colorReset))

	return player_names
}

/*
This function gets a play from the user and stamps it on the board.
*/
func (game *Game) get_and_stamp_piece(player string) string {
	reader := bufio.NewReader(os.Stdin)
	valid := false
	stat := ""
	var trimmedPlay string

	for !valid {
		fmt.Printf("%v! make your play : ", player)
		play, _ := reader.ReadString('\n')
		trimmedPlay = strings.TrimSpace(play)
		valid, stat = game.is_valid(trimmedPlay)
		if !valid {
			if stat == "invalid" {
				fmt.Println("Invalid play. try again.")
			} else {
				fmt.Println("spot taken!")
			}
		} else {
			game.board[(trimmedPlay[0]-'0')-1][(trimmedPlay[1]-'0')-1] = game.players[player]
		}
	}
	return trimmedPlay
}

/*
This function converts a play to a corresponding coordinate on the 2d representation of the game board.
*/
func convert_to_coords(play string) (uint8, uint8) {
	var ascii_one uint8
	ascii_one = 49

	row := play[0] - ascii_one
	column := play[1] - ascii_one

	return row, column
}

/*
This function checks if a play is valid.
*/
func (game *Game) is_valid(play string) (bool, string) {

	var zero uint8
	var max_unit uint8
	var max_tens uint8

	max_tens, max_unit, zero = 5, 9, 0

	if len(play) == 2 {
		row, column := convert_to_coords(play)
		if row >= zero && row < max_tens {
			if column >= zero && column < max_unit {
				if game.board[row][column] == " " {
					return true, "valid"
				} else {
					return false, "taken"
				}
			}
		}
	}
	return false, "invalid"
}

/*
This function prints the game board.
*/
func (game Game) print_game_board() {
	fmt.Println()
	fmt.Println("  1   2   3   4   5   6   7   8   9")
	row_label := 1

	for i := 0; i < ROW; i++ {
		for j := 0; j < COLUMN; j++ {
			if j == 0 {
				fmt.Printf("%d", row_label)
				row_label++
			}
			fmt.Printf(" %v ", game.board[i][j])
			if j != COLUMN-1 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i != ROW-1 {
			fmt.Println("  ----------------------------------")
		}
	}
	fmt.Println()
}

func main() {
	NewGame().play_game()
}
