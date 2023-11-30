package service

import (
	"Ticketing/entity"
	"context"
)

type NotificationUsecase interface {
	GetAllNotification(ctx context.Context) ([]*entity.Notification, error)
	CreateNotification(ctx context.Context, Notification *entity.Notification) error
}

type NotificationRepository interface {
	GetAllNotification(ctx context.Context) ([]*entity.Notification, error)
	CreateNotification(ctx context.Context, Notification *entity.Notification) error
}

type NotificationService struct {
	Repository NotificationRepository
}

func NewNotificationService(Repository NotificationRepository) *NotificationService {
	return &NotificationService{Repository: Repository}
}

// Get All Notification ketika di get maka status notifikasi akan berubah menjadi true
func (s *NotificationService) GetAllNotification(ctx context.Context) ([]*entity.Notification, error) {
	return s.Repository.GetAllNotification(ctx)
}

// func untuk create notification
func (s *NotificationService) CreateNotification(ctx context.Context, Notification *entity.Notification) error {
	return s.Repository.CreateNotification(ctx, Notification)
}