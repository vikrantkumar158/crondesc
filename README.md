
# CronDesc

---

This application parses a standard cron string and expands each field to show the times at which it will run.

## Usage

```bash
go run main.go "*/15 0 1,15 * 1-5 /usr/bin/find"
```
<p align="center">OR</p>

```bash
go build
./crondesc "*/15 0 1,15 * 1-5 /usr/bin/find"
```

The input should be a standard cron string with five time fields (minute, hour, day of month, month, and day of week) plus a command.

For example, the input `*/15 0 1,15 * 1-5 /usr/bin/find` will yield the following output:

```
minute        0 15 30 45
hour          0
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   1 2 3 4 5
command       /usr/bin/find
```

## Features

* Parses each field of a standard cron string and expands it to a list of times at which it will run.
* Handles the '*' format (e.g., "*" for every month).
* Handles the '*/step' format (e.g., "*/15" for every 15 minutes).
* Handles ranges (e.g., "1-5" for 1, 2, 3, 4, 5).
* Handles comma-separated values (e.g., "1,15" for 1 and 15).
* Handles single values (e.g., "0" for 0th hour)

## Limitations

* Special time strings such as "@yearly" are not handled.
* Validation for the given input is constrained to check whether the cron string has exactly six fields and whether the values are in correct ranges. Advanced validation of syntax is not included in this application.

## Setup Instructions
To use this application, make sure you have Go installed on your machine.

* On Mac OS, you can install Go with Homebrew: `brew install go`.
* On Linux, you can install Go using `sudo apt-get install golang-go` on Ubuntu or the equivalent command for your specific distribution.

## Acknowledgement

The crondesc is a basic command line application.
