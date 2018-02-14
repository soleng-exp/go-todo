package Models

import "errors"

type User struct {
	Username string
	Password string
}

func GetUserByAuth(user User) (User, error) {

	if user.Username == "jon" && user.Password == "shhh!" {
		return user, nil
	}

	return user, errors.New("User not found")
}
