package user

import (
	"fmt"
	"sort"
	"time"
)

type (
	User struct {
		ID         string    `json:"id"`
		Name       string    `json:"name"`
		Data       []int     `json:"data"`
		LastUpdate time.Time `json:"lastUpdate"`
	}

	UserList []User
)

func (l UserList) Len() int {
	return len(l)
}

func (l UserList) Less(i, j int) bool {
	return l[i].LastUpdate.After(l[j].LastUpdate)
}

func (l UserList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

var (
	users     = make([]User, 0)
	counterID int
)

func GetAllUsers() []User {
	sort.Sort(UserList(users))
	return users
}

func FindUserByID(id string) *User {
	for _, v := range users {
		if v.ID == id {
			return &User{
				ID:   id,
				Name: v.Name,
				Data: v.Data,
			}
		}
	}
	return nil
}

func CheckUsernameExist(name string) bool {
	for _, v := range users {
		if v.Name == name {
			return true
		}
	}
	return false
}

func CreateUser(name string) {
	counterID++
	users = append(users, User{
		ID:         fmt.Sprintf("%06d", counterID),
		Name:       name,
		Data:       []int{},
		LastUpdate: time.Now(),
	})
}

func UpdateUserData(id string, data int) bool {
	for i := range users {
		if users[i].ID == id {
			users[i].Data = append(users[i].Data, data)
			users[i].LastUpdate = time.Now()
			return true
		}
	}
	return false
}

func DeleteUser(id string) {
	index := -1
	for i, v := range users {
		if v.ID == id {
			index = i
			break
		}
	}
	if index != -1 {
		users[index] = users[len(users)-1]
		users[len(users)-1] = User{}
		users = users[:len(users)-1]
	}
}
