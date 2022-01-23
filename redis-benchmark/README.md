# 分布式缓存作业

1. 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

```
$ redis-benchmark -t set,get -d 10 -q
SET: 92421.44 requests per second
GET: 92421.44 requests per second

$ redis-benchmark -t set,get -d 20 -q
SET: 93720.71 requests per second
GET: 95147.48 requests per second

$ redis-benchmark -t set,get -d 50 -q
SET: 91074.68 requests per second
GET: 95510.98 requests per second

$ redis-benchmark -t set,get -d 100 -q
SET: 94517.96 requests per second
GET: 93283.58 requests per second

$ redis-benchmark -t set,get -d 200 -q
SET: 94696.97 requests per second
GET: 94428.70 requests per second

$ redis-benchmark -t set,get -d 1000 -q
SET: 85836.91 requests per second
GET: 88028.16 requests per second

$ redis-benchmark -t set,get -d 5000 -q
SET: 86880.97 requests per second
GET: 81632.65 requests per second
```

2. 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

* 写入前内存信息

```
$ redis-cli info memory
# Memory
used_memory:865048
used_memory_human:844.77K
used_memory_rss:5304320
used_memory_rss_human:5.06M
used_memory_peak:866664
used_memory_peak_human:846.35K
used_memory_peak_perc:99.81%
used_memory_overhead:803064
used_memory_startup:803064
used_memory_dataset:61984
used_memory_dataset_perc:100.00%
allocator_allocated:1147168
allocator_active:1486848
allocator_resident:3850240
total_system_memory:2084085760
total_system_memory_human:1.94G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.30
allocator_frag_bytes:339680
allocator_rss_ratio:2.59
allocator_rss_bytes:2363392
rss_overhead_ratio:1.38
rss_overhead_bytes:1454080
mem_fragmentation_ratio:6.60
mem_fragmentation_bytes:4501240
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:0
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
```

* 写入 kv 数据

```
$ go build -o write-redis main.go
$ ./write-redis -d 100 -n 10000
```

* 写入后内存信息

```
$ redis-cli info memory
# Memory
used_memory:3046720
used_memory_human:2.91M
used_memory_rss:7585792
used_memory_rss_human:7.23M
used_memory_peak:12563088
used_memory_peak_human:11.98M
used_memory_peak_perc:24.25%
used_memory_overhead:1705208
used_memory_startup:803064
used_memory_dataset:1341512
used_memory_dataset_perc:59.79%
allocator_allocated:3133728
allocator_active:3485696
allocator_resident:8159232
total_system_memory:2084085760
total_system_memory_human:1.94G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.11
allocator_frag_bytes:351968
allocator_rss_ratio:2.34
allocator_rss_bytes:4673536
rss_overhead_ratio:0.93
rss_overhead_bytes:-573440
mem_fragmentation_ratio:2.54
mem_fragmentation_bytes:4601040
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:0
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
```
