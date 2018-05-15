// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 5.2.
package vtm

import (
	"encoding/json"
)

type ListenIpStatistics struct {
	Statistics struct {
		MaxConn         *int `json:"max_conn"`
		BytesIn         *int `json:"bytes_in"`
		BytesOutHi      *int `json:"bytes_out_hi"`
		CurrentConn     *int `json:"current_conn"`
		BytesOut        *int `json:"bytes_out"`
		BytesOutLo      *int `json:"bytes_out_lo"`
		BytesInLo       *int `json:"bytes_in_lo"`
		TotalConn       *int `json:"total_conn"`
		BytesInHi       *int `json:"bytes_in_hi"`
		TotalRequestsHi *int `json:"total_requests_hi"`
		TotalRequests   *int `json:"total_requests"`
		TotalRequestsLo *int `json:"total_requests_lo"`
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
