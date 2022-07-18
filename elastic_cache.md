# Amazon ElastiCache
Managed Redis or Mencached Database instance

## Features
* Caches are in-memory databases with really high performance, low latency
* Helps reduce load off of databases for read intensive workloads
* Helps make your application stateless
* Cache must have an invalidation strategy to make sure only the most current data is used in there.
* Some good use cases for that would be User Session Store and Shopping Cart

## Redis vs Memcached
* Redis is Multi AZ, whereas memcached is multi-node
* Redis Read Replicates have high availability
* Redis has Data Durability using AOF persistence
* Redis backup and restore features
* Redis auth

## Strategies

### Lazy Loading / Cache-Aside / Lazy Population
1. The application looks up the cache
2. If the record is found, then use it
3. Else read from the database, and update the cache after it

#### Pros
* Only requested data is cached
* Node failures are not fatal

#### Cons
* Cache miss penalty that results in 3 round trips
* Stale data

### Write Through
1. The cache is always updated for all reads
2. For every right to the DB, the cache is also updated

#### Pros
* Data in cache is never stale, reads are quick
* Write penalty

#### Cons
* Missing Data until it is added / updated in the DB.
* Cache churn


### Cache Evictions
* Cache eviction can occur in three ways: 
	* You delete the item explicitly in the cache
	* Item is evicted because the memory is full
	* You set an item item-to-live (or TTL)
* If too many evictions happen due to memory, you should scale up or out

