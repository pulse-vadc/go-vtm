// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 4.0.
package vtm

import (
	"encoding/json"
)

type GlobalsStatistics struct {
	Statistics struct {
		TotalTransactions                 *int `json:"total_transactions"`
		TotalBytesOut                     *int `json:"total_bytes_out"`
		TotalDnsResponses                 *int `json:"total_dns_responses"`
		NumberDnsaRequests                *int `json:"number_dnsa_requests"`
		TotalBytesIn                      *int `json:"total_bytes_in"`
		SslCipherAesGcmDecrypts           *int `json:"ssl_cipher_aes_gcm_decrypts"`
		SslSessionIdDiskCacheMiss         *int `json:"ssl_session_id_disk_cache_miss"`
		SysMemInUse                       *int `json:"sys_mem_in_use"`
		SslClientCertExpired              *int `json:"ssl_client_cert_expired"`
		SslCipherDhAgreements             *int `json:"ssl_cipher_dh_agreements"`
		SslCipherDsaSigns                 *int `json:"ssl_cipher_dsa_signs"`
		SslSessionIdMemCacheHit           *int `json:"ssl_session_id_mem_cache_hit"`
		TotalCurrentConn                  *int `json:"total_current_conn"`
		SslCipher3DesEncrypts             *int `json:"ssl_cipher_3des_encrypts"`
		SslCipherRsaDecrypts              *int `json:"ssl_cipher_rsa_decrypts"`
		SslHandshakeTLSv1                 *int `json:"ssl_handshake_t_l_sv1"`
		HourlyPeakSslConnectionsPerSecond *int `json:"hourly_peak_ssl_connections_per_second"`
		EventsSeen                        *int `json:"events_seen"`
		TotalBytesOutLo                   *int `json:"total_bytes_out_lo"`
		SslClientCertNotSent              *int `json:"ssl_client_cert_not_sent"`
		SysCpuIdlePercent                 *int `json:"sys_cpu_idle_percent"`
		NumberDnsptrRequests              *int `json:"number_dnsptr_requests"`
		NumberSnmpGetNextRequests         *int `json:"number_snmp_get_next_requests"`
		SysCpuUserBusyPercent             *int `json:"sys_cpu_user_busy_percent"`
		SslCipherDesDecrypts              *int `json:"ssl_cipher_des_decrypts"`
		TotalBytesOutHi                   *int `json:"total_bytes_out_hi"`
		SslCipherEcdsaSigns               *int `json:"ssl_cipher_ecdsa_signs"`
		NumberDnsptrCacheHits             *int `json:"number_dnsptr_cache_hits"`
		DataMemoryUsage                   *int `json:"data_memory_usage"`
		SslCipherAesGcmEncrypts           *int `json:"ssl_cipher_aes_gcm_encrypts"`
		SslCipherEcdhGenerates            *int `json:"ssl_cipher_ecdh_generates"`
		NumberSnmpBadRequests             *int `json:"number_snmp_bad_requests"`
		HourlyPeakRequestsPerSecond       *int `json:"hourly_peak_requests_per_second"`
		AnalyticsTransactionsMemoryUsage  *int `json:"analytics_transactions_memory_usage"`
		SslCipherRc4Decrypts              *int `json:"ssl_cipher_rc4_decrypts"`
		UpTime                            *int `json:"up_time"`
		SslSessionIdMemCacheMiss          *int `json:"ssl_session_id_mem_cache_miss"`
		TotalBytesInHi                    *int `json:"total_bytes_in_hi"`
		NumberSnmpGetBulkRequests         *int `json:"number_snmp_get_bulk_requests"`
		NumberSnmpGetRequests             *int `json:"number_snmp_get_requests"`
		SslCipherAesDecrypts              *int `json:"ssl_cipher_aes_decrypts"`
		HourlyPeakBytesInPerSecond        *int `json:"hourly_peak_bytes_in_per_second"`
		SslCipherDhGenerates              *int `json:"ssl_cipher_dh_generates"`
		TimeLastConfigUpdate              *int `json:"time_last_config_update"`
		SysMemBuffered                    *int `json:"sys_mem_buffered"`
		SysMemSwapTotal                   *int `json:"sys_mem_swap_total"`
		NumberDnsaCacheHits               *int `json:"number_dnsa_cache_hits"`
		SslCipherEcdsaVerifies            *int `json:"ssl_cipher_ecdsa_verifies"`
		TotalConn                         *int `json:"total_conn"`
		SslCipherDecrypts                 *int `json:"ssl_cipher_decrypts"`
		TotalBadDnsPackets                *int `json:"total_bad_dns_packets"`
		DataEntries                       *int `json:"data_entries"`
		SslHandshakeTLSv11                *int `json:"ssl_handshake_t_l_sv11"`
		AnalyticsTransactionsExported     *int `json:"analytics_transactions_exported"`
		SslCipher3DesDecrypts             *int `json:"ssl_cipher_3des_decrypts"`
		SslCipherDesEncrypts              *int `json:"ssl_cipher_des_encrypts"`
		SslConnections                    *int `json:"ssl_connections"`
		SslSessionIdDiskCacheHit          *int `json:"ssl_session_id_disk_cache_hit"`
		SslClientCertRevoked              *int `json:"ssl_client_cert_revoked"`
		SslCipherAesEncrypts              *int `json:"ssl_cipher_aes_encrypts"`
		SysMemSwapped                     *int `json:"sys_mem_swapped"`
		NumberChildProcesses              *int `json:"number_child_processes"`
		SslCipherEcdhAgreements           *int `json:"ssl_cipher_ecdh_agreements"`
		SslCipherRsaEncryptsExternal      *int `json:"ssl_cipher_rsa_encrypts_external"`
		SysCpuSystemBusyPercent           *int `json:"sys_cpu_system_busy_percent"`
		SslCipherRc4Encrypts              *int `json:"ssl_cipher_rc4_encrypts"`
		SslCipherRsaEncrypts              *int `json:"ssl_cipher_rsa_encrypts"`
		SslHandshakeSslv2                 *int `json:"ssl_handshake_sslv2"`
		SslCipherRsaDecryptsExternal      *int `json:"ssl_cipher_rsa_decrypts_external"`
		HourlyPeakBytesOutPerSecond       *int `json:"hourly_peak_bytes_out_per_second"`
		TotalBackendServerErrors          *int `json:"total_backend_server_errors"`
		NumberSnmpUnauthorisedRequests    *int `json:"number_snmp_unauthorised_requests"`
		SslCipherEncrypts                 *int `json:"ssl_cipher_encrypts"`
		AnalyticsTransactionsDropped      *int `json:"analytics_transactions_dropped"`
		SslCipherDsaVerifies              *int `json:"ssl_cipher_dsa_verifies"`
		SslHandshakeTLSv12                *int `json:"ssl_handshake_t_l_sv12"`
		SslClientCertInvalid              *int `json:"ssl_client_cert_invalid"`
		SysMemFree                        *int `json:"sys_mem_free"`
		TotalBytesInLo                    *int `json:"total_bytes_in_lo"`
		NumIdleConnections                *int `json:"num_idle_connections"`
		TotalRequests                     *int `json:"total_requests"`
		SysMemTotal                       *int `json:"sys_mem_total"`
		SysCpuBusyPercent                 *int `json:"sys_cpu_busy_percent"`
		SslHandshakeSslv3                 *int `json:"ssl_handshake_sslv3"`
		SysFdsFree                        *int `json:"sys_fds_free"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetGlobalsStatistics() (*GlobalsStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/4.0/status/local_tm/statistics/globals")
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
