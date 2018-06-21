// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 4.0.
package vtm

import (
	"encoding/json"
)

type NodesPerPoolNodeStatistics struct {
	Statistics struct {
		PktsToNode         *int    `json:"pkts_to_node"`
		TotalConn          *int    `json:"total_conn"`
		IdleConns          *int    `json:"idle_conns"`
		BytesToNodeLo      *int    `json:"bytes_to_node_lo"`
		PktsFromNode       *int    `json:"pkts_from_node"`
		ResponseMin        *int    `json:"response_min"`
		L4StatelessBuckets *int    `json:"l4_stateless_buckets"`
		BytesToNode        *int    `json:"bytes_to_node"`
		State              *string `json:"state"`
		Failures           *int    `json:"failures"`
		ResponseMean       *int    `json:"response_mean"`
		CurrentRequests    *int    `json:"current_requests"`
		BytesToNodeHi      *int    `json:"bytes_to_node_hi"`
		BytesFromNode      *int    `json:"bytes_from_node"`
		PktsToNodeLo       *int    `json:"pkts_to_node_lo"`
		PooledConn         *int    `json:"pooled_conn"`
		NodePort           *int    `json:"node_port"`
		PktsToNodeHi       *int    `json:"pkts_to_node_hi"`
		PktsFromNodeLo     *int    `json:"pkts_from_node_lo"`
		NewConn            *int    `json:"new_conn"`
		CurrentConn        *int    `json:"current_conn"`
		BytesFromNodeHi    *int    `json:"bytes_from_node_hi"`
		BytesFromNodeLo    *int    `json:"bytes_from_node_lo"`
		Errors             *int    `json:"errors"`
		PktsFromNodeHi     *int    `json:"pkts_from_node_hi"`
		ResponseMax        *int    `json:"response_max"`
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
