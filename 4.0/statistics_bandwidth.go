// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 4.0.
package vtm

import (
	"encoding/json"
)

type BandwidthStatistics struct {
	Statistics struct {
		Maximum     *int `json:"maximum"`
		PktsDropHi  *int `json:"pkts_drop_hi"`
		BytesDropHi *int `json:"bytes_drop_hi"`
		Guarantee   *int `json:"guarantee"`
		BytesOutLo  *int `json:"bytes_out_lo"`
		PktsDrop    *int `json:"pkts_drop"`
		BytesDropLo *int `json:"bytes_drop_lo"`
		PktsDropLo  *int `json:"pkts_drop_lo"`
		BytesOutHi  *int `json:"bytes_out_hi"`
		BytesDrop   *int `json:"bytes_drop"`
		BytesOut    *int `json:"bytes_out"`
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
