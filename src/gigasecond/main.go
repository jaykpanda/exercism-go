package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your birthday in dd-mm-yyyy")
	input, _ := reader.ReadString('\n')
	day, _ := strconv.Atoi(input[0:2])
	month, _ := strconv.Atoi(input[3:5])
	year, _ := strconv.Atoi(input[6:10])
	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	dday := dob.Add(time.Second * 1000000000)
	fmt.Printf("You will live your Gigasecond on this planet on : %v", dday)
}
