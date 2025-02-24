package main 

import (
  "fmt"
  "learning-golang/mypackages/greetings"
  "learning-golang/mypackages/feedback"
)


func main() {

	var name = "Allan"
	var age = 36

    fmt.Println(greetings.GreetingScroll(name))
	fmt.Println(feedback.DrinkingAge(age))
	fmt.Println(feedback.DrivingAge(age))
	
}
