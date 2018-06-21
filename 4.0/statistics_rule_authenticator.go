// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 4.0.
package vtm

import (
	"encoding/json"
)

type RuleAuthenticatorStatistics struct {
	Statistics struct {
		Passes   *int `json:"passes"`
		Requests *int `json:"requests"`
		Fails    *int `json:"fails"`
		Errors   *int `json:"errors"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetRuleAuthenticatorStatistics(name string) (*RuleAuthenticatorStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/4.0/status/local_tm/statistics/rule_authenticators/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(RuleAuthenticatorStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
