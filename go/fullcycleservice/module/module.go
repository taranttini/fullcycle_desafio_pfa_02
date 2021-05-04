package module

import "database/sql/driver"

// module
type Module struct {
	ModuleID int    `json:"moduleId"`
	Name     string `json:"name"`
	Active   bool   `json:"active"`
}

type ModuleFilter struct {
	NameFilter   string `json:"name"`
	ActiveFilter string `json:"active"`
}

type BitBool bool

// Value implements the driver.Valuer interface,
// and turns the BitBool into a bitfield (BIT(1)) for MySQL storage.
func (b BitBool) Value() (driver.Value, error) {
	if b {
		return true, nil
	} else {
		return false, nil
	}
}
