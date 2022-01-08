package main

import (
	"fmt"
	"github.com/milhamh95/erys/db"
)

// use case
func FindAndSetUserAge(username string, age int) error {
	var (
		user db.User
		err  error
	)

	user, err = db.FindUser(username)
	if err != nil {
		return fmt.Errorf("find user in db: %v", err)
	}

	if err = db.SetUserAge(user, age); err != nil {
		return fmt.Errorf("set user age in db: %v", err)
	}

	return nil
}

func main() {
	// handler
	if err := FindAndSetUserAge("bob@example.com", 21); err != nil {
		fmt.Printf("find / set user age: %v", err)
		return
	}

	fmt.Println("successfully updated user's age")
}
