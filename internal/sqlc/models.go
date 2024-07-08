// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlc

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
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

type StatusClient string

const (
	StatusClientActivo   StatusClient = "activo"
	StatusClientInactivo StatusClient = "inactivo"
)

func (e *StatusClient) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = StatusClient(s)
	case string:
		*e = StatusClient(s)
	default:
		return fmt.Errorf("unsupported scan type for StatusClient: %T", src)
	}
	return nil
}

type NullStatusClient struct {
	StatusClient StatusClient `json:"status_client"`
	Valid        bool         `json:"valid"` // Valid is true if StatusClient is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStatusClient) Scan(value interface{}) error {
	if value == nil {
		ns.StatusClient, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.StatusClient.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStatusClient) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.StatusClient), nil
}

type Asesor struct {
	AsesorID  int32            `json:"asesor_id"`
	About     string           `json:"about"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	DeletedAt pgtype.Timestamp `json:"deleted_at"`
}

type Client struct {
	ClientID  int32            `json:"client_id"`
	GivenName string           `json:"given_name"`
	Surname   string           `json:"surname"`
	Photo     pgtype.Text      `json:"photo"`
	Email     string           `json:"email"`
	Password  string           `json:"password"`
	Status    StatusClient     `json:"status"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	DeletedAt pgtype.Timestamp `json:"deleted_at"`
}

type ClientRole struct {
	ClientRoleID int64            `json:"client_role_id"`
	ClientID     int32            `json:"client_id"`
	RoleID       int32            `json:"role_id"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
	DeletedAt    pgtype.Timestamp `json:"deleted_at"`
}

type DirectoryInstitution struct {
	DirectoryInstitutionID int32            `json:"directory_institution_id"`
	InstitutionID          int32            `json:"institution_id"`
	DirectoryID            int32            `json:"directory_id"`
	CreatedAt              pgtype.Timestamp `json:"created_at"`
	UpdatedAt              pgtype.Timestamp `json:"updated_at"`
	DeletedAt              pgtype.Timestamp `json:"deleted_at"`
}

type DirectoryRole struct {
	DirectoryRoleID int64            `json:"directory_role_id"`
	DirectoryID     pgtype.Int4      `json:"directory_id"`
	RoleID          pgtype.Int4      `json:"role_id"`
	CreatedAt       pgtype.Timestamp `json:"created_at"`
	UpdatedAt       pgtype.Timestamp `json:"updated_at"`
	DeletedAt       pgtype.Timestamp `json:"deleted_at"`
}

type DirectoryTree struct {
	DirectoryID int64            `json:"directory_id"`
	ParentID    pgtype.Int4      `json:"parent_id"`
	Name        pgtype.Text      `json:"name"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
	DeletedAt   pgtype.Timestamp `json:"deleted_at"`
}

type Document struct {
	DocumentID  int64            `json:"document_id"`
	DirectoryID int32            `json:"directory_id"`
	FormatID    pgtype.Int4      `json:"format_id"`
	Name        string           `json:"name"`
	FileRute    string           `json:"file_rute"`
	Status      Status           `json:"status"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
	DeletedAt   pgtype.Timestamp `json:"deleted_at"`
}

type DocumentClient struct {
	DocumentClientID int64            `json:"document_client_id"`
	ClientID         int32            `json:"client_id"`
	DocumentID       int32            `json:"document_id"`
	Action           Action           `json:"action"`
	CreatedAt        pgtype.Timestamp `json:"created_at"`
	UpdatedAt        pgtype.Timestamp `json:"updated_at"`
	DeletedAt        pgtype.Timestamp `json:"deleted_at"`
}

type Format struct {
	FormatID        int32            `json:"format_id"`
	UpdatedFormatID pgtype.Int4      `json:"updated_format_id"`
	AsesorID        int32            `json:"asesor_id"`
	FormatName      string           `json:"format_name"`
	Description     string           `json:"description"`
	Extension       Extensions       `json:"extension"`
	Version         string           `json:"version"`
	CreatedAt       pgtype.Timestamp `json:"created_at"`
	UpdatedAt       pgtype.Timestamp `json:"updated_at"`
	DeletedAt       pgtype.Timestamp `json:"deleted_at"`
}

type Institution struct {
	InstitutionID   int32            `json:"institution_id"`
	AsesorID        pgtype.Int4      `json:"asesor_id"`
	InstitutionName string           `json:"institution_name"`
	Logo            pgtype.Text      `json:"logo"`
	Description     string           `json:"description"`
	CreatedAt       pgtype.Timestamp `json:"created_at"`
	UpdatedAt       pgtype.Timestamp `json:"updated_at"`
	DeletedAt       pgtype.Timestamp `json:"deleted_at"`
}

type InstitutionClient struct {
	ClientID      int32            `json:"client_id"`
	InstitutionID int32            `json:"institution_id"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
	DeletedAt     pgtype.Timestamp `json:"deleted_at"`
}

type InstitutionService struct {
	InstitutionID int32            `json:"institution_id"`
	ServiceID     int32            `json:"service_id"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
	DeletedAt     pgtype.Timestamp `json:"deleted_at"`
}

type Permission struct {
	PermissionID int32            `json:"permission_id"`
	Name         string           `json:"name"`
	Code         string           `json:"code"`
	Description  string           `json:"description"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
	DeletedAt    pgtype.Timestamp `json:"deleted_at"`
}

type Role struct {
	RoleID      int32            `json:"role_id"`
	RoleName    string           `json:"role_name"`
	Description string           `json:"description"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
	DeletedAt   pgtype.Timestamp `json:"deleted_at"`
}

type RolePermission struct {
	RolePermissionID int64            `json:"role_permission_id"`
	RoleID           int32            `json:"role_id"`
	PermissionID     int32            `json:"permission_id"`
	CreatedAt        pgtype.Timestamp `json:"created_at"`
	UpdatedAt        pgtype.Timestamp `json:"updated_at"`
	DeletedAt        pgtype.Timestamp `json:"deleted_at"`
}

type Service struct {
	ServiceID int32            `json:"service_id"`
	Name      string           `json:"name"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	DeletedAt pgtype.Timestamp `json:"deleted_at"`
}

type Standard struct {
	StandardID int32            `json:"standard_id"`
	ServiceID  int32            `json:"service_id"`
	Name       string           `json:"name"`
	Complexity pgtype.Text      `json:"complexity"`
	Modality   string           `json:"modality"`
	Article    string           `json:"article"`
	Section    string           `json:"section"`
	Paragraph  pgtype.Text      `json:"paragraph"`
	Criteria   string           `json:"criteria"`
	Comply     pgtype.Bool      `json:"comply"`
	Applys     pgtype.Bool      `json:"applys"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
	UpdatedAt  pgtype.Timestamp `json:"updated_at"`
	DeletedAt  pgtype.Timestamp `json:"deleted_at"`
}
