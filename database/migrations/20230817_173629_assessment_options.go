package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AssessmentOptions_20230817_173629 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AssessmentOptions_20230817_173629{}
	m.Created = "20230817_173629"

	migration.Register("AssessmentOptions_20230817_173629", m)
}

// Run the migrations
func (m *AssessmentOptions_20230817_173629) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE assessment_options(`assessment_option_id` int(11) NOT NULL AUTO_INCREMENT,`assessment_id` int(11) DEFAULT NULL,`assessment_option` varchar(255) NOT NULL,`active` int(11) DEFAULT 1,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,PRIMARY KEY (`assessment_option_id`), FOREIGN KEY (assessment_id) REFERENCES assessments(assessment_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *AssessmentOptions_20230817_173629) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `assessment_options`")
}
