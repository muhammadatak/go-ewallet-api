package service

import (
	"context"
	"log"
	"time"

	"github.com/muhammadatak/go-e-wallet/domain"
	"github.com/muhammadatak/go-e-wallet/dto"
	"github.com/muhammadatak/go-e-wallet/internal/util"
)

type fdsService struct {
	ipCheckerService   domain.IpCheckerService
	loginLogRepository domain.LoginLogRepository
}

func NewFds(ipCheckerService domain.IpCheckerService,
	loginLogRepository domain.LoginLogRepository) domain.FdsService {
	return &fdsService{
		ipCheckerService:   ipCheckerService,
		loginLogRepository: loginLogRepository,
	}
}
func (fd *fdsService) IsAuthorized(ctx context.Context, ip string, userId int64) bool {
	locationCheck, err := fd.ipCheckerService.Query(ctx, ip)
	if err != nil || locationCheck == (dto.IpChecker{}) {
		return false
	}

	newAccess := domain.LoginLog{
		UserId:       userId,
		IsAuthorized: false,
		IpAddress:    ip,
		Timezone:     locationCheck.Timezone,
		Lat:          locationCheck.Lat,
		Lon:          locationCheck.Lon,
		AccessTime:   time.Now(),
	}

	lastLogin, err := fd.loginLogRepository.FindLastAuthorized(ctx, userId)

	if err != nil {
		_ = fd.loginLogRepository.Insert(ctx, &newAccess)
		return false
	}

	if lastLogin == (domain.LoginLog{}) {
		newAccess.IsAuthorized = true
		_ = fd.loginLogRepository.Insert(ctx, &newAccess)
		return true
	}

	distanceHour := newAccess.AccessTime.Sub(lastLogin.AccessTime)
	distanceChange := util.GetDistance(lastLogin.Lat, newAccess.Lat, lastLogin.Lon, newAccess.Lon)
	log.Printf("hour: %f, distance:%f", distanceHour.Hours(), distanceChange)

	if distanceChange/distanceHour.Hours() > 400 {
		_ = fd.loginLogRepository.Insert(ctx, &newAccess)
		return false
	}

	newAccess.IsAuthorized = true
	_ = fd.loginLogRepository.Insert(ctx, &newAccess)
	return true
}
