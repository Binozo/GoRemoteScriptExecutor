#!/bin/sh

if [ "$(id -u)" -ne 0 ]; then
        echo 'This script must be run by root' >&2
        exit 1
fi

read -p "Do you want to install GoRemoteScriptExecutor (y/n)?" choice
case "$choice" in
  y|Y ) echo "Installing...";;
  n|N ) exit;;
  * ) echo "invalid" && exit;;
esac

arch=$(dpkg --print-architecture)
echo "Searching executable for architecture $arch..."

releasesJson=$(curl -s https://api.github.com/repos/Binozo/GoRemoteScriptExecutor/releases/latest)
# Getting Tag name
tagName=$(echo $releasesJson | grep -o -P '(?<="tag_name": ").*(?=", "target_commitish)')
echo "Latest release is $tagName"

downloadUrl="https://github.com/Binozo/GoRemoteScriptExecutor/releases/download/$tagName/goremotescriptexecutor_$arch"
path="/usr/local/bin/goremotescriptexecutor/"
echo "Downloading $downloadUrl to $path..."
filename="goremotescriptexecutor"
mkdir -p $path
wget "$downloadUrl" -O "$path$filename"
chmod +x "$path$filename"

echo "A password is required. If you make a http request you will need to pass it as Authorization header."
stty -echo # Disable echoing
read -p "Please enter a password: " password
stty echo # Enable echoing
printf '\n'

# Generating password
cd $path
$path$filename -set-password "$password"

realuser="${SUDO_USER:-${USER}}"

echo "Installing systemd service..."
echo "
[Unit]
Description=GoRemoteScriptExecutor
ConditionPathExists=$path$filename
After=network.target

[Service]
User=$realuser

WorkingDirectory=$path
ExecStart=$path$filename
Restart=on-failure
RestartSec=10

[Install]
WantedBy=multi-user.target
" > /etc/systemd/system/goremotescriptexecutor.service

systemctl daemon-reload
systemctl enable goremotescriptexecutor.service
systemctl start goremotescriptexecutor.service

echo "Done!"