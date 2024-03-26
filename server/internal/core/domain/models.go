// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package letstour

import (
	"database/sql"
	"encoding/json"
)

type Address struct {
	ID         []byte
	Street     sql.NullString
	City       sql.NullString
	Province   sql.NullString
	PostalCode sql.NullString
	Country    sql.NullString
	IsoCountry sql.NullString
	Latitude   sql.NullFloat64
	Longitude  sql.NullFloat64
}

type Attraction struct {
	ID          []byte
	Name        string
	Description string
	Schedule    json.RawMessage
	Type        string
	Likes       sql.NullInt32
	Isdeleted   sql.NullBool
	Createdat   sql.NullTime
	Updatedat   sql.NullTime
	Deletedat   sql.NullTime
	Enabled     sql.NullBool
}

type Company struct {
	ID        []byte
	Name      string
	Address   string
	Schedule  json.RawMessage
	Isdeleted sql.NullBool
	Createdat sql.NullTime
	Updatedat sql.NullTime
	Deletedat sql.NullTime
	Enabled   sql.NullBool
}

type Post struct {
	ID          []byte
	Title       string
	PostContent json.RawMessage
	Likes       sql.NullInt32
	Isdeleted   sql.NullBool
	Createdat   sql.NullTime
	Updatedat   sql.NullTime
	Deletedat   sql.NullTime
	Enabled     sql.NullBool
	Fkuserid    sql.NullString
}

type PostComment struct {
	ID        []byte
	Comment   string
	Likes     int32
	Isdeleted sql.NullBool
	Createdat sql.NullTime
	Updatedat sql.NullTime
	Deletedat sql.NullTime
	Enabled   sql.NullBool
}

type Review struct {
	ID          []byte
	Title       string
	Score       int32
	Description string
	Likes       int32
	Isdeleted   sql.NullBool
	Createdat   sql.NullTime
	Updatedat   sql.NullTime
	Deletedat   sql.NullTime
	Enabled     sql.NullBool
}

type User struct {
	ID        []byte
	Username  string
	Fullname  string
	Email     string
	Password  string
	Isdeleted sql.NullBool
	Createdat sql.NullTime
	Updatedat sql.NullTime
	Deletedat sql.NullTime
	Enabled   sql.NullBool
}
