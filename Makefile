HOST=localhost
PORT=8080
DOMAIN=$(HOST):$(PORT)

run:
	go run main.go

test:
	curl -i $(DOMAIN)/api/todo/2|grep "404 Not Found"
	curl -i -X POST $(DOMAIN)/api/todo -d title="hoge" -d description="fuga"|grep "200 OK"
	curl -i $(DOMAIN)/api/todo/2|grep "200 OK"
	curl $(DOMAIN)/api/todo/2|grep '"done":false'
	curl -i -X PUT $(DOMAIN)/api/todo/2 -d done="true"|grep "200 OK"
	curl $(DOMAIN)/api/todo/2|grep '"done":true'
	echo "ok"

