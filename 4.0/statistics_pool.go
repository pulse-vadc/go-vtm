// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 4.0.
package vtm

import (
	"encoding/json"
)

type PoolStatistics struct {
	Statistics struct {
		BwLimitBytesDropLo *int    `json:"bw_limit_bytes_drop_lo"`
		Nodes              *int    `json:"nodes"`
		QueueTimeouts      *int    `json:"queue_timeouts"`
		Persistence        *string `json:"persistence"`
		BwLimitPktsDrop    *int    `json:"bw_limit_pkts_drop"`
		TotalConn          *int    `json:"total_conn"`
		MeanQueueTime      *int    `json:"mean_queue_time"`
		BytesOut           *int    `json:"bytes_out"`
		Algorithm          *string `json:"algorithm"`
		BytesInHi          *int    `json:"bytes_in_hi"`
		BytesIn            *int    `json:"bytes_in"`
		SessionMigrated    *int    `json:"session_migrated"`
		MaxQueueTime       *int    `json:"max_queue_time"`
		BytesOutHi         *int    `json:"bytes_out_hi"`
		BwLimitPktsDropLo  *int    `json:"bw_limit_pkts_drop_lo"`
		BwLimitPktsDropHi  *int    `json:"bw_limit_pkts_drop_hi"`
		BytesOutLo         *int    `json:"bytes_out_lo"`
		Draining           *int    `json:"draining"`
		Disabled           *int    `json:"disabled"`
		MinQueueTime       *int    `json:"min_queue_time"`
		ConnsQueued        *int    `json:"conns_queued"`
		State              *string `json:"state"`
		BwLimitBytesDropHi *int    `json:"bw_limit_bytes_drop_hi"`
		BwLimitBytesDrop   *int    `json:"bw_limit_bytes_drop"`
		BytesInLo          *int    `json:"bytes_in_lo"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetPoolStatistics(name string) (*PoolStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/4.0/status/local_tm/statistics/pools/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(PoolStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
