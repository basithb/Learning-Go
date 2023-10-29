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

	user2 := User{2, "Ash"}

	fmt.Printf("User 1 information: %v\n", user1)
	fmt.Printf("User 2 information: %v", user2)
}

// structs with methods

func (u User) PrettyPrint() string {
	userID := fmt.Sprintf("%d", u.id)
	userName := u.name
	combinedInfo := string(userID) + " " + userName

	return combinedInfo

}

func TestStructsWithMethods() {
	user3 := User{id: 3, name: "Primeagen"}
	formattedMessage := user3.PrettyPrint()

	fmt.Printf("Formatted user infomation: %v", formattedMessage)
}
