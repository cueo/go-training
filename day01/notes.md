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