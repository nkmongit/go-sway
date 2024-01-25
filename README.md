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
- Creating & Executing Functions. -Controlling Execution with Control
  Structures.

```go
func main() {
    fmt.Print("Hello World")
}
```

- In the above we are calling a so called function `Print` which is simply a
  command we are executing.
- This command here is outputting 'Hello World' text here to the standard output
  which turns out to be CLI.
- It is a built-in command provided by the "fmt" package that we are importing.
- And the value we are passing inside that `Print` command / function we are
  passing `string` value.
- We can also write string in between using the backticks ``.
- `main()` function gets executed automatically by Go.

```go
import "fmt"
```

- This package is a part of the Go Lang, we don't have to install this package
  externally.

```go
package main
```

- This package clause is called at the top of the file [mandatory].
- Every Go file needs such a package instruction inside of it.

- When writing Go Code we spilt our code across packages, we must have at least
  one package per application / per go program.
- But we can also have multiple packages and a single package can split across
  multiple files.
- We can also have multiple packages in one go project.
- Idea behind this packages to simply is to organize the code.
- We need these packages because once we work with multiple packages, we can use
  features from `package B` to `pacakge A`.
- We can add external packages or find those at this website
  [GO Packages](https://pkg.go.dev/).

![Packages Working](/go-basics/images/packages-info.png)

`The Importance of the name 'main' in package`

- This is a special package name which tells Go that this package will be the
  main entry point of the application we are building.
- And this matters because will not run our program using `go run main.go`.
- Instead we use `go build`
- We get an error while running this command.
- It says
  `go: cannot find main module, to create a module there, run: go mod init`.

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
- We always have to convert variables that are not compaitable with the
  variables that we are doing calculations with.

```go
func main() {
  var investmentAmount = 1000
  var expectedReturnRate = 5.5
  var years = 10
  var futureValue = float64(investmentAmount) * math.Pow((1+expectedReturnRate/100), float64(years))
}
```

- Here in the above program we had to change the type of `investmentAmount` i.e
  `int` to the `float64` because we were using a `float64` value i.e
  `expectedReturnRate`.
- And we should not convert a `float64` type into an `int` because it will loose
  some values (means data loss).

### Type Conversion & Explicit Type Assignment

- We can override go's inferred value by adding an explicit
  `type annotation / type assignment` after the variable name.

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

And if we don't add an explicit type we can omit `var` keyword and write like
this, adding a walrus operator ':='.

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

- The fmt.Scan() function is a great function for easily fetching & using user
  input through the command line.

- But this function also has an important limitation: You can't (easily) fetch
  multi-word input values. Fetching text that consists of more than a single
  word is tricky with this function.

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

Making use of backtick character we can achieve multiline strings But can't make
use of '\n' to break lines it will print '\n' too

```go
fmt.Printf(`Future Value: %.1f
 Future Value (adjusted for inflation): %.1f`, futureValue, futureRealValue)
```

### Understanding Functions

Functions are used to make the code reusable and achieve DRY principle in the
code. Here we are implementing a function that takes a string as output and
prints it in the console.

```go
// Here we passed a parameter text i.e type of string
func outputText(text string) {
  fmt.Printf(text)
}

outputText("This is a string")
```

`Return Values and Variable Scope`

- In Go we can return multiple values, unlike other programming languages.
- Any variables or constants that are declared in a function are scoped to that
  function and are only available in that function only.

```go
// we have to also tell what is the return type
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
// this function returns two values separated by a comma
fv := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
rfv := fv / math.Pow(1+inflationRate/100, years)
return fv, rfv
}
```

And then when we are calling that function we should be storing the two returned
values.

```go
futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
```

`Alternative Return Value Syntax`

We can define the variable names while assigning the return types and just make
use of return keyword

```go
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64,rfv float64) {
// this function returns two values separated by a comma
fv = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
rfv = fv / math.Pow(1+inflationRate/100, years)
return
}
```

### Control Structures

Control statement consists of `if`, `else if` and `else`, they are used for
changing the flow of the GO program.

```go
var age int = 18

if age >= 18 {
  fmt.Println("You allowed to drive bike")
} else if age >= 24 {
  fmt.Println("You are allowed to drive car")
} else {
  fmt.Println("You are not allowed to drive anything, but cycle")
}
```

### Conditional Statements

Go only supports `for` loop unlike other programming languages that have `do`,
`do while`, but easily mimic those features with just the `for` loop here. And
also have features like `break` and `continue`.

```go
// this will print number 0 to 5
for i := 0; i >= 5; i++ {
  fmt.Print(i)
}
```

```go
// this will run infinitely
// mimicing while loop
for {
  fmt.Print("GO IS FUN")
}
```

```go
for isTrue {
  fmt.Print("GO IS BETTER AND EASY")
}
```

### Switch Statements

### Working with Files

`Writing in a File`

To write in a file we have to use a package from Go standard library called `OS`

```go
import "os"

func writeBalanceToFile(balance float64) {
  // name of the file, data written to a file [collection of bytes], desc of permission
  balanceText := fmt.Sprint(balance)
  // 0644 is a file permission
  os.WriteFile("balance.txt", []byte(balanceText), 0644)
}
```

The above function takes a float64 as the parameter, then stores it in the
variable `balanceText` and then converts into the byte array and writes it into
the file `balance.txt`, we even provide the file permission.

``Reading from File`

```go
func getBalance() {
  data, _ := os.ReadFile("balance.txt")
  balanceText := string(data)
  balance, _ := strconv.ParseFloat(balanceText, float64)
  return balance
}
```

The above function is used to read from a file and returns a byte of array and
err if we get one, we have to store those values inside it, if we don't want to
use `err` value somewhere we can replace it with underscore `_`. Then we would
use strconv which is used to convert the string data to other datatype. And the
strconv would again return two values data and err.

`Handling Errors`

In Go Lang errors don't crash the program, in the above example if we remove
that file balance.txt Go won't stop our program but it will return an empty byte
collection if it fails to find that file.

So in the above example if remove the file balance.txt and run the code, Go will
return an empty byte collection, then that gets converted into string and then
while parsing that empty string gets converted into 0 and it returns the
value 0.

```go
func getBalanceFromFile() float64 {
 data, err := os.ReadFile(accountBalanceFile)
 if err != nil {
  return 1000
 }

 balanceText := string(data)
 balance, _ := strconv.ParseFloat(balanceText, 64)
 return balance
}
```

In the above example we are writing err variable instead of `_` so ReadFile will
return a value nil of EOF, and if it returns nil then that means we have no
error.

`Returning a Custom Error`

```go
const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
 data, err := os.ReadFile(accountBalanceFile)
 if err != nil {
  return 1000, errors.New("Failed to find balance file")
 }

 balanceText := string(data)
 balance, err := strconv.ParseFloat(balanceText, 64)
 if err != nil {
  return 1000, errors.New("Failed to parse stored balance value")
 }
 return balance, nil
}
```

In the above example we are creating a custom error which makes use of the
errors package, and also a predefined parameter called error in the function
definition.

`Time to Panic`

Certain situations can come when we need to stop executing when getting an
error. We can simply add `return` at the end of the `if` statement to return
from our program.

Another way we can exit out of that function we can make use of the `panic()`
function, this function can output an error and stops the execution.

```go
panic("Can't continue, sorry")
```

## Working with Packages

- Splitting Code Across Multiple Files.
- Splitting Files Across Multiple Packages.
- Importing & Using Custom Packages.

### Splitting Code Across Files in The Same Package

In the folder go-control-structure folder we had a file called bank.go, which
belonged to a package called main, and created another file, that belongs to the
same main package, and make use of function provided by communication.go.

### Importing and Exporting from Different Pacakges

To export any variable and function we have to write that function or variable
name with first letter as capital this is the way of telling Go that we want to
export this function or variable.

```go
package fileops

