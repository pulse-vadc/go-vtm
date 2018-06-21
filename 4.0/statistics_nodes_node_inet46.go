// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 4.0.
package vtm

import (
	"encoding/json"
)

type NodesNodeInet46Statistics struct {
	Statistics struct {
		TotalConn       *int    `json:"total_conn"`
		IdleConns       *int    `json:"idle_conns"`
		Port            *int    `json:"port"`
		BytesToNodeLo   *int    `json:"bytes_to_node_lo"`
		BytesFromNodeHi *int    `json:"bytes_from_node_hi"`
		ResponseMean    *int    `json:"response_mean"`
		NewConn         *int    `json:"new_conn"`
		BytesToNode     *int    `json:"bytes_to_node"`
		BytesFromNode   *int    `json:"bytes_from_node"`
		BytesFromNodeLo *int    `json:"bytes_from_node_lo"`
		State           *string `json:"state"`
		Failures        *int    `json:"failures"`
		ResponseMin     *int    `json:"response_min"`
		CurrentRequests *int    `json:"current_requests"`
		BytesToNodeHi   *int    `json:"bytes_to_node_hi"`
		Errors          *int    `json:"errors"`
		ResponseMax     *int    `json:"response_max"`
		CurrentConn     *int    `json:"current_conn"`
		PooledConn      *int    `json:"pooled_conn"`
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
