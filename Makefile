HOST=localhost
PORT=8080
DOMAIN=$(HOST):$(PORT)

run:
	go run main.go

test:
	curl -i $(DOMAIN)/todo/1|grep "404 Not Found"
	curl -X POST $(DOMAIN)/todo -d title="hoge" -d contents="fuga"|grep ok
	curl -i $(DOMAIN)/todo/1|grep "200 OK"
	curl $(DOMAIN)/todo/1|grep '"title":"hoge"'
	curl -X PUT $(DOMAIN)/todo/1 -d title="piyo" -d contents="fuga"
	curl $(DOMAIN)/todo/1|grep '"title":"piyo"'
	curl -X DELETE $(DOMAIN)/todo/1|grep ok
	curl -i $(DOMAIN)/todo/1|grep "404 Not Found"
	echo "ok"

