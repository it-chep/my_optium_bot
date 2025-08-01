// Code generated by xo. DO NOT EDIT.
// Package xo contains the types for schema 'public'.
package xo

import (
	"database/sql"
	"reflect"
)

// PatientDoctor represents a row from 'public.patient_doctor'.
type PatientDoctor struct {
	DoctorTg  sql.NullInt64 `db:"doctor_tg" json:"doctor_tg"`   // doctor_tg
	PatientTg sql.NullInt64 `db:"patient_tg" json:"patient_tg"` // patient_tg
	PatientID sql.NullInt64 `db:"patient_id" json:"patient_id"` // patient_id
	ChatID    sql.NullInt64 `db:"chat_id" json:"chat_id"`       // chat_id
}

// zeroPatientDoctor zero value of dto
var zeroPatientDoctor = PatientDoctor{}

func (t PatientDoctor) IsEmpty() bool {
	return reflect.DeepEqual(t, zeroPatientDoctor)
}
