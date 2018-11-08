package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var arr [4]int32

func printArr() {
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4

	for idx, val := range arr {
		fmt.Println(idx, val)
	}
}

func makeSlice() {
	var slice []byte
	fmt.Println(slice)

	slice = append(slice, 'a')
	fmt.Println(slice)
}

func makeMap() {
	var m = make(map[string]int)
	m["1"] = 1
	m["2"] = 1
	m["3"] = 1
	m["1"] = 2
	fmt.Println(m)

	delete(m, "1")
	fmt.Println(m)

	for key, val := range m {
		fmt.Println(key, val)
	}
}

type Movie struct {
	Title string `json:"title"`
	Year int `json:"released"`
	Color bool `ison:"color"`
	Actors []string `json:"actors"`
}

func makeJson() {
	movies := []Movie {
		{ Title: "title1", Year: 1, Color: true, Actors: []string{ "1" } },
		{ Title: "title2", Year: 2, Color: true, Actors: []string{ "2" } },
		{ Title: "title3", Year: 3, Color: false, Actors: []string{ "3" } },
	}

	data, err := json.Marshal(movies)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(0)
	}
	fmt.Printf("%s\n", data)

	var copy []Movie
	err = json.Unmarshal(data, &copy)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(0)
	}
	fmt.Printf("%v\n", copy)	

	data, err = json.MarshalIndent(movies, "/", "    ")
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(0)
	}
	fmt.Printf("%s\n", data)
}

