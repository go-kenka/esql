package uitls

import "testing"

func TestPkgPath(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{p: "data"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PkgPath(tt.args.p); got != tt.want {
				t.Errorf("PkgPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
