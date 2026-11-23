package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/muhammadatak/go-e-wallet/domain"
	"github.com/muhammadatak/go-e-wallet/dto"
)

type ipCheckerService struct {
}

func NewIpChecker() domain.IpCheckerService {
	return &ipCheckerService{}
}

// Query implements [domain.IpCheckerService].
func (i *ipCheckerService) Query(ctx context.Context, ip string) (dto.IpChecker, error) {
	url := fmt.Sprintf("http://ip-api.com/json/%s?fields=status,lat,lon,timezone,query", ip)
	resp, err := http.Get(url)
	if err != nil {
		return dto.IpChecker{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dto.IpChecker{}, err
	}

	checker := dto.IpChecker{}

	err = json.Unmarshal(body, &checker)
	if err != nil {
		return dto.IpChecker{}, err
	}

	return checker, nil
}
