package service

//NOTE :
// FOLDER INI UNTUK MENANGANI LOGIC DAN MEMANGGIL REPOSITORY
import (
	"context"

	"Ticketing/entity"
)

// interface untuk service
// untuk memanngil repository
type UserUsecase interface {
	GetAll(ctx context.Context) ([]*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	GetUserByID(ctx context.Context, id int64) (*entity.User, error)
	Delete(ctx context.Context, id int64) error
	UpdateUserSelf(ctx context.Context, user *entity.User) error
	UpdateUserBalance(ctx context.Context, user *entity.User) error
}

// interface untuk repository
// untuk memanggil repository
// GetAll = untuk menampilkan semua data user, dan itu harus sama dengan yang ada di repository
type UserRepository interface {
	GetAll(ctx context.Context) ([]*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	GetUserByID(ctx context.Context, id int64) (*entity.User, error)
	Delete(ctx context.Context, id int64) error
	UpdateUserSelf(ctx context.Context, user *entity.User) error
	UpdateUserBalance(ctx context.Context, user *entity.User) error
}

// code di line 23 merupakan dependency injection, karena repository tidak langsung di panggil.
// karena repository dipanggil melalui code pada line 18
type UserService struct {
	repository UserRepository
}

// func untuk UserRepository
func NewUserService(repository UserRepository) *UserService {
	return &UserService{repository}
}

// func dibawah ini untuk type user usecase
// ini untuk menampilkan data user
// untuk memanggil repository
func (s *UserService) GetAll(ctx context.Context) ([]*entity.User, error) {
	return s.repository.GetAll(ctx)
}

// func dibawah ini untuk type user usecase
// ini untuk membuat data user
func (s *UserService) CreateUser(ctx context.Context, user *entity.User) error {
	return s.repository.CreateUser(ctx, user)
}

// untuk update data user
func (s *UserService) UpdateUser(ctx context.Context, user *entity.User) error {
	return s.repository.UpdateUser(ctx, user)
}

// untuk get user by id
func (s *UserService) GetUserByID(ctx context.Context, id int64) (*entity.User, error) {
	return s.repository.GetUserByID(ctx, id)
}

// untuk delete by id
func (s *UserService) Delete(ctx context.Context, id int64) error {
	return s.repository.Delete(ctx, id)
}

// Update User Self
func (s *UserService) UpdateUserSelf(ctx context.Context, user *entity.User) error {
	return s.repository.UpdateUser(ctx, user)
}

// func update saldo user by id
func (s *UserService) UpdateUserBalance(ctx context.Context, user *entity.User) error {
	return s.repository.UpdateUserBalance(ctx, user)
}