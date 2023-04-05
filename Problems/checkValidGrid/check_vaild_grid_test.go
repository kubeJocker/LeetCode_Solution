package checkValidGrid

import "testing"

func Test_checkValidGrid(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{grid: [][]int{{0, 11, 16, 5, 20}, {17, 4, 19, 10, 15}, {12, 1, 8, 21, 6}, {3, 18, 23, 14, 9}, {24, 13, 2, 7, 22}}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValidGrid(tt.args.grid); got != tt.want {
				t.Errorf("checkValidGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
