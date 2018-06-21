// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 5.2.
package vtm

import (
	"encoding/json"
)

type PerNodeSlmPerNodeServiceLevelStatistics struct {
	Statistics struct {
		ResponseMax  *int `json:"response_max"`
		ResponseMean *int `json:"response_mean"`
		TotalNonConf *int `json:"total_non_conf"`
		NodePort     *int `json:"node_port"`
		TotalConn    *int `json:"total_conn"`
		ResponseMin  *int `json:"response_min"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetPerNodeSlmPerNodeServiceLevelStatistics(name string) (*PerNodeSlmPerNodeServiceLevelStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/5.2/status/local_tm/statistics/per_node_slm/per_node_service_level/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(PerNodeSlmPerNodeServiceLevelStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
