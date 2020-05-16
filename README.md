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

## Contants

Constants must be initialized during declaration. Unlike JavaScript, constants cannot be assigned during `runtime`, they must be initialized at `compile time`.

```go
const pi = 3.1415
```

```go
const x = 18

fmt.Println(x + 3) // 21
fmt.Println(x + 1.2) // 19.2
fmt.Println(float32(x) + 1.2) // 19.2
```

Same as defining a constant outside main

```go
const x = 18

func main() {
    fmt.Println(x + 3)            // 6
    fmt.Println(x + 1.2)          // 4.2
    fmt.Println(float32(x) + 1.2) // 4.2
}
```

### Multiple constants definition

```go
const (
    firstName = "John"
    age       = 45
)

func main() {
    fmt.Println(firstName, age)
}

```

> Result

```sh
John 45
```

### IOTA

```go
const (
    i = iota
    j = iota
)

func main() {
    fmt.Println(i, j)
}
```

> Result

```sh
0 1
```

### IOTA Operations

```go
const (
    i = 2 << iota
    j = iota + 6
    k
)

const (
    l = iota
)

func main() {
    fmt.Println(i, j, k, l)
}
```

> Result

```sh
2 7 8 0
```

---

## Collections

### Array

```go
var arr [3]int
arr[0] = 1
arr[1] = 2
arr[2] = 3

fmt.Println(arr)
fmt.Println(arr[1])
```

> Result

```sh
[1 2 3]
2
```

> Euqivalent to

```go
arr := [3]int{1, 2, 3}

fmt.Println(arr)
fmt.Println(arr[1])
```

### Slice

`Slice` from `Array`

```go
arr := [3]int{1, 2, 3}
slice := arr[:]
fmt.Println(slice)
```

> Result

```sh
[1, 2, 3]
```

> Slice maintains a copy of Array it was derived from. Any changes, reflect in both

```go
arr := [3]int{1, 2, 3}
slice := arr[:]

arr[2] = 42
slice[2] = 18

fmt.Println(array)
fmt.Println(slice)
```

> Result

```sh
[1, 2, 18] [1, 2, 18]
```

#### Append

```go
slice := []int{1, 2, 3}
fmt.Println(slice)

slice = append(slice, 4)
fmt.Println(slice)
```

> Result

```sh
[1, 2, 3]
[1, 2, 3, 4]
```

#### Sub sets

```go
slice := []int{1, 2, 3, 4, 5, 6}
fmt.Println(slice)
fmt.Println(slice[:])
fmt.Println(slice[1:])
fmt.Println(slice[:4])
fmt.Println(slice[2:3])
```

> Result

```sh
[1 2 3 4 5 6]
[1 2 3 4 5 6]
[2 3 4 5 6]
[1 2 3 4]
[3]
```

---

## Map

Map is a dictionary of key-value pair

```go
m := map[string]int{"foo": 18}
fmt.Println(m)
fmt.Println(m["foo"])

m["foo"] = 42
fmt.Println(m["foo"])

delete(m, "foo")
fmt.Println(m)
```

> Result

```sh
map[foo:18]
18
42
map[]
0
```

---
