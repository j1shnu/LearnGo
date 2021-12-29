package helper

import (
	"errors"
	"fmt"
	"os"
	"regexp"
)

type UserDetails struct {
	FirstName  string
	SecondName string
	EmailID    string
	Age        uint16
	JobType    string
}

func Query(q string) string {
	var val string
	fmt.Print(q)
	fmt.Scan(&val)
	return val
}

func (user UserDetails) Validate() (bool, error) {
	if len(user.FirstName) < 2 {
		return false, errors.New("invalid first name..! should have atleast 2 chars")
	}
	if len(user.SecondName) < 2 {
		return false, errors.New("invalid second name..! should have atleast 2 chars")
	}
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(user.EmailID) {
		return false, errors.New("invalid email id")
	}
	if len(user.JobType) < 2 {
		return false, errors.New("invalid job type..! should have atleast 2 chars")
	}
	if user.Age < 10 {
		return false, errors.New(
			fmt.Sprint("you are ", user.Age, ", too young to apply for a job, We don't support child labour"),
		)
	}
	return true, nil
}

func ErrorHandler(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
