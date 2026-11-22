package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/muhammadatak/go-e-wallet/domain"
	"github.com/muhammadatak/go-e-wallet/dto"
	"github.com/muhammadatak/go-e-wallet/internal/util"
)

type transactionService struct {
	accountRepository     domain.AccountRepository
	transactionRepository domain.TransactionRepository
	cacheRepository       domain.CacheRepository
}

func NewTransaction(accountRepository domain.AccountRepository,
	transactionRepository domain.TransactionRepository,
	cacheRepository domain.CacheRepository) domain.TransactionService {
	return &transactionService{
		accountRepository:     accountRepository,
		transactionRepository: transactionRepository,
		cacheRepository:       cacheRepository,
	}
}

// TransferInquiry implements [domain.TransactionService].
func (t *transactionService) TransferInquiry(ctx context.Context, req dto.TransferInquiryReq, userId int64) (dto.TransferInquiryRes, error) {
	myAccount, err := t.accountRepository.FindByUserId(ctx, userId)

	if err != nil {
		return dto.TransferInquiryRes{}, err
	}

	if myAccount == (domain.Account{}) {
		return dto.TransferInquiryRes{}, domain.ErrNotFound
	}

	dofAccount, err := t.accountRepository.FindByAccountNumber(ctx, req.AccountNumber)

	if err != nil {
		return dto.TransferInquiryRes{}, err
	}

	if dofAccount == (domain.Account{}) {
		return dto.TransferInquiryRes{}, domain.ErrNotFound
	}

	if myAccount.Balance < req.Amount {
		return dto.TransferInquiryRes{}, domain.ErrInsufficientBalance
	}

	inquiryKey := util.GenerateId()

	jsonData, _ := json.Marshal(req)
	_ = t.cacheRepository.Set(inquiryKey, jsonData)

	return dto.TransferInquiryRes{
		InquiryKey: inquiryKey,
	}, nil
}

// TransferExecute implements [domain.TransactionService].
func (t *transactionService) TransferExecute(ctx context.Context, req dto.TransferExecuteReq, userId int64) error {
	val, err := t.cacheRepository.Get(req.InquiryKey)
	if err != nil {
		return err
	}

	var reqInq dto.TransferInquiryReq
	_ = json.Unmarshal(val, &reqInq)
	if reqInq == (dto.TransferInquiryReq{}) {
		return domain.ErrNotFound
	}

	myAccount, err := t.accountRepository.FindByUserId(ctx, userId)

	if err != nil {
		return err
	}

	dofAccount, err := t.accountRepository.FindByAccountNumber(ctx, reqInq.AccountNumber)

	if dofAccount == (domain.Account{}) {
		return domain.ErrNotFound
	}

	debitTransaction := domain.Transaction{
		AccountId:           myAccount.Id,
		SofNumber:           myAccount.AccountNumber,
		DofNumber:           dofAccount.AccountNumber,
		TransactionType:     "D",
		Amount:              reqInq.Amount,
		TransactionDateTime: time.Now(),
	}

	err = t.transactionRepository.Insert(ctx, &debitTransaction)

	if err != nil {
		return err
	}

	creditTransaction := domain.Transaction{
		AccountId:           dofAccount.Id,
		SofNumber:           myAccount.AccountNumber,
		DofNumber:           dofAccount.AccountNumber,
		TransactionType:     "C",
		Amount:              reqInq.Amount,
		TransactionDateTime: time.Now(),
	}

	err = t.transactionRepository.Insert(ctx, &creditTransaction)

	if err != nil {
		return err
	}

	myAccount.Balance -= reqInq.Amount

	err = t.accountRepository.Update(ctx, &myAccount)
	if err != nil {
		return err
	}

	dofAccount.Balance += reqInq.Amount

	err = t.accountRepository.Update(ctx, &dofAccount)
	if err != nil {
		return err
	}

	return nil
}
