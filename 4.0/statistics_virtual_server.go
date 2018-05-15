// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 4.0.
package vtm

import (
	"encoding/json"
)

type VirtualServerStatistics struct {
	Statistics struct {
		BwLimitBytesDropLo   *int    `json:"bw_limit_bytes_drop_lo"`
		GzipBytesSavedHi     *int    `json:"gzip_bytes_saved_hi"`
		BwLimitPktsDropHi    *int    `json:"bw_limit_pkts_drop_hi"`
		GzipBytesSaved       *int    `json:"gzip_bytes_saved"`
		HttpRewriteCookie    *int    `json:"http_rewrite_cookie"`
		CurrentConn          *int    `json:"current_conn"`
		TotalHttpRequestsLo  *int    `json:"total_http_requests_lo"`
		TotalDgram           *int    `json:"total_dgram"`
		TotalRequestsHi      *int    `json:"total_requests_hi"`
		SipTotalCalls        *int    `json:"sip_total_calls"`
		BytesInHi            *int    `json:"bytes_in_hi"`
		ConnectionFailures   *int    `json:"connection_failures"`
		BytesOutHi           *int    `json:"bytes_out_hi"`
		TotalHttpRequestsHi  *int    `json:"total_http_requests_hi"`
		ConnectTimedOut      *int    `json:"connect_timed_out"`
		HttpRewriteLocation  *int    `json:"http_rewrite_location"`
		MaxConn              *int    `json:"max_conn"`
		MaxDurationTimedOut  *int    `json:"max_duration_timed_out"`
		TotalHttp2RequestsHi *int    `json:"total_http2_requests_hi"`
		Port                 *int    `json:"port"`
		PktsInHi             *int    `json:"pkts_in_hi"`
		TotalHttp2Requests   *int    `json:"total_http2_requests"`
		TotalHttp2RequestsLo *int    `json:"total_http2_requests_lo"`
		GzipBytesSavedLo     *int    `json:"gzip_bytes_saved_lo"`
		CertStatusResponses  *int    `json:"cert_status_responses"`
		BwLimitBytesDrop     *int    `json:"bw_limit_bytes_drop"`
		KeepaliveTimedOut    *int    `json:"keepalive_timed_out"`
		TotalUdpUnreachables *int    `json:"total_udp_unreachables"`
		DirectReplies        *int    `json:"direct_replies"`
		HttpCacheHitRate     *int    `json:"http_cache_hit_rate"`
		PktsOutHi            *int    `json:"pkts_out_hi"`
		PktsOut              *int    `json:"pkts_out"`
		TotalTcpReset        *int    `json:"total_tcp_reset"`
		BwLimitPktsDrop      *int    `json:"bw_limit_pkts_drop"`
		UdpTimedOut          *int    `json:"udp_timed_out"`
		TotalConn            *int    `json:"total_conn"`
		Gzip                 *int    `json:"gzip"`
		TotalHttp1RequestsLo *int    `json:"total_http1_requests_lo"`
		BytesOut             *int    `json:"bytes_out"`
		TotalRequests        *int    `json:"total_requests"`
		BytesIn              *int    `json:"bytes_in"`
		TotalHttpRequests    *int    `json:"total_http_requests"`
		TotalHttp1RequestsHi *int    `json:"total_http1_requests_hi"`
		HttpCacheLookups     *int    `json:"http_cache_lookups"`
		TotalRequestsLo      *int    `json:"total_requests_lo"`
		HttpCacheHits        *int    `json:"http_cache_hits"`
		SipRejectedRequests  *int    `json:"sip_rejected_requests"`
		BwLimitPktsDropLo    *int    `json:"bw_limit_pkts_drop_lo"`
		DataTimedOut         *int    `json:"data_timed_out"`
		BytesOutLo           *int    `json:"bytes_out_lo"`
		TotalHttp1Requests   *int    `json:"total_http1_requests"`
		Protocol             *string `json:"protocol"`
		PktsIn               *int    `json:"pkts_in"`
		ConnectionErrors     *int    `json:"connection_errors"`
		CertStatusRequests   *int    `json:"cert_status_requests"`
		PktsInLo             *int    `json:"pkts_in_lo"`
		ProcessingTimedOut   *int    `json:"processing_timed_out"`
		PktsOutLo            *int    `json:"pkts_out_lo"`
		BwLimitBytesDropHi   *int    `json:"bw_limit_bytes_drop_hi"`
		Discard              *int    `json:"discard"`
		BytesInLo            *int    `json:"bytes_in_lo"`
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
