// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"gf-vue3-admin-server/app/dao/internal"
	"gf-vue3-admin-server/app/model"
)

// roleDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type roleDao struct {
	*internal.RoleDao
}

var (
	// Role is globally public accessible object for table admin_role operations.
	Role = &roleDao{
		internal.Role,
	}
)

// Fill with you ideas below.
func (r *roleDao) FindByNameOrAlias(name, alias string) (role *model.Role, err error) {
	if len(name) > 0 {
		r.Where(r.Columns.RoleName+"like", "%"+name+"%")
	}
	if len(alias) > 0 {
		r.Or(r.Columns.Alias+"like", "%"+alias+"%")
	}
	err = r.Scan(&role)
	return
}
