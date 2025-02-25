# overview
- Object-oriented programming is a programming paradigm which uses the idea of “objects” to represent data and methods. 
   - Go does not strictly support object orientation but is a lightweight object Oriented language. 
   - Object Oriented Programming in Golang is different from that in other languages like C++ or Java

- A structure or struct in Golang is a user-defined type that allows to group/combine items of possibly different types into a single type. Any real-world entity which has some set of properties/fields can be represented as a struct. 

 Structs in Golang are user-defined types that hold just the state and not the behavior. 

- This concept is generally compared with the classes in object-oriented programming. It can be termed as a lightweight class that does not support inheritance but supports composition. 

- For Example, an address has a name, street, suburb, state, Pincode. It makes sense to group these three properties into a single structure address as shown below.

type Address struct {
      name string 
      street string
      suburb string
      state string
      code int

}


We can also make them compact by combining the various fields of the same type as shown in the below example:

type Address struct {
    name, street, suburb, state string
    Pincode int
}

- In the above, the type keyword introduces a new type. It is followed by the name of the type (Address) and the keyword struct to illustrate that we’re defining a struct. 

- The struct contains a list of various fields inside the curly braces. Each field has a name and a type.

# define a structure
var a Address

The above code creates a variable of a type Address which is by default set to zero. 

For a struct, zero means all the fields are set to their corresponding zero value. So the fields name, street, suburb, state are set to “”, and Pincode is set to 0. 

You can also initialize a variable of a struct type using a struct literal as shown below:

var a = Address{"18", "Ikumbi", "Mlolongo", "Nairobi", 252636}

NOTE:
  - Always pass the field values in the same order in which they are declared in the struct.
  - You can’t initialize only a subset of fields with the above syntax.
  - Go also supports the name: value syntax for initializing a struct (the order of fields is irrelevant when using this syntax). And this allows you to initialize only a subset of fields. All the uninitialized fields are set to their corresponding zero value. Example:

  var a = Address{Name:”18”, street:”Ikumbi”, state:”Nairobi”, Pincode:252636} // state:””  .. suburb left blank

# accessing struct field
Ref : sample-code/struct.go


# function in struct
We can add functions to the struct that can add behavior.
Ref : sample-code/struct.go

# advantages of using structures in Go
1. Encapsulation: Structures allow you to encapsulate related data together, making it easier to manage and modify the data.
2. Code organization: Structures help to organize code in a logical way, which makes it easier to read and maintain.
3. Flexibility: Structures allow you to define custom types with their own behavior, making it easier to work with complex data.
4. Type safety: Structures provide type safety by allowing you to define the type of each field, which helps to prevent errors caused by assigning the wrong type of value.
5. Efficiency: Structures in Go are very efficient, both in terms of memory usage and performance.

# disadvantages of using structures in Go:
1. Complexity: Structures can make code more complex, especially if the structures have a large number of fields or methods.
2. Boilerplate code: When defining large structures with many fields, it can be time-consuming to write out all of the field names and types.
3. Inheritance: Go does not support inheritance, which can make it more difficult to work with large hierarchies of related data.
4. Immutability: Go structures are mutable by default, which can make it more difficult to enforce immutability in your code.

Overall, the advantages of using structures in Go typically outweigh the disadvantages, as they provide a powerful tool for managing and working with complex data. However, as with any programming technique, it’s important to use structures judiciously and be aware of their limitations.