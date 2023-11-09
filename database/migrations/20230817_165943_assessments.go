package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Assessments_20230817_165943 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Assessments_20230817_165943{}
	m.Created = "20230817_165943"

	migration.Register("Assessments_20230817_165943", m)
}

// Run the migrations
func (m *Assessments_20230817_165943) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE assessments(`assessment_id` int(11) NOT NULL AUTO_INCREMENT,`assessment_title` varchar(255) NOT NULL,`instruction` longtext  NOT NULL,`active` int(11) DEFAULT 1,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,PRIMARY KEY (`assessment_id`))")
}

// Reverse the migrations
func (m *Assessments_20230817_165943) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `assessments`")
}
