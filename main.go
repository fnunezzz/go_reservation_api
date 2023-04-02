package main

func main() {
	storage, err := newPostgresConn()

	if err != nil {
		panic(err)
	}
	server := NewServer(":3000", storage)
	server.Run()
}