package repository

import "checkpoint-2/internal/domain"

type PatientRepository interface {
	Post(domain.CreatePatient) error
	Get(int) (domain.Patient, error)
	GetAll() ([]domain.Patient, error)
	Put(int, domain.UpdatePatient) error
	Patch(int, domain.PatchPatientName) error
	Delete(int) error
}
