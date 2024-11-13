
# GoLang Quick Reference Guide 🐹

## 📚 General Concepts

### Null and Nil
- `null` is referred to as `nil` in Go.
- Data type descriptions in Go come **after** the variable name, unlike some other languages.

### No Inheritance
- Go does **not** support inheritance, nor does it have `super` or `parent` keywords.

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

## Additional Notes 📝

- **Anonymous Functions**: Define and call immediately:
  ```go
  func() { ... }()
  ```
- Ensure **synchronization** when using Goroutines by using WaitGroups, or else the `main` function may exit prematurely.

---
