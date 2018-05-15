// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 5.2.
package vtm

import (
	"encoding/json"
)

type CacheSslCacheStatistics struct {
	Statistics struct {
		Lookups    *int `json:"lookups"`
		HitRate    *int `json:"hit_rate"`
		Oldest     *int `json:"oldest"`
		Hits       *int `json:"hits"`
		Misses     *int `json:"misses"`
		Entries    *int `json:"entries"`
		EntriesMax *int `json:"entries_max"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetCacheSslCacheStatistics() (*CacheSslCacheStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/5.2/status/local_tm/statistics/cache/ssl_cache")
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(CacheSslCacheStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
