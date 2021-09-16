#!/bin/bash

# CAUTION: This script is executed as root. It is used to deploy a new vps.

user_name=$1

function exitln() {
    echo $1
    exit 1
}

# check environments
if [[ -z "${PORTNUM}" ]]; then
    exitln "PORTNUM not set"
fi

if [[ -z "${user_name}" ]]; then
    user_name="wm"
fi

# 1. add a new user
useradd -s /bin/bash -d "/home/${user_name}" -m "${user_name}"

# 2. change ssh port and restart service
sed -i.bak '/Port 22/s/^/#/' /etc/ssh/sshd_config
echo "Port ${PORTNUM}" >> /etc/ssh/sshd_config
systemctl restart sshd

# 3. forbidden root login
