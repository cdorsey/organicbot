pwd
ls
echo "${SSH_KEY}" > ssh_key
chmod 400 ./ssh_key
scp -i ./ssh_key ./organicbot $REMOTE_FILE