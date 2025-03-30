package entity

import "time"

type Staff struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Role      StaffRole `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type StaffRole int

const (
	RoleBarista StaffRole = iota
	RoleCashier
	RoleManager
)

func (s StaffRole) String() string {
	switch s {
	case RoleBarista:
		return "barista"
	case RoleCashier:
		return "cashier"
	case RoleManager:
		return "manager"
	default:
		return "unknown"
	}
}

func (s StaffRole) IsValid() bool {
	switch s {
	case RoleBarista, RoleCashier, RoleManager:
		return true
	}
	return false
}
