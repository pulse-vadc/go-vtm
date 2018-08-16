// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

package vtm

import (
	"os"
	"testing"
)

func TestNewVirtualTrafficManager(t *testing.T) {
	realBaseUrl := os.Getenv("VTM_BASE_URL")
	realUsername := os.Getenv("VTM_USERNAME")
	realPassword := os.Getenv("VTM_PASSWORD")

	connectionData := []struct {
		url       string
		username  string
		password  string
		reachable bool
	}{
		//		{"https://9.8.7.6:9070/api", "admin", "password", false},
		{realBaseUrl, realUsername, "not" + realPassword, false},
		{realBaseUrl, "not" + realUsername, realPassword, false},
		{realBaseUrl, realUsername, realPassword, true},
	}

	for _, conn := range connectionData {
		_, reachable, _ := NewVirtualTrafficManager(conn.url, conn.username, conn.password, false, true)
		if reachable != conn.reachable {
			t.Errorf("Reachability Error: %s with %s:%s - expected %t", conn.url, conn.username, conn.password, conn.reachable)
		}
	}
}
