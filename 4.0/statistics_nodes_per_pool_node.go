// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 4.0.
package vtm

import (
	"encoding/json"
)

type NodesPerPoolNodeStatistics struct {
	Statistics struct {
		ResponseMean       *int    `json:"response_mean"`
		PooledConn         *int    `json:"pooled_conn"`
		PktsToNodeHi       *int    `json:"pkts_to_node_hi"`
		BytesToNodeHi      *int    `json:"bytes_to_node_hi"`
		CurrentConn        *int    `json:"current_conn"`
		BytesFromNode      *int    `json:"bytes_from_node"`
		TotalConn          *int    `json:"total_conn"`
		BytesToNode        *int    `json:"bytes_to_node"`
		PktsFromNode       *int    `json:"pkts_from_node"`
		PktsToNode         *int    `json:"pkts_to_node"`
		PktsToNodeLo       *int    `json:"pkts_to_node_lo"`
		L4StatelessBuckets *int    `json:"l4_stateless_buckets"`
		ResponseMax        *int    `json:"response_max"`
		IdleConns          *int    `json:"idle_conns"`
		CurrentRequests    *int    `json:"current_requests"`
		ResponseMin        *int    `json:"response_min"`
		Errors             *int    `json:"errors"`
		Failures           *int    `json:"failures"`
		NewConn            *int    `json:"new_conn"`
		BytesFromNodeLo    *int    `json:"bytes_from_node_lo"`
		NodePort           *int    `json:"node_port"`
		State              *string `json:"state"`
		BytesToNodeLo      *int    `json:"bytes_to_node_lo"`
		PktsFromNodeHi     *int    `json:"pkts_from_node_hi"`
		PktsFromNodeLo     *int    `json:"pkts_from_node_lo"`
		BytesFromNodeHi    *int    `json:"bytes_from_node_hi"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetNodesPerPoolNodeStatistics(name string) (*NodesPerPoolNodeStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/4.0/status/local_tm/statistics/nodes/per_pool_node/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(NodesPerPoolNodeStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
