// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 4.0.
package vtm

import (
	"encoding/json"
)

type PoolStatistics struct {
	Statistics struct {
		BwLimitBytesDropHi *int    `json:"bw_limit_bytes_drop_hi"`
		TotalConn          *int    `json:"total_conn"`
		BytesOut           *int    `json:"bytes_out"`
		Persistence        *string `json:"persistence"`
		Draining           *int    `json:"draining"`
		BwLimitPktsDropLo  *int    `json:"bw_limit_pkts_drop_lo"`
		BytesOutLo         *int    `json:"bytes_out_lo"`
		BwLimitPktsDropHi  *int    `json:"bw_limit_pkts_drop_hi"`
		State              *string `json:"state"`
		MaxQueueTime       *int    `json:"max_queue_time"`
		BytesInLo          *int    `json:"bytes_in_lo"`
		MinQueueTime       *int    `json:"min_queue_time"`
		Nodes              *int    `json:"nodes"`
		MeanQueueTime      *int    `json:"mean_queue_time"`
		BytesOutHi         *int    `json:"bytes_out_hi"`
		QueueTimeouts      *int    `json:"queue_timeouts"`
		Disabled           *int    `json:"disabled"`
		BytesInHi          *int    `json:"bytes_in_hi"`
		BwLimitBytesDropLo *int    `json:"bw_limit_bytes_drop_lo"`
		SessionMigrated    *int    `json:"session_migrated"`
		BytesIn            *int    `json:"bytes_in"`
		ConnsQueued        *int    `json:"conns_queued"`
		Algorithm          *string `json:"algorithm"`
		BwLimitBytesDrop   *int    `json:"bw_limit_bytes_drop"`
		BwLimitPktsDrop    *int    `json:"bw_limit_pkts_drop"`
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
