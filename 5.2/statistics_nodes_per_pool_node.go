// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 5.2.
package vtm

import (
	"encoding/json"
)

type NodesPerPoolNodeStatistics struct {
	Statistics struct {
		PooledConn         *int    `json:"pooled_conn"`
		Failures           *int    `json:"failures"`
		NodePort           *int    `json:"node_port"`
		BytesToNodeHi      *int    `json:"bytes_to_node_hi"`
		BytesFromNodeHi    *int    `json:"bytes_from_node_hi"`
		ResponseMin        *int    `json:"response_min"`
		ResponseMax        *int    `json:"response_max"`
		Errors             *int    `json:"errors"`
		State              *string `json:"state"`
		BytesFromNodeLo    *int    `json:"bytes_from_node_lo"`
		PktsToNodeLo       *int    `json:"pkts_to_node_lo"`
		ResponseMean       *int    `json:"response_mean"`
		PktsFromNodeLo     *int    `json:"pkts_from_node_lo"`
		CurrentRequests    *int    `json:"current_requests"`
		BytesFromNode      *int    `json:"bytes_from_node"`
		IdleConns          *int    `json:"idle_conns"`
		L4StatelessBuckets *int    `json:"l4_stateless_buckets"`
		TotalConn          *int    `json:"total_conn"`
		PktsToNodeHi       *int    `json:"pkts_to_node_hi"`
		CurrentConn        *int    `json:"current_conn"`
		PktsFromNodeHi     *int    `json:"pkts_from_node_hi"`
		NewConn            *int    `json:"new_conn"`
		PktsFromNode       *int    `json:"pkts_from_node"`
		PktsToNode         *int    `json:"pkts_to_node"`
		BytesToNodeLo      *int    `json:"bytes_to_node_lo"`
		BytesToNode        *int    `json:"bytes_to_node"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetNodesPerPoolNodeStatistics(name string) (*NodesPerPoolNodeStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/5.2/status/local_tm/statistics/nodes/per_pool_node/" + name)
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
