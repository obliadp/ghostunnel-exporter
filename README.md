# Ghostunnel exporter

Annoyed with having to deal with certs for metrics.

This opens up port 8081, and accepts ghostunnels metrics POST request.

Any GET will produce prometheus metrics available from the latest POST.

```bash
./main &
Starting server for testing HTTP POST...

ghostunnel-v1.4.1-darwin-amd64-with-pkcs11 client \
	--listen localhost:8080 \
	--target some.api.horse:443 \
	--disable-authentication \
	--metrics-url http://localhost:8081

[45922] 2019/09/17 17:42:15.380226 starting ghostunnel in client mode
[45922] 2019/09/17 17:42:15.380326 metrics enabled; reporting metrics via POST to http://localhost:8081
[45922] 2019/09/17 17:42:15.528455 using target address api.some.horse:443
[45922] 2019/09/17 17:42:15.532116 listening for connections on localhost:8081
```

produces:

```
curl localhost:8081/metrics -D-
HTTP/1.1 200 OK
Date: Tue, 17 Sep 2019 15:29:39 GMT
Content-Type: text/plain; charset=utf-8
Transfer-Encoding: chunked

# Ghostunnel Runtime Cgo Calls
ghostunnel_runtime_cgo_calls: 24
# Ghostunnel Runtime Mem Heap Idle
ghostunnel_runtime_mem_heap_idle: 6.2038016e+07
# Ghostunnel Runtime Mem Heap Released
ghostunnel_runtime_mem_heap_released: 0
# Ghostunnel Accept Timeout
ghostunnel_accept_timeout: 0
# Ghostunnel Runtime Mem Stack Inuse
ghostunnel_runtime_mem_stack_inuse: 819200
# Ghostunnel Accept Error
ghostunnel_accept_error: 0
# Ghostunnel Conn Lifetime Count
ghostunnel_conn_lifetime_count: 1
# Ghostunnel Conn Lifetime Min
ghostunnel_conn_lifetime_min: 1.41933037e+08
# Ghostunnel Conn Lifetime Max
ghostunnel_conn_lifetime_max: 1.41933037e+08
# Ghostunnel Conn Lifetime Mean
ghostunnel_conn_lifetime_mean: 1.41933037e+08
# Ghostunnel Conn Lifetime 50 Percentile
ghostunnel_conn_lifetime_50_percentile: 1.41933037e+08
# Ghostunnel Conn Lifetime 75 Percentile
ghostunnel_conn_lifetime_75_percentile: 1.41933037e+08
# Ghostunnel Conn Lifetime 95 Percentile
ghostunnel_conn_lifetime_95_percentile: 1.41933037e+08
# Ghostunnel Conn Lifetime 99 Percentile
ghostunnel_conn_lifetime_99_percentile: 1.41933037e+08
# Ghostunnel Runtime Mem Heap Sys
ghostunnel_runtime_mem_heap_sys: 6.6289664e+07
# Ghostunnel Runtime Goroutines
ghostunnel_runtime_goroutines: 9
# Ghostunnel Conn Open
ghostunnel_conn_open: 0
# Ghostunnel Accept Success
ghostunnel_accept_success: 1
# Ghostunnel Runtime Mem Total Alloc
ghostunnel_runtime_mem_total_alloc: 6.90224e+06
# Ghostunnel Runtime Mem Stack Sys
ghostunnel_runtime_mem_stack_sys: 819200
# Ghostunnel Accept Total
ghostunnel_accept_total: 1
# Ghostunnel Runtime Mem Heap Alloc
ghostunnel_runtime_mem_heap_alloc: 2.199048e+06
# Ghostunnel Runtime Mem Mallocs
ghostunnel_runtime_mem_mallocs: 137626
# Ghostunnel Runtime Mem Alloc
ghostunnel_runtime_mem_alloc: 2.199048e+06
# Ghostunnel Runtime Mem Sys
ghostunnel_runtime_mem_sys: 7.25486e+07
# Ghostunnel Runtime Mem Gc Num Gc
ghostunnel_runtime_mem_gc_num_gc: 4
# Ghostunnel Runtime Mem Gc Duration Count
ghostunnel_runtime_mem_gc_duration_count: 4
# Ghostunnel Runtime Mem Gc Duration Min
ghostunnel_runtime_mem_gc_duration_min: 0
# Ghostunnel Runtime Mem Gc Duration Max
ghostunnel_runtime_mem_gc_duration_max: 0
# Ghostunnel Runtime Mem Gc Duration Mean
ghostunnel_runtime_mem_gc_duration_mean: 0
# Ghostunnel Runtime Mem Gc Duration 50 Percentile
ghostunnel_runtime_mem_gc_duration_50_percentile: 0
# Ghostunnel Runtime Mem Gc Duration 75 Percentile
ghostunnel_runtime_mem_gc_duration_75_percentile: 0
# Ghostunnel Runtime Mem Gc Duration 95 Percentile
ghostunnel_runtime_mem_gc_duration_95_percentile: 0
# Ghostunnel Runtime Mem Gc Duration 99 Percentile
ghostunnel_runtime_mem_gc_duration_99_percentile: 0
# Ghostunnel Runtime Mem Heap Inuse
ghostunnel_runtime_mem_heap_inuse: 4.251648e+06
# Ghostunnel Runtime Mem Frees
ghostunnel_runtime_mem_frees: 110025
# Ghostunnel Runtime Mem Heap Objects
ghostunnel_runtime_mem_heap_objects: 27601
# Ghostunnel Runtime Mem Lookups
ghostunnel_runtime_mem_lookups: 0
# Ghostunnel Conn Handshake Count
ghostunnel_conn_handshake_count: 0
# Ghostunnel Conn Handshake Min
ghostunnel_conn_handshake_min: 0
# Ghostunnel Conn Handshake Max
ghostunnel_conn_handshake_max: 0
# Ghostunnel Conn Handshake Mean
ghostunnel_conn_handshake_mean: 0
# Ghostunnel Conn Handshake 50 Percentile
ghostunnel_conn_handshake_50_percentile: 0
# Ghostunnel Conn Handshake 75 Percentile
ghostunnel_conn_handshake_75_percentile: 0
# Ghostunnel Conn Handshake 95 Percentile
ghostunnel_conn_handshake_95_percentile: 0
# Ghostunnel Conn Handshake 99 Percentile
ghostunnel_conn_handshake_99_percentile: 0
# Ghostunnel Runtime Mem Gc Cpu Fraction
ghostunnel_runtime_mem_gc_cpu_fraction: 3.3814443776699805e-06
```
