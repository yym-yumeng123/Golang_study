package main

func main() {
	ch := make(chan int, 10)
	close(ch)
	close(ch)
}
