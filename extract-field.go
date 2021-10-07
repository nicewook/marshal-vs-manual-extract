package extract

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
}
type UserEssential struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewUsers(num int) []User {
	var users []User
	for i := 0; i < num; i++ {
		name := fmt.Sprintf("name-%d", i)
		user := User{
			i,
			name,
			i + 10,
			i + 20,
			i + 5,
		}
		users = append(users, user)
	}
	return users
}

func NewUsersPointer(num int) []*User {
	var users []*User
	for i := 0; i < num; i++ {
		name := fmt.Sprintf("name-%d", i)
		user := &User{
			i,
			name,
			i + 10,
			i + 20,
			i + 5,
		}
		users = append(users, user)
	}
	return users
}

func ExtractByMarshalling() {
	var userEssentials []UserEssential
	users := NewUsers(10_000)

	b, _ := json.Marshal(users)
	json.Unmarshal(b, &userEssentials)
}

func ExtractManually() {
	var userEssentials []UserEssential
	users := NewUsers(10_000)

	for _, user := range users {
		userEssential := UserEssential{
			user.ID,
			user.Name,
		}
		userEssentials = append(userEssentials, userEssential)
	}
}

func ExtractManuallyPointer() {
	var userEssentials []UserEssential
	users := NewUsersPointer(10_000)

	for _, user := range users {
		userEssential := UserEssential{
			user.ID,
			user.Name,
		}
		userEssentials = append(userEssentials, userEssential)
	}
}
