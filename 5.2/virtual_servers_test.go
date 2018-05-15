// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

package vtm

import (
	"reflect"
	"testing"
)

func getVirtualServer(name string) *VirtualServer {
	vs, _ := testTm.GetVirtualServer(name)
	return vs
}

func TestNewVirtualServer(t *testing.T) {
	vs := testTm.NewVirtualServer("Test-VS-1", "discard", 1234)
	vs, err := vs.Apply()
	if err != nil {
		t.Errorf("Failed to create valid virtual server: %#v", err)
	}
}

func TestGetVirtualServer(t *testing.T) {
	// Test getting an existing virtual server
	vs, _ := testTm.GetVirtualServer("Test-VS-1")
	if *vs.Basic.Port != 1234 {
		t.Errorf("Virtual server has incorrect port %d", vs.Basic.Port)
	}

	// Test getting an non-existant virtual server
	_, err := testTm.GetVirtualServer("MadeUpVS")
	if err == nil {
		t.Errorf("Getting non-existent VS didn't error")
	} else if err.ErrorId != "resource.not_found" {
		t.Errorf("Incorrect error ID (%s) returned", err.ErrorId)
	}
}

func TestListVirtualServers(t *testing.T) {
	// Check target virtual server does actually exist
	vs := getVirtualServer("Test-VS-1")
	if vs == nil {
		t.Errorf("Virtual server to delete doesn't exist")
		return
	}

	vsList, err := testTm.ListVirtualServers()
	if err != nil {
		t.Errorf("Error listing virtual servers: %s", err.ErrorId)
	}
	testVsFound := false
	for _, vs := range *vsList {
		if vs == "Test-VS-1" {
			testVsFound = true
		}
	}
	if testVsFound == false {
		t.Errorf("Failed to list test virtual server")
	}
}

func TestApply(t *testing.T) {
	// Check correct start values
	vs := getVirtualServer("Test-VS-1")
	if *vs.Basic.Port != 1234 {
		t.Errorf("Virtual Server port not correctly initialized to 1234")
		return
	}
	/*
		INDIVIDUAL DATA TYPE TESTING - STRING
	*/
	// Test setting enum string (valid)
	*vs.Basic.Protocol = "client_first"
	_, err := vs.Apply()
	if err != nil {
		t.Errorf("Error applying valid protocol: %s", err.ErrorId)
	}
	vsUpdated := getVirtualServer("Test-VS-1")
	if *vsUpdated.Basic.Protocol != "client_first" {
		t.Errorf("Virtual server protocol was not set to client_first")
	}
	// Test setting enum string (invalid)
	*vs.Basic.Protocol = "wizard_magic"
	_, err = vs.Apply()
	if err == nil {
		t.Errorf("No error applying invalid protocol")
	}

	/*
		INDIVIDUAL DATA TYPE TESTING - UNSIGNED INT
	*/
	// Test setting bounded unsigned int (valid)
	vs = getVirtualServer("Test-VS-1")
	*vs.Basic.Port = 4321
	_, err = vs.Apply()
	if err != nil {
		t.Errorf("Error applying valid port: %s", err.ErrorId)
	}
	vsUpdated = getVirtualServer("Test-VS-1")
	if *vsUpdated.Basic.Port != 4321 {
		t.Errorf("Virtual Server port was not set to 4321")
	}
	// Test setting bounded unsigned int (invalid - too low)
	*vs.Basic.Port = 0
	_, err = vs.Apply()
	if err == nil {
		t.Errorf("Error applying invalid (too low) port: %s", err.ErrorId)
	} else if err.ErrorId != "resource.validation_error" {
		t.Errorf("Incorrect error ID (%s) returned (expected resource.validation_error)", err.ErrorId)
	}
	// Test setting bounded unsigned int (invalid - too high)
	*vs.Basic.Port = 70000
	_, err = vs.Apply()
	if err == nil {
		t.Errorf("Error applying invalid (too high) port: %s", err.ErrorId)
	} else if err.ErrorId != "resource.validation_error" {
		t.Errorf("Incorrect error ID (%s) returned (expected resource.validation_error)", err.ErrorId)
	}

	/*
		INDIVIDUAL DATA TYPE TESTING - STRING LIST (NON-UNIQUE VALUES)
	*/
	// Test setting list to non-empty (valid)
	vs = getVirtualServer("Test-VS-1")
	rulesList := []string{"rule1", "rule2", "rule1"}
	*vs.Basic.RequestRules = rulesList
	_, err = vs.Apply()
	if err != nil {
		t.Errorf("Failed to set request rules to valid list: %s", err.ErrorId)
	}
	vsUpdated = getVirtualServer("Test-VS-1")
	if !reflect.DeepEqual(*vsUpdated.Basic.RequestRules, rulesList) {
		t.Errorf("Request rules list was not correctly set: %s", *vsUpdated.Basic.RequestRules)
	}
	// Test setting list to empty (valid)
	*vs.Basic.RequestRules = []string{}
	_, err = vs.Apply()
	if err != nil {
		t.Errorf("Failed to set request rules to valid list: %s", err.ErrorId)
	}
	vsUpdated = getVirtualServer("Test-VS-1")
	if !reflect.DeepEqual(*vsUpdated.Basic.RequestRules, []string{}) {
		t.Errorf("Request rules list was not correctly set to empty")
	}
}

func TestDelete(t *testing.T) {
	// Check target virtual server does actually exist
	vs := getVirtualServer("Test-VS-1")
	if vs == nil {
		t.Errorf("Virtual server to delete doesn't exist")
		return
	}

	// Test deleting an existing virtual server
	err := testTm.DeleteVirtualServer("Test-VS-1")
	if err != nil {
		t.Errorf("Failed to delete virtual server: %s", err.ErrorId)
	}
	vs = getVirtualServer("Test-VS-1")
	if vs != nil {
		t.Errorf("Deleted virtual server still appears to exist")
	}

	// Test deleting a non-existing virtual server
	err = testTm.DeleteVirtualServer("MadeUpVS")
	if err == nil {
		t.Errorf("Attempt to delete non-existant virtual server erroneously succeeded")
	}
}
