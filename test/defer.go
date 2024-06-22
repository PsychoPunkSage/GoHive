package main

import "fmt"

type user1 struct {
	admin   bool
	name    string
	contact int
}

const (
	lognotfound = "log not found"
	logadmin    = "log admin"
	logdeleted  = "log deleted"
)

func logAndDelete(users map[string]user1, name string) (log string) {
	defer delete(users, name)
	user, ok := users[name]
	if !ok {
		return lognotfound
	}
	if user.admin {
		return logadmin
	}
	return logdeleted
}

func defer_func() {
	// Create a map of user1
	users := map[string]user1{
		"John": {admin: false, name: "John", contact: 123456789},
		"Jane": {admin: true, name: "Jane", contact: 987654321},
		"Mike": {admin: false, name: "Mike", contact: 456123789},
	}

	name := "Jane"
	logMessage := logAndDelete(users, name)
	fmt.Println(logMessage)
	fmt.Println("Remaining users:", users)

	name = "Janik"
	logMessage = logAndDelete(users, name)

	fmt.Println(logMessage)
	fmt.Println("Remaining users:", users)
}
