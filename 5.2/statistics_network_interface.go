// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 5.2.
package vtm

import (
	"encoding/json"
)

type NetworkInterfaceStatistics struct {
	Statistics struct {
		RxPackets  *int `json:"rx_packets"`
		RxBytesHi  *int `json:"rx_bytes_hi"`
		TxBytesHi  *int `json:"tx_bytes_hi"`
		TxErrors   *int `json:"tx_errors"`
		TxBytesLo  *int `json:"tx_bytes_lo"`
		RxErrors   *int `json:"rx_errors"`
		Collisions *int `json:"collisions"`
		RxBytesLo  *int `json:"rx_bytes_lo"`
		TxPackets  *int `json:"tx_packets"`
		TxBytes    *int `json:"tx_bytes"`
		RxBytes    *int `json:"rx_bytes"`
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