import (
 "errors"
 "fmt"
 "os"
 "strconv"
)

var Something int = 9

func GetFloatFromFile(fileName string) (float64, error) {
 data, err := os.ReadFile(fileName)
 if err != nil {
  return 1000, errors.New("Failed to find file")
 }

 valueText := string(data)
 value, err := strconv.ParseFloat(valueText, 64)
 if err != nil {
  return 1000, errors.New("Failed to parse stored value")
 }
 return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
 valueText := fmt.Sprint(value)
 os.WriteFile(fileName, []byte(valueText), 0644)
}
```

### Third Party Packages

Now we gonna add a package that is not part of Go but an external package that
generate random data for us.

`go get github.com/Pallinder/go-randomdata`

Will have to run the above command in our folder where we have `go.mod` file,
where it store the package information.

If in case we want to install the packages of go will have to run the `go get`
command to install the packages mentioned in the `go.mod` file.

To make use of the installed package we have to import it in our go file.

```go
import {
"github.com/Pallinder/go-randomdata v1.2.0"
}

func main() {
  fmt.Println(randomdata.PhoneNumber())
}
```

Above will generate a random phone number and print it.

## Understanding Pointers

- What are Pointers?
- Why does this feature exist?
- How to work with Pointers?

### WHAT ARE POINTERS?

Variables that store value addresses instead of values.

Lets say in your code you have created a variable `age := 32`, in that case that
value 32 will be stored the computer's memory and that value in memory
automatically gets an address, and this address is required to the computer to
retrieve that value and work with it.

And a pointer then is a variable where you don't store a value but where you
instead use this special ampersand operator `&` to get and store the address of
stored value.

`agePointer := &age`

So now `agePointer` would contain the address as a value.

### WHY POINTERS?

- Working with pointers we can avoid unnecessary value copies.

  - For a very large & complex values, this may take up too much memory space
    unnecessary.
  - By default. Go creates a copy when passing values to function.
  - In Go when we pass a variable to the function so that it can work with it.
  - Go creates a copy of that value of that variable and it gets a new address
    in the memory.
  - Until the function the function execution is done, and the copied value will
    eventually be cleaned up with Garbage Collector.
  - Garbage Collector removes unused values from memory.
  - With pointers, only one value is stored in memory ( and the address is
    passed around)
  - Instead passing a pointer as argument to a function no copy will be created
    in the memory.

```go
age := 32

