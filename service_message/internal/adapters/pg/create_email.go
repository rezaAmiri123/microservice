package pg

import (
	"context"
	"fmt"

	// sqlPg "github.com/lib/pq"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/service_message/internal/domain/message"
)

const createEmail = `INSERT INTO emails (user_id, from_email, to_email, subject, body) VALUES ($1, $2, $3, $4, $5) RETURNING *`

func (r *PGMessageRepository) CreateEmail(ctx context.Context, arg *message.CreateEmailParams) (*message.Email, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGMessageRepository.CreateEmail")
	defer span.Finish()

	var m message.Email
	if err := r.DB.QueryRowxContext(
		ctx,
		createEmail,
		&arg.UserID,
		&arg.From,
		// sqlPg.Array(&arg.To),
		&arg.To,
		&arg.Subject,
		&arg.Body,
		// ).Scan(&res); err != nil {
	).StructScan(&m); err != nil {
		return nil, fmt.Errorf("postgres connot create email: %w", err)
	}

	return &m, nil
}
