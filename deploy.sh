pwd
ls
echo "${SSH_KEY}" > ~/.ssh/id_rsa
chmod 400 ~/.ssh/id_rsa
scp -B ./organicbot $REMOTE_FILE