func isAllowed(&age) {
  if age > 18 {
    fmt.Println("Allowed to drive")
  } else {
    fmt.Println("Not allowed to drive")
  }
}
```

In the above we are doing the same passing the pointer (address) as parameter.

- Use pointers to directly mutate values.
  - Pass a pointer (address) instead of a value to a function.
  - The function can then directly edit the underlying value - no return value
    is required.
  - Can lead to less code (But also to less understandable code or unexpected
    behaviours)

`Writing Code without Pointers`

`A Pointer's Null Value`

All values in Go have a so-called "Null Value" - i.e., the value that's set as a
default if no value is assigned to a variable.

For example, the null value of an int variable `0`. Of a float64, it would be
0.0. Of a string it's `""`. For a pointer, it's nil - special avlue built-into
Go.

`nil` represents the absence of an address value - i.e., a pointer poiinting at
no address / no value in memory.

## Structs & Custom Types

- What are Structs?
- Creating and Using Structs
- Adding Methods to Structs

Structs are used to group multiple values of similar types in a one single
value.

`Defining a User Defined Struct`

```go
import (
    "time"
)
type User struct {
    firstName striing
    lastName string
    birthDate string
    createdAt: time.Time
}
```

### Instantiating Structs & Structs Literal

Making use of the custom structs that we created above.

```go
// creating a variable type of User struct
var appUser User
// then instantiating the struct `User{}` this is also called as 'Composite Literal' | 'Struct Literal'
appUser = User{
    firstName: userFirstName,
    lastName: userLastName,
    birthDate: userBirthDate
}

fmt.Println(appUser.firstName)
```

`Methods in Structs`

```go
func getUserData(promptText string) string {
 fmt.Print(promptText)
 var value string
 fmt.Scan(&value)
 return value
}
```

Special kind of function in structs they are called as constructor function.
It's just an utility function that used for creating a struct, so that we don't
have to. This is just a pattern to follow, not built in for Go.

```go
// constructor function - utility function that creates a struct
func newUser(firstName, lastName, birthDate string) *User {
 return &User{
  firstName: firstName,
  lastName:  lastName,
  birthDate: birthDate,
  createdAt: time.Now(),
 }
}
```

`Using constructor function for validation`

