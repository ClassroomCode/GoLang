package ttt

import "errors"

var (
	ErrOutOfBounds = errors.New("Move out of bounds")
	ErrDoubleMove  = errors.New("Player has already made a move")
)

// section: types
type Player int

const (
	PlayerX Player = iota
	PlayerO
)

type Board [][]rune

type GameState struct {
	Board      Board
	NextPlayer Player
}

// section: types

// section: apply_move_sig
func (g *GameState) ApplyMove(x, y int, player Player) error {
	// section: apply_move_sig
	if x >= len(g.Board) || y >= len(g.Board) {
		return ErrOutOfBounds
	}

	if player != g.NextPlayer {
		return ErrDoubleMove
	}
	return nil
}
