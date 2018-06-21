// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 4.0.
package vtm

import (
	"encoding/json"
)

type VirtualServerStatistics struct {
	Statistics struct {
		TotalHttp2Requests   *int    `json:"total_http2_requests"`
		ConnectionFailures   *int    `json:"connection_failures"`
		HttpRewriteLocation  *int    `json:"http_rewrite_location"`
		BwLimitBytesDropHi   *int    `json:"bw_limit_bytes_drop_hi"`
		GzipBytesSaved       *int    `json:"gzip_bytes_saved"`
		HttpCacheHitRate     *int    `json:"http_cache_hit_rate"`
		TotalHttp2RequestsHi *int    `json:"total_http2_requests_hi"`
		BwLimitPktsDropLo    *int    `json:"bw_limit_pkts_drop_lo"`
		BytesOutLo           *int    `json:"bytes_out_lo"`
		BytesInHi            *int    `json:"bytes_in_hi"`
		MaxDurationTimedOut  *int    `json:"max_duration_timed_out"`
		HttpCacheLookups     *int    `json:"http_cache_lookups"`
		CertStatusResponses  *int    `json:"cert_status_responses"`
		BytesInLo            *int    `json:"bytes_in_lo"`
		PktsIn               *int    `json:"pkts_in"`
		TotalUdpUnreachables *int    `json:"total_udp_unreachables"`
		KeepaliveTimedOut    *int    `json:"keepalive_timed_out"`
		TotalHttpRequestsLo  *int    `json:"total_http_requests_lo"`
		MaxConn              *int    `json:"max_conn"`
		BytesOut             *int    `json:"bytes_out"`
		BwLimitBytesDrop     *int    `json:"bw_limit_bytes_drop"`
		TotalHttp1Requests   *int    `json:"total_http1_requests"`
		TotalHttp2RequestsLo *int    `json:"total_http2_requests_lo"`
		UdpTimedOut          *int    `json:"udp_timed_out"`
		DataTimedOut         *int    `json:"data_timed_out"`
		BwLimitPktsDrop      *int    `json:"bw_limit_pkts_drop"`
		BwLimitBytesDropLo   *int    `json:"bw_limit_bytes_drop_lo"`
		PktsInHi             *int    `json:"pkts_in_hi"`
		PktsInLo             *int    `json:"pkts_in_lo"`
		Gzip                 *int    `json:"gzip"`
		TotalConn            *int    `json:"total_conn"`
		Port                 *int    `json:"port"`
		ConnectionErrors     *int    `json:"connection_errors"`
		PktsOutLo            *int    `json:"pkts_out_lo"`
		TotalTcpReset        *int    `json:"total_tcp_reset"`
		PktsOut              *int    `json:"pkts_out"`
		PktsOutHi            *int    `json:"pkts_out_hi"`
		CertStatusRequests   *int    `json:"cert_status_requests"`
		GzipBytesSavedHi     *int    `json:"gzip_bytes_saved_hi"`
		TotalRequests        *int    `json:"total_requests"`
		HttpCacheHits        *int    `json:"http_cache_hits"`
		SipRejectedRequests  *int    `json:"sip_rejected_requests"`
		Discard              *int    `json:"discard"`
		Protocol             *string `json:"protocol"`
		TotalHttp1RequestsHi *int    `json:"total_http1_requests_hi"`
		TotalRequestsHi      *int    `json:"total_requests_hi"`
		TotalRequestsLo      *int    `json:"total_requests_lo"`
		BwLimitPktsDropHi    *int    `json:"bw_limit_pkts_drop_hi"`
		BytesOutHi           *int    `json:"bytes_out_hi"`
		GzipBytesSavedLo     *int    `json:"gzip_bytes_saved_lo"`
		HttpRewriteCookie    *int    `json:"http_rewrite_cookie"`
		ProcessingTimedOut   *int    `json:"processing_timed_out"`
		CurrentConn          *int    `json:"current_conn"`
		TotalHttpRequests    *int    `json:"total_http_requests"`
		TotalHttpRequestsHi  *int    `json:"total_http_requests_hi"`
		BytesIn              *int    `json:"bytes_in"`
		DirectReplies        *int    `json:"direct_replies"`
		SipTotalCalls        *int    `json:"sip_total_calls"`
		TotalDgram           *int    `json:"total_dgram"`
		ConnectTimedOut      *int    `json:"connect_timed_out"`
		TotalHttp1RequestsLo *int    `json:"total_http1_requests_lo"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetVirtualServerStatistics(name string) (*VirtualServerStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/4.0/status/local_tm/statistics/virtual_servers/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(VirtualServerStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
