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
- Anonymous Structs
- Embedded Structs
- Methods
## Ch5 Interface

