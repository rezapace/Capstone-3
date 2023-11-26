package repository

import (
	"Ticketing/entity"
	"context"
	"errors"

	"gorm.io/gorm"
)

type RegistrationRepository struct {
	db *gorm.DB
}

func NewRegistrationRepository(db *gorm.DB) *RegistrationRepository {
	return &RegistrationRepository{
		db: db,
	}
}

func (r *RegistrationRepository) CreateUser(ctx context.Context, name, email, number, password string) (*entity.User, error) {
	newUser := entity.NewUser(name, email, number, password)

	// Pastikan email unik
	existingUser, err := r.GetByEmail(ctx, email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user with that email already exists")
	}

	// Simpan pengguna baru ke database
	if err := r.db.WithContext(ctx).Create(newUser).Error; err != nil {
		return nil, err
	}

	return newUser, nil
}

func (r *RegistrationRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
