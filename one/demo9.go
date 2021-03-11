// 结构体
package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	Dob       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	aaaa()
}

func aaaa() {
	var dilbert Employee
	dilbert.Salary = 5000
	dilbert.Name = "tom"
	position := &dilbert.Position
	*position = "Senior " + *position
	dilbert.Dob.Unix()

	fmt.Println(dilbert)
}
