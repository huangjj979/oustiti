package greeting

import "testing"

func TestGreet(t *testing.T) {
	type args struct {
		somebody string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"greet Fang",
			args{"Fang"},
			"Hello Fang!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greet(tt.args.somebody); got != tt.want {
				t.Errorf("Greet() = %v, want %v", got, tt.want)
			}
		})
	}
}
