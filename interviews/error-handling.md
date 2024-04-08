# Error handling in Go

Error handling in Go works via an error type, meaning that a function can return multiple values and one of those can be an error type. This means that you will need to check if the return value for the error is nil, if it's not nil then there is an error and you should handle it.

Here is an example:
```go
func CarName(car string) (string, error) {
    if car == "Tesla" {
        return "", errors.New("Only ICE cars are allowed")
    }
    
    return car
}

car, err := CarName("Lambo")
if err != nil {
    log.Fatal(err)
}
```