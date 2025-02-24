package feedback

func DrinkingAge(age int) string {
	if age < 18 {
		return "Too young to drink"
	} else if age >= 18 && age < 80 {
		return "Drink responsibly"
	} else {
		return "Advanced age. Check your health first?"
	}

}