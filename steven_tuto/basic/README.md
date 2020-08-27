# Golang for ruby developer

## 1. "Hello World"

### Ruby
  ```ruby
  puts "Hello World"
  ```
### Go
  ```go
  package main

  import "fmt"

  func main(){
    fmt.Println("Hello Wolrd")
  }
  ```

  In go, a complete program is created by linking a single, unimported package called the main package with all the packages it imports, transitively. The main package must have package name main and declare a function main that takes no arguments and returns no value.

  ```
  func main() { … }

  ```
  Program execution begins by initializing the main package and then invoking the function main. When that function invocation returns, the program exits. It does not wait for other (non-main) goroutines to complete.

  The language specification does not give special meaning to the name main outside of this context. The name main is not a reserved name.

  It's OK to declare a main function in non-main packages. In such cases, it's just a function named main.

## 2. Comandline

### Ruby
  ```bash
  $ ruby hello_world.rb
  ```
### Go
  ```bash
  $ go run main.go
  ```

## 3. Comment code

### Ruby

  ```ruby
  puts "Commenting Code in Ruby"
  # this is single line comment in ruby
  =begin
  This is multiline 
  Comment in ruby
  =end
  ```
### Go
  ```go
  package main

  import "fmt"

  func main(){
    fmt.Println("Commenting in Go")
    // this is single line comment in Go
    /* this is multiline
    Comment in Go */
  }
  ```

## 4. Declare variables

### Ruby

  ```ruby
  a = 1
  b
  b = "hello"
  ```
### Go
  ```go
    var a int = 1
    // OR
    // this dynamically declares type for variable a as int.
    var a = 1
    // this dynamically defines variable a and declares its type as int.
    a := 1

    // declares variable a as type int
    a := 1

    // declares variable b as type string
    b := "hello"
  ```
## 5. Data Types

  Some types in go

  ```go
    var a bool = true
    var b int = 1
    // In go the string should be declared using double quote.
    // Using single quote unlike in Ruby can be used only for single character for its byte representation
    var c string = "hello world" 
    var d float32 = 1.222
    var x complex128 = cmplx.Sqrt(-5 + 12i)
  ```
  ##

  In Ruby , it is possible to assign a different data type to a variable that already contains data, but this can't be done in Go.

### Ruby
  ```ruby
  a = 1
  a = "Hello"
  ```
### Go
  ```go
  a := 1
  a = "Hello"
  // Gives error: cannot use "hello"(type string) as type int is assignment
  ```
## 6.Hash/Map

  Like Ruby, it is possible to define hashes in Go, but with a different name, it is map. The syntax is map[string]string (data type for key and for the value)

### Ruby

  ```ruby
  #Ruby:
  hash = {name:  'Nikita', lastname: 'Acharya'}
  # Access data assigned to name key
  name = hash[:name]
  ```
### Go
  ```go
  //Go
  hash := map[string]string{"name":  "Nikita", "lastname": "Acharya"}
  // Access data assigned to name key
  name := hash["name"]
  ```
  And, it is possible to check if key exists:

  ```go
  //GO
  name := map[string]string{name: "Nikita"}
  lastname, ok := name["lastName"]
  if ok{
    fmt.Printf("Last Name: %v", lastname)
  }else{
    fmt.Println("Last Name is missing")
  }

  ```
  If `name` has already the key `lastname`, `ok` will be assigned `true` and `lastname` will be assigned the value of `name["lastname"]`

## 7. Array

In Go, we have also array, but we need to declare its length.

### Ruby
  ```ruby
  #Ruby 
  array = [1,2,3]
  ```
### Go
  ```Go
  //Go
  array := [3]int{1,2,3}
  names := [2]string{"Nikita", "Aneeta"}
  ```
### *Slice

  Array in Go is limited in the case the value of Array changed at runtime. And Array doesn't provide subarray capability either. Therefore, Go provide a data type called Slice. The Slice stores the elements sequentially and can be expanded at any time. The Slice declaration is similar to Array, but we don't need to define a length for it.

  ```go
  var b []int
  ```

### Add element to Array/Slice

In Ruby, we use operator `+` to add the new elements into array.
  ```ruby
  #Ruby
  numbers = [1, 2]
  a  = numbers + [3, 4]
  #-> [1, 2, 3, 4]
  ```
In Go, we use `append` function.
  ```go
  //Go
  //Go
  numbers := []int{1,2}
  numbers = append(numbers, 3, 4)
  //->[1 2 3 4]
  ```

### Sub Slicing Array

  ```ruby 
  #Ruby 
  number2 = [1, 2, 3, 4]
  slice1 = number2[2..-1] # -> [3, 4]
  slice2 = number2[0..2] # -> [1, 2, 3]
  slice3 = number2[1..3] # -> [2, 3, 4]

  ```

  ```go
  //Go
  // initialize a slice with 4 len, and values
  number2 = []int{1,2,3,4}
  fmt.Println(numbers) // -> [1 2 3 4]
  // create sub slices
  slice1 := number2[2:]
  fmt.Println(slice1) // -> [3 4]
  slice2 := number2[:3]
  fmt.Println(slice2) // -> [1 2 3]
  slice3 := number2[1:4]
  fmt.Println(slice3) // -> [2 3 4]
  ```
### Copy Slice/Array

  In Ruby, it is possible to copy an array by performing assignment.

  ```ruby
  #Ruby
  array1 = [1, 2, 3, 4]
  array2 = array2 # -> [1, 2, 3, 4]
  ```

  But we can do that in Go. We must initiate a new array, then use the `copy` function.

  ```go
  //Go
  // Make a copy of array1
  array1 := []int{1,2,3,4} 
  array2 := make(int[],4)
  copy(array2, array1)
  ```

  Note: The number of elements to be copied depends on the length of the target array that we declare.

  ```go
  a := []int{1,2,3,4}
  b := make([]int, 2)
  copy(b, a) // copy a to b
  fmt.Println(b)
  //=> [1 2]
  ```