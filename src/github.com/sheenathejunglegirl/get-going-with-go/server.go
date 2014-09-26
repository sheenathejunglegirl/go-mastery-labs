package main

import (
  "github.com/go-martini/martini"
  "github.com/codegangsta/martini-contrib/render"
  "database/sql"
  "os"
  "log"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
  os.Remove("./questions.db")

  db, err := sql.Open("sqlite3", "./questions.db")
  handleError(err)
  defer db.Close()

  sqlStmt := `
  create table questions (id integer not null primary key, description text, wrong_answer1 text, wrong_answer2 text, right_answer text, answered_correctly boolean);
  delete from questions;
  insert into questions values(1, "description", "wrong", "wrong", "right", 0);
  insert into questions values(2, "description", "wrong", "wrong", "right", 0);
  insert into questions values(3, "description", "wrong", "wrong", "right", 1);
  insert into questions values(4, "description", "wrong", "wrong", "right", 0);
  `
  _, err = db.Exec(sqlStmt)
  handleError(err)

  m := martini.Classic()
  m.Use(render.Renderer())

  m.Get("/questions/unanswered/random", func(r render.Render) {
    question, err := getUnansweredQuestion(db)
    handleError(err)
    totalAnsweredCorrectly, err := totalQuestionsAnsweredCorrectly(db)
    handleError(err)
    total, err := totalQuestions(db)
    handleError(err)
    questionPage := QuestionPage{question, total, totalAnsweredCorrectly}
    r.HTML(200, "question", questionPage)
  })

  m.Run()
}

func handleError(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
