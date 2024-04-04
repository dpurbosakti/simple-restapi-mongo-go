mongorun:
		docker run -d -p 27017:27017 --name mongodb mongo

mongostart:
		docker start mongodb

mongostop:
		docker stop mongodb

.PHONY: mongorun mongostart mongostop