package main

import (
	"testing"
)

func TestCronParser(t *testing.T) {
	tests := []struct {
		name       string
		crn        cron
		wantString string
	}{
		{"test string 1", cron{minutes: "1-5/2", hours: "9-17", dayOfMonth: "*", month: "*", dayOfWeek: "*", command: "/usr/bin/find"},
			`minute        1 3 5
		hour          9 10 11 12 13 14 15 16 17
		day of month  1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31
		month         1 2 3 4 5 6 7 8 9 10 11 12
		day of week   0 1 2 3 4 5 6
		command       /usr/bin/find`},
		{"test string 2", cron{minutes: "*", hours: "*", dayOfMonth: "*", month: "*", dayOfWeek: "*", command: "/usr/bin/find"},
			`minute        0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 
			hour          0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23
			day of month  1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31
			month         1 2 3 4 5 6 7 8 9 10 11 12
			day of week   0 1 2 3 4 5 6
			command       /usr/bin/find`},
		{"test string 3", cron{minutes: "*/15", hours: "0", dayOfMonth: "1,15", month: "*", dayOfWeek: "1-5", command: "/usr/bin/find"},
			`minute        0 15 30 45 
			hour          0 
			day of month  1 15 
			month         1 2 3 4 5 6 7 8 9 10 11 12 
			day of week   1 2 3 4 5 
			command       /usr/bin/find`},
		{"test string 4", cron{minutes: "*/15", hours: "0", dayOfMonth: "1,15", month: "*", dayOfWeek: "1-5", command: "/usr/bin/find"},
			`Invalid cron message, Please check and retry!`},
		{"test string 5", cron{minutes: "1-5/2,9-14/3", hours: "0", dayOfMonth: "1,15", month: "*", dayOfWeek: "1-5", command: "/usr/bin/find"},
			`minute        1 3 5 9 12 
			hour          0
			day of month  1 15
			month         1 2 3 4 5 6 7 8 9 10 11 12
			day of week   1 2 3 4 5
			command       /usr/bin/find`},
		{"test string 6", cron{minutes: "*/15", hours: "0", dayOfMonth: "1,15", month: "JAN", dayOfWeek: "1-5", command: "/usr/bin/find"},
			`minute        0 15 30 45 
			hour          0
			day of month  1 15
			month         JAN
			day of week   1 2 3 4 5
			command       /usr/bin/find`},
		{"test string 7", cron{minutes: "*/15", hours: "0/", dayOfMonth: "1,15", month: "JAN", dayOfWeek: "1-5", command: "/usr/bin/find"},
			`Invalid cron message, Please check and retry!`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotString := tt.crn.parseCron()

			if gotString == tt.wantString {
				t.Errorf("CronParser() gotString = %v, wantString %v", gotString, tt.wantString)
				return
			}
		})
	}

}
