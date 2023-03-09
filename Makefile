debug:
	go run cmd/server/main.go -provider-uri 'proxy://?provider=whosonfirst://&minimum=5&pool=boltdb://whosonfirst%3Fdsn=whosonfirst.db'
