package test

import (
	"Chess/model"
	"Chess/service"
	"testing"
)

func Test_checkBishopMove(t *testing.T) {
	type args struct {
		x1           int
		y1           int
		x2           int
		y2           int
		color        int
		checkerBoard *model.Chess
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		// TODO: Add test cases.
		{
			name: "Valid bishop move",
			args: args{
				x1:           3,
				y1:           4,
				x2:           6,
				y2:           7,
				color:        1,
				checkerBoard: &model.Chess{},
			},
			want:  true,
			want1: "Bishop moves in a diagonal pattern.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := service.CheckBishopMove(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.color, tt.args.checkerBoard)
			if got != tt.want {
				t.Errorf("checkBishopMove() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("checkBishopMove() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_checkKingMove(t *testing.T) {
	type args struct {
		x1           int
		y1           int
		x2           int
		y2           int
		color        int
		checkerBoard *model.Chess
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		// TODO: Add test cases.
		{
			name: "Valid king move",
			args: args{
				x1:           4,
				y1:           4,
				x2:           5,
				y2:           4,
				color:        1,
				checkerBoard: &model.Chess{},
			},
			want:  true,
			want1: "King moves one step horizontally.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := service.CheckKingMove(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.color, tt.args.checkerBoard)
			if got != tt.want {
				t.Errorf("checkKingMove() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("checkKingMove() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_checkKnightMove(t *testing.T) {
	type args struct {
		x1           int
		y1           int
		x2           int
		y2           int
		color        int
		checkerBoard *model.Chess
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		// TODO: Add test cases.
		{
			name: "Valid knight move",
			args: args{
				x1:           2,
				y1:           1,
				x2:           4,
				y2:           2,
				color:        1,
				checkerBoard: &model.Chess{},
			},
			want:  true,
			want1: "Knight moves in an L-shape.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := service.CheckKnightMove(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.color, tt.args.checkerBoard)
			if got != tt.want {
				t.Errorf("checkKnightMove() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("checkKnightMove() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_checkPawnMove(t *testing.T) {
	type args struct {
		x1           int
		y1           int
		x2           int
		y2           int
		color        int
		checkerBoard *model.Chess
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		// TODO: Add test cases.
		{
			name: "Valid pawn move",
			args: args{
				x1:           1,
				y1:           1,
				x2:           1,
				y2:           2,
				color:        1,
				checkerBoard: &model.Chess{},
			},
			want:  true,
			want1: "Pawn moves one step forward.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := service.CheckPawnMove(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.color, tt.args.checkerBoard)
			if got != tt.want {
				t.Errorf("checkPawnMove() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("checkPawnMove() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_checkQueenMove(t *testing.T) {
	type args struct {
		x1           int
		y1           int
		x2           int
		y2           int
		color        int
		checkerBoard *model.Chess
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		// TODO: Add test cases.
		{
			name: "Valid queen move",
			args: args{
				x1:           4,
				y1:           4,
				x2:           7,
				y2:           7,
				color:        1,
				checkerBoard: &model.Chess{},
			},
			want:  true,
			want1: "Queen moves in a straight or diagonal pattern.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := service.CheckQueenMove(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.color, tt.args.checkerBoard)
			if got != tt.want {
				t.Errorf("checkQueenMove() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("checkQueenMove() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_checkRookMove(t *testing.T) {
	type args struct {
		x1           int
		y1           int
		x2           int
		y2           int
		color        int
		checkerBoard *model.Chess
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		// TODO: Add test cases.
		//{
		//	name: "Valid rook move",
		//	args: args{
		//		x1:           1,
		//		y1:           1,
		//		x2:           8,
		//		y2:           1,
		//		color:        1,
		//		checkerBoard: &model.Chess{},
		//	},
		//	want:  true,
		//	want1: "Rook moves in a straight line.",
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := service.CheckRookMove(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.color, tt.args.checkerBoard)
			if got != tt.want {
				t.Errorf("checkRookMove() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("checkRookMove() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
