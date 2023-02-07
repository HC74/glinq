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

### *Methods*
#### filter methods
- Instance(slice) Initialize the glinq instance.
- Where(selector) Filter the contents of the slice according to the func expression.
- Distinct(selector) Perform deduplication operations based on the return value of the expression.
- FirstWhere(selector) Get the first element that satisfies the condition according to the condition of the expression.

### Aggregation methods

- Max(selector) Returns the minimum element of this slice according to the provided Comparator.
- Min(selector) Returns the maximum element of this slice according to the provided Comparator.

### Ohters methods

- Any() Determines whether a slice contains any elements.
- AnyN(selector) Determine whether there is an element in the slice that meets the condition according to the expression. 
- ToList() return the final result set.
- FlatMapI(selector) Convert to map type key as int,The key is returned by the selector expression.
- FlatMapS(selector) Same as the FlatMapI method, but this time the key is a string.
- ForEach(selector) iterate over the current slice.

---

### *Example*
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
