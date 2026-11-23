package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/muhammadatak/go-e-wallet/domain"
	"github.com/muhammadatak/go-e-wallet/dto"
)

type topUpService struct {
	notificationService   domain.NotificationService
	midtransService       domain.MidtransService
	topUpRepository       domain.TopUpRepository
	accountRepository     domain.AccountRepository
	transactionRepository domain.TransactionRepository
}

func NewTopUp(notificationService domain.NotificationService, midtransService domain.MidtransService, topUpRepository domain.TopUpRepository, accountRepository domain.AccountRepository, transactionRepository domain.TransactionRepository) domain.TopUpService {
	return &topUpService{
		notificationService:   notificationService,
		midtransService:       midtransService,
		topUpRepository:       topUpRepository,
		accountRepository:     accountRepository,
		transactionRepository: transactionRepository,
	}
}

// InitializedTopUp implements [domain.TopUpService].
func (t *topUpService) InitializedTopUp(ctx context.Context, req dto.TopUpReq) (dto.TopUpRes, error) {
	topup := domain.TopUp{
		Id:     uuid.NewString(),
		UserId: req.UserId,
		Status: 0,
		Amount: req.Amount,
	}
	err := t.midtransService.GenerateUrl(ctx, &topup)
	if err != nil {
		return dto.TopUpRes{}, err
	}

	err = t.topUpRepository.Insert(ctx, &topup)
	if err != nil {
		return dto.TopUpRes{}, err
	}

	return dto.TopUpRes{
		SnapUrl: topup.SnapUrl,
	}, nil
}

// ConfirmedTopUp implements [domain.TopUpService].
func (t *topUpService) ConfirmedTopUp(ctx context.Context, id string) error {
	topup, err := t.topUpRepository.FindById(ctx, id)
	if err != nil {
		return err
	}
	if topup == (domain.TopUp{}) {
		return domain.ErrNotFound
	}

	account, err := t.accountRepository.FindByUserId(ctx, topup.UserId)
	if err != nil {
		return err
	}

	if account == (domain.Account{}) {
		return domain.ErrNotFound
	}

	err = t.transactionRepository.Insert(ctx, &domain.Transaction{
		AccountId:           account.Id,
		SofNumber:           "00",
		DofNumber:           account.AccountNumber,
		TransactionType:     "C",
		Amount:              topup.Amount,
		TransactionDateTime: time.Now(),
	})

	if err != nil {
		return err
	}

	account.Balance += topup.Amount
	err = t.accountRepository.Update(ctx, &account)
	if err != nil {
		return err
	}

	data := map[string]string{
		"amount": fmt.Sprintf("%.2f", topup.Amount),
	}

	_ = t.notificationService.Insert(ctx, account.UserId, "TOPUP_SUCCESS", data)

	return nil
}
