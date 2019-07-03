package main

import (
	"fmt"
	"strings"
)

type fbRule struct {
	Number   int
	Response string
}

type fbResponse struct {
	Result         string
	Message        string
	NextQuestion   string
	Rules          []fbRule
	Numbers        []int
	Grade          string
	ElapsedSeconds int
}

func main() {
	for nextQues := firstQuestion("https://api.noopschallenge.com/fizzbot/questions/1"); nextQues != ""; {
		nextQues = doChallenge("https://api.noopschallenge.com" + nextQues)
	}

}

func firstQuestion(url string) string {
	var result fbResponse
	jsonGet(url, &result)

	fmt.Printf("Message:\n--------------------\n%v\n--------------------\n\n", result.Message)

	ansResult := postAnswer(url, "Go")

	fmt.Printf("Result: %v\n", ansResult.Result)
	fmt.Printf("Message:\n--------------------\n%v\n--------------------\n\n", ansResult.Message)
	return ansResult.NextQuestion
}

func answer(rules []fbRule, numbers []int) string {
	ans := ""
	for _, i := range numbers {
		a := ""
		for _, r := range rules {
			if i%r.Number == 0 {
				a += r.Response
			}
		}
		if a == "" {
			a = fmt.Sprintf("%d", i)
		}
		ans += " " + a
	}
	return strings.Trim(ans, " ")
}

func postAnswer(url string, ans string) fbResponse {
	fmt.Printf("Answer: %v\n\n", ans)

	content := jsonType{
		"answer": ans,
	}

	var result fbResponse
	jsonPost(url, content, &result)
	return result
}

func doChallenge(url string) string {
	var result fbResponse
	jsonGet(url, &result)

	fmt.Printf("--------------------\n%v\nRules: %v\nNumbers: %v\n--------------------\n\n", result.Message, result.Rules, result.Numbers)

	ansResult := postAnswer(url, answer(result.Rules, result.Numbers))

	fmt.Printf("Result: %v\n", ansResult.Result)
	if ansResult.Result == "interview complete" {
		fmt.Printf("--------------------\nGrade: %v\nTime: %v\n%v\n--------------------\n\n", ansResult.Grade, ansResult.ElapsedSeconds, ansResult.Message)
	} else {
		fmt.Printf("--------------------\n%v\n--------------------\n\n", ansResult.Message)
	}
	return ansResult.NextQuestion
}
