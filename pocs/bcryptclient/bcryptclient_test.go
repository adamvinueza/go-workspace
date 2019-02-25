package bcryptclient

import "testing"

type Test struct {
	Name           string
	UserID         string
	Password       string
	HashedPassword string
	Authenticates  bool
}

func TestBCryptClient(t *testing.T) {
	tests := []Test{
		{
			Name:           "first test",
			UserID:         "0",
			Password:       "password",
			HashedPassword: "$2a$12$yq6qOGgCBzUFukaXIPuEUeBcaglmV35B0FKrps.fInksuzfLsF/iC",
			Authenticates:  true,
		},
	}

	for _, j := range tests {
		got := authenticate(j.UserID, j.Password)
		if (got == nil) != j.Authenticates {
			t.Errorf("Test %s: expected authentication value of '%v', got '%v'", j.Name, j.Authenticates, got.Error())
		}
	}
}
