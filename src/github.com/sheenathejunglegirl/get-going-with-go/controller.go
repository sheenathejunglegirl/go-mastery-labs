package main

import (
  "database/sql"
  "log"
  _ "github.com/mattn/go-sqlite3"
)

func getUnansweredQuestion(db *sql.DB) (Question, error) {
  rows, err := db.Query("select id, description, wrong_answer1, wrong_answer2, right_answer from questions where answered_correctly = 0;")
  handleError(err)
  defer rows.Close()

  questions := make([]Question, 0)
  for rows.Next() {
    question := Question{}
    rows.Scan(&question.Id, &question.Description, &question.WrongAnswer1, &question.WrongAnswer2, &question.RightAnswer)
    log.Print(question)
    questions = append(questions, question)
  }
  return questions[0], nil
}

func markAsAnsweredCorrectly(db *sql.DB, id int) (error) {
  return nil
}

func totalQuestionsAnsweredCorrectly(db *sql.DB) (int, error) {
  var count int
  db.QueryRow("SELECT count(*) from questions where answered_correctly = 1;").Scan(&count)

  return count, nil
}

func totalQuestions(db *sql.DB) (int, error) {
  var count int
  db.QueryRow("SELECT count(*) from questions;").Scan(&count)

  return count, nil
}
