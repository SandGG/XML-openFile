package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type student struct {
	Name   string `xml:"name"`
	Number int    `xml:"number"`
	Quali  quali  `xml:"quali"`
}

type quali struct {
	First  int `xml:"first"`
	Second int `xml:"second"`
	Third  int `xml:"third"`
}

//array of students
type students struct {
	StdArray []student `xml:"student"`
}

func main() {
	file, err := os.Open("./student.xml")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	readXML(file)
}

func readXML(file *os.File) {
	var studentsVar students
	var b, errC = ioutil.ReadAll(file)
	if errC != nil {
		log.Fatal(errC)
	}
	xml.Unmarshal(b, &studentsVar)

	for i := 0; i < len(studentsVar.StdArray); i++ {
		fmt.Println("Student name: ", studentsVar.StdArray[i].Name)
		fmt.Println("Student number: ", studentsVar.StdArray[i].Number)
		fmt.Println("First qualification: ", studentsVar.StdArray[i].Quali.First)
		fmt.Println("Second qualification: ", studentsVar.StdArray[i].Quali.Second)
		fmt.Println("Third qualification: ", studentsVar.StdArray[i].Quali.Third)
	}
}
