package movies

import (
	"maps"
	"testing"
)

func TestValidateMovie(t *testing.T) {
	tests := []struct {
		name  string
		movie Movie
		want  map[string]string
	}{
		{
			name: "Valid",
			movie: Movie{
				Title:       "validTitle",
				ReleaseYear: 2026,
				Genre:       "validGenre",
				Runtime:     90,
			},
			want: make(map[string]string),
		},
		{
			name: "EmptyString",
			movie: Movie{
				Title:       "",
				ReleaseYear: 2026,
				Genre:       "",
				Runtime:     90,
			},
			want: map[string]string{
				"title": "title must be provided",
				"genre": "genre must be provided",
			},
		},
		{
			name: "EmptyInt",
			movie: Movie{
				Title: "validTitle",
				Genre: "validGenre",
			},
			want: map[string]string{
				"release_year": "release year must be provided",
				"runtime":      "runtime must be provided",
			},
		},
		{
			name: "NegativeRuntime",
			movie: Movie{
				Title:       "validTitle",
				ReleaseYear: 2026,
				Genre:       "validGenre",
				Runtime:     -1,
			},
			want: map[string]string{
				"runtime": "runtime must be a positive integer",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validationErrors := validateMovie(tt.movie)
			if !maps.Equal(validationErrors, tt.want) {
				t.Errorf("got %v, want %v", validationErrors, tt.want)
			}
		})
	}
}
