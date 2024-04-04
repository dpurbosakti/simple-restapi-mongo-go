mongorun:
		docker run -d -p 27017:27017 --name mongodb mongo

mongostart:
		docker start mongodb

mongostop:
		docker stop mongodb

mongoexec:
		docker exec -it mongodb bash
## to run mongo cli simply type mongosh

test:
		go test ./... -v

.PHONY: mongorun mongostart mongostop mongoexec test