```go
// validations
func newUser(firstName, lastName, birthDate string) *User {
 if firstName == "" || lastName == "" || birthDate == "" {
 return nil, errors.New("First name, last name and birthdate are required")
   }
 return &User{
  firstName: firstName,
  lastName:  lastName,
  birthDate: birthDate,
  createdAt: time.Now(),
 }
}
```

`Struct Embedding`

This means we can build a new struct that builds up on existing struct.

```go
// struct embedding
type Admin struct {
 email    string
 password string
 User
}
```

```go
// constructor function for Admin
func NewAdmin(email, password string) Admin {
 return Admin{
  email:    email,
  password: password,
  User: User{
   firstName: "ADMIN",
   lastName:  "ADMIN",
   birthDate: "___",
   createdAt: time.Now(),
  },
 }
}
```

We can also use this `type` keyword to assign an alias, to other built-in types.

Creating my own type, maybe `string` making use of `type` keyword.

```go
// assigned an alias to the string built in type
type customString string

// adding methods to custom type, because we can't use it on local type
func (text customString) log() {
 fmt.Println(text)
}

func main() {
 var name customString
 name = "Nishant"

 name.log()
}
```

Practice Project - `Todo that can write into file and spit JSON`

In this project we are using the struct type to define the structure of the
Todos.

```go
type Note struct {
 Title     string
 Content   string
 CreatedAt time.Time
}
```

The main obstacle while creating this application, was taking input from the
user, while we could use the `fmt.Scanln()` method but it won't take more than
one word as a string, so we wiil be using the `bufio.NewReader()`.

```go
func getUserInput(promt string) string {
 fmt.Print(promt)
 reader := bufio.NewReader(os.Stdin)

 text, err := reader.ReadString('\n')
 if err != nil {
  return ""
 }

 text = strings.TrimSuffix(text, "\n")
 text = strings.TrimSuffix(text, "\r")

 return text
}
```

We also had to trim out the `newline` and `\r` because the windows include both
`\n` and `\r`.

### Struct Tags

The struct tags are eventually `metadata`, that we can add into our struct
fields. But in here we have to give the metadata the JSON format because we
making use of `json.Marshall()` to convert our struct type to JSON data.

```go
type Note struct {
 Title     string    `json:"title"`
 Content   string    `json:"content"`
 CreatedAt time.Time `json:"created_at"`
}
```

In the above code we are providing the metadata to the struct type, so whenever
the we are converting the `struct` into the JSON format it will search for the
JSON metadata.

## Interfaces

Lets say in this above application we just don't want to allow users to store
notes but also todos. And todo also needs to be `struct` but unlike note it
should not have a title, createdAt instead it should just have content.

```go
// creating an interface
type saver interface {
 Save() error
}
```

Here we have created an interface, which is called as saver, because it is
pattern in the Go Lang that if we there's only one method definition inside the
interface, the interface name should be same as that one method inside the
interface with in the end appended with "er" string.

Now that we have an interface, if any of the struct has the method named
`Save()` will be able to use that function without re-defining the method for
each and every struct.

`The Special "Any Value Allowed" Type`

```go
func printSomething(value interface{}) {
 fmt.Println(value)
}
```

Here we taking the parameter value of `any` type by providing an empty
`interface{}` to it and alternatively we can also write `any` keyword.

```go
func printSomethinAgain(value any) {
 fmt.Println(value)
}
```

Looking the type of the value:

```go
func printSomething(value interface{}) {
 // checking the type of `value`
 switch value.(type) {
 case string:
  fmt.Println("Integer:", value)
 case float64:
  fmt.Println("Float:", value)
 default:
  fmt.Printf("%T:%v\n", value, value)
 }
}
```

Alternative syntax for looking up a value type. No matter if the value is struct
or any other value.

```go
typedValue, ok := value.(int)

if ok {
 fmt.Println("Integer:",typedValue + 1)
}
```

`ok` would return a bool value, if the type is `int` for the value then it would
return `true` else `false`.

But this flexibility comes with a downside where if we are passing down two
values through parameters and we assign type as any and then return them while
doing add using '+'. It will throw an error because meanwhile we can add
strings, numbers, float but not structs.

