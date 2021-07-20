SSH_LOC="pi@zephyr-arcade.local"

deploy:
	GOOS=linux GOARCH=arm go build -o bin/arcade .
	rsync -avhut --progress --delete ./bin/ $(SSH_LOC):~/bin/
	ssh $(SSH_LOC) << echo "~/bin/arcade"
