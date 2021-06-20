package controllers

import (
	"database/sql"
	"net/http"

	"github.com/To-Doc/WebServices/models"
	"github.com/To-Doc/WebServices/utils"
	"github.com/gin-gonic/gin"
)

func GetDoc(dB *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var search models.Search
		var patient models.Patient
		var patients []models.Patient
		if err := c.ShouldBindJSON(&search); err != nil {
			utils.RespondWithError(http.StatusBadRequest, "invalid format", c)
			return
		}
    
    if search.Key == "email" {
		var patientId int
		errC := dB.QueryRow("SELECT p.patient_name, p.age, g.gender FROM patient p INNER JOIN user u ON u.user_id=p.user_id INNER JOIN gender_lookup g ON g.gender_id=p.gender WHERE u.email=?", search.Email).Scan(&patientId)
		if errC == sql.ErrNoRows {
			utils.RespondWithError(http.StatusBadRequest, "Sorry Patient not found", c)
			return
		}
    
     if user.Email != "" {
			err = dB.QueryRow("SELECT password FROM user WHERE email=?", user.Email).Scan(&password)
			if err != nil {
				utils.RespondWithError(http.StatusInternalServerError, "unable to let you in", c)
				return
			}
			if !(password == user.Password) {
				utils.RespondWithError(http.StatusBadRequest, "invalid credentials", c)
				return
			}
		}
		if err != nil {
			utils.RespondWithError(http.StatusInternalServerError, "unexepected error occurred", c)
			return
		}
    JSON(http.StatusOK, gin.H{
			"success": "true",
			"message": "successfully logged in",
		})

	}
}
    
    
    
