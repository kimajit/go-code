package Ajit

import "testing"

var userExistsMock func(email string) bool

type preCheckMock struct{}

func (u preCheckMock) userExists(email string) bool {
	return userExistsMock(email)
}
func TestNewUser(t *testing.T) {
	t.Parallel()
	user := User{
		Name:     "ajit kumar",
		Email:    "ajithkumar.sinha@srsconsultinginc.com",
		UserName: "ajit",
	}
	regCond = preCheckMock{}
	userExistsMock = func(email string) bool {
		return false
	}
	err1 := NewUser(user)
	if err1 != nil {
		t.Fatal(err1)
	}

	userExistsMock = func(email string) bool {
		return true
	}
	err2 := NewUser(user)
	if err2 == nil {
		t.Errorf("throw an error got nil")
	}

}
