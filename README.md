## Features of Go
* Simple language => only 25 keywords
* Statically typed
* Out of the box:
  * Memory management
  * Garbage collection
* Compiled language
* No virtual machine
* Faster than most other languages, slower than C
* Not OO
* Data and behavior scattered (?)
* Build high concurrent applications
* Users of Go: Docker, k8s, Terraform


* Pointers in Go
  * Similar to references in Java
  * Cannot perform arithmetic operations using pointers like C


* Library support
  * Lots of libraries for testing, benchmarking, profiling
  * Specialized libraries for concurrency, race detections


* No classes
* struct, arrays and maps

---

# Day 1

## Tooling support
* Launch GoDoc on port 6060
```go
godoc -http :6060
```


## Rules
* Local variables have to be used
* Local variables need not be initialized
  * Gets initialized to default
* `if` does not require parantheses
* else or else if block MUST start inline with if block
* no pre-increment operator


## Data types
* byte
* int
* int64 (javaLong)
* float32 (float32)
* float64 (javaDouble)
* uint32 (unsignedInt)


* Type inference
  * When declared, type is automatically inferred
```go
var x = 10
fmt.Printf("%T\n", x)
```

* Define your own type
```go
type Currency float64
var dollar Currency = 60.9
fmt.Printf("$%v", dollar)
```

* Constants
```go
const PI float64 = 3.14
println("Value of pi =", PI)
```


## Access
* Capitalized functions are public (exported)
* Non-capitalized are private
* are all const (in caps) public?


## Slicing
* Creates a shallow copy


## Map
* Unordered

---

# Day 2

## Functions
* Functions can return multiple values, eg:
```go
func x(y int) (int, int) {
    return y*y, y*y*y
}
```
* Use `_` to ignore return values
* Named return values
```go
func sum(x int, y int) (sum int) {
	sum = x + y
}
```
* Variadic function
```go
func sum(numbers ...int) (total int) {
    
}
```

## Higher order function
* Functions can be assigned to a variables
* Function can be passed as a pointer to another function
* Function can return a pointer to another function

## Pointers
* No pointer arithmetic
* Pointers are used to pass by reference (Go is pass by value)
* Holds the address to the variable
```go
x := 10
var addr *int := &x  // addr := &x
```
* Use `new()` to initialize pointers
```go
var strAddr *string
fmt.Println(strAddr)  // <nil>
// fmt.Println(*strAddr)  // invalid memory address or nil pointer dereference

// noinspection GoVarAndConstTypeMayBeOmitted
var strAddr2 *string = new(string)  // similar to malloc in C
``` 

## `struct`
```go
type Cat struct {
    name string
    age int
    cuteness int64
}
```
* `struct` is passed by value
* variables need to be capitalized for it to be visible outside the package (file ?)

## `interface`
* Interface is an abstraction over type
* Kind of analogous to generics in Java
