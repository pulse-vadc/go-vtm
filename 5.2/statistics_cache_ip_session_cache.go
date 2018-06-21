// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 5.2.
package vtm

import (
	"encoding/json"
)

type CacheIpSessionCacheStatistics struct {
	Statistics struct {
		Misses     *int `json:"misses"`
		Lookups    *int `json:"lookups"`
		EntriesMax *int `json:"entries_max"`
		Entries    *int `json:"entries"`
		Oldest     *int `json:"oldest"`
		Hits       *int `json:"hits"`
		HitRate    *int `json:"hit_rate"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetCacheIpSessionCacheStatistics() (*CacheIpSessionCacheStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/5.2/status/local_tm/statistics/cache/ip_session_cache")
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(CacheIpSessionCacheStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
