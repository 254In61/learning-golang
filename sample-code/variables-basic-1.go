package main 

import "fmt"

func main() {

	var name_1 string
	var age_1 int
	var height_1 float64

	var name_2 = "Allan Maseghe"
	var age_2 = 43
	var height_2 = 180.2

    fmt.Println(name_1) // A blank output since variable is zero and it is a string
	fmt.Println(age_1) // output = 0
	fmt.Println(height_1) // output = 0

	fmt.Println(name_2) 
	fmt.Println(age_2)
	fmt.Println(height_2)

	// Multiple variables of different types 
    // are declared and initialized in the single line 
    var myvariable1, myvariable2, myvariable3 = 2, "GFG", 67.56 
  
    // Display the value and  
    // type of the variables 
    fmt.Printf("The value of myvariable1 is : %d\n", myvariable1) 
  
    fmt.Printf("The type of myvariable1 is : %T\n", myvariable1) 
  
    fmt.Printf("\nThe value of myvariable2 is : %s\n", myvariable2) 
  
	fmt.Printf("The type of myvariable2 is : %T\n", myvariable2) 
  
    fmt.Printf("\nThe value of myvariable3 is : %f\n", myvariable3) 
  
    fmt.Printf("The type of myvariable3 is : %T\n", myvariable3) 
} 
	

