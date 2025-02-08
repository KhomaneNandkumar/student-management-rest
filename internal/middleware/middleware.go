package middleware

import (
	"example.com/student-management/internal/models"
	"github.com/projectdiscovery/gologger"
	"gorm.io/gorm"
)

//function to get all students

func GetAllStudents(DB *gorm.DB) []models.Student {
	var students []models.Student
	if result := DB.Find(&students); result.Error != nil {
		gologger.Error().Msgf("error Fetching all Students: %v", result.Error)
	}
	return students
}

// function to get selective student by its student id

func GetStudntByID(DB *gorm.DB, id int) models.Student {
	var student models.Student
	if result := DB.First(&student, id); result.Error != nil {
		gologger.Error().Msgf("Error While Fetching The Student By Given Student ID:%v", result.Error)
	}
	return student
}

// Add New Student TO The Database

func AddNewStudent(DB *gorm.DB, student models.Student) error {
	if result := DB.Create(&student); result.Error != nil {
		gologger.Error().Msgf("Error While Adding New Student %v", result.Error)
	}
	return nil
}

// Function To Add Existing Record In The Students

func UpdateStudent(DB *gorm.DB, id int, input models.Student) error {
	var student models.Student
	if result := DB.First(&student, id); result.Error != nil {
		gologger.Error().Msgf("Error While Updating The Student Information %v", result.Error)
	}
	if input.FirstName != "" {
		student.FirstName = input.FirstName
	} else if input.LastName != "" {
		student.LastName = input.LastName
	} else if input.Age != 0 {
		student.Age = input.Age
	}
	DB.Save(&student)
	return nil
}

// function to delete a selective record from students

func DeleteStudent(DB *gorm.DB, id int) error {
	var student models.Student
	if result := DB.First(&student, id); result.Error != nil {
		gologger.Error().Msgf("Error Deleting Student %v", result.Error)
		return result.Error
	}
	DB.Delete(&student)
	return nil
}
