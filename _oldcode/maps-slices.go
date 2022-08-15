package main

import (
	"log"
	"sort"
)

type MonthName struct {
	Portuguese string
	English    string
}

func main() {
	monthMap := make(map[string]MonthName)
	jan := MonthName{"Janeiro", "January"}
	monthMap["jan"] = jan
	monthMap["feb"] = MonthName{"Fevereiro", "February"}
	log.Println(monthMap["jan"])
	log.Println(monthMap["jan"].English)
	log.Println(monthMap["feb"].Portuguese)

	var weekdays []string
	weekdays = append(weekdays, "sunday")
	weekdays = append(weekdays, "monday")
	weekdays = append(weekdays, "tuesday", "wednesday", "thursday")
	log.Println(weekdays)
	log.Println(weekdays[2:4])
	log.Println(weekdays[2:])
	sort.Strings(weekdays)
	log.Println("sorted weekdays", weekdays)

	names := []string{"cat", "dog", "fish"} // declare + initialize shorthand
	log.Println(names)
}
