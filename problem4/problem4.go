package main

import "fmt"

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""
	// your code here

	for _, value := range s.name {
		if value >= 97 && value <= 122 {
			value += 3
			if value > 122 {
				value -= 26
			}

		}

		nameEncode += string(value)
	}
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""
	// your code here
	for _, value := range s.nameEncode {
		if value >= 97 && value <= 122 {
			value -= 3
			if value < 97 {
				value += 26
			}

		}

		nameDecode += string(value)
	}
	return nameDecode
}

func main() {
	var menu int
	var a = student{}
	var c Chiper = &a
	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	fmt.Scan(&menu)
	if menu == 1 {
		fmt.Print("\nInput Student's Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of Student's Name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student's Encode Name : ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of Student's Name " + a.nameEncode + " is : " + c.Decode())
	} else {
		fmt.Println("Wrong input name menu")
	}
}
