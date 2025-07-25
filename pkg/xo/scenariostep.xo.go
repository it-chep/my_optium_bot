// Code generated by xo. DO NOT EDIT.
// Package xo contains the types for schema 'public'.
package xo

import (
	"database/sql"
	"reflect"
)

// ScenarioStep represents a row from 'public.scenario_steps'.
type ScenarioStep struct {
	ID         int          `db:"id" json:"id"`                   // id
	ScenarioID int          `db:"scenario_id" json:"scenario_id"` // scenario_id
	StepOrder  int          `db:"step_order" json:"step_order"`   // step_order
	Content    string       `db:"content" json:"content"`         // content
	IsFinal    sql.NullBool `db:"is_final" json:"is_final"`       // is_final
}

// zeroScenarioStep zero value of dto
var zeroScenarioStep = ScenarioStep{}

func (t ScenarioStep) IsEmpty() bool {
	return reflect.DeepEqual(t, zeroScenarioStep)
}
