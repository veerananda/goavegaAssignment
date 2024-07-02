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

	// use fields so that all the whitespaces in the middle are handled, even if the user gives double space
	// between two elements, the program handles it
	res := strings.Fields(input)

	// if all the elements are not equal to 6, this could be a wrong cron message as per standard notion
	if len(res) != 6 {
		fmt.Println("Invalid cron pattern")
		return
	}

	// update the fields to the structure
	crn := cron{
		minutes:    res[0],
		hours:      res[1],
		dayOfMonth: res[2],
		month:      res[3],
		dayOfWeek:  res[4],
		command:    res[5],
	}

	finalResult := cronParser(crn) // intermediate function to call the interface function using the created structure data
	fmt.Println(finalResult)
}

func cronParser(cr parse) string {
	output := cr.parseCron() // call the interface method
	return output
}

func (c cron) parseCron() string {
	// calling parse method for each type like minutes, hours, months etc with the hard coded values of the range of that type
	// eg: minutes least value could be 0 and highest value could be 59
	// Also collecting the boolean to identify if that element type passed parsing successfully
	resMin, resMinBool := parseMessage(0, 59, c.minutes)
	resHours, resHoursBool := parseMessage(0, 23, c.hours)
	resDayOfMonth, resDayOfMonthBool := parseMessage(1, 31, c.dayOfMonth)
	resMonth, resMonthBool := parseMessage(1, 12, c.month)
	resDayOfWeek, resDayOfWeekBool := parseMessage(0, 6, c.dayOfWeek)

	var resCommand []string
	resCommand = append(resCommand, c.command)
	var result string
	// Once the result is obtained check if all the values are parsed correctly, even one value is incorrect cron pattern
	// the complete cron message is invalid
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

// function helps to format data and returns the data as string
func formatData(str []string, typeOfData string) string {
	inp := fmt.Sprintf("%-14s", typeOfData)
	for _, v := range str {
		inp = inp + v + " "
	}
	return inp
}

// parseMessage that takes in limits of the type and the string to be parsed
// returns slice of strings and bool value if the message of the type is parsed successfully or not
func parseMessage(lLimit int, uLimit int, str string) ([]string, bool) {
	var output []string

	var input []string
	// check for all the possible data user wants to enter, this is generally done using "," and parse each substring
	// obtained after splitting with respect to ","
	if strings.Contains(str, ",") {
		input = strings.Split(str, ",")
	} else {
		input = append(input, str)
	}
	for _, v := range input {
		// the next priority symbol to check is "/".
		// after splitting check if the range has valid data
		// if the data is missing on the left side of "/" or the right side of "/" consider as Invalid cron message
		if strings.Contains(v, "/") {
			tempRes := strings.Split(v, "/")
			if len(tempRes) < 2 {
				output = nil
				return output, false
			}
			var uLim int
			var lLim int
			// if the range is indicated by *, entire range needs stepping.
			// generally this is the first element of the string if at all user enters the complete range using "*"
			if tempRes[0] == "*" {
				uLim = uLimit
				lLim = lLimit
			} else if strings.Contains(tempRes[0], "-") { // check for "-" to traverse the range
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
		} else if strings.Contains(v, "-") { // when the string doesn't contain "/", next priority is "-"
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
		} else if strings.Contains(v, "*") { // last priority in symbols is given to *, as it is used as idle element.
			if len(v) == 1 {
				for i := lLimit; i <= uLimit; i++ {
					output = append(output, strconv.Itoa(i))
				}
			} else {
				output = nil
				return output, false
			}
		} else { // if the string doesn't fall in any of the if cases, it must a single number, word to define days and months
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
					} else {
						output = nil
						return output, false
					}
				}
			}
		}
	}
	// if at all user enter multiple combination values which are kind of ok and those lead to printing to a value more than once,
	// hence to remove the duplicates this function call is used.
	// This function is not really useful if the user follows cron rules, just to be extra careful this is being used.
	output = removeDuplicates(output)
	return output, true
}

// function to remove duplicates in the output generated by parseMessage
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
