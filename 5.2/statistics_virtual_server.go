// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 5.2.
package vtm

import (
	"encoding/json"
)

type VirtualServerStatistics struct {
	Statistics struct {
		AuthSamlResponsesHi         *int    `json:"auth_saml_responses_hi"`
		BwLimitBytesDrop            *int    `json:"bw_limit_bytes_drop"`
		CurrentConn                 *int    `json:"current_conn"`
		SslTicketResumed            *int    `json:"ssl_ticket_resumed"`
		SslNewSession               *int    `json:"ssl_new_session"`
		SslTicketReceived           *int    `json:"ssl_ticket_received"`
		GzipBytesSavedHi            *int    `json:"gzip_bytes_saved_hi"`
		SslCacheResumedHi           *int    `json:"ssl_cache_resumed_hi"`
		SslTicketReceivedLo         *int    `json:"ssl_ticket_received_lo"`
		SslTicketExpiredLo          *int    `json:"ssl_ticket_expired_lo"`
		SslCacheSavedHi             *int    `json:"ssl_cache_saved_hi"`
		AuthSessionsUsedHi          *int    `json:"auth_sessions_used_hi"`
		SslTicketRejectedHi         *int    `json:"ssl_ticket_rejected_hi"`
		AuthSessionsRejectedLo      *int    `json:"auth_sessions_rejected_lo"`
		SslCacheLookup              *int    `json:"ssl_cache_lookup"`
		TotalHttp1Requests          *int    `json:"total_http1_requests"`
		BytesOutLo                  *int    `json:"bytes_out_lo"`
		SslTicketResumedHi          *int    `json:"ssl_ticket_resumed_hi"`
		ProcessingTimedOut          *int    `json:"processing_timed_out"`
		SslCacheSaved               *int    `json:"ssl_cache_saved"`
		Port                        *int    `json:"port"`
		AuthSessionsCreatedLo       *int    `json:"auth_sessions_created_lo"`
		UdpTimedOut                 *int    `json:"udp_timed_out"`
		Gzip                        *int    `json:"gzip"`
		HttpRewriteCookie           *int    `json:"http_rewrite_cookie"`
		TotalRequestsLo             *int    `json:"total_requests_lo"`
		SslCacheRejected            *int    `json:"ssl_cache_rejected"`
		PktsOut                     *int    `json:"pkts_out"`
		BytesOutHi                  *int    `json:"bytes_out_hi"`
		KeepaliveTimedOut           *int    `json:"keepalive_timed_out"`
		SslCacheMissHi              *int    `json:"ssl_cache_miss_hi"`
		TotalHttp2RequestsHi        *int    `json:"total_http2_requests_hi"`
		AuthSamlResponsesAcceptedLo *int    `json:"auth_saml_responses_accepted_lo"`
		SslTicketReceivedHi         *int    `json:"ssl_ticket_received_hi"`
		AuthSamlRedirects           *int    `json:"auth_saml_redirects"`
		SslCacheMissLo              *int    `json:"ssl_cache_miss_lo"`
		ConnectionFailures          *int    `json:"connection_failures"`
		BytesInLo                   *int    `json:"bytes_in_lo"`
		SslCacheResumedLo           *int    `json:"ssl_cache_resumed_lo"`
		AuthSamlResponsesRejectedHi *int    `json:"auth_saml_responses_rejected_hi"`
		BytesInHi                   *int    `json:"bytes_in_hi"`
		SslCacheRejectedHi          *int    `json:"ssl_cache_rejected_hi"`
		AuthSessionsCreatedHi       *int    `json:"auth_sessions_created_hi"`
		TotalRequests               *int    `json:"total_requests"`
		AuthSessionsRejected        *int    `json:"auth_sessions_rejected"`
		BytesIn                     *int    `json:"bytes_in"`
		AuthSessionsCreated         *int    `json:"auth_sessions_created"`
		SslTicketKeyNotFound        *int    `json:"ssl_ticket_key_not_found"`
		SslTicketIssuedHi           *int    `json:"ssl_ticket_issued_hi"`
		SslTicketKeyNotFoundLo      *int    `json:"ssl_ticket_key_not_found_lo"`
		ConnectTimedOut             *int    `json:"connect_timed_out"`
		BwLimitPktsDrop             *int    `json:"bw_limit_pkts_drop"`
		AuthSessionsRejectedHi      *int    `json:"auth_sessions_rejected_hi"`
		SslCacheMiss                *int    `json:"ssl_cache_miss"`
		AuthSamlResponsesAcceptedHi *int    `json:"auth_saml_responses_accepted_hi"`
		SslCacheSavedLo             *int    `json:"ssl_cache_saved_lo"`
		GzipBytesSavedLo            *int    `json:"gzip_bytes_saved_lo"`
		MaxConn                     *int    `json:"max_conn"`
		HttpCacheLookups            *int    `json:"http_cache_lookups"`
		AuthSamlResponsesLo         *int    `json:"auth_saml_responses_lo"`
		MaxDurationTimedOut         *int    `json:"max_duration_timed_out"`
		TotalHttp2RequestsLo        *int    `json:"total_http2_requests_lo"`
		SipTotalCalls               *int    `json:"sip_total_calls"`
		SslTicketRejectedLo         *int    `json:"ssl_ticket_rejected_lo"`
		Protocol                    *string `json:"protocol"`
		PktsOutHi                   *int    `json:"pkts_out_hi"`
		TotalHttpRequestsLo         *int    `json:"total_http_requests_lo"`
		AuthSessionsUsedLo          *int    `json:"auth_sessions_used_lo"`
		BwLimitPktsDropHi           *int    `json:"bw_limit_pkts_drop_hi"`
		AuthSamlRedirectsLo         *int    `json:"auth_saml_redirects_lo"`
		PktsInLo                    *int    `json:"pkts_in_lo"`
		AuthSamlResponsesRejectedLo *int    `json:"auth_saml_responses_rejected_lo"`
		CertStatusRequests          *int    `json:"cert_status_requests"`
		BwLimitBytesDropLo          *int    `json:"bw_limit_bytes_drop_lo"`
		AuthSamlResponses           *int    `json:"auth_saml_responses"`
		TotalUdpUnreachables        *int    `json:"total_udp_unreachables"`
		SslTicketExpiredHi          *int    `json:"ssl_ticket_expired_hi"`
		TotalConn                   *int    `json:"total_conn"`
		PktsIn                      *int    `json:"pkts_in"`
		TotalHttpRequests           *int    `json:"total_http_requests"`
		CertStatusResponses         *int    `json:"cert_status_responses"`
		TotalHttp1RequestsHi        *int    `json:"total_http1_requests_hi"`
		PktsOutLo                   *int    `json:"pkts_out_lo"`
		HttpCacheHits               *int    `json:"http_cache_hits"`
		AuthSamlResponsesAccepted   *int    `json:"auth_saml_responses_accepted"`
		AuthSamlRedirectsHi         *int    `json:"auth_saml_redirects_hi"`
		SslCacheLookupLo            *int    `json:"ssl_cache_lookup_lo"`
		SslTicketIssued             *int    `json:"ssl_ticket_issued"`
		HttpRewriteLocation         *int    `json:"http_rewrite_location"`
		Discard                     *int    `json:"discard"`
		SslTicketExpired            *int    `json:"ssl_ticket_expired"`
		BwLimitPktsDropLo           *int    `json:"bw_limit_pkts_drop_lo"`
		SslNewSessionLo             *int    `json:"ssl_new_session_lo"`
		TotalHttpRequestsHi         *int    `json:"total_http_requests_hi"`
		SslCacheLookupHi            *int    `json:"ssl_cache_lookup_hi"`
		DirectReplies               *int    `json:"direct_replies"`
		HttpCacheHitRate            *int    `json:"http_cache_hit_rate"`
		SslTicketRejected           *int    `json:"ssl_ticket_rejected"`
		ConnectionErrors            *int    `json:"connection_errors"`
		TotalTcpReset               *int    `json:"total_tcp_reset"`
		SslTicketIssuedLo           *int    `json:"ssl_ticket_issued_lo"`
		SslCacheRejectedLo          *int    `json:"ssl_cache_rejected_lo"`
		AuthSamlResponsesRejected   *int    `json:"auth_saml_responses_rejected"`
		BytesOut                    *int    `json:"bytes_out"`
		SslTicketKeyNotFoundHi      *int    `json:"ssl_ticket_key_not_found_hi"`
		PktsInHi                    *int    `json:"pkts_in_hi"`
		SslNewSessionHi             *int    `json:"ssl_new_session_hi"`
		TotalDgram                  *int    `json:"total_dgram"`
		AuthSessionsUsed            *int    `json:"auth_sessions_used"`
		TotalHttp1RequestsLo        *int    `json:"total_http1_requests_lo"`
		TotalHttp2Requests          *int    `json:"total_http2_requests"`
		SipRejectedRequests         *int    `json:"sip_rejected_requests"`
		BwLimitBytesDropHi          *int    `json:"bw_limit_bytes_drop_hi"`
		SslCacheResumed             *int    `json:"ssl_cache_resumed"`
		TotalRequestsHi             *int    `json:"total_requests_hi"`
		DataTimedOut                *int    `json:"data_timed_out"`
		GzipBytesSaved              *int    `json:"gzip_bytes_saved"`
		SslTicketResumedLo          *int    `json:"ssl_ticket_resumed_lo"`
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
