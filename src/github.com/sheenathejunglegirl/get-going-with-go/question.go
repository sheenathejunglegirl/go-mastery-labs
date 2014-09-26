package main

import(
  "fmt"
)

type QuizQuestionPage struct {
  UnansweredQuestion Question
  Total int
  TotalAnsweredCorrectly int
}

type Question struct {
  Id  int
  Description string
  WrongAnswer1 string
  WrongAnswer2 string
  RightAnswer string
}

func (question *Question) String() string {
    return fmt.Sprintf("%s - %s", question.Description, question.RightAnswer)
}
