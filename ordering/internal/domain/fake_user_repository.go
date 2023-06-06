package domain

import "context"

type FakeUserRepository struct {
}

func NewFakeUserRepository() *FakeUserRepository {
	return &FakeUserRepository{}
}

var _ UserRepository = (*FakeUserRepository)(nil)

func (r *FakeUserRepository) Authorize(ctx context.Context, userID string) error {
	return nil
}
