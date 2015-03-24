// Test cases for OnTime.

package main

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Printf("The current time is: %s", time.Now())
	os.Exit(m.Run())
}
