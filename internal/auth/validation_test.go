package auth

import (
	"maps"
	"strings"
	"testing"
)

func TestValidateRegistration(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		password string
		want     map[string]string
	}{
		{
			name:     "Valid",
			email:    "validemail@gmail.com",
			password: "validPassword",
			want:     make(map[string]string),
		},
		{
			name:     "EmptyCredentials",
			email:    "",
			password: "",
			want: map[string]string{
				"email":    "email cannot be empty",
				"password": "password cannot be empty",
			},
		},
		{
			name:     "EmailNoDomain",
			email:    "invalidemail",
			password: "validPassword",
			want: map[string]string{
				"email": "must be a valid email",
			},
		},
		{
			name:     "EmailNoUser",
			email:    "@gmail.com",
			password: "validPassword",
			want: map[string]string{
				"email": "must be a valid email",
			},
		},
		{
			name:     "ShortPassword",
			email:    "validemail@gmail.com",
			password: "invalid",
			want: map[string]string{
				"password": "password must be at least 8 characters long",
			},
		},
		{
			name:     "LongPassword",
			email:    "validemail@gmail.com",
			password: strings.Repeat("é", 37),
			want: map[string]string{
				"password": "password must not be more than 72 bytes long",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validationErrors := validateRegistration(tt.email, tt.password)
			if !maps.Equal(validationErrors, tt.want) {
				t.Errorf("got %v, want %v", validationErrors, tt.want)
			}
		})
	}
}
