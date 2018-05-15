// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 4.0.
package vtm

import (
	"encoding/json"
)

type NetworkInterfaceStatistics struct {
	Statistics struct {
		TxBytes    *int `json:"tx_bytes"`
		TxBytesHi  *int `json:"tx_bytes_hi"`
		TxErrors   *int `json:"tx_errors"`
		RxBytes    *int `json:"rx_bytes"`
		RxBytesLo  *int `json:"rx_bytes_lo"`
		RxPackets  *int `json:"rx_packets"`
		TxPackets  *int `json:"tx_packets"`
		Collisions *int `json:"collisions"`
		RxErrors   *int `json:"rx_errors"`
		RxBytesHi  *int `json:"rx_bytes_hi"`
		TxBytesLo  *int `json:"tx_bytes_lo"`
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
