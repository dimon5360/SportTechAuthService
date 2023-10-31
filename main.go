package main

import (
	"fmt"
	"main/storage/storage"
	"main/storage/utils"
)

func main() {

	env := utils.Init()
	env.Load("env/app.env")
	env.Load("env/db.env")

	fmt.Println("Storage service v." + env.Value("VERSION_APP"))

	storage := storage.CreateStorage()
	storage.Init(env.Value("CONNECTION_STRING") + "?sslmode=disable")

	user := storage.GetUserById("1")
	fmt.Println(user)
}
