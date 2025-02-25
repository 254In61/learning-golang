// Golang program to show how to
// declare and define the struct
 
package main
 
import "fmt"

// defining the struct
type Address struct {
    street   string
    suburb   string
    code     int
}

// We can add functions to the struct that can add behavior to it as shown below:
func (addr Address) print_details(){
    fmt.Printf("Street %s is located in suburb named %s.\n", addr.street, addr.suburb)
    fmt.Printf("Suburb %s is identified by code %d.\n", addr.suburb, addr.code)
}

func main(){
    // Declaring a variable of a `struct` type
    // All the struct fields are initialized with their zero value
    var a Address 
    fmt.Println(a) 
    
    /*
    Output :

    $ go run main.go 
      {  0}

    */

    // Declaring and initializing a struct using a struct literal
    a1 := Address{"Clermont Street", "Dadenong", 3623}
    fmt.Println("Address1: ", a1)

    // Accessing struct fields using the dot operator
    fmt.Println("Street Name: ", a1.street)
    fmt.Println("Post Code: ", a1.code)

    /*
    Output :
    
    $ go run main.go 
      Address1:  {Clermont Street Dadenong 3623}

    */
    
    // Naming fields while initializing a struct
    a2 := Address{street: "Anikaa st", suburb: "California", code: 277001}
    fmt.Println("Address2: ", a2)
    // Accessing struct fields using the dot operator
    fmt.Println("Street Name: ", a2.street)
    fmt.Println("Post Code: ", a2.suburb)
    
    /*
    Output :

    $ go run main.go 
      Address2:  {Anikaa st California 277001}

    */

    // Uninitialized fields are set to their corresponding zero-value
    a3 := Address{suburb: "Victoria-Park"}
    fmt.Println("Address3: ", a3)

     /*
    Output :

    $ go run main.go 
      Address3:  { Victoria-Park 0}

    */

    // using functions to add behavior to a struct
    add_1 := Address{"Clermont Street" , "Holmview", 4301}
    add_1.print_details()

    // modify address details
    add_1.code = 7971
    add_1.print_details()
    

}
 
