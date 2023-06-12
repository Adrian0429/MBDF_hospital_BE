package services

import (
	"github.com/Caknoooo/golang-clean_template/repository"
)

type ReservasiService interface {
}

type reservasiService struct {
	reservasiRepo repository.ReservasiRepository
}

func NewReservasiService(rs repository.ReservasiRepository) ReservasiService {
	return &reservasiService{
		reservasiRepo: rs,
	}
}
