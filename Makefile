docker-build:
	docker build -t termfolio .

docker-run:
	docker run -d -it -p 2222:42069 termfolio