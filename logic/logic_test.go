package logic

import (
	"fmt"
	"testing"
)

func TestProcessCronString(t *testing.T) {
	tests := []struct {
		name string
		args string
		err  error
	}{
		{
			name: "Process Cron String: Happy Path",
			args: "*/15 0 1,15 * 1-5 /usr/bin/find",
			err:  nil,
		},
		{
			name: "Process Cron String: Invalid String",
			args: "find",
			err:  fmt.Errorf("insufficient number of arguments in cron string"),
		},
		{
			name: "Process Cron String: Invalid Minute",
			args: "70 0 1,15 * 1-5 /usr/bin/find",
			err:  fmt.Errorf("value 70 is outside of the expected range"),
		},
		{
			name: "Process Cron String Invalid Hour",
			args: "*/15 25 1,15 * 1-5 /usr/bin/find",
			err:  fmt.Errorf("value 25 is outside of the expected range"),
		},
		{
			name: "Process Cron String Invalid Day Of The Month",
			args: "*/15 0 1,39 * 1-5 /usr/bin/find",
			err:  fmt.Errorf("value 39 is outside of the expected range"),
		},
		{
			name: "Process Cron String Invalid Month",
			args: "*/15 0 1,15 13 1-5 /usr/bin/find",
			err:  fmt.Errorf("value 13 is outside of the expected range"),
		},
		{
			name: "Process Cron String Invalid Day Of The Week",
			args: "*/15 0 1,15 * 1-9 /usr/bin/find",
			err:  fmt.Errorf("value 8 is outside of the expected range"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualErr := ProcessCronString(tt.args)
			if actualErr != nil && tt.err != nil && actualErr.Error() != tt.err.Error() {
				t.Errorf("ProcessCronString() got = %v, want %v", actualErr, tt.err)
			}
		})
	}
}
