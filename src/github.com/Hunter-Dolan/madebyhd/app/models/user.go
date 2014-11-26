package models

import "../../config/globals"

type User struct {
	Id   int64
	Name string
}

func FindUserByName(name string) User {
	user := User{}

	global.DB.Connection.Where("name = ?", name).First(&user)
	return user
}
