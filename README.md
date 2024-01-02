# LEARNING GO

`WHAT ARE WE GONNA LEARN?`

- Essentials of Go Lang
- Practice Project
- Arrays, Slices, Maps
- Structs
- Packages
- Concurrency
- Functions
- Interfaces
- Pointers

## Go Essentials

Values, Basic Types & Core Language Features

- Understanding the Key Components of a Go Program.
- Working with Values & Types.
- Creating & Executing Functions.
- Controlling Execution with Control Structures.

```go
func main() {
    fmt.Print("Hello World")
}
```

- In the above we are calling a so called function `Print` which is
  simply a command we are executing.
- This command here is outputting 'Hello World' text here to the standard output
  which turns out to be CLI.
- It is a built-in command provided by the "fmt" package that we are importing.
- And the value we are passing inside that `Print` command / function we are passing `string` value.
- We can also write string in between using the backticks ``.
- `main()` function gets executed automatically by Go.

```go
import "fmt"
```

- This package is a part of the Go Lang, we don't have to install this package externally.

```go
package main
```

- This package clause is called at the top of the file [mandatory].
- Every Go file needs such a package instruction inside of it.

- When writing Go Code we spilt our code across packages, we must have at least one package per application / per go program.
- But we can also have multiple packages and a single package can split across multiple files.
- We can also have multiple packages in one go project.
- Idea behind this packages to simply is to organize the code.
- We need these packages because once we work with multiple packages, we can use features from `package B` to `pacakge A`.
- We can add external packages or find those at this website [GO Packages](https://pkg.go.dev/).

![Packages Working](/go-basics/images/packages-info.png)

`The Importance of the name 'main' in package`

- This is a special package name which tells Go that this package will be the main entry point of the application we are building.
- And this matters because will not run our program using `go run main.go`.
- Instead we use `go build`
- We get an error while running this command.
- It says `go: cannot find main module, to create a module there, run: go mod init`.

```go
go mod init example.com/first-app
```

`main() function in Go`

- You cannot have two main() functions belonging to the same package.
- It is the starting point of any Go file.

### Working with Variables, Values & Operators

```go
// investmentAmount is of type int
var investmentAmount = 1000
```

- We can declare variables using `var` keyword and assign any value.
- We always have to convert variables that are not compaitable with the variables that we are doing calculations with.

```go
func main() {
  var investmentAmount = 1000
  var expectedReturnRate = 5.5
  var years = 10
  var futureValue = float64(investmentAmount) * math.Pow((1+expectedReturnRate/100), float64(years))
}
```

- Here in the above program we had to change the type of `investmentAmount` i.e `int` to the `float64` because we were using a `float64` value i.e `expectedReturnRate`.
- And we should not convert a `float64` type into an `int` because it will loose some values (means data loss).

### Type Conversion & Explicit Type Assignment

- We can override go's inferred value by adding an explicit `type annotation / type assignment` after the variable name.

```go
var investmentAmount float64 = 1000;
```

- Here we explicitly changed the type of the `investmentAmount` value.

`Every value in Go has a specific type`

- `int`
  - A number `without` decimal places
  - -5, 12, 101, -58, 0, 1
- `float64`
  - A number `with` decimal places
  - -15.9, 0.0, 23.12, 18.1, 122.55
- `string`
  - A text value
  - "Hello"
  - Strings are created with double qoutes ("") or backticks (``).
  - The latter syntax allows you to write across multiple lines.
- `bool`
  - A value that's either `true` or `false`
  - true, false

And if we don't add an explicit type we can omit `var` keyword and write like this, adding a walrus operator ':='.

```go
func main() {
  expectedReturnRate := 5.5
}
```

`Can also declare and define multiple variables in the same line.`

```go
var investmentAmount, expectedReturnRate float64 = 1000, 5.5
```

`Adding constant values in Go`

```go
const inflationRate float64 = 2.5
```

As we know that we cannot reassign the value of constant.

`Fetching input from the user`

```go
fmt.Scan(&investmentAmount)
```

`Improved User Input`

```go
fmt.Print("Enter the value for Investment Amount: ")
fmt.Scan(&investmentAmount)
```

### fmt.Scan() Limitations

- The fmt.Scan() function is a great function for easily fetching & using user input through the command line.

- But this function also has an important limitation: You can't (easily) fetch multi-word input values. Fetching text that consists of more than a single word is tricky with this function.

`Formatting Strings / Output`

```go
fmt.Println("Future Value: ", futureValue)
```

We can also make use of the formatting strings

```go
fmt.Printf("Future value: %v \nFuture value (adjusted for inflation): %v", futureValue, futureRealValue)
```

`Formatting floats in Strings`

```go
fmt.Printf("Future value: %.0f \nFuture value (adjusted for inflation): %1f\n", futureValue, futureRealValue)
```

`Creating formatted string`

This formatted string would return a string.

```go
formattedFV := fmt.Sprintf("Future value: %.1f\n", futureValue)
formattedRFV := fmt.Sprintf("Future value (adjusted for inflation): %.2f", futureRealValue)
```

`Building Multiline Strings`

Making use of backtick character we can achieve multiline strings
But can't make use of '\n' to break lines it will print '\n' too

```go
fmt.Printf(`Future Value: %.1f
 Future Value (adjusted for inflation): %.1f`, futureValue, futureRealValue)
```

### Understanding Functions
