### Basic Go Syntax

1. **Variables and Types:**
   - **Declaration:** In Go, you can declare a variable using the `var` keyword. For example, `var x int` declares a variable `x` of type `int`.
   - **Initialization:** You can also initialize a variable at the time of declaration, like `var y int = 10`.
   - **Short Declaration:** Go also supports a shorter form for declaring and initializing variables: `z := 20`. This infers the type of `z` from the value assigned to it.

2. **Basic Data Types:**
   - **Integers:** `int`, `uint` (unsigned), along with sized versions like `int8`, `int64`.
   - **Floating Point:** `float32`, `float64`.
   - **Boolean:** `bool`, representing true or false.
   - **String:** `string`, a sequence of characters.

3. **Control Structures:**
   - **If-Else:** Similar to other languages but without parentheses. Example:
     ```go
     if x > 0 {
         fmt.Println("x is positive")
     } else {
         fmt.Println("x is non-positive")
     }
     ```
   - **For Loop:** Go's only looping construct. It can be used like a traditional for-loop, while-loop, or infinite loop.
     ```go
     for i := 0; i < 10; i++ {
         fmt.Println(i)
     }
     ```
   - **Switch:** A simpler way to write multiple `if-else` conditions.
     ```go
     switch day {
     case "Monday":
         fmt.Println("Start of the week")
     case "Friday":
         fmt.Println("Weekend is here")
     default:
         fmt.Println("It's a weekday")
     }
     ```

4. **Functions:**
   - Functions are declared with the `func` keyword.
   - Example of a simple function:
     ```go
     func add(x int, y int) int {
         return x + y
     }
     ```
   - Functions can return multiple values as well.

### Exercise: Simple Calculator

To reinforce these concepts, let's write a simple calculator program. This program will:
   - Take two numbers as input.
   - Perform addition, subtraction, multiplication, and division.

Here's a skeleton of the program:

```go
package main

import (
    "fmt"
)

func main() {
    var num1, num2 float64
    fmt.Println("Enter two numbers: ")
    fmt.Scanln(&num1, &num2)

    // Perform arithmetic operations and display the results
    // Add, subtract, multiply, and divide num1 and num2
}
```

Your task is to implement the arithmetic operations inside the `main` function. Try to use functions for each operation. This exercise will help you get comfortable with basic Go syntax and functions.

How are you finding Go so far? Do you need any clarifications on these concepts?