# go-uid-whosonfirst-proxy

Work in progress

## Documentation

Documentation is incomplete

## Example

```
$> go run cmd/server/main.go -provider-uri 'proxy://?provider=whosonfirst://&minimum=5'
2023/03/08 12:57:24 💬 Listening for requests at http://localhost:8080
2023/03/08 12:57:29 pool length: 5
2023/03/08 12:57:34 pool length: 5
2023/03/08 12:57:39 pool length: 3
2023/03/08 12:57:44 refill poll w/ 2 integers and 10 workers
2023/03/08 12:57:44 time to refill the pool with 5 integers (success: 5 failed: 0): 20.001307335s (pool length is now 3)
2023/03/08 12:57:44 pool length: 3
2023/03/08 12:57:49 pool length: 5
```

```
$> curl localhost:8080/
1813516673
```

## See also

* https://github.com/aaronland/go-uid
* https://github.com/aaronland/go-uid-server
* https://github.com/aaronland/go-uid-proxy
* https://github.com/aaronland/go-uid-whosonfirst