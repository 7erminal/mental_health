package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type UserAssessmentChoices_20230817_183121 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UserAssessmentChoices_20230817_183121{}
	m.Created = "20230817_183121"

	migration.Register("UserAssessmentChoices_20230817_183121", m)
}

// Run the migrations
func (m *UserAssessmentChoices_20230817_183121) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE user_assessment_choices(`user_assessment_choice_id` bigInt(11) NOT NULL AUTO_INCREMENT,`assessment_question_id` int(11) DEFAULT NULL,`assessment_choice` int(11) DEFAULT NULL,`active` int(11) DEFAULT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` bigInt(11),`modified_by` int(11) DEFAULT NULL,PRIMARY KEY (`user_assessment_choice_id`), FOREIGN KEY (assessment_question_id) REFERENCES assessment_questions(assessment_question_id) ON UPDATE CASCADE ON DELETE NO ACTION, FOREIGN KEY (assessment_choice) REFERENCES assessment_options(assessment_option_id) ON UPDATE CASCADE ON DELETE NO ACTION, FOREIGN KEY (created_by) REFERENCES users(user_id) ON UPDATE CASCADE ON DELETE CASCADE)")
}

// Reverse the migrations
func (m *UserAssessmentChoices_20230817_183121) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `user_assessment_choices`")
}
