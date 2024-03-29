package controller

import (
	"net/http"
	"salutemp/backend/src/model"
	"strconv"

	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Serve() *gin.Engine
}

type PgController struct {
	model.Model
}

func (pg *PgController) Serve() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())
	r.GET("/v1/medications/:medID", func(c *gin.Context) {
		id := c.Param("medID")
		intId, err := strconv.Atoi(id)

		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, pg.Medication(intId))
	})

	r.GET("/v1/medications/", func(c *gin.Context) {
		meds, err := pg.AllMedications()
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
			return
		}
		c.JSON(http.StatusOK, meds)
	})

	r.POST("/v1/addmedications", func(c *gin.Context) {
		var med model.Medication

		if err := c.BindJSON(&med); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal medication")
			return
		}
		fmt.Println("Bye, ")

		fmt.Println(med)

		insertedMed, err := pg.AddMedication(med)
		fmt.Println(insertedMed)
		fmt.Println(err)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Failed to add a medication")
			panic(err)
		}

		c.JSON(http.StatusOK, insertedMed)
	})

	r.DELETE("/v1/medications/:medID", func(c *gin.Context) {
		id := c.Param("medID")
		intID, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid medID")
			return
		}

		err = pg.DeleteMedication(intID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to delete medication")
			return
		}

		c.JSON(http.StatusOK, "Medication deleted successfully")
	})

	r.PUT("/v1/medications/:medID", func(c *gin.Context) {
		var med model.Medication

		if err := c.BindJSON(&med); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal medication")
			return
		}

		id := c.Param("medID")
		intID, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid medID")
			return
		}

		med.MedicationID = intID

		err = pg.EditMedication(med)

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to edit medication")
			return
		}

		c.JSON(http.StatusOK, "Medication edited successfully")
	})

	//user routes

	r.GET("v1/userexists/:email", func(c *gin.Context) {
		email := c.Param("email")

		// Retrieve the user.
		user, err := pg.GetUserByEmail(email)
		if err != nil {
			// Handle the error, log it, or return an appropriate response.
			c.JSON(http.StatusNotFound, gin.H{"error": "Something went wrong when finding this user"})
			return
		}

		if user != nil {
			c.JSON(http.StatusOK, gin.H{"message": "This user was found", "user": user})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "User not found"})
		}
	})

	r.GET("/v1/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		_, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid ID")
			return
		}
		c.JSON(http.StatusOK, pg.User(id))
	})

	r.GET("/v1/users/", func(c *gin.Context) {
		patients, err := pg.AllUsers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
			return
		}
		c.JSON(http.StatusOK, patients)
	})

	r.POST("/v1/addusers", func(c *gin.Context) {
		var patient model.User

		if err := c.BindJSON(&patient); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal user")
			return
		}

		insertedPatient, err := pg.AddUser(patient)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Failed to add a user")
			panic(err)
		}

		if err != nil {
			c.JSON(http.StatusBadRequest, "Failed to add a patient")
			panic(err)
		}

		c.JSON(http.StatusOK, insertedPatient)
	})

	r.DELETE("/v1/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		// intID, err := strconv.Atoi(id)

		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, "Invalid ID")
		// 	return
		// }

		err := pg.DeleteUser(id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to delete user")
			return
		}

		c.JSON(http.StatusOK, "User deleted successfully")
	})

	r.PUT("/v1/users/:id", func(c *gin.Context) {
		var user model.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal user")
			return
		}

		id := c.Param("id")
		// intID, err := strconv.Atoi(id)

		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, "Invalid ID")
		// 	return
		// }

		user.UserID = id

		var err = pg.EditUser(user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to edit user")
			return
		}

		c.JSON(http.StatusOK, "User edited successfully")
	})

	//user devices
	// user device routes

	r.GET("/v1/userdevices/:id", func(c *gin.Context) {
		id := c.Param("id")
		intID, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid ID")
			return
		}
		c.JSON(http.StatusOK, pg.UserDevice(int(intID)))
	})

	r.GET("/v1/userdevices/", func(c *gin.Context) {
		userDevices, err := pg.AllUserDevices()
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
			return
		}
		c.JSON(http.StatusOK, userDevices)
	})

	r.POST("/v1/adduserdevices", func(c *gin.Context) {
		var userDevice model.UserDevice

		if err := c.BindJSON(&userDevice); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal user device")
			return
		}

		insertedUserDevice, err := pg.AddUserDevice(userDevice)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Failed to add a user device")
			panic(err)
		}

		c.JSON(http.StatusOK, insertedUserDevice)
	})

	r.DELETE("/v1/userdevices/:id", func(c *gin.Context) {
		id := c.Param("id")
		intID, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid ID")
			return
		}

		err = pg.DeleteUserDevice(int(intID))

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to delete user device")
			return
		}

		c.JSON(http.StatusOK, "User device deleted successfully")
	})

	r.PUT("/v1/userdevices/:id", func(c *gin.Context) {
		var userDevice model.UserDevice

		if err := c.BindJSON(&userDevice); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal user device")
			return
		}

		id := c.Param("id")
		intID, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid ID")
			return
		}

		userDevice.UserDeviceID = int(intID)

		err = pg.EditUserDevice(userDevice)

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to edit user device")
			return
		}

		c.JSON(http.StatusOK, "User device edited successfully")
	})

	//stored medication routes

	r.GET("/v1/storedmedications/:id", func(c *gin.Context) {
		id := c.Param("id")
		intID, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid ID")
			return
		}

		storedMedication, err := pg.StoredMedication(intID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
			return
		}

		c.JSON(http.StatusOK, storedMedication)
	})

	r.GET("/v1/storedmedications/", func(c *gin.Context) {
		storedMedications, err := pg.AllStoredMedications()

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
			panic(err)
		}

		c.JSON(http.StatusOK, storedMedications)
	})

	r.POST("/v1/addstoredmedications", func(c *gin.Context) {
		var storedMedication model.StoredMedication

		if err := c.BindJSON(&storedMedication); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal stored medication")
			return
		}

		insertedMedication, err := pg.AddStoredMedication(storedMedication)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Failed to add stored medication")
			panic(err)
		}

		c.JSON(http.StatusOK, insertedMedication)
	})

	r.DELETE("/v1/storedmedications/:id", func(c *gin.Context) {
		id := c.Param("id")
		intID, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid ID")
			return
		}

		err = pg.DeleteStoredMedication(int(intID))

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to delete stored medication")
			return
		}

		c.JSON(http.StatusOK, "Stored medication deleted successfully")
	})

	r.PUT("/v1/storedmedications/:id", func(c *gin.Context) {
		var storedMedication model.StoredMedication

		if err := c.BindJSON(&storedMedication); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal stored medication")
			return
		}

		id := c.Param("id")
		intID, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid ID")
			return
		}

		storedMedication.StoredMedicationID = int(intID)

		err = pg.EditStoredMedication(storedMedication)

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to edit stored medication")
			return
		}

		c.JSON(http.StatusOK, "Stored medication edited successfully")
	})

	r.GET("/v1/storedmedications/user/:id", func(c *gin.Context) {
		id := c.Param("id")

		storedMedication, err := pg.GetAllStoredMedsFromDBByUser(id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
			return
		}

		c.JSON(http.StatusOK, storedMedication)
	})

	//alerts

	r.GET("/v1/alerts/:id", func(c *gin.Context) {
		id := c.Param("id")
		intID, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid ID")
			return
		}

		alert, err2 := pg.Alert(intID)

		if err2 != nil {
			c.JSON(http.StatusBadRequest, "Invalid Alert")
			return
		}

		c.JSON(http.StatusOK, alert)
	})

	r.GET("/v1/alerts/", func(c *gin.Context) {
		alerts, err := pg.AllAlerts()

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
			return
		}

		c.JSON(http.StatusOK, alerts)
	})

	r.POST("/v1/addalerts", func(c *gin.Context) {
		var alert model.Alert

		if err := c.BindJSON(&alert); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal alert")
			return
		}

		addedAlert, err := pg.AddAlert(alert)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Failed to add alert")
			panic(err)
		}

		c.JSON(http.StatusOK, addedAlert)
	})

	r.DELETE("/v1/alerts/:id", func(c *gin.Context) {
		id := c.Param("id")
		intID, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid ID")
			return
		}

		err = pg.DeleteAlert(int(intID))

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to delete alert")
			return
		}

		c.JSON(http.StatusOK, "Alert deleted successfully")
	})

	r.PUT("/v1/alerts/:id", func(c *gin.Context) {
		var alert model.Alert

		if err := c.BindJSON(&alert); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal alert")
			return
		}

		id := c.Param("id")
		intID, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid ID")
			return
		}

		alert.WarningID = int(intID)

		err = pg.EditAlert(alert)

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to edit alert")
			return
		}

		c.JSON(http.StatusOK, "Alert edited successfully")
	})

	r.GET("/v1/statusreports/:eventtime/:storedmedicationid", func(c *gin.Context) {
		eventTimeString := c.Param("eventtime")
		storedMedicationIDStr := c.Param("storedmedicationid")
		storedMedicationID, err := strconv.Atoi(storedMedicationIDStr)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid event time or stored medication ID")
			return
		}

		eventTime, err := time.Parse(time.RFC3339, eventTimeString)

		event, err := pg.StatusReport(eventTime, storedMedicationID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
			return
		}

		c.JSON(http.StatusOK, event)
	})

	r.GET("/v1/statusreports/", func(c *gin.Context) {
		events, err := pg.AllStatusReports()

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
			return
		}

		c.JSON(http.StatusOK, events)
	})

	r.POST("/v1/addstatusreports", func(c *gin.Context) {
		var event model.StatusReport

		if err := c.BindJSON(&event); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal condition event")
			return
		}

		_, err := pg.AddStatusReport(event)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Failed to add condition event")
			panic(err)
		}

		c.JSON(http.StatusOK, "Condition event added successfully")
	})

	r.DELETE("/v1/statusreports/:eventtime/:storedmedicationid", func(c *gin.Context) {
		eventTimeString := c.Param("eventtime")
		storedMedicationIDStr := c.Param("storedmedicationid")
		storedMedicationID, err := strconv.Atoi(storedMedicationIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid event time or stored medication ID")
			return
		}

		eventTime, err := time.Parse(time.RFC3339, eventTimeString)

		err = pg.DeleteStatusReport(eventTime, storedMedicationID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to delete condition event")
			return
		}

		c.JSON(http.StatusOK, "Condition event deleted successfully")
	})

	r.PUT("/v1/statusreports/:eventtime/:storedmedicationid", func(c *gin.Context) {
		eventTimeParam := c.Param("eventtime")
		storedMedicationIDParam := c.Param("storedmedicationid")

		eventTime, err := time.Parse(time.RFC3339, eventTimeParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid event time format")
			return
		}

		storedMedicationID, err := strconv.Atoi(storedMedicationIDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid stored medication ID")
			return
		}

		var event model.StatusReport
		if err := c.BindJSON(&event); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal status report")
			return
		}

		event.EventTime = eventTime
		event.StoredMedicationID = storedMedicationID

		// Call the function to edit the status report in the database using eventTime and storedMedicationID
		err = pg.EditStatusReport(event)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to edit status report")
			return
		}

		c.JSON(http.StatusOK, "Status report edited successfully")
	})

	// Endpoint to get status reports from the last 24 hours
	r.GET("/v1/statusreports/recent/:storedmedicationid", func(c *gin.Context) {
		storedMedicationIDParam := c.Param("storedmedicationid")

		storedMedicationID, err := strconv.Atoi(storedMedicationIDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid stored medication ID")
			return
		}

		events, err := pg.GetAllStatusReportsLast24Hrs(storedMedicationID)

		if err != nil {
			// Respond with an internal server error and the error message
			c.JSON(http.StatusInternalServerError, gin.H{"Failed to unmarshal status reports": err.Error()})
			return
		}

		// Respond with the filtered status reports
		c.JSON(http.StatusOK, events)
	})

	r.GET("/v1/medicationconstraints/:medicationid/:conditiontype", func(c *gin.Context) {
		medicationIDStr := c.Param("medicationid")
		medicationID, err := strconv.Atoi(medicationIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid medication ID")
			return
		}

		conditionType := c.Param("conditiontype")

		constraint, err := pg.MedicationConstraint(medicationID, conditionType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
			return
		}

		c.JSON(http.StatusOK, constraint)
	})

	r.GET("/v1/medicationconstraints/", func(c *gin.Context) {
		constraints, err := pg.AllMedicationConstraints()

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
			panic(err)
		}

		c.JSON(http.StatusOK, constraints)
	})

	r.POST("/v1/addmedicationconstraints", func(c *gin.Context) {
		var constraint model.MedicationConstraint

		if err := c.BindJSON(&constraint); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal medication constraint")
			return
		}

		_, err := pg.AddMedicationConstraint(constraint)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Failed to add medication constraint")
			panic(err)
		}

		c.JSON(http.StatusOK, "Medication constraint added successfully")
	})

	r.DELETE("/v1/medicationconstraints/:medicationid/:conditiontype", func(c *gin.Context) {
		medicationIDStr := c.Param("medicationid")
		medicationID, err := strconv.Atoi(medicationIDStr)

		conditionType := c.Param("conditiontype")

		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid medication ID")
			return
		}

		err = pg.DeleteMedicationConstraint(medicationID, conditionType)

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to delete medication constraint")
			return
		}

		c.JSON(http.StatusOK, "Medication constraint deleted successfully")
	})

	r.PUT("/v1/medicationconstraints/:medicationid/:conditiontype", func(c *gin.Context) {
		medicationIDStr := c.Param("medicationid")
		_, err := strconv.Atoi(medicationIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid medication ID")
			return
		}

		var constraint model.MedicationConstraint
		if err := c.BindJSON(&constraint); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal medication constraint")
			return
		}

		// Call the function to update the medication constraint in the database
		err = pg.EditMedicationConstraint(constraint)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to update medication constraint")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Medication constraint updated successfully",
		})
	})

	r.GET("/v1/allusermedicationswithconstraint/:userId", func(c *gin.Context) {
		userId := c.Param("userId")

		constraint, err := pg.GetAllUserMedicationsWithConstraint(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
			return
		}

		c.JSON(http.StatusOK, constraint)
	})

	r.GET("/v1/medicationconstraints/storedmedication/:storedmedication", func(c *gin.Context) {
		storedMedication := c.Param("storedmedication")
		storedMedicationId, err := strconv.Atoi(storedMedication)
		constraints, err := pg.AllMedicationConstraintsByStoredMedication(storedMedicationId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
			return
		}
		c.JSON(http.StatusOK, constraints)
	})

	// expo_notification_token routes

	r.GET("/v1/expo_notification_tokens/:user_id", func(c *gin.Context) {
		userID := c.Param("user_id")

		token, err := pg.ExpoNotificationToken(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to retrieve Expo notification token")
			return
		}

		c.JSON(http.StatusOK, token)
	})

	r.GET("/v1/expo_notification_tokens/", func(c *gin.Context) {
		tokens, err := pg.AllExpoNotificationTokens()
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
			return
		}
		c.JSON(http.StatusOK, tokens)
	})

	r.POST("/v1/add_expo_notification_token", func(c *gin.Context) {
		var token model.ExpoNotificationToken

		if err := c.BindJSON(&token); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal expo_notification_token")
			return
		}

		insertedToken, err := pg.AddExpoNotificationToken(token)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Failed to add an expo_notification_token")
			panic(err)
		}

		c.JSON(http.StatusOK, insertedToken)
	})

	r.DELETE("/v1/expo_notification_tokens/:user_id", func(c *gin.Context) {
		userID := c.Param("user_id")

		err := pg.DeleteExpoNotificationToken(userID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to delete expo_notification_token")
			return
		}

		c.JSON(http.StatusOK, "Expo Notification Token deleted successfully")
	})

	r.PUT("/v1/expo_notification_tokens/:user_id", func(c *gin.Context) {
		var token model.ExpoNotificationToken

		if err := c.BindJSON(&token); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal expo_notification_token")
			return
		}

		userID := c.Param("user_id")

		token.UserID = userID

		err := pg.EditExpoNotificationToken(token)

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Failed to edit expo_notification_token")
			return
		}

		c.JSON(http.StatusOK, "Expo Notification Token edited successfully")
	})

	return r
}
