package models

import (
	"gorm.io/gorm"
	"time"
)

type Student struct {
	ID        int       `json:"id" *gorm:"primaryKey"`
	Nama      string    `json:"nama" form:"nama"`
	NIM       string    `json:"nim" form:"nim"`
	Jurusan   string    `json:"jurusan" form:"jurusan"`
	Alamat    string    `json:"alamat" form:"alamat"`
	Foto      string    `json:"foto" form:"foto"`
	CreatedAt time.Time `*gorm:"autoCreateTime"`
	UpdatedAt time.Time `*gorm:"autoUpdateTime"`
}

func CreateStudentData(db *gorm.DB, student Student) (Student, error) {
	if err := db.Create(&student).Error; err != nil {
		return Student{}, err
	}
	return student, nil
}

func GetStudent(db *gorm.DB) ([]Student, error) {
	var student []Student
	if err := db.Find(&student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func GetStudentByID(db *gorm.DB, studentId int) (Student, error) {
	var student Student
	if err := db.Where("id = ?", studentId).First(&student).Error; err != nil {
		return Student{}, err
	}
	return student, nil
}

func UpdateStudent(db *gorm.DB, student *Student) (*Student, error) {
	student.UpdatedAt = time.Now()
	if err := db.Model(&Student{}).Where("id = ?", student.ID).Updates(student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func DeleteStudent(db *gorm.DB, mahasiswaID int) error {
	if err := db.Delete(&Student{}, mahasiswaID).Error; err != nil {
		return err
	}
	return nil
}
