SSH_LOC="pi@zephyr-arcade.local"

deploy:
	GOOS=linux GOARCH=arm go build -o bin/arcade .
	rsync -avhut --progress --delete ./bin/ $(SSH_LOC):~/bin/
	ssh $(SSH_LOC) << echo "~/bin/arcade"

connect:
	sed -i '' -e '$ d' ~/.ssh/known_hosts
	ssh $(SSH_LOC) -p 3000
