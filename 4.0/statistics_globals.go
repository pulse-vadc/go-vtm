// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 4.0.
package vtm

import (
	"encoding/json"
)

type GlobalsStatistics struct {
	Statistics struct {
		SslSessionIdMemCacheHit           *int `json:"ssl_session_id_mem_cache_hit"`
		SslCipherEcdsaSigns               *int `json:"ssl_cipher_ecdsa_signs"`
		AnalyticsTransactionsMemoryUsage  *int `json:"analytics_transactions_memory_usage"`
		TotalBytesIn                      *int `json:"total_bytes_in"`
		SysMemBuffered                    *int `json:"sys_mem_buffered"`
		SslCipherAesDecrypts              *int `json:"ssl_cipher_aes_decrypts"`
		NumberDnsaCacheHits               *int `json:"number_dnsa_cache_hits"`
		SslCipherAesGcmDecrypts           *int `json:"ssl_cipher_aes_gcm_decrypts"`
		AnalyticsTransactionsExported     *int `json:"analytics_transactions_exported"`
		SslCipherRsaEncryptsExternal      *int `json:"ssl_cipher_rsa_encrypts_external"`
		HourlyPeakBytesInPerSecond        *int `json:"hourly_peak_bytes_in_per_second"`
		NumberSnmpGetRequests             *int `json:"number_snmp_get_requests"`
		SslHandshakeTLSv1                 *int `json:"ssl_handshake_t_l_sv1"`
		SslCipherRsaEncrypts              *int `json:"ssl_cipher_rsa_encrypts"`
		SslSessionIdDiskCacheMiss         *int `json:"ssl_session_id_disk_cache_miss"`
		NumberSnmpBadRequests             *int `json:"number_snmp_bad_requests"`
		UpTime                            *int `json:"up_time"`
		TotalBadDnsPackets                *int `json:"total_bad_dns_packets"`
		HourlyPeakRequestsPerSecond       *int `json:"hourly_peak_requests_per_second"`
		SslCipherDhAgreements             *int `json:"ssl_cipher_dh_agreements"`
		HourlyPeakBytesOutPerSecond       *int `json:"hourly_peak_bytes_out_per_second"`
		NumberSnmpGetNextRequests         *int `json:"number_snmp_get_next_requests"`
		EventsSeen                        *int `json:"events_seen"`
		AnalyticsTransactionsDropped      *int `json:"analytics_transactions_dropped"`
		SslClientCertInvalid              *int `json:"ssl_client_cert_invalid"`
		TotalBytesInLo                    *int `json:"total_bytes_in_lo"`
		TotalBytesOut                     *int `json:"total_bytes_out"`
		NumberSnmpGetBulkRequests         *int `json:"number_snmp_get_bulk_requests"`
		SysMemSwapTotal                   *int `json:"sys_mem_swap_total"`
		SslCipherDesDecrypts              *int `json:"ssl_cipher_des_decrypts"`
		SslClientCertRevoked              *int `json:"ssl_client_cert_revoked"`
		NumberDnsptrCacheHits             *int `json:"number_dnsptr_cache_hits"`
		SslHandshakeSslv3                 *int `json:"ssl_handshake_sslv3"`
		SslCipherEcdhAgreements           *int `json:"ssl_cipher_ecdh_agreements"`
		SslCipher3DesEncrypts             *int `json:"ssl_cipher_3des_encrypts"`
		SslCipherRsaDecrypts              *int `json:"ssl_cipher_rsa_decrypts"`
		SslHandshakeSslv2                 *int `json:"ssl_handshake_sslv2"`
		SslCipherEcdhGenerates            *int `json:"ssl_cipher_ecdh_generates"`
		TotalBytesOutLo                   *int `json:"total_bytes_out_lo"`
		SysMemFree                        *int `json:"sys_mem_free"`
		TotalBackendServerErrors          *int `json:"total_backend_server_errors"`
		SslCipherDsaSigns                 *int `json:"ssl_cipher_dsa_signs"`
		SslSessionIdMemCacheMiss          *int `json:"ssl_session_id_mem_cache_miss"`
		TotalBytesOutHi                   *int `json:"total_bytes_out_hi"`
		TotalConn                         *int `json:"total_conn"`
		DataEntries                       *int `json:"data_entries"`
		SslCipherRc4Encrypts              *int `json:"ssl_cipher_rc4_encrypts"`
		NumberDnsptrRequests              *int `json:"number_dnsptr_requests"`
		SslCipherDesEncrypts              *int `json:"ssl_cipher_des_encrypts"`
		TotalBytesInHi                    *int `json:"total_bytes_in_hi"`
		TotalDnsResponses                 *int `json:"total_dns_responses"`
		SysFdsFree                        *int `json:"sys_fds_free"`
		SslCipherDsaVerifies              *int `json:"ssl_cipher_dsa_verifies"`
		SslClientCertNotSent              *int `json:"ssl_client_cert_not_sent"`
		SslCipherRsaDecryptsExternal      *int `json:"ssl_cipher_rsa_decrypts_external"`
		SysMemInUse                       *int `json:"sys_mem_in_use"`
		SslCipherAesGcmEncrypts           *int `json:"ssl_cipher_aes_gcm_encrypts"`
		SslHandshakeTLSv12                *int `json:"ssl_handshake_t_l_sv12"`
		TotalCurrentConn                  *int `json:"total_current_conn"`
		SysCpuSystemBusyPercent           *int `json:"sys_cpu_system_busy_percent"`
		SslCipherAesEncrypts              *int `json:"ssl_cipher_aes_encrypts"`
		SslCipherDhGenerates              *int `json:"ssl_cipher_dh_generates"`
		SysCpuIdlePercent                 *int `json:"sys_cpu_idle_percent"`
		SslClientCertExpired              *int `json:"ssl_client_cert_expired"`
		SslConnections                    *int `json:"ssl_connections"`
		NumberSnmpUnauthorisedRequests    *int `json:"number_snmp_unauthorised_requests"`
		SysMemTotal                       *int `json:"sys_mem_total"`
		SslSessionIdDiskCacheHit          *int `json:"ssl_session_id_disk_cache_hit"`
		NumberChildProcesses              *int `json:"number_child_processes"`
		SysMemSwapped                     *int `json:"sys_mem_swapped"`
		SslCipherRc4Decrypts              *int `json:"ssl_cipher_rc4_decrypts"`
		SslCipher3DesDecrypts             *int `json:"ssl_cipher_3des_decrypts"`
		NumberDnsaRequests                *int `json:"number_dnsa_requests"`
		SslHandshakeTLSv11                *int `json:"ssl_handshake_t_l_sv11"`
		HourlyPeakSslConnectionsPerSecond *int `json:"hourly_peak_ssl_connections_per_second"`
		DataMemoryUsage                   *int `json:"data_memory_usage"`
		TotalTransactions                 *int `json:"total_transactions"`
		SslCipherDecrypts                 *int `json:"ssl_cipher_decrypts"`
		SysCpuBusyPercent                 *int `json:"sys_cpu_busy_percent"`
		SslCipherEcdsaVerifies            *int `json:"ssl_cipher_ecdsa_verifies"`
		NumIdleConnections                *int `json:"num_idle_connections"`
		SysCpuUserBusyPercent             *int `json:"sys_cpu_user_busy_percent"`
		TimeLastConfigUpdate              *int `json:"time_last_config_update"`
		SslCipherEncrypts                 *int `json:"ssl_cipher_encrypts"`
		TotalRequests                     *int `json:"total_requests"`
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
