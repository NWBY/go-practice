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