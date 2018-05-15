// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 4.0.
package vtm

import (
	"encoding/json"
)

type NodesNodeInet46Statistics struct {
	Statistics struct {
		ResponseMean    *int    `json:"response_mean"`
		PooledConn      *int    `json:"pooled_conn"`
		Failures        *int    `json:"failures"`
		NewConn         *int    `json:"new_conn"`
		IdleConns       *int    `json:"idle_conns"`
		CurrentConn     *int    `json:"current_conn"`
		BytesFromNodeLo *int    `json:"bytes_from_node_lo"`
		BytesFromNode   *int    `json:"bytes_from_node"`
		TotalConn       *int    `json:"total_conn"`
		BytesToNode     *int    `json:"bytes_to_node"`
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

func (vtm VirtualTrafficManager) GetNodesNodeInet46Statistics(name string) (*NodesNodeInet46Statistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/4.0/status/local_tm/statistics/nodes/node_inet46/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(NodesNodeInet46Statistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
