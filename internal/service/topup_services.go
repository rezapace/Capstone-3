package service

import (
    "context"
    "Ticketing/entity"
    "Ticketing/internal/repository"
    "github.com/midtrans/midtrans-go"
    "github.com/midtrans/midtrans-go/coreapi"
)

type TopupService interface {
    CreateTopup(ctx context.Context, topup entity.Topup) (entity.Topup, error)
    CreateMidtransCharge(orderID string, amount int64) (*coreapi.ChargeResponse, error)
    // TopupSaldo(ctx context.Context, topup entity.Topup) (entity.Topup, error)
}

type topupService struct {
    topupRepository repository.TopupRepository
}

// type TopupRepository interface {
//     TopupSaldo(ctx context.Context, topup entity.Topup) (entity.Topup, error)
// }

func NewTopupService(topupRepository repository.TopupRepository) *topupService {
    return &topupService{topupRepository}
}

func (s *topupService) CreateTopup(ctx context.Context, topup entity.Topup) (entity.Topup, error) {
    return s.topupRepository.InsertTopup(ctx, topup)
}

func (s *topupService) CreateMidtransCharge(orderID string, amount int64) (*coreapi.ChargeResponse, error) {
    c := coreapi.Client{}
    c.New("YOUR-SERVER-KEY", midtrans.Sandbox) // Ganti dengan server key Anda

    chargeReq := &coreapi.ChargeReq{
        PaymentType: coreapi.PaymentTypeBankTransfer, // Sesuaikan dengan jenis pembayaran
        TransactionDetails: midtrans.TransactionDetails{
            OrderID:  orderID,
            GrossAmt: amount,
        },
        // Tambahkan detail lainnya sesuai kebutuhan
    }

    return c.ChargeTransaction(chargeReq)
}

// topup saldo sederhana
// func (s *topupService) TopupSaldo(ctx context.Context, topup entity.Topup) (entity.Topup, error) {
//     return s.topupRepository.TopupSaldo(ctx, topup)
// }