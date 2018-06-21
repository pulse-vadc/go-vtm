// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 4.0.
package vtm

import (
	"encoding/json"
)

type CacheWebCacheStatistics struct {
	Statistics struct {
		HitsLo                   *int `json:"hits_lo"`
		LookupsLo                *int `json:"lookups_lo"`
		UrlStoreSize             *int `json:"url_store_size"`
		Entries                  *int `json:"entries"`
		LookupsHi                *int `json:"lookups_hi"`
		UrlStoreAllocated        *int `json:"url_store_allocated"`
		MissesHi                 *int `json:"misses_hi"`
		UrlStoreTotalFailures    *int `json:"url_store_total_failures"`
		HitsHi                   *int `json:"hits_hi"`
		Lookups                  *int `json:"lookups"`
		MaxEntries               *int `json:"max_entries"`
		MemUsed                  *int `json:"mem_used"`
		MissesLo                 *int `json:"misses_lo"`
		UrlStoreFree             *int `json:"url_store_free"`
		UrlStoreTotalAllocations *int `json:"url_store_total_allocations"`
		Misses                   *int `json:"misses"`
		UrlStoreTotalFrees       *int `json:"url_store_total_frees"`
		Oldest                   *int `json:"oldest"`
		HitRate                  *int `json:"hit_rate"`
		Hits                     *int `json:"hits"`
		MemMaximum               *int `json:"mem_maximum"`
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
