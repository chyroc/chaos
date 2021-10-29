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

func TestContainAny(t *testing.T) {
	type args struct {
		s          string
		substrList []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// nil
		{"1", args{"", nil}, false},
		{"1", args{"1", nil}, false},
		{"1", args{"12", nil}, false},

		// []string{}
		{"2", args{"", []string{}}, false},
		{"2", args{"1", []string{}}, false},
		{"2", args{"12", []string{}}, false},

		// []string{}
		{"3", args{"", []string{""}}, true},
		{"3", args{"1", []string{""}}, true},

		// []string{}
		{"4", args{"", []string{"x"}}, false},
		{"4", args{"x", []string{"x"}}, true},
		{"4", args{"xx", []string{"x"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainAny(tt.args.s, tt.args.substrList); got != tt.want {
				t.Errorf("ContainAny() = %v, want %v", got, tt.want)
			}
		})
	}
}
