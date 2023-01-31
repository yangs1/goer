package main

func main() {
	s := NewService("127.0.0.1", 8080)
	s.Start()
}
