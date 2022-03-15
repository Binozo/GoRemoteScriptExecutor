# Remote Script Executor
A Service which works 🤝 with [GitHubActionRemoteScriptExecutor](https://github.com/Binozo/GitHubActionRemoteScriptExecutor) \
Developed to automatically run scripts on target machines through GitHub Actions.

## Setup
1. Download the [main](https://github.com/Binozo/GoRemoteScriptExecutor/tree/master/main/main) executable and the systemd Service [file](https://github.com/Binozo/GoRemoteScriptExecutor/tree/master/remoteScriptExecutor.service)
2. Move the systemd Service file to ``/etc/systemd/system/``
3. Create a Folder named ``remoteScriptExecutor`` in the Home Directory (`/home/ubuntu/`)
4. Copy the main file to ``/home/ubuntu/remoteScriptExecutor/``
5. Create the ``creds.txt`` file in the same directory
6. **Write a password into that file** (Leaving it empty removes the password checking)
7. Create a shell script file in ``/home/ubuntu/remoteScriptExecutor/``
8. Run ``sudo systemctl daemon-reload && sudo systemctl start remoteScriptExecutor``

(If you want to change the password you have to restart the service in order to load the new password)

##Usage

Take a look at [this](https://github.com/Binozo/GitHubActionRemoteScriptExecutor)