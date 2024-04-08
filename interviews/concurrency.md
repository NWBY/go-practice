# Concurrency

### How do you implement concurrency in Go?

In Go, concurrency is implemented using goroutines and channels. Goroutines are lightweight threads that can be created with the go keyword. They allow concurrent execution of functions.

Channels, on the other hand, are used for communication and synchronization between goroutines. They can be created using make and can be used to send and receive values.

To start a goroutine, simply prefix a function call with the go keyword. This will create a new goroutine that executes the function concurrently. Channels can be used to share data between goroutines, allowing them to communicate and synchronize.

By using goroutines and channels, Go provides a simple and efficient way to implement concurrency in programs.

### Difference between buffered and unbuffered channels in Go?

Unbuffered Channel has no capacity initially, but Buffered Channel has a capacity.

Unbuffered Channel will block the goroutine whenever it is empty and waiting to be filled. But Buffered Channel will also block the goroutine either when it is empty and waiting to be filled or it's on its full-capacity and there's a statement that want to fill the channel.
