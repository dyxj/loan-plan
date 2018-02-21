package main

import (
	"fmt"
	"github.com/dyxj/loan-plan/loan"
	"time"
)

func main() {
	var loanAmount int64 = 5000 * 100
	nominalRate := float64(5.0)
	var duration int = 24
	startDateStr := "2018-01-01T00:00:01Z"
	dateFormat := "2006-01-02T15:04:05Z"
	startDate, err := time.Parse(dateFormat, startDateStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(loanAmount)
	fmt.Println(nominalRate)
	fmt.Println(duration)
	fmt.Println(startDateStr)
	fmt.Println(startDate)
	fmt.Println("-----------------")
	slRM, err := loan.GenPlan(loanAmount, nominalRate, duration, startDate)
	fmt.Println("-----------------")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(slRM))

	for _, v := range slRM {
		fmt.Println(*v)
	}
}
