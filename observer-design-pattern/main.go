package main

func main() {
	shirtItem := newItem("Nike shirt")
	observerFirst := &customer{id: "test@example.com"}
	observerSecond := &customer{id: "test2@example.com"}
	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)
	shirtItem.updateAvailability()
}
