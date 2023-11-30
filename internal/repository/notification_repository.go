package repository

import (
	"context"
	"Ticketing/entity"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{
		db: db,
	}
}

// get all notification
func (r *NotificationRepository) GetAllNotification(ctx context.Context) ([]*entity.Notification, error) {
	Notifications := make([]*entity.Notification, 0)
	result := r.db.WithContext(ctx).Find(&Notifications)
	if result.Error != nil {
		return nil, result.Error
	}
	return Notifications, nil
}
// create notification
func (r *NotificationRepository) CreateNotification(ctx context.Context, Notification *entity.Notification) error {
	result := r.db.WithContext(ctx).Create(&Notification)
	if result.Error != nil {
		return result.Error
	}
	return nil
}



