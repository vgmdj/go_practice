package abstract_factory

import (
	"fmt"
	"testing"
)

func TestUserDaoFactory(t *testing.T) {
	factory := UserDaoFactoryImpl{}

	s := factory.GetUserGroupDao().GetUserGroup()
	fmt.Print(s) //GetUserGroup
	fmt.Printf("\n")
	s2 := factory.GetUserProfileDao().GetUserProfile()
	fmt.Print(s2) //GetUserProfile
}
