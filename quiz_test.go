package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultflags(t *testing.T) {
	expectedFileName := "problems.csv"
	expectedTimer := 30
	actualFileName, actualTimer := readFlags()
	assert.Equal(t, expectedFileName, actualFileName)
	assert.Equal(t, expectedTimer, actualTimer)
}

func TestQuizArrays(t *testing.T) {
	expectedQuestions := []string{"5+5", "7+3", "1+1", "8+3", "1+2", "8+6",
		"3+1", "1+4", "5+1", "2+3", "3+3", "2+4", "5+2"}
	expectedAnswers := []string{"10", "10", "2", "11", "3", "14",
		"4", "5", "6", "5", "6", "6", "7"}
	actualQuestions, actualAnswers := getQuestions("problems.csv")
	assert.Equal(t, expectedQuestions, actualQuestions)
	assert.Equal(t, expectedAnswers, actualAnswers)
}

func TestEndGame(t *testing.T) {
	expected := "You scored 0 out of 0\n"
	actual := endGame(0, 0)
	assert.Equal(t, actual, expected)
}
