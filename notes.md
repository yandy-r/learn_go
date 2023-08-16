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
- Most data in `Go` is read and written as a sequence of `bytes`

### Maps

- A `map` is a collection of key-value pairs.
- Teh map type is written as `map[keyType]valueType`

```go
var nilMap map[string]int // -> nil
emptyMap := map[string]int{} // -> empty not the same as nil, this has a length of 0

// create a map with initial values
// The keys are string and the values are slices of strings
teams := map[string][]string {
    "Orcas": []string{"Fred", "Ralph", "Bijou"},
    "Lions": []string{"Sarah", "Peter", "Billie"},
    "Kittens": []string{"Waldo", "Raul", "Ze"},
}

// If you know how many key-value pairs you intend to put in the map, but don’t know the exact values,
// you can use make to create a map with a default size:
ages := make(map[int][]string, 10)
```

- The key for a map can be any comparable type. This means you cannot use a slice or a map as the key for a map.
- > Use a map when the order of elements doesn’t matter. Use a slice when the order of elements is important.

### Structs

When you have related data that you want to group together, you should define a struct.

> If you already know an object-oriented language, you might be wondering about the difference between classes and
> structs. The difference is simple: Go doesn’t have classes, because it doesn’t have inheritance. This doesn’t mean
> that Go doesn’t have some of the features of object-oriented languages, it just does things a little differently.

The syntax for defining a `struct` is as follows:

```go
type person struct {
 name string
 age int
 pet string
}

// Using the struct type
julia := person{
 "Julia",
 40,
 "dog",
}
```

## Chapter 7

### Pointers

Pointers are used to share data across different functions and methods. Pointers allow you to share data without having
to copy it. The `zero` value for a pointer is `nil`.

Go’s pointer syntax is partially borrowed from C and C++. Since Go has a garbage collector, most of the pain of memory
management is removed. Furthermore, some of the tricks that you can do with pointers in C and C++, including pointer
arithmetic, are not allowed in Go.

The & is the address operator. It precedes a value type and returns the address of the memory location where the value
is stored:

```go
x := "hello"
pointerToX := &x
```

The \* is the indirection operator. It precedes a variable of pointer type and returns the pointed-to value. This is
called dereferencing:

```go
x := 10
pointerToX := &x
fmt.Println(pointerToX) // prints a memory address (example 0xc00001a0e8)
fmt.Println(*pointerToX) // prints 10

z := 5 + *pointerToX
fmt.Println(z) // prints 15
```

Before dereferencing a pointer, you must make sure that the pointer is non-nil. Your program will panic if you attempt
to dereference a nil pointer:

```go
var x *int

fmt.Println(x == nil) // prints true
fmt.Println(*x) // panics
```

Not being able to take the address of a constant is sometimes inconvenient. If you have a struct with a field of a
pointer to a primitive type, you can’t assign a literal directly to the field:

```go
type person struct {
  FristName string
  MiddleName *string
  LastName string
}

p := person{
  FristName: "John",
  MiddleName: "Perry", // This line won't compile
  LastName: "Doe",
}
```

One way around this problem is to write a helper funtion that takes in a boolean, numeric, or string type and returns a
pointer to that type.

```go
func stringp(s string) *string {
  return &s
}

p := person{
  FristName : "John",
  MiddleName: stringp("Perry"),
  LastName: "Doe",
}
```

Why does this work? When we pass a constant to a function, the constant is copied to a parameter, which is a variable.
Since it’s a variable, it has an address in memory. The function then returns the variable’s memory address.

#### Pointers Indicate Mutable Parameters

The lack of immutable declarations in Go might seem problematic, but the ability to choose between value and pointer
parameter types addresses the issue. As the Software Construction course materials go on to explain: “[U]sing mutable
objects is just fine if you are using them entirely locally within a method, and with only one reference to the object.”
Rather than declare that some variables and parameters are immutable, Go developers use pointers to indicate that a
parameter is mutable.

If a pointer is passed to a function, the function gets a copy of the pointer. This still points to the original data,
which means that the original data can be modified by the called function.

There are a couple of related implications of this.

##### 1st

