package usecase

import (
	"checkpoint-2/internal/domain"
)

type Dentist interface {
	Post(domain.CreateDentist) error
	Get(int) (domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	Put(int, domain.Dentist) error
	Patch(int, domain.Dentist) error
	Delete(int) error
}
