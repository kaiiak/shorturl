package util

import "testing"

func TestConvertTo62(t *testing.T) {
	type args struct {
		i int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple",
			args: args{
				i: 1,
			},
			want: "1",
		},
		{
			name: "multi",
			args: args{
				i: 123,
			},
			want: "1Z",
		},
		{
			name: "big",
			args: args{
				i: 100000000,
			},
			want: "6LAze",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertTo62(tt.args.i); got != tt.want {
				t.Errorf("ConvertTo62() = %v, want %v", got, tt.want)
			}
		})
	}
}
