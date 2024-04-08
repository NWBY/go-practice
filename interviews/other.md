# Other common topics

### What is the init function?

The "init" function is a special function in Go that is used to initialize global variables or perform any other setup tasks needed by a package before it is used. The init function is called automatically when the package is first initialized. Its execution order within a package is not guaranteed.

Multiple init functions can be defined within a single package and even within a single file. This allows for modular and flexible initialization of package-level resources. Overall, the init function plays a crucial role in ensuring that packages are correctly initialized and ready to use when they are called.

```go
func init() {
  fmt.Println("This will get called on main initialization")
}

func main() {
  fmt.Println("My Wonderful Go Program")
}
```

Notice in this above example, we’ve not explicitly called the init() function anywhere within our program. Go handles the execution for us implicitly and thus the above program should provide output that looks like this:
```go
$ go run test.go
This will get called on main initialization
My Wonderful Go Program
Awesome, so with this working, we can start to do cool things such as variable initialization.
```

As we can see the init function get automatically called and only gets called once

### What is the use of defer?

The defer statement in Golang is used to postpone the execution of a function until the surrounding function completes. It is often used when you want to make sure some cleanup tasks are performed before exiting a function, regardless of errors or other conditions.

A good example is when reading and writing files, we don't want to close the file until we're done either reading or writing to it. If we don't defer we'd need to keep opening the file.

### What are closures?

Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it.

```go
package main

import "fmt"

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
func main() {
    nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())
}
```

### Type assertions

Type assertions in Go are done by using the `.(<type)` on a variable which is an interface type and then using the return values which are `<value>, ok` in an if statement.

For example:
```go
var val interface{} = "Sam"

str, ok := val.(string)

if ok {
    fmt.Println("It's a string")
} else {
    fmt.Println("It's not a string")
}

