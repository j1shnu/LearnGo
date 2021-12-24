package main

import (
	"fmt"
	"time"
)

// Simple func with param
func hi(name string) {
	fmt.Printf("Hi %v ....!\n", name)
}

// Func with return
func bye(name string, f func(string)) string {
	f(name)
	val2 := fmt.Sprintf("Bye %v ...!\n", name)
	return val2
}

// https://www.epochconverter.com/ - get no.of days to bday
func getUnixTime(date string) int64 {
	dayTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(fmt.Sprintf("Error while parsing birth date:- %v", err))
	}
	return dayTime.Unix()
}
func daysToBday(month int8, day int8) string {
	today := time.Now().Unix()
	bday := fmt.Sprintf("%v-%v-%v", time.Now().Year(), month, day)
	bdayTimeUnix := getUnixTime(bday)
	if today > bdayTimeUnix {
		bday = fmt.Sprintf("%v-%v-%v", time.Now().Year()+1, month, day)
		bdayTimeUnix = getUnixTime(bday)
		fmt.Println("This year Birthday is over.")
	}
	days := bdayTimeUnix - today
	return fmt.Sprintf("%v days to BirthDay", days/86400)
}

func main() {
	// hi("Bro")
	// hi("Buddy")
	// fmt.Println(bye("Bro", hi))
	fmt.Println(daysToBday(11, 30))

}
