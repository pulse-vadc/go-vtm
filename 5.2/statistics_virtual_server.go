// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 5.2.
package vtm

import (
	"encoding/json"
)

type VirtualServerStatistics struct {
	Statistics struct {
		AuthSessionsCreatedHi       *int    `json:"auth_sessions_created_hi"`
		SipTotalCalls               *int    `json:"sip_total_calls"`
		UdpTimedOut                 *int    `json:"udp_timed_out"`
		HttpRewriteLocation         *int    `json:"http_rewrite_location"`
		SslCacheMissLo              *int    `json:"ssl_cache_miss_lo"`
		TotalHttp2Requests          *int    `json:"total_http2_requests"`
		ConnectionFailures          *int    `json:"connection_failures"`
		SslTicketReceivedHi         *int    `json:"ssl_ticket_received_hi"`
		SslTicketExpiredHi          *int    `json:"ssl_ticket_expired_hi"`
		SslNewSessionHi             *int    `json:"ssl_new_session_hi"`
		SslTicketResumedLo          *int    `json:"ssl_ticket_resumed_lo"`
		TotalHttp2RequestsLo        *int    `json:"total_http2_requests_lo"`
		SslCacheResumed             *int    `json:"ssl_cache_resumed"`
		SslCacheRejected            *int    `json:"ssl_cache_rejected"`
		ProcessingTimedOut          *int    `json:"processing_timed_out"`
		BytesInLo                   *int    `json:"bytes_in_lo"`
		PktsInLo                    *int    `json:"pkts_in_lo"`
		Gzip                        *int    `json:"gzip"`
		SslCacheLookupHi            *int    `json:"ssl_cache_lookup_hi"`
		SslCacheResumedLo           *int    `json:"ssl_cache_resumed_lo"`
		AuthSessionsRejectedLo      *int    `json:"auth_sessions_rejected_lo"`
		DirectReplies               *int    `json:"direct_replies"`
		GzipBytesSavedLo            *int    `json:"gzip_bytes_saved_lo"`
		PktsOut                     *int    `json:"pkts_out"`
		Protocol                    *string `json:"protocol"`
		CurrentConn                 *int    `json:"current_conn"`
		TotalHttp1RequestsHi        *int    `json:"total_http1_requests_hi"`
		SslNewSessionLo             *int    `json:"ssl_new_session_lo"`
		HttpCacheHits               *int    `json:"http_cache_hits"`
		TotalHttpRequestsHi         *int    `json:"total_http_requests_hi"`
		SslTicketKeyNotFoundLo      *int    `json:"ssl_ticket_key_not_found_lo"`
		AuthSamlRedirectsHi         *int    `json:"auth_saml_redirects_hi"`
		SslTicketRejectedHi         *int    `json:"ssl_ticket_rejected_hi"`
		AuthSamlResponsesAccepted   *int    `json:"auth_saml_responses_accepted"`
		ConnectTimedOut             *int    `json:"connect_timed_out"`
		AuthSamlRedirects           *int    `json:"auth_saml_redirects"`
		AuthSamlResponses           *int    `json:"auth_saml_responses"`
		TotalHttp2RequestsHi        *int    `json:"total_http2_requests_hi"`
		CertStatusResponses         *int    `json:"cert_status_responses"`
		KeepaliveTimedOut           *int    `json:"keepalive_timed_out"`
		AuthSamlResponsesAcceptedHi *int    `json:"auth_saml_responses_accepted_hi"`
		TotalHttp1Requests          *int    `json:"total_http1_requests"`
		SslTicketExpiredLo          *int    `json:"ssl_ticket_expired_lo"`
		AuthSessionsCreated         *int    `json:"auth_sessions_created"`
		SslTicketRejectedLo         *int    `json:"ssl_ticket_rejected_lo"`
		AuthSessionsRejectedHi      *int    `json:"auth_sessions_rejected_hi"`
		Discard                     *int    `json:"discard"`
		SslTicketResumed            *int    `json:"ssl_ticket_resumed"`
		SslTicketKeyNotFound        *int    `json:"ssl_ticket_key_not_found"`
		SslCacheSavedLo             *int    `json:"ssl_cache_saved_lo"`
		AuthSamlResponsesHi         *int    `json:"auth_saml_responses_hi"`
		SslTicketIssued             *int    `json:"ssl_ticket_issued"`
		TotalHttpRequestsLo         *int    `json:"total_http_requests_lo"`
		BytesOutHi                  *int    `json:"bytes_out_hi"`
		ConnectionErrors            *int    `json:"connection_errors"`
		TotalRequestsLo             *int    `json:"total_requests_lo"`
		AuthSessionsRejected        *int    `json:"auth_sessions_rejected"`
		SslNewSession               *int    `json:"ssl_new_session"`
		BwLimitPktsDrop             *int    `json:"bw_limit_pkts_drop"`
		TotalHttp1RequestsLo        *int    `json:"total_http1_requests_lo"`
		PktsOutLo                   *int    `json:"pkts_out_lo"`
		SslTicketKeyNotFoundHi      *int    `json:"ssl_ticket_key_not_found_hi"`
		TotalHttpRequests           *int    `json:"total_http_requests"`
		BwLimitBytesDropLo          *int    `json:"bw_limit_bytes_drop_lo"`
		MaxConn                     *int    `json:"max_conn"`
		SslCacheLookupLo            *int    `json:"ssl_cache_lookup_lo"`
		SslTicketRejected           *int    `json:"ssl_ticket_rejected"`
		SslCacheSavedHi             *int    `json:"ssl_cache_saved_hi"`
		TotalTcpReset               *int    `json:"total_tcp_reset"`
		GzipBytesSavedHi            *int    `json:"gzip_bytes_saved_hi"`
		GzipBytesSaved              *int    `json:"gzip_bytes_saved"`
		SslTicketIssuedHi           *int    `json:"ssl_ticket_issued_hi"`
		BytesOut                    *int    `json:"bytes_out"`
		BytesInHi                   *int    `json:"bytes_in_hi"`
		HttpCacheLookups            *int    `json:"http_cache_lookups"`
		TotalDgram                  *int    `json:"total_dgram"`
		AuthSamlResponsesRejected   *int    `json:"auth_saml_responses_rejected"`
		TotalConn                   *int    `json:"total_conn"`
		BwLimitBytesDrop            *int    `json:"bw_limit_bytes_drop"`
		BwLimitBytesDropHi          *int    `json:"bw_limit_bytes_drop_hi"`
		PktsOutHi                   *int    `json:"pkts_out_hi"`
		BwLimitPktsDropLo           *int    `json:"bw_limit_pkts_drop_lo"`
		SslCacheRejectedLo          *int    `json:"ssl_cache_rejected_lo"`
		TotalRequestsHi             *int    `json:"total_requests_hi"`
		SslCacheResumedHi           *int    `json:"ssl_cache_resumed_hi"`
		SslTicketReceived           *int    `json:"ssl_ticket_received"`
		PktsInHi                    *int    `json:"pkts_in_hi"`
		SslTicketResumedHi          *int    `json:"ssl_ticket_resumed_hi"`
		AuthSamlRedirectsLo         *int    `json:"auth_saml_redirects_lo"`
		AuthSessionsCreatedLo       *int    `json:"auth_sessions_created_lo"`
		TotalRequests               *int    `json:"total_requests"`
		Port                        *int    `json:"port"`
		SslTicketIssuedLo           *int    `json:"ssl_ticket_issued_lo"`
		TotalUdpUnreachables        *int    `json:"total_udp_unreachables"`
		AuthSessionsUsed            *int    `json:"auth_sessions_used"`
		SslCacheSaved               *int    `json:"ssl_cache_saved"`
		SipRejectedRequests         *int    `json:"sip_rejected_requests"`
		HttpCacheHitRate            *int    `json:"http_cache_hit_rate"`
		MaxDurationTimedOut         *int    `json:"max_duration_timed_out"`
		AuthSamlResponsesRejectedHi *int    `json:"auth_saml_responses_rejected_hi"`
		SslTicketExpired            *int    `json:"ssl_ticket_expired"`
		HttpRewriteCookie           *int    `json:"http_rewrite_cookie"`
		SslCacheLookup              *int    `json:"ssl_cache_lookup"`
		CertStatusRequests          *int    `json:"cert_status_requests"`
		SslCacheRejectedHi          *int    `json:"ssl_cache_rejected_hi"`
		AuthSamlResponsesAcceptedLo *int    `json:"auth_saml_responses_accepted_lo"`
		SslCacheMiss                *int    `json:"ssl_cache_miss"`
		AuthSessionsUsedLo          *int    `json:"auth_sessions_used_lo"`
		BytesOutLo                  *int    `json:"bytes_out_lo"`
		DataTimedOut                *int    `json:"data_timed_out"`
		AuthSessionsUsedHi          *int    `json:"auth_sessions_used_hi"`
		BytesIn                     *int    `json:"bytes_in"`
		PktsIn                      *int    `json:"pkts_in"`
		BwLimitPktsDropHi           *int    `json:"bw_limit_pkts_drop_hi"`
		AuthSamlResponsesLo         *int    `json:"auth_saml_responses_lo"`
		AuthSamlResponsesRejectedLo *int    `json:"auth_saml_responses_rejected_lo"`
		SslCacheMissHi              *int    `json:"ssl_cache_miss_hi"`
		SslTicketReceivedLo         *int    `json:"ssl_ticket_received_lo"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetVirtualServerStatistics(name string) (*VirtualServerStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/5.2/status/local_tm/statistics/virtual_servers/" + name)
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
