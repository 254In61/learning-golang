package feedback

func DrivingAge(age int) string {
	if age < 16 {
		return "Too young to be driving"
	} else if age >= 16 && age < 80 {
		return "Drive safely"
	} else {
		return "Advanced age. Constant drive tests?"
	}

}