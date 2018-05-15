// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 4.0.
package vtm

import (
	"encoding/json"
)

type NodesNodeStatistics struct {
	Statistics struct {
		ResponseMean    *int    `json:"response_mean"`
		PooledConn      *int    `json:"pooled_conn"`
		Failures        *int    `json:"failures"`
		NewConn         *int    `json:"new_conn"`
		CurrentConn     *int    `json:"current_conn"`
		BytesFromNodeLo *int    `json:"bytes_from_node_lo"`
		TotalConn       *int    `json:"total_conn"`
		Port            *int    `json:"port"`
		State           *string `json:"state"`
		BytesToNodeLo   *int    `json:"bytes_to_node_lo"`
		ResponseMax     *int    `json:"response_max"`
		BytesToNodeHi   *int    `json:"bytes_to_node_hi"`
		CurrentRequests *int    `json:"current_requests"`
		BytesFromNodeHi *int    `json:"bytes_from_node_hi"`
		ResponseMin     *int    `json:"response_min"`
		Errors          *int    `json:"errors"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetNodesNodeStatistics(name string) (*NodesNodeStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/4.0/status/local_tm/statistics/nodes/node/" + name)
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
