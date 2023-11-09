package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type MotivationalQuotes_20230817_182929 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &MotivationalQuotes_20230817_182929{}
	m.Created = "20230817_182929"

	migration.Register("MotivationalQuotes_20230817_182929", m)
}

// Run the migrations
func (m *MotivationalQuotes_20230817_182929) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE motivational_quotes(`motivational_quote_id` int(11) NOT NULL AUTO_INCREMENT,`motivational_quote` varchar(255) NOT NULL,`active` int(11) DEFAULT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,PRIMARY KEY (`motivational_quote_id`))")
}

// Reverse the migrations
func (m *MotivationalQuotes_20230817_182929) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `motivational_quotes`")
}
