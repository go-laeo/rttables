package rttables

import (
	"bytes"
	"os"
	"testing"
)

func TestPatchWriter(t *testing.T) {
	type args struct {
		rows []Table
	}
	b, err := os.ReadFile("./testdata/standard.rt_tables")
	if err != nil {
		t.Errorf("testdata reads failed: %w", err)
		return
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		{
			name: "PatchWriter should generate success",
			args: args{
				rows: standardRtTables,
			},
			wantW: string(b),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := PatchWriter(w, tt.args.rows); (err != nil) != tt.wantErr {
				t.Errorf("PatchWriter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("PatchWriter() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
