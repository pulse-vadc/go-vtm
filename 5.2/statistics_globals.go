// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 5.2.
package vtm

import (
	"encoding/json"
)

type GlobalsStatistics struct {
	Statistics struct {
		NumberDnsptrCacheHits             *int `json:"number_dnsptr_cache_hits"`
		SslHandshakeSslv2                 *int `json:"ssl_handshake_sslv2"`
		SslCipherAesGcmDecrypts           *int `json:"ssl_cipher_aes_gcm_decrypts"`
		HourlyPeakRequestsPerSecond       *int `json:"hourly_peak_requests_per_second"`
		NumberSnmpBadRequests             *int `json:"number_snmp_bad_requests"`
		SysCpuUserBusyPercent             *int `json:"sys_cpu_user_busy_percent"`
		TotalBytesIn                      *int `json:"total_bytes_in"`
		SslClientCertNotSent              *int `json:"ssl_client_cert_not_sent"`
		SslCipher3DesEncrypts             *int `json:"ssl_cipher_3des_encrypts"`
		SslCipherAesGcmEncrypts           *int `json:"ssl_cipher_aes_gcm_encrypts"`
		SslCipherDhAgreements             *int `json:"ssl_cipher_dh_agreements"`
		SslCipherRc4Encrypts              *int `json:"ssl_cipher_rc4_encrypts"`
		SslCipher3DesDecrypts             *int `json:"ssl_cipher_3des_decrypts"`
		SslCipherDesDecrypts              *int `json:"ssl_cipher_des_decrypts"`
		SslCipherRc4Decrypts              *int `json:"ssl_cipher_rc4_decrypts"`
		SysMemTotal                       *int `json:"sys_mem_total"`
		SslClientCertInvalid              *int `json:"ssl_client_cert_invalid"`
		SslSessionIdDiskCacheMiss         *int `json:"ssl_session_id_disk_cache_miss"`
		SslSessionIdMemCacheMiss          *int `json:"ssl_session_id_mem_cache_miss"`
		SslCipherEcdhGenerates            *int `json:"ssl_cipher_ecdh_generates"`
		SysCpuSystemBusyPercent           *int `json:"sys_cpu_system_busy_percent"`
		TotalBytesInLo                    *int `json:"total_bytes_in_lo"`
		SslCipherRsaEncryptsExternal      *int `json:"ssl_cipher_rsa_encrypts_external"`
		SslConnections                    *int `json:"ssl_connections"`
		NumberSnmpUnauthorisedRequests    *int `json:"number_snmp_unauthorised_requests"`
		AnalyticsTransactionsExported     *int `json:"analytics_transactions_exported"`
		SslClientCertExpired              *int `json:"ssl_client_cert_expired"`
		UpTime                            *int `json:"up_time"`
		SslCipherDesEncrypts              *int `json:"ssl_cipher_des_encrypts"`
		TotalBytesInHi                    *int `json:"total_bytes_in_hi"`
		TotalCurrentConn                  *int `json:"total_current_conn"`
		SslCipherDhGenerates              *int `json:"ssl_cipher_dh_generates"`
		TotalConn                         *int `json:"total_conn"`
		SslClientCertRevoked              *int `json:"ssl_client_cert_revoked"`
		SslCipherRsaDecrypts              *int `json:"ssl_cipher_rsa_decrypts"`
		HourlyPeakSslConnectionsPerSecond *int `json:"hourly_peak_ssl_connections_per_second"`
		NumberDnsaCacheHits               *int `json:"number_dnsa_cache_hits"`
		SslCipherEcdsaVerifies            *int `json:"ssl_cipher_ecdsa_verifies"`
		NumberSnmpGetRequests             *int `json:"number_snmp_get_requests"`
		SysMemBuffered                    *int `json:"sys_mem_buffered"`
		TotalBadDnsPackets                *int `json:"total_bad_dns_packets"`
		SysMemSwapTotal                   *int `json:"sys_mem_swap_total"`
		SslCipherEcdhAgreements           *int `json:"ssl_cipher_ecdh_agreements"`
		SslSessionIdDiskCacheHit          *int `json:"ssl_session_id_disk_cache_hit"`
		NumberChildProcesses              *int `json:"number_child_processes"`
		SslCipherDecrypts                 *int `json:"ssl_cipher_decrypts"`
		SslCipherDsaSigns                 *int `json:"ssl_cipher_dsa_signs"`
		SslCipherEncrypts                 *int `json:"ssl_cipher_encrypts"`
		SysCpuBusyPercent                 *int `json:"sys_cpu_busy_percent"`
		SslHandshakeTlsv1                 *int `json:"ssl_handshake_tlsv1"`
		HourlyPeakBytesInPerSecond        *int `json:"hourly_peak_bytes_in_per_second"`
		TotalBytesOut                     *int `json:"total_bytes_out"`
		SysMemInUse                       *int `json:"sys_mem_in_use"`
		NumberSnmpGetNextRequests         *int `json:"number_snmp_get_next_requests"`
		SslCipherRsaDecryptsExternal      *int `json:"ssl_cipher_rsa_decrypts_external"`
		HourlyPeakBytesOutPerSecond       *int `json:"hourly_peak_bytes_out_per_second"`
		SslCipherDsaVerifies              *int `json:"ssl_cipher_dsa_verifies"`
		TotalRequests                     *int `json:"total_requests"`
		SslCipherEcdsaSigns               *int `json:"ssl_cipher_ecdsa_signs"`
		TotalBackendServerErrors          *int `json:"total_backend_server_errors"`
		NumberDnsptrRequests              *int `json:"number_dnsptr_requests"`
		SslHandshakeTlsv11                *int `json:"ssl_handshake_tlsv11"`
		SysFdsFree                        *int `json:"sys_fds_free"`
		EventsSeen                        *int `json:"events_seen"`
		SslSessionIdMemCacheHit           *int `json:"ssl_session_id_mem_cache_hit"`
		SslCipherRsaEncrypts              *int `json:"ssl_cipher_rsa_encrypts"`
		DataEntries                       *int `json:"data_entries"`
		SysCpuIdlePercent                 *int `json:"sys_cpu_idle_percent"`
		TotalTransactions                 *int `json:"total_transactions"`
		SysMemFree                        *int `json:"sys_mem_free"`
		SysMemSwapped                     *int `json:"sys_mem_swapped"`
		SslCipherAesDecrypts              *int `json:"ssl_cipher_aes_decrypts"`
		TotalDnsResponses                 *int `json:"total_dns_responses"`
		TimeLastConfigUpdate              *int `json:"time_last_config_update"`
		NumberSnmpGetBulkRequests         *int `json:"number_snmp_get_bulk_requests"`
		DataMemoryUsage                   *int `json:"data_memory_usage"`
		SslCipherAesEncrypts              *int `json:"ssl_cipher_aes_encrypts"`
		NumIdleConnections                *int `json:"num_idle_connections"`
		SslHandshakeSslv3                 *int `json:"ssl_handshake_sslv3"`
		NumberDnsaRequests                *int `json:"number_dnsa_requests"`
		AnalyticsTransactionsDropped      *int `json:"analytics_transactions_dropped"`
		TotalBytesOutHi                   *int `json:"total_bytes_out_hi"`
		TotalBytesOutLo                   *int `json:"total_bytes_out_lo"`
		SslHandshakeTlsv12                *int `json:"ssl_handshake_tlsv12"`
		AnalyticsTransactionsMemoryUsage  *int `json:"analytics_transactions_memory_usage"`
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
