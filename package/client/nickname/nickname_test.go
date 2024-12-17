package nickname_test

import (
	"testing"

	"github.com/Jakub-Pazio/Wren/package/client/nickname"
)

func TestNewNick(t *testing.T) {
	t.Run("creating valid nickname", func(t *testing.T) {
		validNick := "jpazio"

		nn, err := nickname.New(validNick)

		if err != nil {
			t.Errorf("should not error but got %q", err)
		}

		if nn != "jpazio" {
			t.Errorf("got %s want %s", nn, "jpazio")
		}
	})

	t.Run("should fail parsing forbidden nickname", func(t *testing.T) {
		cases := []struct {
			name          string
			nickname      string
			expectedError error
		}{
			{"dot in username", "j.pazio", nickname.ForbiddenCharError},
			{"empty username", "", nickname.EmptyNicknameError},
			{"starting with '%' character", "%jpazio", nickname.ForbiddenStartingCharError},
		}

		for _, tc := range cases {
			_, err := nickname.New(tc.nickname)
			if err == nil {
				t.Errorf("%v: should return error %v but did not", tc.name, err)
			}
			if err != tc.expectedError {
				t.Errorf("%v: got %v want %v", tc.name, err, tc.expectedError)
			}
		}
	})
}
