package repository

import "checkpoint-2/internal/domain"

type DentistRepository interface {
	Post(domain.CreateDentist) error
	Get(int) (domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	Put(int, domain.UpdateDentist) error
	Patch(int, domain.Dentist) error
	Delete(int) error
}
