# Concurrency

### How do you implement concurrency in Go?

In Go, concurrency is implemented using goroutines and channels. Goroutines are lightweight threads that can be created with the go keyword. They allow concurrent execution of functions.

Channels, on the other hand, are used for communication and synchronization between goroutines. They can be created using make and can be used to send and receive values.

To start a goroutine, simply prefix a function call with the go keyword. This will create a new goroutine that executes the function concurrently. Channels can be used to share data between goroutines, allowing them to communicate and synchronize.

By using goroutines and channels, Go provides a simple and efficient way to implement concurrency in programs.

### Difference between buffered and unbuffered channels in Go?

Unbuffered Channel has no capacity initially, but Buffered Channel has a capacity.

Unbuffered Channel will block the goroutine whenever it is empty and waiting to be filled. But Buffered Channel will also block the goroutine either when it is empty and waiting to be filled or it's on its full-capacity and there's a statement that want to fill the channel.

### What is a channel?

In Go, a channel is a data structure that allows goroutines (concurrent functions) to communicate and synchronize with each other. It can be thought of as a conduit through which you can pass values between goroutines. A channel has a specific type that indicates the type of values that can be sent and received on it.

Channels can be used to implement synchronization between goroutines and data sharing. They provide a safe and efficient way to coordinate the flow of information, ensuring that goroutines can send and receive data in a controlled and synchronized manner.

### How do you use a channel?

A channel is made using the `make()` function and closed using the `close()` function. It's common to defer the closing of a channel.
```go
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
```

