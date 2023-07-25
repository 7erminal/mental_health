package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type RenameMaritalStatusColumnInUsersTable_20230725_162944 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &RenameMaritalStatusColumnInUsersTable_20230725_162944{}
	m.Created = "20230725_162944"

	migration.Register("RenameMaritalStatusColumnInUsersTable_20230725_162944", m)
}

// Run the migrations
func (m *RenameMaritalStatusColumnInUsersTable_20230725_162944) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE users CHANGE `migration_status` `marital_status` varchar(20)")
}

// Reverse the migrations
func (m *RenameMaritalStatusColumnInUsersTable_20230725_162944) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
