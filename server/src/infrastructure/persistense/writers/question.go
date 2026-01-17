package writers

import (
	"adaptivetesting/src/domain/aggregates/question"
	"context"

	"github.com/jackc/pgx/v5"
)

func SaveQuestion(ctx context.Context, txn pgx.Tx, question *question.Question) error {
	model := question.ToPersistense()

	_, err := txn.Exec(ctx, `
		insert into 
		questions (id, by_topic_id, by_course_id, text)
		values ($1, $2, $3, $4)
		on conflict (id)
		do update set 
			by_topic_id = excluded.by_topic_id,
			by_course_id = excluded.by_course_id,
			text = excluded.text;
		`,
		model.ID,
		model.ByTopicID,
		model.ByCourseID,
		model.Text,
	)

	if err != nil {
		return err
	}

	for _, answer := range model.Answers {
		_, err := txn.Exec(ctx, `
			insert into 
			answers (by_question_id, text, is_correct, serial_number)
			values ($1, $2, $3, $4)
			on conflict (id)
			do update set 
				text = excluded.text,
				is_correct = excluded.is_correct,
				serial_number = excleded.serial_number;
			`,
			model.ID,
			answer.Text,
			answer.IsCorrect, answer.SerialNumber,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
