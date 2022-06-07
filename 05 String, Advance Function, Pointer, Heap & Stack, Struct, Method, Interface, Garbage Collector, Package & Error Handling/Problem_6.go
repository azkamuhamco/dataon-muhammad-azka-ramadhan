package main
import (
	"fmt"
	"unicode"
)

type student struct {
	name	   string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""

	for i:=0; i<len(s.name); i++ {
		if isUpper(s.name) {
			nameEncode += string(int('Z') - int(s.name[i]) + int('A'))
		} else if isLower(s.name) {
			nameEncode += string(int('z') - int(s.name[i]) + int('a'))
		} else {
			nameEncode += string(int(s.name[i]))
		}
	}
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""

	for i:=0; i<len(s.nameEncode); i++ {
		if isUpper(s.nameEncode) {
			nameDecode += string(int('Z') - int(s.nameEncode[i]) + int('A'))
		} else if isLower(s.nameEncode) {
			nameDecode += string(int('z') - int(s.nameEncode[i]) + int('a'))
		} else {
			nameDecode += string(int(s.nameEncode[i]))
		}
	}
	return nameDecode
}

func isUpper(s string) bool {
    for _, r := range s {
        if !unicode.IsUpper(r) && unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

func isLower(s string) bool {
    for _, r := range s {
        if !unicode.IsLower(r) && unicode.IsLetter(r) {
            return false
        }
    }
    return true
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