```go
func add(a, b interface {}) {
 return a + b
}
```

We can do a work around on this

```go
func add(a, b interface {}) {
 aInt, aIsInt := a.(int)
 bInt, bIsInt := b.(int)

 if aIsInt && bIsInt {
  return aInt + bInt
 }
}
```

As we have returned a value as an int type but the value returned can be of any
type and for that we can also return `any` type or `interface{}`.

```go
func add(a, b interface{}) any {
 aInt, aIsInt := a.(int)
 bInt, bIsInt := b.(int)

 if aIsInt && bIsInt {
  return (aInt + bInt)
 }
 return nil
}
```

<!--TODO : WE WILL UNDERSTAND THESE LATER ON -->

## Generics

## Arrays, Slices & Maps

`Storing Data in Collections`

- Understanding Arrays & Array Limitations
- Understanding & Using Slices
- Slices Deep Dive
- Understanding & Using Maps

We can store data in collections using `struct` type.

```go
type TemperatureData struct {
 day1 float64
 day2 float64
 day3 float64
 day4 float64
}
```

But it can be cumbersome what if are trying to store a lot of data that have
same datatype.

Example: `Weather Data` - 25.93, 17.81, 55.3, 42.1

To solve this issue we will be using the array.

An array of "weather data points" (array of float64 values)

```go
{25.93, 17.91, 55.8, 42.1}
```

`Creating an Array`

```go
package main

import "fmt"

func main() {
 fmt.Println("ARRAYS")

 prices := [4]float64{10.99, 9.99, 45.99, 20.0}

 fmt.Println(prices)
}
```

`Creating arrays in another way`

```go
var productNames [4]string
```

And if we print this array of string which hasn't been intialized with any
string values yet. Will get an empty array printed instead of an error.

We can assign values with the `index` position, we can also assign while we
intialize the array:

```go
var productNames [4]string = [4]string{"A Book"}
productNames[2] = "A Carpet"
```

Here the first value is being set as "A Book" and second would be "A Carpet",
there will be some empty slot but that is because the array has the length of
`4` and we have inserted only two values.

`Slices in Array | Slicing in Array`

Lets say we want only [9.99, 45.99] from the array to get this we can make use
of slices.

To achieve this we can do:

```go
featuredPrices := prices[1:3]
```

Here we have used the `index` positions of our array and from the index position
`1` to excluding `3` it will return the array back.

But it won't modify the original array in anyway, it will return a new array.

But if you would change the values in the array that is copied over from
`prices` array to `featuredPrices` arrays, changes occurred in `featuredPrices`
would affect the `prices` too.

```go
prices := [4]float64{10.7, 4.5, 45.00, 20.8}
fmt.Println(prices) // 10.7 4.5 45 20

featuredPrices := prices[1:]

fmt.Println(featuredPrices) // 4.5 45 20

featuredPrices[0] = 3.4 // 3.4 45 20

fmt.Println(prices) // 10.7 3.4 45 20
```

This happens because in Go, the values assigned of the array `prices` into the
another variable is just referencing to the `prices` variable in the memory,
it's not creating another variable for this.

- And this is the reason Go saves a lot of memory.
- Go also saves some `metadata` for the slices that we get, these are the
  `length` and the `capacity`, and in the Go Lang has these features built in,
  like `len()` and `cap()`.

The len() gives us the number of items in the slice or array, meanwhile the
cap() is bit different.

`Creating 'DYNAMIC ARRAYS' using slices`

- Creating a dynamic array is easy in Go, you just don't specify the number of elements you want in array.

```go
prices := []float64{10.6, 8.0, 8.7}
```

Here you can see that we haven't passed how many elements should the array have.
But to add more elements in this dynamic array we can't just put `prices[3] = 12`
this won't work as the array we defined already have `3` elements.

To add more items in this `dynamic` array we have to make use of the `append` method.
Which takes two values one is `slice` and other is the value that you want to add this will add the value in the last index and increase the length of the array.
It also returns a new slice, that we have added to our previous slice.

```go
prices = append(prices, 9.0)
```

This `append` method won't change the original values of the slice but it returns new slice with the value appended to it.
With this new value we could either assign these new values to the `old` slice or give it to the a new variable.
