# intro
- Go/Golang is a procedural and statically typed programming language having the syntax similar to C programming language.
  Strong typing: Go is a statically typed language, which helps catch errors at compile time rather than at runtime.
  Go is a statically typed language, which means that the type of a variable must be declared before it can be used.

- It was developed in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson at Google but launched in 2009 as an open-source programming language. 
  
- Programs are assembled by using packages, for efficient management of dependencies.
  Go has a standard library that provides support for a wide range of functionality, including networking, encryption, and file handling.
  
- Supports environment adopting patterns alike to dynamic languages. For eg., type inference (y := 0 is a valid declaration of a variable y of type int).

- Go is designed to be simple, efficient, and easy to learn, making it a popular choice for building scalable network services, web applications, and command-line tools.

- Go is known for its support for concurrency, which is the ability to run multiple tasks simultaneously. 
  Concurrency is achieved in Go through the use of Goroutines and Channels, which allow you to write code that can run multiple operations at the same time. 
  This makes Go an ideal choice for building high-performance and scalable network services, as well as for solving complex computational problems.

- Another important feature of Go is its garbage collection, which automatically manages memory for you. 
  This eliminates the need for manual memory management, reducing the likelihood of memory leaks and other bugs that can arise from manual memory management.

- Go is often used for building large-scale distributed systems and high-performance applications.

- Cross-platform support: Go can be compiled to run on many different platforms, including Windows, Linux, and macOS.

- Fast compile times: Go has a fast compiler, which makes it easy to iterate quickly during development.

# how to run
- Ensure go is insalled.
  $ sudo apt update
  $ sudo apt install golang-go -y

- Verify the installation : $ go version

- Run the Go File : $ go run main.go

- Build and Run (Optional)
  To create a binary executable: 
   $ go build -o myapp main.go  # Build an executable file called myapp.
   $ ./myapp  # Runs the executable file.

# Syntax
- Check sample-code/1-basic-syntax.go


# Identifiers

- In programming languages, identifiers are used for identification purposes. In other words, identifiers are the user-defined names of the program components. In the Go language, an identifier can be a variable name, function name, constant, statement label, package name, or type. 

Example:

package main
import "fmt"

func main() {

 var name = "GeeksforGeeks"
  
}


There is a total of three identifiers available in the above example:
 main: Name of the package
 main: Name of the function
 name: Name of the variable

- Check out rules for identifiers : https://www.geeksforgeeks.org/identifiers-in-go-language/?ref=next_article

- The identifier represented by the underscore character(_) is known as a blank identifier. 
  It is used as an anonymous placeholder instead of a regular identifier, and it has a special meaning in declarations, as an operand, and in assignments.

# importing external Identifiers from a package/file

## overview
- The identifier which is allowed to access it from another package is known as the exported identifier. The exported identifiers are those identifiers which obey the following conditions:

  1. The first character of the exported identifier’s name should be in the Unicode upper case letter.
  
  2. The identifier should be declared in the package block or be a variable, function, type, or method name within that package.

- package line at the top reffers to the directory which contains identities to be imported.

## how to?
- Run the following command in your project root directory to ensure it's a Go module:
   $ go mod init learning-golang
     go: creating new go.mod: module learning-golang
     go: to add module requirements and sums:
        go mod tidy
  
  This creates a file go.mod that includes your module path.
  
   module learning-golang

   go 1.18

- Your package should be inside your Go module directory
  NB : The import refferences the folder dir name and NOT the file where the code is written!

 $ tree
.
├── README.md
├── go.mod
├── main.go
├── mypackages
│   └── greetings
│       └── greetings.go
├── sample-code
│   └── 1-basic-syntax.go
└── theory
    └── introduction.md

4 directories, 6 files

3 directories, 6 files

-   Ensure that your main.go file imports the package correctly. 
    If your module is named myproject, you should import it like this:

package main 

import (
  "fmt"
  "learning-golang/mypackages/greetings"
)

var LocalGreeting = "Hello, local file greetings"

func main() { 

    fmt.Println(LocalGreeting)
	fmt.Println(greetings.AllanGreeting)
	fmt.Println(greetings.NancyGreeting)
	
}

# golangci-lint
https://golangci-lint.run/welcome/install/#local-installation

https://golangci-lint.run/welcome/quick-start/

 $ sudo snap install golangci-lint --classic
 golangci-lint 1.64.5 from GolangCI-Lint (golangci✓) installed

 $ golangci-lint --version
 golangci-lint has version 1.64.5 built with go1.24.0 from 0a603e49 on 2025-02-13T21:19:55Z


## enabled or disabled linters?
  $ golangci-lint help linters

## run the linter

