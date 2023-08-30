package logic

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"crondesc/consts"
	"crondesc/dto"
)

func ProcessCronString(cronString string) error {
	var err error

	req, err := convertCronStringToCronDescRequestDto(cronString)
	if err != nil {
		fmt.Printf("Error encountered for Request Conversion: %v\n", err.Error())
		return err
	}

	res := dto.CronDescResponse{Command: req.Command}
	res.Minute, err = expandField(req.Minute, consts.MinuteLowerUpperLimit[0], consts.MinuteLowerUpperLimit[1])
	if err != nil {
		fmt.Printf("Error encountered for Minute Conversion: %v\n", err.Error())
		return err
	}

	res.Hour, err = expandField(req.Hour, consts.HourLowerUpperLimit[0], consts.HourLowerUpperLimit[1])
	if err != nil {
		fmt.Printf("Error encountered for Hour Conversion: %v\n", err.Error())
		return err
	}

	res.DayOfTheMonth, err = expandField(req.DayOfTheMonth, consts.DayOfMonthLowerUpperLimit[0], consts.DayOfMonthLowerUpperLimit[1])
	if err != nil {
		fmt.Printf("Error encountered for Day Of The Month Conversion: %v\n", err.Error())
		return err
	}

	res.Month, err = expandField(req.Month, consts.MonthLowerUpperLimit[0], consts.MonthLowerUpperLimit[1])
	if err != nil {
		fmt.Printf("Error encountered for Month Conversion: %v\n", err.Error())
		return err
	}

	res.DayOfTheWeek, err = expandField(req.DayOfTheWeek, consts.DayOfWeekLowerUpperLimit[0], consts.DayOfWeekLowerUpperLimit[1])
	if err != nil {
		fmt.Printf("Error encountered for Day Of The Week Conversion: %v\n", err.Error())
		return err
	}

	printCronDescription(res)
	return nil
}

func expandField(cronValue string, lowerLimit, upperLimit int) (cronDesc string, err error) {
	isValid := func(val int) error {
		if val < lowerLimit || val > upperLimit {
			return errors.New("is outside of the expected range")
		}
		return nil
	}

	switch {
	case cronValue == "*":
		// Every value in range if wildcard was used
		for i := lowerLimit; i <= upperLimit; i++ {
			cronDesc = cronDesc + strconv.Itoa(i) + " "
		}
	case strings.Contains(cronValue, "*/"):
		// Every nth value for the */n format
		stepStr := strings.Split(cronValue, "*/")[1]
		step, _ := strconv.Atoi(stepStr)
		for i := lowerLimit; i <= upperLimit; i += step {
			if err := isValid(i); err != nil {
				return cronDesc, fmt.Errorf("value %d %v", i, err)
			}
			cronDesc = cronDesc + strconv.Itoa(i) + " "
		}
	case strings.Contains(cronValue, "-"):
		// Range of values for the a-b format
		rangeParts := strings.Split(cronValue, "-")
		start, _ := strconv.Atoi(rangeParts[0])
		end, _ := strconv.Atoi(rangeParts[1])
		for i := start; i <= end; i++ {
			if err := isValid(i); err != nil {
				return cronDesc, fmt.Errorf("value %d %v", i, err)
			}
			cronDesc = cronDesc + strconv.Itoa(i) + " "
		}
	case strings.Contains(cronValue, ","):
		// List of values for the a,b,c format
		listParts := strings.Split(cronValue, ",")
		for _, part := range listParts {
			val, _ := strconv.Atoi(part)
			if err := isValid(val); err != nil {
				return cronDesc, fmt.Errorf("value %d %v", val, err)
			}
			cronDesc = cronDesc + strconv.Itoa(val) + " "
		}
	default:
		// Single value
		val, _ := strconv.Atoi(cronValue)
		if err := isValid(val); err != nil {
			return cronDesc, fmt.Errorf("value %d %v", val, err)
		}
		cronDesc = cronDesc + strconv.Itoa(val) + " "
	}
	return
}

func printCronDescription(res dto.CronDescResponse) {
	fmt.Printf("minute\t\t\t %s\n", res.Minute)
	fmt.Printf("hour\t\t\t %s\n", res.Hour)
	fmt.Printf("day of month\t\t %s\n", res.DayOfTheMonth)
	fmt.Printf("month\t\t\t %s\n", res.Month)
	fmt.Printf("day of week\t\t %s\n", res.DayOfTheWeek)
	fmt.Printf("command\t\t\t %s\n", res.Command)
}

func convertCronStringToCronDescRequestDto(cronString string) (dto.CronDescResponse, error) {
	res := dto.CronDescResponse{}

	cron := strings.Split(cronString, " ")
	if len(cron) < 6 {
		return res, fmt.Errorf("insufficient number of arguments in cron string")
	}

	res.Minute = cron[0]
	res.Hour = cron[1]
	res.DayOfTheMonth = cron[2]
	res.Month = cron[3]
	res.DayOfTheWeek = cron[4]
	res.Command = cron[5]
	return res, nil
}
