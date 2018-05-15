// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 4.0.
package vtm

import (
	"encoding/json"
)

type CacheWebCacheStatistics struct {
	Statistics struct {
		HitsHi                   *int `json:"hits_hi"`
		LookupsHi                *int `json:"lookups_hi"`
		LookupsLo                *int `json:"lookups_lo"`
		MemMaximum               *int `json:"mem_maximum"`
		HitRate                  *int `json:"hit_rate"`
		Misses                   *int `json:"misses"`
		MissesHi                 *int `json:"misses_hi"`
		UrlStoreTotalFailures    *int `json:"url_store_total_failures"`
		Entries                  *int `json:"entries"`
		HitsLo                   *int `json:"hits_lo"`
		Lookups                  *int `json:"lookups"`
		UrlStoreTotalAllocations *int `json:"url_store_total_allocations"`
		UrlStoreSize             *int `json:"url_store_size"`
		MissesLo                 *int `json:"misses_lo"`
		Oldest                   *int `json:"oldest"`
		UrlStoreFree             *int `json:"url_store_free"`
		UrlStoreAllocated        *int `json:"url_store_allocated"`
		MaxEntries               *int `json:"max_entries"`
		UrlStoreTotalFrees       *int `json:"url_store_total_frees"`
		Hits                     *int `json:"hits"`
		MemUsed                  *int `json:"mem_used"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetCacheWebCacheStatistics() (*CacheWebCacheStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/4.0/status/local_tm/statistics/cache/web_cache")
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
