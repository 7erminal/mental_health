package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToUsersTable_20230725_143703 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToUsersTable_20230725_143703{}
	m.Created = "20230725_143703"

	migration.Register("AddColumnToUsersTable_20230725_143703", m)
}

// Run the migrations
func (m *AddColumnToUsersTable_20230725_143703) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE users ADD COLUMN `migration_status` varchar(20) DEFAULT NULL AFTER `id_number`")
}

// Reverse the migrations
func (m *AddColumnToUsersTable_20230725_143703) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
