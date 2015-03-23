// Test cases for OnTime.

package ontime

import (
	"testing"
	"fmt"
	)

func Test01mb(t *testing.T) {
	if MetToBin(0) != 0 {
		t.Fail()
	}
	fmt.Println("0met = 0bin")
}

func Test02mb(t *testing.T) {
	if MetToBin(50000) != 32768 {
		t.Fail()
	}
	fmt.Println("50000met = 32768bin")
}

//func TestMain(m *testing.M) {
//	fmt.Printf("The current time is: %s", time.Now())
//	os.Exit(m.Run())
//}
