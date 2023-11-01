package main

import (
	"fmt"
	"main/storage/storage"
	"main/storage/utils"
)

/// TODO:
/// 1. write and compile protobuf
/// 2. include grpc
/// 3. receive and process queries and transfer responses
/// 4. connect to kafka

func main() {

	env := utils.Init()
	env.Load("env/app.env")
	env.Load("env/db.env")

	fmt.Println("Auth service v." + env.Value("VERSION_APP"))

	storage := storage.CreateStorage()
	storage.Init(env.Value("CONNECTION_STRING") + "?sslmode=disable")

	user := storage.GetUserById("1")
	fmt.Println(user)
}
