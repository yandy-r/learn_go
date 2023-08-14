# Golang Notes

## Variables

- Variables are declared using the `var` keyword
- Variables can also be declared using the `:=` operator
- Multiple variables can be declared in a single line using the `var` keyword
- The shorthand operator `:=` can only be used inside of functions

```go
func main() {
  var name string = "John Doe"
  var age int = 30
  var isCool = true
  isCool = false
  // Shorthand
  name := "John Doe"

  // Multiple variables using the shorthand operator
  size := 1.3
  name, email := "John Doe", "
}
```

## Chapter 3: Composite Types

### Arrays

**Arrays** are rarely used directly in `Go`. All elements in the array must be of the same type.

```go
// Declare an array of size 3 and type int

// initialises all elements to zero (0)
var x [3]int

// Declare an array of size 3 and type with initial values
var y [3]int = [3]int{1, 2, 3}

// if specifying the values, the size number can be omitted
var z [...]int{3, 5, 9, 34, 23}
```

You cannot read or write past the end of an array, doing so is a compilie time error. You cant' use type conversion to
convert arrays of different sizes and identical types.

Best practices is not to use `Arrays` directly and instead use `Slices`.

### Slices

Slices are used most of the time when a data structure to hold a sequennce of values is required. Slices can grow as
needed when the capacity is reached and resizing is required.

- Slices are assigned consecutive memory locations/addresses which makes them quick to read and write to.
- When declaring a slice, the size is not specified.
- Just like arrays all elements in a slice must be of the same type.
- Just like arrays you can't read or write past the end of a slice or use negative indexes.

```go
// Declares a slice of type float64
var x = []float64{1, 2, 3, 4, 5}
```

- A slice is not comparable.
- It is a compile time error to use the `==` operator to see if two slices are equal.
- The only thin a slice can be compared with is `nil`

#### len

The `len` builtin function returns the `length` of a slice. This builtin also works on `strings` and other types. When a
`nil` slice is passed to the `len` builtin, it returns `0`.

```go
a := []int{1, 2, 3, 4, 5}

fmt.Println(len(a)) // -> 5
```

#### append

- The builtin `append` function is used to grow slilces:
- The `append` function takes in at least two arguments, the first is the slice to append to and the rest are the values
  to append.
- Because it's pass by value, returns a copy of the passed in slice and must be assigned to a variable.
- It is a compoile-time error if you forget to assign the value returned from `append`.

```go
var a []int

a = append(a, 10)

// To append one slice to anther, use the `...` operator
b := []int{20, 40, 45, 73}
a = append(a, b...)
```

#### Capacity

- Every slice has a `capacity`, which is the number of consecutive memory locations reserved.
- The capacity can be longer than the length of the slice.
- When values are added and the length increases past the `capcity`, the `go` runtime creates a new slice with a larger
  capicty. The values from the original slice are copied tot he new slice and the new values are added to the end. The
  new `slice` is then returned.

NOTE: The `go` runtime is compiled into every program to provide memory management and other features. No additional
libraries or virtual machines such as the `Java JVM` are reuired to run a `go` binary.

- Becuase it takes time to allocate new membory everytime a `slice` grows, the `go` runtime takes the following steps.
- Double the `capacity` when it is less tahn **1024**.
- then grow it by at least **25%** afterward.
- The builtin `cap` function returns the current capacity of a slice.

```go
var x []int
fmt.Println(x, len(x), cap(x)) // -> [] 0 0
x = append(x, 10)
fmt.Println(x, len(x), cap(x)) // -> [10] 1 1
x = append(x, 20)
fmt.Println(x, len(x), cap(x)) // -> [10 20] 2 2
x = append(x, 30)
fmt.Println(x, len(x), cap(x)) // -> [10 20 30] 3 4
x = append(x, 40)
fmt.Println(x, len(x), cap(x)) // -> [10 20 30 40] 4 4
x = append(x, 50)
fmt.Println(x, len(x), cap(x)) // -> [10 20 30 40 50] 5 8
```

#### make

- The builtin `make` function is used to create slices, maps and channels.
- The make function allows you to specify the type, length and, optionally, the capacity of the slice.

```go
// careate an int slice with a length of 5
x := make([]int, 5)
// also has a capacity of 5 since the value was obmitted.
// all elements are initialised to zero (0)
```

- One commono mistake is to try and append to a slice created with `make`.

```go
x := make([]int, 5)
x = append(x, 10)
```

- Because `make` initializes all elements to `0` and has both `length` and `capacity` the `10` is appended to the end of
  the slice and the `length` is doubled.
- Capacity cna also be specified while using the make function.

```go
x := make([]int, 5, 10)
```

- The above code creates a slice of length `5` and capacity `10`.
- Specifying a `capcity` that's less than the `length` is a compile-time error.

#### Slicing Slices

- Slicing a slice consists of a starting offset and ending offset separated by a colon `:`.
- If the starting offset is left off, it defaults to `0`.
- If the ending offset is left off, it defaults to the length of the slice.

```go
x := []int{1, 2, 3, 4, 5}
y := x[:2] // -> [1, 2]
z := x[1:] // -> [2, 3, 4, 5]
d := x[1:3] // -> [2, 3]
e := x[:] // -> [1, 2, 3, 4, 5]

fmt.Println("x: ", x)
fmt.Println("y: ", y)
fmt.Println("z: ", z)
fmt.Println("d: ", d)
fmt.Println("e: ", e)
```

##### Sharing Storage

When you take a slice from a slice, the new slice shares the same underlying storage as the original slice. This means
that you now have two copies of the same data.

```go
x := []int{1, 2, 3, 4, 5}
y := x[:2] // -> [1, 2]
z := x[1:] // -> [2, 3, 4, 5]

x[1] = 20
y[0] = 10
z[1] = 30

fmt.Println("x: ", x) // -> [10, 20, 30, 4, 5]
fmt.Println("y: ", y) // -> [10, 20]
fmt.Println("z: ", z) // -> [20, 30, 4]
```

#### copy

- To make an independent copy of a slice, use the builtin `copy` function.
- The `copy` function takes in two arguments, the first is destination slice and the second is the source slice.
- The `copy` fucntion also returns the number of elements copied.

```go
x := int{1, 2, 3, 4, 5}
y := make([]int, 3) // -> [0, 0, 0]
num := copy(y, x) // -> y: = [1, 2, 3] and num: = 3
```

### Strings, Runes and Bytes

- Under the covers, `go` uses a sequence of `bytes` to represent a string.
- Some libraries and the `for-range` loop assume the `string` is composed of `UTF-8` encoded points.
- > According to the language specification, `Go` source code is always written in UTF-8. Unless you use hexadecimal
  > escapes in a string literal, your string literals are written in UTF-8.
