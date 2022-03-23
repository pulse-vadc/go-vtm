// Copyright (C) 2018-2022, Pulse Secure, LLC.
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

package vtm

import (
	"testing"
)

func TestSetRule(t *testing.T) {
	rule := "log.info('Hello World!');"
	err := testTm.SetRule("Test-Rule-1", rule)
	if err != nil {
		t.Errorf("Failed to set rule: %s", err.ErrorId)
	}
}

func TestListRules(t *testing.T) {
	ruleList, err := testTm.ListRules()
	if err != nil {
		t.Errorf("Failed to list rules: %s", err.ErrorId)
	}
	testRuleFound := false
	for _, rule := range *ruleList {
		if rule == "Test-Rule-1" {
			testRuleFound = true
		}
	}
	if testRuleFound == false {
		t.Errorf("Failed to list test rule")
	}
}

func TestGetRule(t *testing.T) {
	// Test get existing rule
	rule, err := testTm.GetRule("Test-Rule-1")
	if err != nil {
		t.Errorf("Error getting existing rule: %s", err.ErrorId)
	} else if rule != "log.info('Hello World!');" {
		t.Errorf("Rule contains wrong value - %s", rule)
	}
	// Test get non-existing rule
	rule, err = testTm.GetRule("MadeUpRule")
	if err == nil {
		t.Errorf("Attempt to get non-existent rule erroneously succeeded: %s", rule)
	}
}

func TestDeleteRule(t *testing.T) {
	// Test deleting non-existant rule
	err := testTm.DeleteRule("MadeUpRule")
	if err == nil {
		t.Errorf("Attempt to delete non-existent rule erroneously succeeded")
	}

	// Test deleting existing rule
	err = testTm.DeleteRule("Test-Rule-1")
	if err != nil {
		t.Errorf("Attempt to delete rule failed: %s", err.ErrorId)
	}
}
