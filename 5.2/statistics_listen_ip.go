// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 5.2.
package vtm

import (
	"encoding/json"
)

type ListenIpStatistics struct {
	Statistics struct {
		BytesOut        *int `json:"bytes_out"`
		TotalRequestsHi *int `json:"total_requests_hi"`
		TotalRequests   *int `json:"total_requests"`
		MaxConn         *int `json:"max_conn"`
		BytesOutHi      *int `json:"bytes_out_hi"`
		TotalConn       *int `json:"total_conn"`
		BytesInHi       *int `json:"bytes_in_hi"`
		BytesOutLo      *int `json:"bytes_out_lo"`
		CurrentConn     *int `json:"current_conn"`
		TotalRequestsLo *int `json:"total_requests_lo"`
		BytesInLo       *int `json:"bytes_in_lo"`
		BytesIn         *int `json:"bytes_in"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetListenIpStatistics(name string) (*ListenIpStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/5.2/status/local_tm/statistics/listen_ips/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(ListenIpStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
