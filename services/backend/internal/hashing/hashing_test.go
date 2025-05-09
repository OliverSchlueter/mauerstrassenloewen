package hashing_test

import (
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/hashing"
	"testing"
)

func TestSHA256(t *testing.T) {
	for _, tc := range []struct {
		in  string
		exp string
	}{
		{in: "test", exp: "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"},
		{in: "my-secure-password", exp: "8024102039a0d259ded6b6c6233199cc663f69df27a9a5338ac684422ed273d3"},
	} {
		t.Run("test", func(t *testing.T) {
			got := hashing.SHA256(tc.in)
			if got != tc.exp {
				t.Errorf("got %s, want %s", got, tc.exp)
			}
		})
	}
}
