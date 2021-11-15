package main

import (
	"fmt"
	"time"
)

func main() {
	unix := time.Unix(1636444068, 0).AddDate(0, 0, -1)
	format := unix.Format("2006-01-02 15:04:05")
	fmt.Println(format)
	before := unix.Before(time.Now())
	fmt.Println(before)
}
