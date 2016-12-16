pwd
ls
echo "${SSH_KEY}" > ssh_key
chmod 400 ./ssh_key
spawn "scp -i ./ssh_key ./organicbot $REMOTE_FILE"

expect "Enter passphrase for key './ssh_key': " { send "\r" }