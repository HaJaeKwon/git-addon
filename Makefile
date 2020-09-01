move:
	go build -o bin/gitmove git-addon/cmd/move

user:
	go build -o bin/gituser git-addon/cmd/user

all:
	move
	user
