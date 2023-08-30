package dto

type CronDescRequest struct {
	Minute        string
	Hour          string
	DayOfTheMonth string
	Month         string
	DayOfTheWeek  string
	Command       string
}

type CronDescResponse struct {
	Minute        string
	Hour          string
	DayOfTheMonth string
	Month         string
	DayOfTheWeek  string
	Command       string
}
