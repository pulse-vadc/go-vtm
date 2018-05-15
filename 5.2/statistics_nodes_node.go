// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 5.2.
package vtm

import (
	"encoding/json"
)

type NodesNodeStatistics struct {
	Statistics struct {
		CurrentRequests *int    `json:"current_requests"`
		NewConn         *int    `json:"new_conn"`
		CurrentConn     *int    `json:"current_conn"`
		Errors          *int    `json:"errors"`
		BytesFromNodeLo *int    `json:"bytes_from_node_lo"`
		PooledConn      *int    `json:"pooled_conn"`
		State           *string `json:"state"`
		ResponseMean    *int    `json:"response_mean"`
		Port            *int    `json:"port"`
		TotalConn       *int    `json:"total_conn"`
		BytesToNodeLo   *int    `json:"bytes_to_node_lo"`
		Failures        *int    `json:"failures"`
		ResponseMin     *int    `json:"response_min"`
		ResponseMax     *int    `json:"response_max"`
		BytesFromNodeHi *int    `json:"bytes_from_node_hi"`
		BytesToNodeHi   *int    `json:"bytes_to_node_hi"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetNodesNodeStatistics(name string) (*NodesNodeStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/5.2/status/local_tm/statistics/nodes/node/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(NodesNodeStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
