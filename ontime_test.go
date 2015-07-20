// Test cases for OnTime.

package ontime

import (
	"os"
	"time"
	"testing"
)

func TestBinTime(t *testing.T) {
	now, _ = time.Parse("3:04pm", "12:34pm")
	if binHour() != 8 {
		t.Error("For binHour, expected 8, got", binHour())
	}
	if binMinute() != 6 {
		t.Error("For binMinute, expected 6, got", binMinute())
	}
}

func TestMain(m *testing.M) {
	//fmt.Printf("The current time is: %s", time.Now())
	os.Exit(m.Run())
}
