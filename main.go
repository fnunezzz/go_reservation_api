package main

import "log"

func main() {
	storage, err := newPostgresConn()

	if err != nil {
		panic(err)
	}

	err = storage.Init()

	if err != nil {
		log.Fatalln(err)
	}

	server := NewServer(":3000", storage)
	server.Run()
}