The first implication is that when you pass a nil pointer to a function, you cannot make the value non-nil. You can only
reassign the value if there was a value already assigned to the pointer. While confusing at first, it makes sense. Since
the memory location was passed to the function via call-by-value, we can’t change the memory address, any more than we
could change the value of an int parameter. We can demonstrate this with the following program:

```go
func failedUpdate(g *int) {
  x := 10
  g = &x
}

func main() {
  var f *int // f is nil

  failedUpdate(f)
  fmt.Println(f) // still nil
}
```

We start with a nil variable f in main. When we call failedUpdate, we copy the value of f, which is nil, into the
parameter named g. This means that g is also set to nil. We then declare a new variable x within failedUpdate with the
value 10. Next, we change g in failedUpdate to point to x. This does not change the f in main, and when we exit
failedUpdate and return to main, f is still nil.

##### 2nd

if you want the value assigned to a pointer parameter to still be there when you exit the function, you must dereference
the pointer and set the value. If you change the pointer, you have changed the copy, not the original. Dereferencing
puts the new value in the memory location pointed to by both the original and the copy. Here’s a short program that
shows how this works:

```go
func failedUpdate(px *int) {
    x2 := 20
    px = &x2
}

func update(px *int) {
    *px = 20
}

func main() {
    x := 10
    failedUpdate(&x)
    fmt.Println(x) // prints 10
    update(&x)
    fmt.Println(x) // prints 20
}
```

In this example, we start with x in main set to 10. When we call failedUpdate, we copy the address of x into the
parameter px. Next, we declare x2 in failedUpdate, set to 20. We then point px in failedUpdate to the address of x2.
When we return to main, the value of x is unchanged. When we call update, we copy the address of x into px again.
However, this time we change the value of what px in update points to, the variable x in main. When we return to main, x
has been changed.

#### Pointers Are a Last Resort

When returning values from a function, you should favor value types. Only use a pointer type as a return type if there
is state within the data type that needs to be modified. When we look at I/O in “io and Friends”, we’ll see that with
buffers for reading or writing data. In addition, there are data types that are used with concurrency that must always
be passed as pointers.

#### The Zero Value Versue No Value

When not working with JSON (or other external protocols), resist the temptation to use a pointer field to indicate no
value. While a pointer does provide a handy way to indicate no value, if you are not going to modify the value, you
should use a value type instead, paired with a boolean.

#### Difference Between Maps and Slices

Avoid using maps for input parameters or return values, especially on public APIs. On an API-design level, maps are a
bad choice because they say nothing about what values are contained within; there’s nothing that explicitly defines what
keys are in the map, so the only way to know what they are is to trace through the code.

Passing a slice to a function has more complicated behavior: any modification to the contents of the slice is reflected
in the original variable, but using append to change the length isn’t reflected in the original variable, even if the
slice has a capacity greater than its length. That’s because a slice is implemented as a struct with three fields: an
int field for length, an int field for capacity, and a pointer to a block of memory.

#### Slices as Buffers

When reading data from an external resource (like a file or a network connection), many languages use code like this:

```ts
r = open_resource()
while r.has_data() {
  data-chunk = r.next_chunk()
  process(data-chunk)
}

close(r)
```

The problem with this pattern is that every time we iterate through that while loop, we allocate another data_chunk even
though each one is only used once. This creates lots of unnecessary memory allocations. Garbage-collected languages
handle those allocations for you automatically, but the work still needs to be done to clean them up when you are done
processing.

In `Go`; Rather than returning a new allocation each time we read from a data source, we create a slice of bytes once
and use it as a buffer to read data from the data source:

```go
file, err := os.Open(fileName)
if err != nil {
  return err
}
defer file.Close()

data := make([]byte, 100)
for {
  count, err := file.Read(data)
  if err != nil {
    return err
  }
  if count == 0 {
    return nil
  }
  process(data[:count])
}
```

We can’t change the length or capacity of a slice when we pass it to a function, but we can change the contents up to
the current length. In this code, we create a buffer of 100 bytes and each time through the loop, we copy the next block
of bytes (up to 100) into the slice. We then pass the populated portion of the buffer to process.
