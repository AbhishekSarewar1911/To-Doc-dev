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
    
    
    
		if search.Key == "specialization" {
			var sp_id int8
			errS := dB.QueryRow("SELECT specialization_id FROM specialization_lookup WHERE specialization=?", search.Name).Scan(&sp_id)
			if errS == sql.ErrNoRows {
				utils.RespondWithError(http.StatusBadRequest, "No Doctor available of given speciality in your area", c)
				return
			}
			rows, err := dB.Query("SELECT doctor_name ,doctor_phoneno ,email ,qualification,longitude,latitude FROM doctor WHERE specialization=? and city=?", sp_id, cityId)
			if err != nil {
				utils.RespondWithError(http.StatusInternalServerError, "No Doctor Found", c)
				return
			} else {
				for rows.Next() {
					err := rows.Scan(&doctor.Name, &doctor.Phone, &doctor.Email, &doctor.Qualifiaction, &doctor.Longitude, &doctor.Latitude)
					if err != nil {
						utils.RespondWithError(http.StatusInternalServerError, err.Error(), c)
						return
					}
					doctors = append(doctors, doctor)
				}
				c.JSON(http.StatusOK, gin.H{
					"success": "true",
					"message": doctors,
				})
			}
		}
		if search.Key == "name" {
			rows, err := dB.Query("SELECT doctor_name ,doctor_phoneno ,email ,qualification,longitude,latitude FROM doctor WHERE doctor_name=? and city=?", search.Name, cityId)
			if err != nil {
				utils.RespondWithError(http.StatusInternalServerError, "No Doctor Found", c)
				return
			} else {
				for rows.Next() {
					err := rows.Scan(&doctor.Name, &doctor.Phone, &doctor.Email, &doctor.Qualifiaction, &doctor.Longitude, &doctor.Latitude)
					if err != nil {
						utils.RespondWithError(http.StatusInternalServerError, err.Error(), c)
						return
					}
					doctors = append(doctors, doctor)
				}
				c.JSON(http.StatusOK, gin.H{
					"success": "true",
					"message": doctors,
				})
			}
		}
	}

}
