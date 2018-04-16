package entities

import (
	"strings"
)

type Course struct {
	ID      	int
	Title   	string
	Content 	string
	Host		string
	HostURL		string
	URL			string
	URLApi  	string
	Price   	string
	Duration	string
	Language	string
	SkillLvl	string
	IDUdemy     int
}

var FoundCourses = []Course{}

func IsContain(title string, originalTitle string) (contains bool) {
	contains = false
	arr := strings.Split(title, " ")
	arr2 := strings.Split(originalTitle, " ")

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr[i] == arr2[j] {
				return true
			}
		}
	}
	return
}
