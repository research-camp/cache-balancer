# Cache Balancer

Load balancer router for cache systems with medium hit ratio.

```
input: string(32)

match with a label in caches (this label could be for example IP or name of that cache)

matching function:
    length difference from that cache
    start from left and check 4 by 4 digits
    cache factor
```
