# Arrays
1. Represent a fixed length and a type 
```go
var a [4]int // Empty array of length 4, containings integers

a[0] = 1 // Adding 1 to index 0
i := a[0] // Selecting index 0 from the array

// Array do not need to be initialized
//a[2] == 0
```

### Ways to initialize
```go
b := [2]string{"Penn", "Teller"}

b := [...]string{"Penn", "Teller"} //Compiler counts the array elements
```

# Slices

```go
// Same as array but leave out the count
letters := []string{"a", "b", "c", "d"}

// or

func make([]T, len, cap) []T

```