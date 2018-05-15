// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 5.2.
package vtm

import (
	"encoding/json"
)

type RuleStatistics struct {
	Statistics struct {
		ExecutionTimeWarnings *int `json:"execution_time_warnings"`
		Retries               *int `json:"retries"`
		PoolSelect            *int `json:"pool_select"`
		Aborts                *int `json:"aborts"`
		Discards              *int `json:"discards"`
		Executions            *int `json:"executions"`
		Responds              *int `json:"responds"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetRuleStatistics(name string) (*RuleStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/5.2/status/local_tm/statistics/rules/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(RuleStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
