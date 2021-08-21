package rttables

import (
	"reflect"
	"testing"
)

func TestParseCustom(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    []Table
		wantErr bool
	}{
		{
			name: "ParseCustom should success parsing standard rt_tables file",
			args: args{
				file: "./testdata/standard.rt_tables",
			},
			want:    standardRtTables,
			wantErr: false,
		},
		{
			name: "ParseCustom should throws ErrMalformedFormat",
			args: args{
				file: "./testdata/malformed.rt_tables",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCustom(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCustom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCustom() = %v, want %v", got, tt.want)
			}
		})
	}
}
