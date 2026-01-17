package writers

import (
	"adaptivetesting/src/application/interfaces"
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/aggregates/group"
	"adaptivetesting/src/domain/aggregates/question"
	"adaptivetesting/src/domain/aggregates/topic"
	"adaptivetesting/src/domain/aggregates/user"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Writer struct {
	pool *pgxpool.Pool
}

type TransactionWriter struct {
	txn pgx.Tx
	ctx context.Context
}

func (this *TransactionWriter) SaveUser(user *user.User) error {
	return SaveUser(this.ctx, this.txn, user)
}

func (this *TransactionWriter) SaveCourse(course *course.Course) error {
	return SaveCourse(this.ctx, this.txn, course)
}

func (this *TransactionWriter) SaveQuestion(question *question.Question) error {
	return SaveQuestion(this.ctx, this.txn, question)
}

func (this *TransactionWriter) SaveGroup(group *group.Group) error {
	return SaveGroup(this.ctx, this.txn, group)
}

func (this *TransactionWriter) SaveTopic(topic *topic.Topic) error {
	return SaveTopic(this.ctx, this.txn, topic)
}

func (txm *Writer) Execute(
	ctx context.Context,
	fn func(w interfaces.ITransactionWriter) error,
) (err error) {
	tx, err := txm.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	if err := fn(&TransactionWriter{ctx: ctx, txn: tx}); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func NewWriter(pool *pgxpool.Pool) *Writer {
	return &Writer{pool: pool}
}
