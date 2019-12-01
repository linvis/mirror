package main

import "mirror/db"

import "fmt"

func main() {
	user, err := db.QueryUserByName("admin")
	fmt.Println(user, err)
}
