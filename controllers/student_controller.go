package controllers

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
	"mahasiswa/config"
	"mahasiswa/models"
	"net/http"
	"strconv"
)

func AddStudent(c echo.Context) error {
	fileHeader, err := c.FormFile("foto")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to get file from form",
			"error":   err.Error(),
		})
	}

	file, err := fileHeader.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to open file",
			"error":   err.Error(),
		})
	}
	defer file.Close()

	cloudinaryService, err := config.InitCloudinary()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to initialize Cloudinary service",
			"error":   err.Error(),
		})
	}
	uploadResult, err := cloudinaryService.Upload.Upload(context.Background(), file, uploader.UploadParams{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to upload file to Cloudinary",
			"error":   err.Error(),
		})
	}

	student := &models.Student{}
	if err := c.Bind(student); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request",
		})
	}

	student.Foto = uploadResult.SecureURL

	createdStudent, err := models.CreateStudentData(config.DB, *student)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create student data",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, createdStudent)
}

func GetStudents(c echo.Context) error {
	student, err := models.GetStudent(config.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get students",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, student)
}

func GetStudentByID(c echo.Context) error {
	db := config.DB
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get student",
			"error":   err.Error(),
		})
	}
	student, err := models.GetStudentByID(db, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get student",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, student)
}

func UpdateStudent(c echo.Context) error {
	studentIDStr := c.Param("id")
	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid student ID",
			"error":   err.Error(),
		})
	}

	student := &models.Student{}
	if err := c.Bind(student); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request",
		})
	}

	student.ID = studentID
	updatedStudent, err := models.UpdateStudent(config.DB, student)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to update student",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, updatedStudent)
}

func DeleteStudent(c echo.Context) error {
	studentIDStr := c.Param("id")
	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid student ID",
			"error":   err.Error(),
		})
	}
	err = models.DeleteStudent(config.DB, studentID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to delete student",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Student deleted successfully",
	})
}
