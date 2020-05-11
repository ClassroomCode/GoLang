package ttt_test

import (
	"testing"

	"github.com/gopherguides/learn/_training/testing/table/src/ttt"
)

// section: happy_path
func Test_ApplyMove_HappyPath(t *testing.T) {
	state := ttt.GameState{
		Board: ttt.Board([][]rune{
			{'X', '-', '-'},
			{'-', '-', '-'},
			{'-', '-', '-'},
		}),
		NextPlayer: ttt.PlayerO,
	}

	err := state.ApplyMove(0, 1, ttt.PlayerO)

	if err != nil {
		t.Fatal("Unexpected err applying move: err:", err)
	}
}

// section: happy_path

// section: out_of_bounds
func Test_ApplyMove_OutOfBounds(t *testing.T) {
	state := ttt.GameState{
		Board: ttt.Board([][]rune{
			{'X', '-', '-'},
			{'-', '-', '-'},
			{'-', '-', '-'},
		}),
		NextPlayer: ttt.PlayerO,
	}

	err := state.ApplyMove(20, 20, ttt.PlayerO)

	if err != ttt.ErrOutOfBounds {
		t.Fatalf("Expected out of bounds err but received %#v\n", err)
	}
}

// section: out_of_bounds

// section: double_moves
func Test_ApplyMove_DoubleMove(t *testing.T) {
	state := ttt.GameState{
		Board: ttt.Board([][]rune{
			{'X', '-', '-'},
			{'-', '-', '-'},
			{'-', '-', '-'},
		}),
		NextPlayer: ttt.PlayerX,
	}

	err := state.ApplyMove(0, 1, ttt.PlayerO)

	if err != ttt.ErrDoubleMove {
		t.Fatalf("Expected double move err but received %#v\n", err)
	}
}

// section: double_moves

// section: table_driven
func Test_TicTacToe_ApplyMove(t *testing.T) {
	gameTests := []struct {
		name        string
		state       ttt.GameState
		xcoord      int
		ycoord      int
		player      ttt.Player
		expectedErr error
	}{
		{
			"happy path",
			ttt.GameState{
				Board: ttt.Board([][]rune{
					{'X', '-', '-'},
					{'-', '-', '-'},
					{'-', '-', '-'},
				}),
				NextPlayer: ttt.PlayerO,
			},
			0, 1,
			ttt.PlayerO,
			nil,
		},
		{
			"out of bounds moves",
			ttt.GameState{
				Board: ttt.Board([][]rune{
					{'X', '-', '-'},
					{'-', '-', '-'},
					{'-', '-', '-'},
				}),
				NextPlayer: ttt.PlayerO,
			},
			20, 20,
			ttt.PlayerO,
			ttt.ErrOutOfBounds,
		},
	}

	for _, tt := range gameTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.state.ApplyMove(tt.xcoord, tt.ycoord, tt.player)
			if err != tt.expectedErr {
				t.Fatalf("Unexpected err: got: %#v, want: %#v\n", err, tt.expectedErr)
			}
		})
	}
}

// section: table_driven
