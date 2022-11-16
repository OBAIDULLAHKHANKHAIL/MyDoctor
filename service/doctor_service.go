package service

import "github.com/obaidullahkhankhail/MyDoctor/entity"

type DoctorService interface {
	FindAllDoctors() []entity.Doctor
	RegisterDoctor(entity.Doctor) entity.Doctor
}

type doctorService struct{
	doctor []entity.Doctor
}

func NewDoctor() DoctorService{
	return &doctorService{}
}

func (service *doctorService) RegisterDoctor(doctor entity.Doctor) entity.Doctor{
	service.doctor = append(service.doctor, doctor)
	return doctor
}

func (service *doctorService) FindAllDoctors() []entity.Doctor{
	return service.doctor
}

