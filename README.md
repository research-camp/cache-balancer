# Cache Balancer

Load balancer router simulator for cache systems with medium hit ratio. This is a simple simulator for distributed caching systems.
It uses statistical operations to send real time requests among varios systems considering the load and hit ratio.

## collaborators

- Amirhossein Najafizadeh (Amirkabir University of Technology)
- Dr. Niloofar Charmchi (Université de Rennes I)

## how to run?

Make sure to have ```Golang``` installed on your system.

```
go run main.go
```

## steps

```
input: string(32)
```

```
match with a label in caches (this label could be for example IP or name of that cache)
```

```
matching function:
    length difference from that cache
    start from left and check 4 by 4 digits
    cache factor
```
