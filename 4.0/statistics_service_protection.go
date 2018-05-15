// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 4.0.
package vtm

import (
	"encoding/json"
)

type ServiceProtectionStatistics struct {
	Statistics struct {
		RefusalRfc2396  *int `json:"refusal_rfc2396"`
		RefusalSize     *int `json:"refusal_size"`
		TotalRefusal    *int `json:"total_refusal"`
		RefusalIp       *int `json:"refusal_ip"`
		RefusalBinary   *int `json:"refusal_binary"`
		RefusalConc10Ip *int `json:"refusal_conc10_ip"`
		RefusalConnRate *int `json:"refusal_conn_rate"`
		LastRefusalTime *int `json:"last_refusal_time"`
		RefusalConc1Ip  *int `json:"refusal_conc1_ip"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetServiceProtectionStatistics(name string) (*ServiceProtectionStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/4.0/status/local_tm/statistics/service_protection/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(ServiceProtectionStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}