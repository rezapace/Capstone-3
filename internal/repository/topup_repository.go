package repository

import (
    "Ticketing/entity"
    "context"
    "gorm.io/gorm"
)

type TopupRepository interface {
    InsertTopup(ctx context.Context, topup entity.Topup) (entity.Topup, error)
}

type topupRepository struct {
    db *gorm.DB
}

func NewTopupRepository(db *gorm.DB) *topupRepository {
    return &topupRepository{db}
}

func (r *topupRepository) InsertTopup(ctx context.Context, topup entity.Topup) (entity.Topup, error) {
    result := r.db.WithContext(ctx).Create(&topup)
    if result.Error != nil {
        return entity.Topup{}, result.Error
    }
    return topup, nil
}

// topup saldo sederhana
// func (r *topupRepository) TopupSaldo(ctx context.Context, topup entity.Topup) (entity.Topup, error) {
//     result := r.db.WithContext(ctx).Create(&topup)
//     if result.Error != nil {
//         return entity.Topup{}, result.Error
//     }
//     return topup, nil
// }
