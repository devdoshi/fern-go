// Generated by Fern. Do not edit.

package api

import (
	uuid "github.com/gofrs/uuid"
	time "time"
)

type Type struct {
	One   int       `json:"one"`
	Two   float64   `json:"two"`
	Three string    `json:"three"`
	Four  bool      `json:"four"`
	Five  int64     `json:"five"`
	Six   time.Time `json:"six"`
	Seven time.Time `json:"seven"`
	Eight uuid.UUID `json:"eight"`
	Nine  []byte    `json:"nine"`
}
