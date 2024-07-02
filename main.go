package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// an interface with a common method that evaluates input message
type parse interface {
	parseCron() string
}

// a struct that holds data of individual elements such as minutes, hours, day of month, month, day of week and command
type cron struct {
	minutes    string
	hours      string
	dayOfMonth string
	month      string
	dayOfWeek  string
	command    string
}

// support variables to validate if the cron is given in letter format for day of week and month
var days = []string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}
var months = []string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <arg1> <arg2> ...")
		return
	}
	input := os.Args[1]

	res := strings.Fields(input)

	if len(res) != 6 {
		fmt.Println("Invalid cron pattern")
		return
	}

	crn := cron{
		minutes:    res[0],
		hours:      res[1],
		dayOfMonth: res[2],
		month:      res[3],
		dayOfWeek:  res[4],
		command:    res[5],
	}

	finalResult := cronParser(crn)
	fmt.Println(finalResult)
}

func cronParser(cr parse) string {
	output := cr.parseCron()
	return output
}

func (c cron) parseCron() string {
	resMin, resMinBool := parseMessage(0, 59, c.minutes)
	resHours, resHoursBool := parseMessage(0, 23, c.hours)
	resDayOfMonth, resDayOfMonthBool := parseMessage(1, 31, c.dayOfMonth)
	resMonth, resMonthBool := parseMessage(1, 12, c.month)
	resDayOfWeek, resDayOfWeekBool := parseMessage(0, 6, c.dayOfWeek)

	var resCommand []string
	resCommand = append(resCommand, c.command)
	var result string
	if resMinBool && resHoursBool && resDayOfMonthBool && resMonthBool && resDayOfWeekBool {
		tempStr := formatData(resMin, "minute")
		result = result + tempStr + "\n"
		tempStr = formatData(resHours, "hour")
		result = result + tempStr + "\n"
		tempStr = formatData(resDayOfMonth, "day of month")
		result = result + tempStr + "\n"
		tempStr = formatData(resMonth, "month")
		result = result + tempStr + "\n"
		tempStr = formatData(resDayOfWeek, "day of week")
		result = result + tempStr + "\n"
		tempStr = formatData(resCommand, "command")
		result = result + tempStr
	} else {
		result = "Invalid cron message, Please check and retry!"
	}
	return result
}

func formatData(str []string, typeOfData string) string {
	inp := fmt.Sprintf("%-14s", typeOfData)
	for _, v := range str {
		inp = inp + v + " "
	}
	return inp
}

func parseMessage(lLimit int, uLimit int, str string) ([]string, bool) {
	var output []string

	var input []string
	if strings.Contains(str, ",") {
		input = strings.Split(str, ",")
	} else {
		input = append(input, str)
	}
	for _, v := range input {
		if strings.Contains(v, "/") {
			tempRes := strings.Split(v, "/")
			if len(tempRes) < 2 {
				output = nil
				return output, false
			}
			var uLim int
			var lLim int
			if tempRes[0] == "*" {
				uLim = uLimit
				lLim = lLimit
			} else if strings.Contains(tempRes[0], "-") {
				tempResOne := strings.Split(tempRes[0], "-")
				if len(tempResOne) < 2 {
					output = nil
					return output, false
				}
				temLowerLimit, err := strconv.Atoi(tempResOne[0])
				if err != nil {
					output = nil
					return output, false
				}
				temUpperLimit, err := strconv.Atoi(tempResOne[1])
				if err != nil {
					output = nil
					return output, false
				}
				if temLowerLimit < temUpperLimit && temLowerLimit >= lLimit && temUpperLimit <= uLimit {
					lLim = temLowerLimit
					uLim = temUpperLimit
				} else {
					output = nil
					return output, false
				}
			} else {
				output = nil
				return output, false
			}
			splitter, err := strconv.Atoi(tempRes[1])
			if err != nil {
				output = nil
				return output, false
			}
			for i := lLim; i <= uLim; {
				output = append(output, strconv.Itoa(i))
				i = i + splitter
			}
		} else if strings.Contains(v, "-") {
			tempRes := strings.Split(v, "-")
			if len(tempRes) < 2 {
				output = nil
				return output, false
			}
			lowerRange, err := strconv.Atoi(tempRes[0])
			if err != nil {
				output = nil
				return output, false
			}
			upperRange, err := strconv.Atoi(tempRes[1])
			if err != nil {
				output = nil
				return output, false
			}
			if lowerRange < upperRange && lowerRange >= lLimit && upperRange <= uLimit {
				for i := lowerRange; i <= upperRange; i++ {
					output = append(output, strconv.Itoa(i))
				}
			} else {
				output = nil
				return output, false
			}
		} else if strings.Contains(v, "*") {
			if len(v) == 1 {
				for i := lLimit; i <= uLimit; i++ {
					output = append(output, strconv.Itoa(i))
				}
			} else {
				output = nil
				return output, false
			}
		} else {
			if len(v) != 0 {
				val, err := strconv.Atoi(v)
				if err != nil {
					if lLimit == 0 && uLimit == 6 {
						found := false
						for _, x := range days {
							if strings.EqualFold(x, v) {
								output = append(output, v)
								found = true
							}
						}
						if !found {
							output = nil
							return output, false
						} else {
							continue
						}
					} else if lLimit == 1 && uLimit == 12 {
						found := false
						for _, y := range months {
							if strings.EqualFold(y, v) {
								output = append(output, v)
								found = true
							}
						}
						if !found {
							output = nil
							return output, false
						} else {
							continue
						}
					} else {
						output = nil
						return output, false
					}
				} else {
					if val >= lLimit && val <= uLimit {
						output = append(output, v)
					}
				}
			}
		}
	}
	output = removeDuplicates(output)
	return output, true
}

func removeDuplicates(out []string) []string {
	tMap := make(map[string]bool)
	tSlice := []string{}

	for _, elem := range out {
		if !tMap[elem] {
			tMap[elem] = true
			tSlice = append(tSlice, elem)
		}
	}
	return tSlice
}
