// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 5.2.
package vtm

import (
	"encoding/json"
)

type NetworkInterfaceStatistics struct {
	Statistics struct {
		RxBytesLo  *int `json:"rx_bytes_lo"`
		RxBytes    *int `json:"rx_bytes"`
		RxPackets  *int `json:"rx_packets"`
		TxBytes    *int `json:"tx_bytes"`
		TxBytesLo  *int `json:"tx_bytes_lo"`
		TxBytesHi  *int `json:"tx_bytes_hi"`
		RxErrors   *int `json:"rx_errors"`
		TxPackets  *int `json:"tx_packets"`
		Collisions *int `json:"collisions"`
		RxBytesHi  *int `json:"rx_bytes_hi"`
		TxErrors   *int `json:"tx_errors"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetNetworkInterfaceStatistics(name string) (*NetworkInterfaceStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/5.2/status/local_tm/statistics/network_interface/" + name)
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
