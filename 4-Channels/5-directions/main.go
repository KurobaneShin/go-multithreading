package main

func main() {
	helloChannel := make(chan string)
	go setter(helloChannel, "Hello")

	msg := getter(helloChannel)
	println(msg)
}

func setter(channel chan<- string, name string) {
	channel <- name
}

func getter(channel <-chan string) string {
	return <-channel
}
