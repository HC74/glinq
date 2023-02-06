# Glinq
基于Go语言的切片操作库
## Installation

When used with Go modules, use the following import path:
```go
go get github.com/HC74/glinq@latest
```
## <u>Quick start</u>
The method of chain call

`Instance(slice)` `.Where(selector)` `.Distinct(selector)`

*Example*
1. We need to initialize a set of data for the following cases
    ```go
   import . "github.com/HC74/glinq"
   type Cat struct {
	   color, name string
	   age         int
    }
   // Initialize four cats
   slice := AsSlice(&Cat{color: "red", name: "flower", age: 3}, &Cat{color: "blue", name: "flower2", age: 5},
		&Cat{color: "green", name: "flower3", age: 2}, &Cat{color: "green", name: "flower4", age: 3})
    ```
   Filter out cats whose color is green and whose age is greater than 2
      ```go
      first, _ := Instance(slice).Where(func(c *Cat) bool {
        return c.color == "green" && c.age > 2
      }).First()
      // First method, get the first element
      fmt.Println(first) // &{green flower4 3}
      ```