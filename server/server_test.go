package server

import "testing"

func TestIsAllowed(t *testing.T) {
	type args struct {
		regexes []string
		name    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty regexes",
			args: args{
				regexes: []string{},
				name:    "test",
			},
			want: true,
		},
		{
			name: "empty name",
			args: args{
				regexes: []string{"test"},
				name:    "",
			},
			want: false,
		},
		{
			name: "match",
			args: args{
				regexes: []string{"test"},
				name:    "test",
			},
			want: true,
		},
		{
			name: "no match",
			args: args{
				regexes: []string{"test"},
				name:    "no-match",
			},
			want: false,
		},
		{
			name: "multiple regexes",
			args: args{
				regexes: []string{"test", "test2"},
				name:    "test",
			},
			want: true,
		},
		{
			name: "multiple regexes no match",
			args: args{
				regexes: []string{"test", "test2"},
				name:    "no-match",
			},
			want: false,
		},
		{
			name: "multiple regexes match",
			args: args{
				regexes: []string{"test", "test2"},
				name:    "test2",
			},
			want: true,
		},
		{
			name: "multiple regexes match",
			args: args{
				regexes: []string{"test", "test2"},
				name:    "test2",
			},
			want: true,
		},
		{
			name: "multiple regexes match",
			args: args{
				regexes: []string{"test", "test2"},
				name:    "test2",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAllowed(tt.args.name, tt.args.regexes)
			if got != tt.want {
				t.Errorf("isAllowed() = %v, want %v (test: %s)", got, tt.want, tt.name)
			}
		})
	}

}
