package db

import (
	"checkpoint-2/internal/application/repository"
	"checkpoint-2/internal/domain"
	"database/sql"
)

type patientDatabase struct {
	db *sql.DB
}

func (d *patientDatabase) Post(patient domain.Patient) error {
	_, err := d.db.Exec(
		`INSERT INTO patients
		(id, name, surname, rg, registry_date)
		VALUES
		(?, ?, ?, ?, ?, ?)`,
		patient.Id,		
		patient.Name,
		patient.Surname,
		patient.RG,
		patient.RegistryDate,
	)
	if err != nil {
		return err
	}
	return nil
}

func (d *patientDatabase) Get(id int) (domain.Patient, error) {
	var patient domain.Patient

	rows, err := d.db.Query(`SELECT * from patients WHERE id=?`, id)
	if err != nil {
		return domain.Patient{}, err
	}
	for rows.Next() {
		err := rows.Scan(
			&patient.Id,
			&patient.Name,
			&patient.Surname,
			&patient.RG,
			&patient.RegistryDate,
		)
		if err != nil {
			return domain.Patient{}, err
		}
	}
	return patient, nil
}

func (d *patientDatabase) GetAll() ([]domain.Patient, error) {
	var patients []domain.Patient

	rows, err := d.db.Query("SELECT id, name, surname, rg, registry_date FROM patients")

	if err != nil {
		return patients, err
	}

	defer rows.Close()

	for rows.Next() {
		var patient domain.Patient

		err := rows.Scan(
			&patient.Id,
			&patient.Name,
			&patient.Surname,
			&patient.RG,
			&patient.RegistryDate,
		)
		if err != nil {
			return patients, err
		}
		patients = append(patients, patient)
	}
	return patients, nil
}

func (d *patientDatabase) Put(id int, dentist domain.Patient) error {
	return nil
}
func (d *patientDatabase) Patch(id int, dentist domain.Patient) error {
	return nil
}
func (d *patientDatabase) Delete(id int) error {
	return nil
}

func NewPatientDatabase(database *sql.DB) repository.PatientRepository {
	return &patientDatabase{database}
}