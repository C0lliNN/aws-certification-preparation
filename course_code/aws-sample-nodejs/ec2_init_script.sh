#!/bin/bash
yum install -y git
git config --system credential.helper '!aws codecommit credential-helper $@'
git config --system credential.UseHttpPath true
git clone https://git-codecommit.us-east-1.amazonaws.com/v1/repos/sample-nodejs /app
cd /app

echo "$(hostname -i)" > static/host.html

curl -sL https://rpm.nodesource.com/setup_14.x | bash -
yum install -y nodejs
npm install
export PORT=80
node index.js
