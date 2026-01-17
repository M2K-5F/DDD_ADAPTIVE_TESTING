package readers

import (
	"adaptivetesting/src/domain/aggregates/question"
	"adaptivetesting/src/domain/identificators"
	"adaptivetesting/src/infrastructure/persistense/models"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type QuestionReader struct {
	pool *pgxpool.Pool
}

func (this *QuestionReader) GetByID(ctx context.Context, questionID identificators.QuestionID) (*question.Question, error) {
	rows, err := this.pool.Query(ctx, `
		select q.id, q.by_topic_id, q.by_course_id, q.text, a.text, a.is_correct, a.serial_number 
		from questions q
		inner join answers a on q.id = a.by_question_id
		where q.id = $1
		`, questionID,
	)
	if err != nil {
		return nil, err
	}

	var questionRows models.QuestionRows
	var answers []*question.AnswerCreate

	for rows.Next() {
		var answerRow models.AnswerRows
		err := rows.Scan(&questionRows.ID, &questionRows.ByTopicID, &questionRows.ByCourseID, &questionRows.Text, &answerRow.Text, &answerRow.IsCorrect, &answerRow.SerialNumber)
		if err != nil {
			return nil, err
		}
		answers = append(answers,
			&question.AnswerCreate{
				Text:         answerRow.Text,
				IsCorrect:    answerRow.IsCorrect,
				SerialNumber: answerRow.SerialNumber,
			})
	}

	questionEntity, err := question.Fabric.Recover(questionRows.ID, questionRows.ByCourseID, questionRows.ByTopicID, questionRows.Text, answers)
	if err != nil {
		return nil, err
	}

	return questionEntity, nil
}
