package main

import (
	"fmt"
	"math"
)

type Student struct {
	name  string
	score int
}

func getStudentInfo(numStudents int) []Student {
	students := make([]Student, numStudents)

	for i := 0; i < numStudents; i++ {
		fmt.Printf("Enter name for student %d: ", i+1)
		fmt.Scanln(&students[i].name)

		var score int
		for {
			fmt.Printf("Enter score for student %d: ", i+1)
			_, err := fmt.Scanln(&score)
			if err != nil {
				fmt.Println("Invalid input. Please enter an integer for the score.")
				continue
			}
			break
		}
		students[i].score = score
	}

	return students
}

func calculateAverageScore(students []Student) float64 {
	totalScore := 0
	for _, student := range students {
		totalScore += student.score
	}
	averageScore := float64(totalScore) / float64(len(students))
	return averageScore
}

func findMinMaxScore(students []Student) (string, int, string, int) {
	minScore := math.MaxInt32
	maxScore := math.MinInt32
	var minStudentName, maxStudentName string

	for _, student := range students {
		if student.score < minScore {
			minScore = student.score
			minStudentName = student.name
		}
		if student.score > maxScore {
			maxScore = student.score
			maxStudentName = student.name
		}
	}

	return minStudentName, minScore, maxStudentName, maxScore
}

func main() {
	numStudents := 5
	students := getStudentInfo(numStudents)

	minStudentName, minScore, maxStudentName, maxScore := findMinMaxScore(students)

	fmt.Println("Minimum score:", minStudentName, minScore)
	fmt.Println("Maximum score:", maxStudentName, maxScore)
}
