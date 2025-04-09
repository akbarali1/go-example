package main

import "fmt"

type User struct {
	Name string
}

func changeNameByValue(u User) {
	u.Name = "Ali"
}

func changeNameByPointer(u *User) {
	u.Name = "Ali"
}

func main() {
	user := User{Name: "Hasan"}

	changeNameByValue(user)
	fmt.Println("By Value:", user.Name) // Hasan

	changeNameByPointer(&user)
	fmt.Println("By Pointer:", user.Name) // Ali
}
