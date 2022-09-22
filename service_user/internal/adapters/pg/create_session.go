package pg

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/service_user/internal/domain/user"
)

const createSession = `INSERT INTO sessions 
						(session_id, username, refresh_token, user_agent, client_ip, is_blocked, expires_at) 
						VALUES ($1, $2, $3, $4, $5, $6, $7)`

// const createSession = `INSERT INTO sessions
//
//	(session_id, username, refresh_token, user_agent, client_ip, is_blocked, expires_at)
//	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING
//	(session_id, username, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at, updated_at)`
const getSessionByID = `SELECT 
								session_id, username, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at, updated_at 
							FROM sessions
								WHERE session_id = $1`

func (r *PGUserRepository) CreateSession(ctx context.Context, arg *user.CreateSessionParams) (*user.Session, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGUserRepository.CreateSession")
	defer span.Finish()

	if err := r.DB.QueryRowxContext(
		ctx,
		createSession,
		arg.SessionID,
		arg.Username,
		arg.RefreshToken,
		arg.UserAgent,
		arg.ClientIp,
		arg.IsBlocked,
		arg.ExpiresAt,
		// ).Scan(&res); err != nil {
	).Err(); err != nil {
		return nil, fmt.Errorf("postgres connot create session: %w", err)
	}

	s := &user.Session{}
	if err := r.DB.GetContext(ctx, s, getSessionByID, arg.SessionID); err != nil {
		return nil, fmt.Errorf("postgres connot get session: %w", err)
	}
	return s, nil
}
