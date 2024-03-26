package main

import (
	"bytes"
	"testing"
)

func TestTowerOfHanoi(t *testing.T) {
	type args struct {
		n    int
		from string
		aux  string
		to   string
		w    *bytes.Buffer
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name: "Case with 2 disks",
			args: args{
				n:    2,
				from: "A",
				aux:  "B",
				to:   "C",
				w:    bytes.NewBuffer(nil),
			},
			wantW: "Disk 1 moved from A to B.\nDisk 2 moved from A to C.\nDisk 1 moved from B to C.\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TowerOfHanoi(tt.args.w, tt.args.n, tt.args.from, tt.args.aux, tt.args.to)
			if gotW := tt.args.w.String(); gotW != tt.wantW {
				t.Errorf("TowerOfHanoi() = %v, want %v", gotW, tt.wantW)
			}
		})
		tt.args.w.Reset()
	}
}
