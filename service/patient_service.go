package service

import "github.com/obaidullahkhankhail/MyDoctor/entity"

type PatientService interface {
	FindAllPatients() []entity.Patient
	RegisterPatient(entity.Patient) entity.Patient
}

type patientService struct{
	patients []entity.Patient
}

func NewPatient() PatientService{
	return &patientService{}
}

func (service *patientService) RegisterPatient(patient entity.Patient) entity.Patient{
	service.patients = append(service.patients, patient)
	return patient
}

func (service *patientService) FindAllPatients() []entity.Patient{
	return service.patients
}

