cat $SSH_KEY > ssh_key
go build -o ./organicbot $GOPATH/src/github.com/cdorsey/organicbot
scp -i ./ssh_key ./organicbot $REMOTE_FILE