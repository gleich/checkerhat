SSH_LOC="pi@zephyr-checkerhat.local"

deploy:
	GOOS=linux GOARCH=arm go build -o bin/checkerhat .
	rsync -avhut --progress --delete ./bin/ $(SSH_LOC):~/bin/
	ssh $(SSH_LOC) << echo "~/bin/checkerhat"

connect:
	sed -i '' -e '$ d' ~/.ssh/known_hosts
	ssh $(SSH_LOC) -p 3000
