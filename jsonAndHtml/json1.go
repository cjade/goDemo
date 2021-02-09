package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
电影评论
*/
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	endata := encode()
	decode(endata)
	fmt.Printf("%s \n", endata)
}

// 解密
func decode(data []byte) {
	// var deData []struct {
	// 	Title string
	// }

	deData := []Movie{}
	if err := json.Unmarshal(data, &deData); err != nil {
		log.Fatalf("json Unmarshal err %s", err)
	}

	fmt.Printf("%v \n", deData)
}

// 编码
func encode() []byte {
	// data, err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		log.Fatalf("json Marshal err %s", err)
	}

	return data
}
