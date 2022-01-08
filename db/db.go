package db

import "fmt"

type User struct {
	ID       string
	Username string
	Age      int
}

func FindUser(username string) (User, error) {
	//return User{}, fmt.Errorf("failed to  find user: not found")
	return User{
		ID:       "aljk87",
		Username: "Goth13",
		Age:      20,
	}, nil

}

func SetUserAge(user User, age int) error {
	return fmt.Errorf("failed to set user age: data is not valid")
}
