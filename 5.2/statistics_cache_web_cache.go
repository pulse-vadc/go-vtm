// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 5.2.
package vtm

import (
	"encoding/json"
)

type CacheWebCacheStatistics struct {
	Statistics struct {
		MaxEntries               *int `json:"max_entries"`
		Misses                   *int `json:"misses"`
		UrlStoreSize             *int `json:"url_store_size"`
		Hits                     *int `json:"hits"`
		HitsLo                   *int `json:"hits_lo"`
		UrlStoreTotalFailures    *int `json:"url_store_total_failures"`
		Lookups                  *int `json:"lookups"`
		MemMaximum               *int `json:"mem_maximum"`
		Oldest                   *int `json:"oldest"`
		HitsHi                   *int `json:"hits_hi"`
		MissesLo                 *int `json:"misses_lo"`
		LookupsHi                *int `json:"lookups_hi"`
		MissesHi                 *int `json:"misses_hi"`
		HitRate                  *int `json:"hit_rate"`
		UrlStoreAllocated        *int `json:"url_store_allocated"`
		Entries                  *int `json:"entries"`
		UrlStoreTotalFrees       *int `json:"url_store_total_frees"`
		MemUsed                  *int `json:"mem_used"`
		LookupsLo                *int `json:"lookups_lo"`
		UrlStoreTotalAllocations *int `json:"url_store_total_allocations"`
		UrlStoreFree             *int `json:"url_store_free"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetCacheWebCacheStatistics() (*CacheWebCacheStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/5.2/status/local_tm/statistics/cache/web_cache")
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(CacheWebCacheStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
