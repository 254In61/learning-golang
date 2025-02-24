package main   
// It contains the package main of the program, which have overall content of the program.
// It is the initial point to run the program, So it is compulsory to write.
/* 
- Every program must start with the package declaration.
- In Go language, packages are used to organize and reuse the code. 
  In Go, there are two types of program available :
   a) executable : programs that we can run directly from the terminal
   b) library.   : package of programs that we can reuse them in our program.

- Here, the package main tells the compiler that the package must compile in the executable program rather than a shared library. 
  It is the starting point of the program and also contains the main function in it.

*/

import "fmt"
//  import statement is used to include external packages that provide additional functionality beyond the built-in language features. 
//  In this case, “fmt” is used to implement formatted Input/Output with functions.

func main() { 
	// main function, it is beginning of execution of program.

    fmt.Println("Hello, World!")
	// fmt.Println() is a standard library function to print something as a output on screen.
	// In this, fmt package has transmitted Println method which is used to display the output.
}
