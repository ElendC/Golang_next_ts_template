# Pointers

```go
x := 10
p := &x // creating poitner to x's memory address

Println(p) // Prints Memory address
Println(*p) // Prints 10

*p = 20         // Modify x through the pointer
Println(x)  // Now x is 20
```

Use pointers to avoid copying, but modify the original


```go
type Person struct {
    Name string
    Age  int
}

// Function to modify Person's Age
func updateAge(p *Person, newAge int) {
    p.Age = newAge // Modify the original Person's Age
}

func main() {
    person := Person{Name: "Alice", Age: 25}

    // Pass a pointer to the struct
    updateAge(&person, 30)

    fmt.Println(person) // Output: {Alice 30}
}
```


Without pointers: Each "new" struct is a separate instance in memory.

With pointers: You work with the same instance, so changes reflect across all references to it. This is more efficient when dealing with large structs or when you need to modify the original.