package rttables

import "testing"

func TestContainsCustom(t *testing.T) {
	type args struct {
		file string
		t    Table
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "ContainsCustom should returns true",
			args: args{
				file: "./testdata/standard.rt_tables",
				t: Table{
					ID:   253,
					Name: "default",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ContainsCustom(tt.args.file, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("ContainsCustom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ContainsCustom() = %v, want %v", got, tt.want)
			}
		})
	}
}
