package main

import (
	"adaptivetesting/src/application/dto/requests"
	question_usercases "adaptivetesting/src/application/usecases/question"
	"adaptivetesting/src/domain/aggregates/identificators"
	"adaptivetesting/src/domain/aggregates/question"
	"adaptivetesting/src/domain/aggregates/topic"
	"adaptivetesting/src/domain/aggregates/user"
	"encoding/json"
	"fmt"
)

type RepoQ struct{}

func (this *RepoQ) Save(question *question.Question) error {
	return nil
}

func (this *RepoQ) GetByID(id identificators.QuestionID) (*question.Question, error) {
	return &question.Question{}, nil
}

type RepoT struct{}

func (this *RepoT) Save(*topic.Topic) error {
	return nil
}

func (this *RepoT) GetByID(id identificators.TopicID) (*topic.Topic, error) {
	return &topic.Topic{}, nil
}

func main() {
	CreateQuestion, err := question_usercases.Fabric(&(RepoT{}), &(RepoQ{}))
	if err != nil {
		panic(err)
	}

	current_user, err := user.NewTeacher("Penis", "12345")
	if err != nil {
		panic(err)
	}

	created_question, err := CreateQuestion.Execute(current_user, &requests.CreateQuestionRequest{
		Text:    "Question Penisaaa",
		Answers: []requests.AnswerRequest{{Text: ".!.kdhvouidmucdifynvdsad", IsCorrect: true}, {Text: "Not Penisofsndfiufhnkduvsduf", IsCorrect: false}},
	})
	if err != nil {
		panic(err)
	}
	jsoned, err := json.Marshal(created_question)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsoned))
}
