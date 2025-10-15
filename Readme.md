
# GoLang Quick Reference Guide 🐹

## 📚 General Concepts

### Null and Nil
- `null` is referred to as `nil` in Go.
- Data type descriptions in Go come **after** the variable name, unlike some other languages.

### Interface
used when we don't know which datatype will function get, in below code we have  

```
func divide(a, b interface{}) (float64, error) {
    af, okA := a.(float64)
    bf, okB := b.(float64)
    if !okA || !okB {
        return 0, errors.New("invalid types, expected float64")
    }
    if bf == 0 {
        return 0, errors.New("division by zero")
    }
    return af / bf, nil
}
```
af, okA:=a.(float64) => here we are changing the a datatype of a, if the operation is successful then okA will be true else false and af is the actual value which will be used in future   

If you wnat to use 'a' only in code then we have 2 options:  
1> overwrite the a with af   
2> write direct: ` a = a.(float64)` but the catch is if the operation fails then whole code would panic and goes down  
### No Inheritance
- Go does **not** support inheritance, nor does it have `super` or `parent` keywords.
- Way around:
  ```
        type Animal struct {
            Name string
        }
        
        func (a Animal) Speak() {
            fmt.Println("Animal speaking...")
        }
        
        type Dog struct {
            Animal // embeds Animal; its methods/fields are *promoted*
        }
        
        func (d Dog) Speak() { // overrides the promoted Speak
            fmt.Println("Dog barking...")
        }
        
        func main() {
            d := Dog{Animal: Animal{Name: "Bruno"}}
        
            d.Speak()         // Dog barking...  (Dog’s method wins)
            d.Animal.Speak()  // Animal speaking... (“super” style call)
        
            // Promoted field access:
            fmt.Println(d.Name) // "Bruno" (from embedded Animal)
        }

  ```
---

## 📄 File Handling in Go

- You can have multiple files in the same package using the same `main` function. 
- If a conflict occurs, adjust by changing one main function and calling it within another.

---

## 🖨️ Print Functions

- **`fmt.Printf`** requires format specifiers such as `%T` (type) or `%v` (value) to display variables.
- **`fmt.Println`** displays variables directly without requiring format specifiers.

---

## 🔤 Data Types and Strings

- **All input data** in Go is treated as `string`.
- Handling large numbers (64-bit and above) often requires packages like `big.Int` or `big.NewInt`.

---

## 🕒 Date & Time Formatting

```go
time.Now().Format("01-02-2006 15:04:05 Monday")
```

---

## 🛑 Error Handling

- Use the `fmt` package instead of `print` and `panic` for better formatting and error management.

---

## 🧩 Working with Arrays and Slices

- `fruits2[:3]` retrieves the **first 3 elements** of an array.
- Use `make` for creating slices, maps, and channels in Go.  
- Arrays:
  ```
      c := [...]string{"go", "is", "fun"}
      var grid [2][3]int
  ```
- maps:
  ```
      m := make(map[string]int)
      // with capacity
      m := make(map[string]int, 10)
  ```
- slices(vectors i.e. dynaimc size):
  ```
      s2 := make([]int, 0, 5)      // len=0 cap=5
      // Full slice expressions (control capacity)
      s5 := arr[1:3:3]             // low:high:max -> len=2 cap=2
  ```
- channels:
  ```
      // Unbuffered channel (sends/receives block until the other side is ready)
      ch := make(chan int)

      // Buffered channel (capacity N; send blocks only when buffer full)
      ch2 := make(chan string, 3)
  ```
  

---

## 📦 Maps and Structs

- **Map Declaration Example**:
  ```go
  var myOnlineData map[string]interface{}
  ```
  Here, `interface{}` allows mapping of any data type.

- **Difference in `println` with Structs**:
  - `%v` shows the struct value.
  - `%+v` shows struct fields **and** values.

---

## 🔀 Control Flow

- Go supports the `goto` command to jump between code blocks.
- The `defer` keyword stores the deferred code, executing it in reverse order.

---

## 🚀 Goroutines and Concurrency

- **Goroutines**: Lightweight threads managed by Go runtime for concurrent execution.
- **Example of Starting a Goroutine**:
  ```go
  go functionName()
  ```
- **WaitGroups (`sync.WaitGroup`)**:
  - Use `wg.Add()` to increase thread count.
  - Use `wg.Done()` to decrease it.

---

## 🔗 Channels for Communication

- **Basic Syntax**:
  - Sending data: `ch <- value`
  - Receiving data: `value := <-ch`
- Channels operate like stacks; data is removed after being retrieved.
- Closing a channel disallows further data pushes.
- **Control Concurrency**
  - These are also used to control number of concurrency taking place i.e. number of goroutines working at a time
  - EXAMPLE: 
    - Number of Goroutine is controlled by fixing the size of channel
    - Each goroutine pushes in stack before starting but if the channel is full
    - Then that goroutine has to wait befoer channel is freed; else if the channel has space
    - Then goroutine don't have to wait and can start immediately after pushing into channel

---

## 📦 Modules and Packages

- **Local Setup**:
  - `go mod init {folder_name}`: Makes the folder's contents available locally.
- **Remote Setup**:
  - `go mod init github.com/username/{folder_name}`: Publishes to GitHub for broader access and Makes the folder's contents available locally.

---

## ✍️ Syntax Quick Reference

### Variables

```go
var {name} int  // Declaration
{name} := value // Short-hand assignment
```

### Maps

```go
{name} := make(map[int]int)
```

### Structs

```go
type User struct{}
```

---
### Packages 

All files in the same folder belong to the same package. You do not need to import files within the same package explicitly.
Example:

```go
project/
├── main.go
├── helpers/
│   ├── math.go
│   ├── string.go
```

- helpers/math.go:
  ```go
  package helpers

  func Add(a, b int) int {
      return a + b
  }
  ```

- helpers/string.go:
  ```go
  package helpers

  func ToUpperCase(s string) string {
      return strings.ToUpper(s)
  }
  ```
- main.go:
  ```go
  package main

  import (
      "project/helpers"
      "fmt"
  )

  func main() {
      fmt.Println(helpers.Add(3, 5))          // Output: 8
      fmt.Println(helpers.ToUpperCase("go")) // Output: GO
  }
  ```

here project is initialised directory during 
```go
  go mod init project
```
but if you used github URL then you would import these folder as
```go
  import (
    "github.com/username/{folder_name}/helpers"
  )
```
above in main folder

---

## Additional Notes 📝

- **Anonymous Functions**: Define and call immediately:
  ```go
  func() { ... }()
  ```
- Ensure **synchronization** when using Goroutines by using WaitGroups, or else the `main` function may exit prematurely.
- Miscellaneous
  ```go
  type DefaultParser struct {
    customField string
  }
  func (d DefaultParser) GetCustomField() string {   // copy of the instance and not actual instance will be changed
    return d.customField
  }
  func (*d DefaultParser) GetCustomField() string {   // actual instance will be modified in this function
    return d.customField
  }
  ```

---

## Receiver

(r *mutationResolver) -> receiver to show that CreateAccount is a method of mutationResolver just like in OOP 
```go
  func (r *mutationResolver) CreateAccount(ctx context.Context, in AccountInput) (*Account, error) {
```




