// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlc

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type Action string

const (
	ActionBorrado     Action = "borrado"
	ActionActualizado Action = "actualizado"
	ActionCreado      Action = "creado"
)

func (e *Action) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Action(s)
	case string:
		*e = Action(s)
	default:
		return fmt.Errorf("unsupported scan type for Action: %T", src)
	}
	return nil
}

type NullAction struct {
	Action Action `json:"action"`
	Valid  bool   `json:"valid"` // Valid is true if Action is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAction) Scan(value interface{}) error {
	if value == nil {
		ns.Action, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Action.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAction) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Action), nil
}

type Extensions string

const (
	ExtensionsPdf  Extensions = ".pdf"
	ExtensionsDoc  Extensions = ".doc"
	ExtensionsDocx Extensions = ".docx"
)

func (e *Extensions) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Extensions(s)
	case string:
		*e = Extensions(s)
	default:
		return fmt.Errorf("unsupported scan type for Extensions: %T", src)
	}
	return nil
}

type NullExtensions struct {
	Extensions Extensions `json:"extensions"`
	Valid      bool       `json:"valid"` // Valid is true if Extensions is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullExtensions) Scan(value interface{}) error {
	if value == nil {
		ns.Extensions, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Extensions.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullExtensions) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Extensions), nil
}

type Status string

const (
	StatusAprobado   Status = "aprobado"
	StatusEnrevision Status = "en revision"
	StatusRechazado  Status = "rechazado"
)

func (e *Status) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Status(s)
	case string:
		*e = Status(s)
	default:
		return fmt.Errorf("unsupported scan type for Status: %T", src)
	}
	return nil
}

type NullStatus struct {
	Status Status `json:"status"`
	Valid  bool   `json:"valid"` // Valid is true if Status is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStatus) Scan(value interface{}) error {
	if value == nil {
		ns.Status, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Status.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Status), nil
}

type Asesor struct {
	AsesorID  int64        `json:"asesor_id"`
	ClientID  int32        `json:"client_id"`
	Username  string       `json:"username"`
	Photo     string       `json:"photo"`
	About     string       `json:"about"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

type Client struct {
	ClientID  int64        `json:"client_id"`
	GivenName string       `json:"given_name"`
	Surname   string       `json:"surname"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

type ClientRole struct {
	ClientRoleID int64        `json:"client_role_id"`
	ClientID     int32        `json:"client_id"`
	RoleID       int32        `json:"role_id"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
	DeletedAt    sql.NullTime `json:"deleted_at"`
}

type DirectoryInstitution struct {
	DirectoryInstitutionID int64        `json:"directory_institution_id"`
	InstitutionID          int32        `json:"institution_id"`
	DirectoryID            int32        `json:"directory_id"`
	CreatedAt              time.Time    `json:"created_at"`
	UpdatedAt              sql.NullTime `json:"updated_at"`
	DeletedAt              sql.NullTime `json:"deleted_at"`
}

type DirectoryRole struct {
	DirectoryRoleID int64         `json:"directory_role_id"`
	DirectoryID     sql.NullInt32 `json:"directory_id"`
	RoleID          sql.NullInt32 `json:"role_id"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       sql.NullTime  `json:"updated_at"`
	DeletedAt       sql.NullTime  `json:"deleted_at"`
}

type DirectoryTree struct {
	DirectoryID int64          `json:"directory_id"`
	ParentID    sql.NullInt32  `json:"parent_id"`
	Name        sql.NullString `json:"name"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	DeletedAt   sql.NullTime   `json:"deleted_at"`
}

type Document struct {
	DocumentID  int64         `json:"document_id"`
	DirectoryID int32         `json:"directory_id"`
	FormatID    sql.NullInt32 `json:"format_id"`
	Url         string        `json:"url"`
	Status      Status        `json:"status"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   sql.NullTime  `json:"updated_at"`
	DeletedAt   sql.NullTime  `json:"deleted_at"`
}

type DocumentClient struct {
	DocumentClientID int64        `json:"document_client_id"`
	ClientID         int32        `json:"client_id"`
	DocumentID       int32        `json:"document_id"`
	Action           Action       `json:"action"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        sql.NullTime `json:"updated_at"`
	DeletedAt        sql.NullTime `json:"deleted_at"`
}

type Format struct {
	FormatID        int64         `json:"format_id"`
	UpdatedFormatID sql.NullInt32 `json:"updated_format_id"`
	AsesorID        int32         `json:"asesor_id"`
	FormatName      string        `json:"format_name"`
	Description     string        `json:"description"`
	Extension       Extensions    `json:"extension"`
	Version         string        `json:"version"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       sql.NullTime  `json:"updated_at"`
	DeletedAt       sql.NullTime  `json:"deleted_at"`
}

type Institution struct {
	InstitutionID   int64          `json:"institution_id"`
	AsesorID        sql.NullInt32  `json:"asesor_id"`
	InstitutionName string         `json:"institution_name"`
	Logo            sql.NullString `json:"logo"`
	Description     string         `json:"description"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       sql.NullTime   `json:"updated_at"`
	DeletedAt       sql.NullTime   `json:"deleted_at"`
}

type InstitutionClient struct {
	InstitutionClientID int64        `json:"institution_client_id"`
	ClientID            int32        `json:"client_id"`
	InstitutionID       int32        `json:"institution_id"`
	CreatedAt           time.Time    `json:"created_at"`
	UpdatedAt           sql.NullTime `json:"updated_at"`
	DeletedAt           sql.NullTime `json:"deleted_at"`
}

type InstitutionService struct {
	InstitutionServicesID int64        `json:"institution_services_id"`
	InstitutionID         int32        `json:"institution_id"`
	ServicesID            int32        `json:"services_id"`
	CreatedAt             time.Time    `json:"created_at"`
	UpdatedAt             sql.NullTime `json:"updated_at"`
	DeletedAt             sql.NullTime `json:"deleted_at"`
}

type Membership struct {
	MembershipID     int64         `json:"membership_id"`
	MembershipTypeID sql.NullInt32 `json:"membership_type_id"`
	CreatedAt        time.Time     `json:"created_at"`
	FinishAt         time.Time     `json:"finish_at"`
	UpdatedAt        sql.NullTime  `json:"updated_at"`
	DeletedAt        sql.NullTime  `json:"deleted_at"`
}

type MembershipType struct {
	MembershipTypeID int64        `json:"membership_type_id"`
	MembershipName   string       `json:"membership_name"`
	Users            int32        `json:"users"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        sql.NullTime `json:"updated_at"`
	DeletedAt        sql.NullTime `json:"deleted_at"`
}

type Permission struct {
	PermissionID   int64        `json:"permission_id"`
	PermissionType string       `json:"permission_type"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      sql.NullTime `json:"updated_at"`
	DeletedAt      sql.NullTime `json:"deleted_at"`
}

type Role struct {
	RoleID    int64        `json:"role_id"`
	RoleName  string       `json:"role_name"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

type RolePermission struct {
	RolePermissionID int64        `json:"role_permission_id"`
	RoleID           int32        `json:"role_id"`
	PermissionID     int32        `json:"permission_id"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        sql.NullTime `json:"updated_at"`
	DeletedAt        sql.NullTime `json:"deleted_at"`
}

type Service struct {
	ServicesID  int64        `json:"services_id"`
	ServiceName string       `json:"service_name"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
}
