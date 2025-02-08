package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/student-management/initializers"
	"example.com/student-management/internal/middleware"
	"example.com/student-management/internal/models"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, "Server Is Up")
}

func NoRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, "Endpoint doesn't exist")
}

func GetAll(c *gin.Context) {
	res := middleware.GetAllStudents(initializers.DB)
	c.JSON(http.StatusOK, res)
}

func GetStudent(c *gin.Context) {
	id := c.Param("id")
	sid, _ := strconv.Atoi(id)
	res := middleware.GetStudntByID(initializers.DB, sid)
	c.JSON(http.StatusOK, res)
}

func AddStudent(c *gin.Context) {
	student := models.Student{}
	err := json.NewDecoder(c.Request.Body).Decode(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request Payload")
		return
	}
	res := middleware.AddNewStudent(initializers.DB, student)
	if res == nil {
		c.JSON(http.StatusOK, "Student Added Successfully")
	} else {
		c.JSON(http.StatusNotFound, "Student Could Not be Added")
	}
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	studentId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Query Parameter")
		return
	}
	student := models.Student{}
	err = json.NewDecoder(c.Request.Body).Decode(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Query Parametrs")
		return
	}

	res := middleware.UpdateStudent(initializers.DB, studentId, student)
	if res == nil {
		c.JSON(http.StatusOK, "Student Updated")
	} else {
		c.JSON(http.StatusNotFound, "Student Coundn't be Updated")
	}
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	sid, _ := strconv.Atoi(id)
	res := middleware.DeleteStudent(initializers.DB, sid)

	if res == nil {
		c.JSON(http.StatusOK, "Student Deleted")
	} else {
		c.JSON(http.StatusNotFound, "Student Not Found")
	}
}
