package domain

type Appointment struct {
	Id int `json:"id"`
	Patient Patient `json:"patient"`
	Dentist Dentist `json:"dentist"`
	AppointmentDate string `json:"appointment_date"`
}