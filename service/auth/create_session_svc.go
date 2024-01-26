package service

import (
	"context"
	"log"

	"github.com/GeraAnggaraPutra/blueprint-go/repository/auth/dto"
	"github.com/GeraAnggaraPutra/blueprint-go/repository/auth/model"
)

func (s *authService) CreateSessionSvc(
	ctx context.Context,
	req model.Session,
) (data model.Session, err error) {
	data, err = s.authQuery.CreateSessionQuery(
		ctx,
		dto.CreateSessionParams{
			GUID:      req.GUID,
			UserGUID:  req.UserGUID,
			IPAddress: req.IPAddress,
			Token:     req.Token,
			UserAgent: req.UserAgent,
			ExpiredAt: req.ExpiredAt,
			CreatedAt: req.CreatedAt,
		},
	)
	if err != nil {
		log.Printf("failed create session with err: %s", err)
		return
	}

	return
}
