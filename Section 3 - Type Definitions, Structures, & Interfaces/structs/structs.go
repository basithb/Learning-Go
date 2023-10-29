package structs

import "fmt"

// structs

type User struct {
	id   int
	name string
}

func TestStructs() {
	var user1 User
	user1 = User{id: 1, name: "Basith"}

	user2 := User{id: 2, name: "Ash"}

	fmt.Printf("User 1 information: %v\n", user1)
	fmt.Printf("User 2 information: %v", user2)
}
