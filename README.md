# Just Go

Go: Getting Started

---

## Sample code

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello World!")
}
```

---

## Key features

- Fast Compilation
- Fully compiled
- Strongly typed
- Concurrent by default
- Garbage collected
- Simplicity

---

## Ideal for

- **Web applications & services**
- Task automation
- ~~Desktop applications~~
- ~~Machine learning~~

---

## Modules

### Generate go.mod

```sh
go mod init github.com/spicycoder/just-go
```

> `go.mod` will have content similar to

```go
module github.com/spicycoder/just-go

go 1.14
```

---

### Run by module

```sh
go run github.com/spicycoder/just-go
```

---

## Data types

```go
// declaration, initialization
var i int
i = 18
fmt.Println(i)

// declaration + initialization
var f float32 = 3.1415
fmt.Println(f)

// implicit initialization
firstName := "John"
fmt.Println(firstName)

// complex data type
c := complex(3, 4)
fmt.Println(c)

real, imaginary := real(c), imag(c)
println(real, imaginary)
```

---

## Pointers

```go
var firstName *string
fmt.Println(firstName)
```

> Result

```sh
<nil>
```

### De-reference

```go
// pointer
var firstName *string = new(string)

// dereference
*firstName = "John"

fmt.Println(firstName)
fmt.Println(*firstName)
```

> Result

```sh
0xc0000881e0
John
```

### Example

```go
firstName := "John"
firstNamePointer := &firstName

fmt.Println(firstNamePointer)
fmt.Println(*firstNamePointer)
fmt.Println(&firstNamePointer)
```

> Result

```sh
0xc0000461f0
John
0xc000006028
```

---
