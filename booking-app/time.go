package main

import ("fmt", "time")

func main() {
	fmt.Println("Welcome to Time Study")
//Study during API creation		
	presentTime := time.Now()
	fmt.Println("The time.Now() is ", presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //we use "01-02-2006 15:04:05 Monday" because it is provided in documentation

	createdTime := time.Date(20202, time.November, 10, 13, 43, 10, 0 , time.UTC) //year time.Month, date, hour, min, sec
	fmt.Println(createdTime)


	//User "go env" in command line for study


	//BUILD FOR DIFFERENT devices
	/* 
	GOOS='windows' go build main.go
	GOOD='darwin' go build main.go
	GOOD='linux' go build main.go
	*/
}