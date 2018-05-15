// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 5.2.
package vtm

import (
	"encoding/json"
)

type PoolStatistics struct {
	Statistics struct {
		BytesOutHi         *int    `json:"bytes_out_hi"`
		SessionMigrated    *int    `json:"session_migrated"`
		BwLimitBytesDrop   *int    `json:"bw_limit_bytes_drop"`
		MinQueueTime       *int    `json:"min_queue_time"`
		BytesOut           *int    `json:"bytes_out"`
		BwLimitPktsDropLo  *int    `json:"bw_limit_pkts_drop_lo"`
		BwLimitBytesDropLo *int    `json:"bw_limit_bytes_drop_lo"`
		MeanQueueTime      *int    `json:"mean_queue_time"`
		Disabled           *int    `json:"disabled"`
		BytesInHi          *int    `json:"bytes_in_hi"`
		Algorithm          *string `json:"algorithm"`
		Persistence        *string `json:"persistence"`
		BytesInLo          *int    `json:"bytes_in_lo"`
		BwLimitPktsDropHi  *int    `json:"bw_limit_pkts_drop_hi"`
		BytesIn            *int    `json:"bytes_in"`
		BwLimitBytesDropHi *int    `json:"bw_limit_bytes_drop_hi"`
		QueueTimeouts      *int    `json:"queue_timeouts"`
		Nodes              *int    `json:"nodes"`
		BytesOutLo         *int    `json:"bytes_out_lo"`
		ConnsQueued        *int    `json:"conns_queued"`
		State              *string `json:"state"`
		BwLimitPktsDrop    *int    `json:"bw_limit_pkts_drop"`
		Draining           *int    `json:"draining"`
		TotalConn          *int    `json:"total_conn"`
		MaxQueueTime       *int    `json:"max_queue_time"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetPoolStatistics(name string) (*PoolStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/5.2/status/local_tm/statistics/pools/" + name)
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
