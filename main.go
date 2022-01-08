package main

import (
	"fmt"
	"github.com/milhamh95/erys/db"
)

func FindUser(username string) (db.User, error) {
	user, err := db.FindUser(username)
	if err != nil {
		return db.User{}, fmt.Errorf("failed to Find User: %v", err)
	}

	return user, nil
}

func SetUserAge(u db.User, age int) error {
	err := db.SetUserAge(u, age)
	if err != nil {
		return fmt.Errorf("failed to SetUserAge: %v", err)
	}
	return db.SetUserAge(u, age)
}

func FindAndSetUserAge(username string, age int) error {
	var (
		user db.User
		err  error
	)

	user, err = FindUser(username)
	if err != nil {
		return fmt.Errorf("failed to FindAndSetUserAge: %v", err)
	}

	if err = SetUserAge(user, age); err != nil {
		return fmt.Errorf("failed to FindAndSetUserAge: %v", err)
	}

	return nil
}

func main() {
	if err := FindAndSetUserAge("bob@example.com", 21); err != nil {
		fmt.Printf("failed finding or updating user: %v", err)
		return
	}

	fmt.Println("successfully updated user's age")
}
