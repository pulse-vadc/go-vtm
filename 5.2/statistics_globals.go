// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 5.2.
package vtm

import (
	"encoding/json"
)

type GlobalsStatistics struct {
	Statistics struct {
		SysMemBuffered                    *int `json:"sys_mem_buffered"`
		SslClientCertInvalid              *int `json:"ssl_client_cert_invalid"`
		SysMemSwapTotal                   *int `json:"sys_mem_swap_total"`
		SysCpuSystemBusyPercent           *int `json:"sys_cpu_system_busy_percent"`
		SslCipherDesDecrypts              *int `json:"ssl_cipher_des_decrypts"`
		SslHandshakeSslv2                 *int `json:"ssl_handshake_sslv2"`
		NumberChildProcesses              *int `json:"number_child_processes"`
		TotalBytesOutLo                   *int `json:"total_bytes_out_lo"`
		TotalBadDnsPackets                *int `json:"total_bad_dns_packets"`
		SslCipherEncrypts                 *int `json:"ssl_cipher_encrypts"`
		AnalyticsTransactionsMemoryUsage  *int `json:"analytics_transactions_memory_usage"`
		SslHandshakeTlsv12                *int `json:"ssl_handshake_tlsv12"`
		SslClientCertRevoked              *int `json:"ssl_client_cert_revoked"`
		TotalBackendServerErrors          *int `json:"total_backend_server_errors"`
		TotalBytesInLo                    *int `json:"total_bytes_in_lo"`
		NumberSnmpGetNextRequests         *int `json:"number_snmp_get_next_requests"`
		TotalRequests                     *int `json:"total_requests"`
		NumberSnmpBadRequests             *int `json:"number_snmp_bad_requests"`
		TotalCurrentConn                  *int `json:"total_current_conn"`
		SslCipherDecrypts                 *int `json:"ssl_cipher_decrypts"`
		TimeLastConfigUpdate              *int `json:"time_last_config_update"`
		SslSessionIdDiskCacheHit          *int `json:"ssl_session_id_disk_cache_hit"`
		SslHandshakeTlsv1                 *int `json:"ssl_handshake_tlsv1"`
		NumberDnsaRequests                *int `json:"number_dnsa_requests"`
		HourlyPeakRequestsPerSecond       *int `json:"hourly_peak_requests_per_second"`
		SslCipherEcdsaVerifies            *int `json:"ssl_cipher_ecdsa_verifies"`
		DataEntries                       *int `json:"data_entries"`
		AnalyticsTransactionsExported     *int `json:"analytics_transactions_exported"`
		NumberSnmpGetBulkRequests         *int `json:"number_snmp_get_bulk_requests"`
		NumberSnmpUnauthorisedRequests    *int `json:"number_snmp_unauthorised_requests"`
		SslCipherRc4Decrypts              *int `json:"ssl_cipher_rc4_decrypts"`
		UpTime                            *int `json:"up_time"`
		SysMemInUse                       *int `json:"sys_mem_in_use"`
		SslConnections                    *int `json:"ssl_connections"`
		SslCipherDsaVerifies              *int `json:"ssl_cipher_dsa_verifies"`
		TotalConn                         *int `json:"total_conn"`
		SysCpuIdlePercent                 *int `json:"sys_cpu_idle_percent"`
		SslCipherAesDecrypts              *int `json:"ssl_cipher_aes_decrypts"`
		HourlyPeakSslConnectionsPerSecond *int `json:"hourly_peak_ssl_connections_per_second"`
		SslCipherEcdsaSigns               *int `json:"ssl_cipher_ecdsa_signs"`
		SysFdsFree                        *int `json:"sys_fds_free"`
		SslCipherAesGcmEncrypts           *int `json:"ssl_cipher_aes_gcm_encrypts"`
		SslSessionIdMemCacheMiss          *int `json:"ssl_session_id_mem_cache_miss"`
		SysMemTotal                       *int `json:"sys_mem_total"`
		DataMemoryUsage                   *int `json:"data_memory_usage"`
		EventsSeen                        *int `json:"events_seen"`
		SslClientCertExpired              *int `json:"ssl_client_cert_expired"`
		SysCpuUserBusyPercent             *int `json:"sys_cpu_user_busy_percent"`
		TotalBytesInHi                    *int `json:"total_bytes_in_hi"`
		SslCipherRsaEncrypts              *int `json:"ssl_cipher_rsa_encrypts"`
		HourlyPeakBytesInPerSecond        *int `json:"hourly_peak_bytes_in_per_second"`
		SslCipherRsaEncryptsExternal      *int `json:"ssl_cipher_rsa_encrypts_external"`
		SslCipherEcdhGenerates            *int `json:"ssl_cipher_ecdh_generates"`
		SslHandshakeTlsv11                *int `json:"ssl_handshake_tlsv11"`
		SslHandshakeSslv3                 *int `json:"ssl_handshake_sslv3"`
		SslCipher3DesDecrypts             *int `json:"ssl_cipher_3des_decrypts"`
		SslSessionIdDiskCacheMiss         *int `json:"ssl_session_id_disk_cache_miss"`
		SslCipherRc4Encrypts              *int `json:"ssl_cipher_rc4_encrypts"`
		SslClientCertNotSent              *int `json:"ssl_client_cert_not_sent"`
		SysCpuBusyPercent                 *int `json:"sys_cpu_busy_percent"`
		SslSessionIdMemCacheHit           *int `json:"ssl_session_id_mem_cache_hit"`
		SysMemFree                        *int `json:"sys_mem_free"`
		TotalBytesOutHi                   *int `json:"total_bytes_out_hi"`
		SslCipherDhAgreements             *int `json:"ssl_cipher_dh_agreements"`
		NumberSnmpGetRequests             *int `json:"number_snmp_get_requests"`
		SslCipherDesEncrypts              *int `json:"ssl_cipher_des_encrypts"`
		TotalBytesIn                      *int `json:"total_bytes_in"`
		SslCipherRsaDecryptsExternal      *int `json:"ssl_cipher_rsa_decrypts_external"`
		TotalTransactions                 *int `json:"total_transactions"`
		TotalDnsResponses                 *int `json:"total_dns_responses"`
		NumberDnsaCacheHits               *int `json:"number_dnsa_cache_hits"`
		SslCipherAesGcmDecrypts           *int `json:"ssl_cipher_aes_gcm_decrypts"`
		SslCipherRsaDecrypts              *int `json:"ssl_cipher_rsa_decrypts"`
		SslCipherDhGenerates              *int `json:"ssl_cipher_dh_generates"`
		SslCipherEcdhAgreements           *int `json:"ssl_cipher_ecdh_agreements"`
		SysMemSwapped                     *int `json:"sys_mem_swapped"`
		NumIdleConnections                *int `json:"num_idle_connections"`
		NumberDnsptrCacheHits             *int `json:"number_dnsptr_cache_hits"`
		SslCipherAesEncrypts              *int `json:"ssl_cipher_aes_encrypts"`
		NumberDnsptrRequests              *int `json:"number_dnsptr_requests"`
		SslCipherDsaSigns                 *int `json:"ssl_cipher_dsa_signs"`
		AnalyticsTransactionsDropped      *int `json:"analytics_transactions_dropped"`
		HourlyPeakBytesOutPerSecond       *int `json:"hourly_peak_bytes_out_per_second"`
		TotalBytesOut                     *int `json:"total_bytes_out"`
		SslCipher3DesEncrypts             *int `json:"ssl_cipher_3des_encrypts"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetGlobalsStatistics() (*GlobalsStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/5.2/status/local_tm/statistics/globals")
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(GlobalsStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
