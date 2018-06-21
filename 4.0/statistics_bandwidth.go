// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 4.0.
package vtm

import (
	"encoding/json"
)

type BandwidthStatistics struct {
	Statistics struct {
		BytesDropLo *int `json:"bytes_drop_lo"`
		BytesDrop   *int `json:"bytes_drop"`
		BytesDropHi *int `json:"bytes_drop_hi"`
		BytesOutHi  *int `json:"bytes_out_hi"`
		BytesOut    *int `json:"bytes_out"`
		Guarantee   *int `json:"guarantee"`
		Maximum     *int `json:"maximum"`
		PktsDropHi  *int `json:"pkts_drop_hi"`
		PktsDrop    *int `json:"pkts_drop"`
		PktsDropLo  *int `json:"pkts_drop_lo"`
		BytesOutLo  *int `json:"bytes_out_lo"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetBandwidthStatistics(name string) (*BandwidthStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/4.0/status/local_tm/statistics/bandwidth/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(BandwidthStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
