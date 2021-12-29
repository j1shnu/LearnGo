package main

import (
	"Job_Finder/helper"
	"fmt"
	"strconv"
)

func main() {

	fmt.Printf("\nWelcome to simple job finder app\n\n")

	firstName := helper.Query("Enter Your First Name:- ")
	secondName := helper.Query("Enter Your Second Name:- ")
	emailID := helper.Query("Enter Your Email ID:- ")
	ageTemp, _ := strconv.Atoi(helper.Query("Enter Your Age:- "))
	age := uint16(ageTemp)
	jobType := helper.Query("Type of Job You're Looking For:- ")

	user := helper.UserDetails{
		FirstName:  firstName,
		SecondName: secondName,
		EmailID:    emailID,
		Age:        age,
		JobType:    jobType,
	}

	isValid, error := user.Validate()
	helper.ErrorHandler(error)
	if isValid {
		fmt.Printf(
			"\nHi %v, Thank you for using our app. We'll notify you in %v once a Job match found.\n", user.FirstName, user.EmailID,
		)
	}
}
