package model
import "time"


type User struct {
    UserID                 int    `json:"user_id"`
    FirstName              string `json:"first_name"`
    LastName               string `json:"last_name"`
    Email                  string `json:"email"`
    PushNotificationEnabled bool   `json:"push_notification_enabled"`
}




type UserDevice struct {
    UserDeviceID int    `json:"user_device_id"`
    UserID       int    `json:"user_id"`
    DeviceID     string `json:"device_id"`
}


type Medication struct {
    MedicationID   int    `json:"medication_id"`
    MedicationName string `json:"medication_name"`
}

type StoredMedication struct {
    StoredMedicationID int     `json:"stored_medication_id"`
    MedicationID       int     `json:"medication_id"`
    UserID             int     `json:"user_id"`
    CurrentTemperature float64 `json:"current_temperature"`
    CurrentHumidity    float64 `json:"current_humidity"`
    CurrentLight       float64 `json:"current_light"`
}

type Alert struct {
    WarningID           int       `json:"warning_id"`
    StoredMedicationID  int       `json:"stored_medication_id"`
    WarningTimestamp    time.Time `json:"warning_timestamp"`
    WarningDescription  string    `json:"warning_description"`
    ConditionType       string    `json:"condition_type"`
}

type StatusReport struct {
    EventTime           time.Time  `json:"event_time"`
    StoredMedicationID  int        `json:"stored_medication_id"`
    Temperature         float64    `json:"temperature"`
    Humidity            float64    `json:"humidity"`
    Light               float64    `json:"light"`
}

type MedicationConstraint struct {
    StoredMedicationID  int     `json:"medication_id"`
    ConditionType string  `json:"condition_type"` 
    MaxThreshold  float64 `json:"max_threshold"`
    MinThreshold  float64 `json:"min_threshold"`
    Duration      string  `json:"duration"` 
}
