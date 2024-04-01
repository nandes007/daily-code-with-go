package main

import "fmt"

func main() {
	user1 := &user{
		name: "John",
		age:  30,
	}
	user2 := &user{
		name: "Doe",
		age:  20,
	}
	userCollection := &userCollection{
		users: []*user{user1, user2},
	}
	iterator := userCollection.createIterator()
	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
