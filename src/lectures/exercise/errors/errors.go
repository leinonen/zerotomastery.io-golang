//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hours, minutes, seconds int
}

type TimeParseError struct {
	msg, input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

func ParseTime(str string) (Time, error) {
	parts := strings.Split(str, ":")
	if len(parts) < 3 {
		return Time{}, &TimeParseError{"unable to parse time string", str}
	}
	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return Time{}, &TimeParseError{"parse hours failed", str}
	}
	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return Time{}, &TimeParseError{"parse minutes failed", str}
	}
	seconds, err := strconv.Atoi(parts[2])
	if err != nil {
		return Time{}, &TimeParseError{"parse seconds failed", str}
	}
	if hours > 24 || hours < 0 {
		return Time{}, &TimeParseError{"invalid hours", str}
	}
	if minutes > 59 || minutes < 0 {
		return Time{}, &TimeParseError{"invalid minutes", str}
	}
	if seconds > 59 || seconds < 0 {
		return Time{}, &TimeParseError{"invalid seconds", str}
	}
	return Time{
		hours:   int(hours),
		minutes: int(minutes),
		seconds: int(seconds),
	}, nil
}
