// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 5.2.
package vtm

import (
	"encoding/json"
)

type NodesNodeInet46Statistics struct {
	Statistics struct {
		BytesToNode     *int    `json:"bytes_to_node"`
		ResponseMean    *int    `json:"response_mean"`
		PooledConn      *int    `json:"pooled_conn"`
		CurrentRequests *int    `json:"current_requests"`
		Failures        *int    `json:"failures"`
		IdleConns       *int    `json:"idle_conns"`
		BytesFromNode   *int    `json:"bytes_from_node"`
		BytesFromNodeHi *int    `json:"bytes_from_node_hi"`
		BytesToNodeHi   *int    `json:"bytes_to_node_hi"`
		TotalConn       *int    `json:"total_conn"`
		ResponseMin     *int    `json:"response_min"`
		CurrentConn     *int    `json:"current_conn"`
		Port            *int    `json:"port"`
		Errors          *int    `json:"errors"`
		NewConn         *int    `json:"new_conn"`
		State           *string `json:"state"`
		BytesFromNodeLo *int    `json:"bytes_from_node_lo"`
		BytesToNodeLo   *int    `json:"bytes_to_node_lo"`
		ResponseMax     *int    `json:"response_max"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetNodesNodeInet46Statistics(name string) (*NodesNodeInet46Statistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/5.2/status/local_tm/statistics/nodes/node_inet46/" + name)
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
