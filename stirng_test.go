package chaos

import (
	"testing"
)

func Test_FindAfterSubstr(t *testing.T) {
	type args struct {
		s      string
		substr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// substr=''
		{"1", args{"", ""}, ""},
		{"1", args{"1", ""}, ""},
		{"1", args{"1/", ""}, ""},

		// s=''
		{"2", args{"", "1"}, ""},
		{"2", args{"", "1/3"}, ""},
		{"2", args{"", "///"}, ""},

		// substr not in s
		{"3", args{"1", "/"}, ""},
		{"3", args{"12", "/"}, ""},
		{"3", args{"123", "/"}, ""},

		// s contain one substr
		{"4", args{"1/", "/"}, ""},
		{"4", args{"12/", "/"}, ""},
		{"4", args{"/1", "/"}, "1"},
		{"4", args{"1/2", "/"}, "2"},
		{"4", args{"1/2", "/"}, "2"},

		// s contain multi substr
		{"4", args{"1//", "/"}, ""},
		{"4", args{"1/2/", "/"}, ""},
		{"4", args{"/2/1", "/"}, "1"},
		{"4", args{"/1/2", "/"}, "2"},
		{"4", args{"/3/1/2", "/"}, "2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindLastAfterSubstr(tt.args.s, tt.args.substr); got != tt.want {
				t.Errorf("FindLastAfterSubstr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CountPrefix(t *testing.T) {
	type args struct {
		s      string
		substr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// substr=''
		{"1", args{"", ""}, 0},
		{"1", args{"1", ""}, 0},
		{"1", args{"1/", ""}, 0},

		// s=''
		{"2", args{"", "1"}, 0},
		{"2", args{"", "1/3"}, 0},
		{"2", args{"", "///"}, 0},

		// substr not in s
		{"3", args{"1", "/"}, 0},
		{"3", args{"12", "/"}, 0},
		{"3", args{"123", "/"}, 0},

		// substr not at prefix s
		{"4", args{"1/", "/"}, 0},
		{"4", args{"1/2", "/"}, 0},
		{"5", args{"123/", "/"}, 0},

		// s has one prefix
		{"5", args{"/1", "/"}, 1},
		{"5", args{"/12", "/"}, 1},
		{"5", args{"/12/3/4", "/"}, 1},
		{"5", args{"/1////", "/"}, 1},

		// start with space
		{"6", args{" /1", "/"}, 0},
		{"6", args{" /12", "/"}, 0},
		{"6", args{" /12/3/4", "/"}, 0},
		{"6", args{" /1////", "/"}, 0},

		// has multi prefix
		{"7", args{"//1//", "/"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountPrefix(tt.args.s, tt.args.substr); got != tt.want {
				t.Errorf("CountPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
