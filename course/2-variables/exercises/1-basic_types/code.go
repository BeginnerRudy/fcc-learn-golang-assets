package main

import "fmt"

func main() {
	// initialize variables here
	var smsSendingLimit int = 9
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
