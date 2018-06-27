#! /usr/bin/env bash
TMP=$HOME/tmp

# from https://www.microsoft.com/net/learn/get-started/linux/ubuntu18-04
curl -sSL https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > $TMP/microsoft.asc.gpg
sudo -A mv $TMP/microsoft.asc.gpg /etc/apt/trusted.gpg.d/

curl -sSL https://packages.microsoft.com/config/ubuntu/18.04/prod.list > $TMP/dotnet.list
sudo -A mv $TMP/dotnet.list /etc/apt/sources.list.d/

sudo chown root:root /etc/apt/trusted.gpg.d/microsoft.asc.gpg
sudo chown root:root /etc/apt/sources.list.d/dotnet.list

sudo apt update
sudo apt install 'dotnet-sdk-2.1'
