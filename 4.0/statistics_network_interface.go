// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 4.0.
package vtm

import (
	"encoding/json"
)

type NetworkInterfaceStatistics struct {
	Statistics struct {
		TxErrors   *int `json:"tx_errors"`
		TxPackets  *int `json:"tx_packets"`
		RxErrors   *int `json:"rx_errors"`
		TxBytes    *int `json:"tx_bytes"`
		RxBytesHi  *int `json:"rx_bytes_hi"`
		RxBytesLo  *int `json:"rx_bytes_lo"`
		TxBytesHi  *int `json:"tx_bytes_hi"`
		RxPackets  *int `json:"rx_packets"`
		RxBytes    *int `json:"rx_bytes"`
		TxBytesLo  *int `json:"tx_bytes_lo"`
		Collisions *int `json:"collisions"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetNetworkInterfaceStatistics(name string) (*NetworkInterfaceStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/4.0/status/local_tm/statistics/network_interface/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(NetworkInterfaceStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
