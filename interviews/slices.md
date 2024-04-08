# Slice

### Difference between arrays and slices

In Go, an array is a fixed-length sequence of elements of the same type. Once an array is defined, the length cannot be changed. On the other hand, a slice is a dynamically-sized, flexible view of an underlying array. It is created with a variable length and can be resized.

Slices are typically used when you need to work with a portion of an array or when you want to pass a subset of an array to a function. Slices provide more flexibility and are widely used in Go for their convenience and efficiency in managing collections of data.

Initialising an array:
```go
arr := [5]int{1, 2, 3, 4, 5}
```

vs initialising a slice:
```go
slice := []int{1, 2, 3, 4, 5}
```

An empty slice can also be created with `make()`
```go
s := make([]byte, 5)
```