- Run linting across all files within the root directory.

learning-golang$ golangci-lint run
mypackages/feedback/drinking.go:6:31: illegal rune literal (typecheck)
        } else if age >= 18 && age < '80' {
                                     ^
main.go:15:17: undefined: greetings (typecheck)
    fmt.Println(greetings.GreetingScroll(name))
                ^
main.go:16:14: undefined: feeback (typecheck)
        fmt.Println(feeback.DrinkingAge(age))
                    ^
main.go:17:34: undefined: ae (typecheck)
        fmt.Println(feedback.DrivingAge(ae))
                                        ^

- Run in specific folder 

learning-golang$ golangci-lint run mypackages/feedback/
mypackages/feedback/drinking.go:6:31: illegal rune literal (typecheck)
        } else if age >= 18 && age < '80' {

# Data Types

## overview
https://www.geeksforgeeks.org/data-types-in-go/
- Data types specify the type of data that a valid Go variable can hold. 

- In Go language, the type is divided into four categories which are as follows:

 1. Basic type: These are further categorized into 3 sub-categories :
    a) - Numbers : Divided into 3 sub-categories
         i) Integers
         ii) Floating-Point Numbers
         iii) Complex numbers
    b) - strings
    c) - booleans
 2. Aggregate type: Array and structs come under this category.
 3. Reference type: Pointers, slices, maps, functions, and channels come under this category.
 4. Interface type

# variables
## overview
- A typical program uses various values that may change during its execution. For Example, a program that performs some operations on the values entered by the user. The values entered by one user may differ from those entered by another user. Hence this makes it necessary to use variables as another user may not use the same values. When a user enters a new value that will be used in the process of operation, can store temporarily in the Random Access Memory of the computer and these values in this part of memory vary throughout the execution and hence another term for this came which is known as Variables.

- So basically, a Variable is a placeholder of the information which can be changed at runtime. And variables allow to Retrieve and Manipulate the stored information.

## Declaring a variable :

1. Using var keyword
    var variable_name type = expression

  - In the above syntax, either type or = expression can be omitted, but NOT both
  - If the '=' expression is omitted, then the variable holds zero-value for the type like zero for the number, false for Booleans, “” for strings, and nil for interface and reference type. So, there is no such concept of an uninitialized variable in Go language.
  - If you remove type, then you are allowed to declare multiple variables of a different type in the single declaration. The type of variables is determined by the initialized values.

2. Using short variable declaration: The local variables which are declared and initialize in the functions are declared by using short variable declaration.

  variable_name:= expression

  In the above expression, the type of the variable is determined by the type of the expression.

# rune
https://www.geeksforgeeks.org/rune-in-golang/


# operators
Operators are the foundation of any programming language. Thus the functionality of the Go language is incomplete without the use of operators. Operators allow us to perform different kinds of operations on operands. In the Go language, operators Can be categorized based on their different functionality:

1. Arithmetic Operators : These are used to perform arithmetic/mathematical operations on operands in Go language
   - Addition: The ‘+’ operator adds two operands. For example, x+y.
   - Subtraction: The ‘-‘ operator subtracts two operands. For example, x-y.
   - Multiplication: The ‘*’ operator multiplies two operands. For example, x*y.
   - Division: The ‘/’ operator divides the first operand by the second. For example, x/y.
   - Modulus: The ‘%’ operator returns the remainder when the first operand is divided by the second. For example, x%y.

2. Relational Operators : Relational operators are used for the comparison of two values.
   - ‘=='(Equal To) operator checks whether the two given operands are equal or not. If so, it returns true. Otherwise, it returns false. For example, 5==5 will return true.
   - ‘!='(Not Equal To) operator checks whether the two given operands are equal or not. If not, it returns true. Otherwise, it returns false. It is the exact boolean complement of the ‘==’ operator. For example, 5!=5 will return false.
   - ‘>'(Greater Than)operator checks whether the first operand is greater than the second operand. If so, it returns true. Otherwise, it returns false. For example, 6>5 will return true.
   - ‘<‘(Less Than)operator checks whether the first operand is lesser than the second operand. If so, it returns true. Otherwise, it returns false. For example, 6<5 will return false.
   - ‘>='(Greater Than Equal To)operator checks whether the first operand is greater than or equal to the second operand. If so, it returns true. Otherwise, it returns false. For example, 5>=5 will return true.
   - ‘<='(Less Than Equal To)operator checks whether the first operand is lesser than or equal to the second operand. If so, it returns true. Otherwise, it returns false. For example, 5<=5 will also return true.
   
3. Logical Operators
4. Bitwise Operators
5. Assignment Operators
6. Misc Operators