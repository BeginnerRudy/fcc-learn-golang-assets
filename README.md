# Assets for "Learn Go" on FreeCodeCamp

This is a snapshot of the code samples for the ["Learn Go" course](https://boot.dev/courses/learn-golang) on [Boot.dev](https://boot.dev) at the time the video for FreeCodeCamp was released on YouTube. If you want the most up-to-date version of the code, please visit the official [Boot.dev course](https://boot.dev/courses/learn-http). Otherwise, if you're looking for the files used in the video, you're in the right place!

* [Course code samples](/course)
* [Project steps](/project)

## License

You are free to use this content and code for personal education purposes. However, you are *not* authorized to publish this content or code elsewhere, whether for commercial purposes or not. 

## Ch1 Why Go?
- Go is compiled to machine code directly, it does not run in a VM like Java, the execution spped is similar with Java or C#.
- When you want to sell python product, the customer would have access to the source code directly, since it is the interpreted lanaguage. If you really want, you could always compile with Cython. This will generate C code, which you can then compile with any C compiler such as GCC.
- Go is a strong type language, we can find type issue during compiling time. 
- Java Program uses JVM for GC, so it consumes resources for JVM. While the Golang support GC as well, but it does not requires a VM, it uses a go-runtime for GC, just like a side car which compiled with your go program. Memory Overhead: Java/C# << Golang << C/CRust. For reference: https://medium.com/@dexwritescode/comparison-between-java-go-and-rust-fdb21bd5fb7c

## Ch2 Variables
- Variable Declaration
    - var [variable name] [type] = [initial value]
        - Can be used inside and outside of functions
        - Variable declaration and value assignment can be done separately
    - [variable name] := [initial value]
        - Can only be used inside functions       
        - Variable declaration and value assignment cannot be done separately (must be done in the same line)
    - v1, v2 := "v1", 12
- Type Casting
    - var1 := type(var2)
    - E.g val1 := int(var2)
- Constant
    - const CONSTNAME type = value
    - const CONSTNAME = value
    - The constant in Golang is different to Java, since the value shoule be known and computed at compile time in Go while the constant value can be computed during runing time in Java. In both language, the constant cannot be modified after the value is assigned. 
- String formartting
    - fmt.Printf -> prints a formatted string to standard out
    - fmt.Sprintf -> returns the formatted string
    - %v, %s, %d, %f
- Conditionals
    - Go does not use parentheses around the condition
    - if condition { statement } else if condition { statement } else { statement }
- The initial section of an if statement
    - if init statement; condition {statement}
    - if length := getLength(email); length < 1 { do something }

## Ch3 Functions
- Functions
    -  Function signature: func functionName(var1 type, var2 type) returnType {}
    - It is different with the C style, since it feels more natural in English, that X is an integer ==> var x int = 10
- Multiple parameters
    - When multiple arguments are of the same type, we only need to declare the type at the last one.
    - func add(x, y int) int { return x + y }
- Multiple return values
    - func addSub(x, y int) (int, int) { return x+y, x-y }
- Named return values
    - func addSub(x, y int) (sum, diff int) { return x+y, x-y }
    - named return values will be automatically returned. 
- Ignore return values
    - We can ignore the return value with under score
    - firstName, _ := getNames()

 ## Ch4 Structs
 - Structs is a collection type in go, it is just a collection of key value pairs. 
    ```
    type struct_name struct {
        member1 datatype;
        member2 datatype;
        member3 datatype;
        ...
    }

    type Person struct {
        name string
        age int
        job string
        salary int
    }

    
    // Pers1 specification
    pers1.name = "Hege"
    pers1.age = 45
    pers1.job = "Teacher"
    pers1.salary = 6000

    // Pass Struct as Function Arguments
    func printPerson(pers Person) {
        fmt.Println("Name: ", pers.name)
        fmt.Println("Age: ", pers.age)
        fmt.Println("Job: ", pers.job)
        fmt.Println("Salary: ", pers.salary)
    }
    ```
- Nested Structs
    ```
    type car struct {
        Make string
        Model string
        Height int
        Width int
        FrontWheel Wheel
        BackWheel Wheel
    }

    type Wheel struct {
        Radius int
        Material string
    }
    ```
- Anonymous Structs
    ```
    // The purpose of using anonymous struct is that we do not want to create multiple instance of the struct
    myCar := struct {
        Make string
        Model string
    } {
        Make: "tesla"
        Model: "model 3"
    }
    ```
- Embedded Structs
    ```
    type car struct {
        make string
        model string
    }

    type truck struct {
        // "car" is embedded, so the definition of a "truck" now also additionally contains all of
        // the fields of the car struct
        car
        bedSize int
    }

    lanesTruck := truck {
        bedSize := 10,
        car: car{
            make: "toyota",
            model: "camry"
        }
    }
    ```
- Methods
    ```
    type type_name struct { }
    func (m type_name) function_name() int {
        //code
    }

    type rect struct {
        width int
        height int
    }

    func (r rect) area() int {
        return r.width * r.height
    }
    ```
## Ch5 Interface
- Interfaces
    ```
    // Creating an interface
    type myinterface interface{
    // Methods
    fun1() int
    fun2() float64
    }

    // Creating an interface
    type tank interface {
        // Methods
        Tarea() float64
        Volume() float64
    }
    
    type myvalue struct {
        radius float64
        height float64
    }
    
    // Implementing methods of
    // the tank interface
    func (m myvalue) Tarea() float64 {
    
        return 2*m.radius*m.height +
            2*3.14*m.radius*m.radius
    }
    
    func (m myvalue) Volume() float64 {
    
        return 3.14 * m.radius * m.radius * m.height
    }

    // Main Method
    func main() {
        // Accessing elements of
        // the tank interface
        var t tank
        t = myvalue{10, 14}
        fmt.Println("Area of tank :", t.Tarea())
        fmt.Println("Volume of tank:", t.Volume())
    }
    ``` 
- Multiple Interfaces
    - Yes, a struct can implement as many as interfaces we want. 
- Naming args
    - Are you required to name the arguments of an interface in order for your code to compile properly? 
        -  No, but we usuallt add it for readability
- Type Assertion 
    - When working with interfaces in Go, every once-in-awhile you'll need access to the underlying type of an interface value. You can cast an interface to its underlying type using a type assertion.
        ```
        v = i.(T)
        // where i is the interface and T is the concrete type. It will panic if the underlying type is not T. To have a safe cast, you use:
        v, ok = i.(T)
        ```
- Type Switch
    ```
    func test(e expense) {
        address, cost := getExpenseReport(e)
        switch e.(type) {
            case email:
                fmt.Printf("Report: The email going to %s will cost: %.2f\n", address, cost)
                fmt.Println("====================================")
            case sms:
                fmt.Printf("Report: The sms going to %s will cost: %.2f\n", address, cost)
                fmt.Println("====================================")
            default:
                fmt.Println("Report: Invalid expense")
                fmt.Println("====================================")
        }
    }
    ```
- Clean Interfaces
    - Interfaces sohuld hava as few methods as possible
    - Interfaces allow you to define a method's behavior once and use it for many different types. 
        - False, just the method signature
## Errors
- Errors
    - Golang returns error explicitly, it is different from try catch block, because it enfore you to think about the error of each function. 
        ```
        // Go programs express errors with error values. An Error is any type that implements the simple built-in error interface. 
        type error interface {
            Error() string
        }
        ```
- custom errors
    - we can implement the built-in error type to customize the error string with our desired data.
    ```
    type divideError struct {
        dividend float64
    }

    func (de divideError) Error() string {
        return fmt.Sprintf("can not divide %v by zero", de.dividend)
    }

    // don't edit below this line
    func divide(dividend, divisor float64) (float64, error) {
        if divisor == 0 {
            // We convert the `divideError` struct to an `error` type by returning it
            // as an error. As an error type, when it's printed its default value
            // will be the result of the Error() method
            return 0, divideError{dividend: dividend}
        }
        return dividend / divisor, nil
    }
    ```
- errors package 
    - The Go standard library provides an "errors" package that makes it easy to deal with errors. 
    ```
    var err error = errors.New("something, went wrong")
    ```
## Loops
- for loop
    ```
    for INITIAL; CONDITION; AFTER{
        // do something
    }

    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }
    ```
- omit condition (endless loop)
    ```
    for INITIAL; ; AFTER {
        // do something FOREVER
    }
    ```
- while loop
    - There is no while loop in Go. We can simulate it via a for loop.
    ```
    for CONDITION {
        // do some stuff while CONDITION is true
    }

    plantHeight := 1

    for plantHeight < 5 {
        fmt.Println("+++");
        plantHeight++
    }
    ```
- continue and break
    - continue and break are similar to the ones in Java and C#.
## Slices
- Arrays
    - In Go, an array has a fixed size. 
        ```
        var array_variable = [size]datatype{elements of array}
        var nums = [5]int{1, 2, 3, 4, 5}
        ```
- Slices
    ```
    numbers := []int{1, 2, 3, 4, 5} // slice
    var nums = [5]int{1, 2, 3, 4, 5} // array
    num := nums[0:2] // slice, the lower index is inclusive while the higher index is exclusive
    foo := nums[:] // slice, all the elements from nums
    ```
- Make
    - The make function can be used to create a slice
        ```
        // func make([]T, len, cap) []T
        mySlice := make([]int, 5, 10)

        // The capacity argument is usually omitted and default to the len
        mySlice := make([]int, 5)

        // If we want to create a slice with specific set of values, we can use a slice literal
        mySlice := []string{"I", "love", "go"}
        ```
- variadic functions 
    - A variadic function receives the variadic arguments as a slice. 
        ```
        func sum(nums ...int) int {
            // nums is just a slice
            for i := 0; i < len(nums); i++{
                num := nums[i]
            }
        }
        ```
- append
    ```
    // append is a variadic thing
    slice = append(slice, oneThing)
    slice = append(slice, oneThing, twoThing)
    slice = append(slice, oneThing, twoThing, threeThings)
    ```
- 2-dimentional slice
    ```
    rows := [][]int{}
    matrix := make([][]int, 0)
    ```
- Range
    ```
    for INDEX, ELEMENT := range SLICE {

    }

    fruits := []string{"apple", "bnana", "grape"}
    for i, fruit := range fruits {
        fmt.Println(i, fruit)
    }
    ```
## Maps
- Map
    ```
    // An Empty map
    map[Key_Type]Value_Type{}

    // Map with key-value pair
    map[Key_Type]Value_Type{key1: value1, ..., keyN: valueN}

    // Example 1
    var mymap map[int]string

    // Create a map use the make function
    make(map[Key_Type]Value_Type, initial_Capacity)
    make(map[Key_Type]Value_Type)

    var My_map = make(map[float64]string)

    // How to iterate over a map?
    map_ := map[int]string{
        90: "Dog",
        91: "Cat",
        92: "Cow"
    }

    for id, pet := range map_ {
        fmt.Println(id, pet)
    }

    // How to add key-value paires in the map
    map_name[key]=value
    map_[100] = "lion"

    // How to retrieve a value realted to a key in the map?
    map_name[key]

    val1 := map_[100]
    val2 := map_[90]

    // How to check the existence of a key in a map?
    pet_name, ok := map_[90]
    _, ok = map_[100]

    How to delete a key from map
    delete(map_name, key)

    delete(map_, 90)
    delete(map_, 100)

    // Map are assined by reference, 2 maps points to the same reference, they one modification will shown from another map

    // Assigned the map into a new variable
    new_map := map_
 
    // Perform modification in new_map
    new_map[96] = "Parrot"
    new_map[98] = "Pig"

    // This change will also show in the map_ variable
    ```
- Nested map
    ```
    func getNameCounts(names []string) map[rune]map[string]int {
        // A map use rune as key, and a map of string, int pair as value
        nameCounts := make(map[rune]map[string]int)
        for _, name := range names {
            firstLetter := rune(0)
            if len(name) != 0 {
                firstLetter = rune(name[0])
            }

            if _, ok := nameCounts[firstLetter]; !ok {
                nameCounts[firstLetter] = make(map[string]int)
            }
            if _, ok := nameCounts[firstLetter][name]; !ok {
                nameCounts[firstLetter][name] = 0
            }
            nameCounts[firstLetter][name]++
        }
        return nameCounts
    }
    ```
## Advanced Functions (SKIP FOR NOW)
- Functions as data
## Pointers
A pointer is a special kind of variable that is not only used to store the memory addresses of other variables but also points where the memory is located and provides ways to find out the value stored at that memory location.

Two important operators which we will use in pointers i.e. 

1. \* Operator also termed as the dereferencing operator used to **declare** pointer variable and **access** the value stored in the address.
2. \& operator termed as address operator used to returns the **address of a variable** or to access the address of a variable to a pointer.
```
// Declarubg a pointer
var pointer_name *Data_Type

var s *string

// Initilization of a pointer
var a = 45 // normal variable declaration
var s *int = &a
```
Important rules
1. The default value or zero-value of a pointer is always nil. Or you can say that an uninitialized pointer will always have a **nil value**.
2. pointers are type specific, you cannot assign int address to a string pointer.
3. To overcome the above mention problem you can use the Type Inference concept of the var keyword. 
    ```
    // using var keyword we are not defining any type with variable
    var y = 458
     
    // taking a pointer variable using var keyword without specifying the type
    var p = &y
 
    fmt.Println("Value stored in y = ", y)
    fmt.Println("Address of y = ", &y)
    fmt.Println("Value stored in pointer variable p = ", p)
    ```
4. You can also use the shorthand (:=) syntax to declare and initialize the pointer variables.
    ```
    // using := operator to declare
    // and initialize the variable
    y := 458
    
    // taking a pointer variable using
    // := by assigning it with the
    // address of variable y
    p := &y

    fmt.Println("Value stored in y = ", y)
    fmt.Println("Address of y = ", &y)
    fmt.Println("Value stored in pointer variable p = ", p)
    ```
5. Dereferencing the Pointer
    ```
    //We can access the pointer value with the * operator
    // using var keyword
    // we are not defining
    // any type with variable
    var y = 458
      
    // taking a pointer variable using
    // var keyword without specifying
    // the type
    var p = &y
  
    fmt.Println("Value stored in y = ", y)
    fmt.Println("Address of y = ", &y)
    fmt.Println("Value stored in pointer variable p = ", p)
 
    // this is dereferencing a pointer
    // using * operator before a pointer
    // variable to access the value stored
    // at the variable at which it is pointing
    fmt.Println("Value stored in y(*p) = ", *p)     <========== Dereferencing the Pointer
    ```
## Local Developments
- The official package management tool is ***Go Module***.
### Package
- There are 2 kinds of packages: main package and other
- Only the file in main package can be compiled to an executable file.
- The customized packages
    - There should be only 1 package under 1 directory while there can be multiple files inside a directory
    - However, there can be packages inside sub-directorise.
- Import rule
    ```
    import (
        "fmt"
        "./service"
        apiNew "./service/api" <--- rename the package
        "github.com/rsj217/service"
    )
    ```
### Module
- https://juejin.cn/post/6869570760738865166
- 还记得一个package是.go文件的集合么? 一个Module就是package(go 包)的集合. 即: 一个Module--多个package; 一个package--多个.go文件.
- 在一个module的树状存储结构的根部, 必须有一个go.mod文件, 该文件定义了如下内容:
    - 此module的module path(也是根目录里的package的import path)
    - 该module的编译所需的依赖module. 每个依赖项的结构都是 依赖module的module path + 版本semantic versioin.

### Golang commands

## Channels and concurrenty
- https://www.geeksforgeeks.org/channel-in-golang/
### Goroutines

## Mutex

## Generics
### Type parameters
```
import "golang.org/x/exp/constraints"

type store[P product] interface {
    sell(P)
}

func GMin[T constraints.Ordered](x, y T) T {
    if x < y {
        return x
    }
    return y
}

// Providing the type argument to GMin, in this case int, is called instantiation. 
x := GMin[int](2, 3)

// After successful instantiation we have a non-generic function that can be called just like any other function. For example, in code like
fmin := GMin[float64]
m := fmin(2.71, 3.14)
```
### Type sets
```
type Ordered interface {
    Integer|Float|~string
}
```
## Go facts
