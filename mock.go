package Ajit

import (
	"fmt"
	"log"
)

type User struct {
	Name     string
	Email    string
	UserName string
}
type precheck interface {
	userExists(string) bool
}
type regiCheck struct{}

func (r regiCheck) userExists(email string) bool {
	return UserAvilable(email)
}

var regCond precheck

func init() {
	regCond = regiCheck{}
}
func NewUser(u User) error {

	avilable := regCond.userExists(u.Email)
	if avilable {
		return fmt.Errorf("email is already '%s'avilable", u.Email)
	}
	log.Println(u.Name)
	return nil

}
