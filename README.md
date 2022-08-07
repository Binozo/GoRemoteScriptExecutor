# GoRemoteScriptExecutor
Securely execute scripts on remote Servers using simple HTTP GET requests.

## Setup
### Automatic
I made a setup script to install GoRemoteScriptExecutor on your Server. It automatically detects your cpu architecture and downloads the correct executable and manages everything for you.

```bash
wget https://raw.githubusercontent.com/Binozo/GoRemoteScriptExecutor/master/install.sh -O install.sh && chmod +x install.sh && sudo ./install.sh
```
### Manual
You can also install GoRemoteScriptExecutor manually.
1. Download the latest executable from the [Releases Page](https://github.com/Binozo/GoRemoteScriptExecutor/releases) and rename it to `goremotescriptexecutor`.
2. Place the executable in the `/opt/goremotescriptexecutor/` directory.
3. Make it executable. `chmod +x /opt/goremotescriptexecutor/goremotescriptexecutor`
4. Add the `goremotescriptexecutor` directory to your user group. \
    `cd /opt/ && chown -R $USER:$USER goremotescriptexecutor`
5. Create a password: `/opt/goremotescriptexecutor/goremotescriptexecutor --set-password <password>`
6. If everything is fine, create a systemd service for the executable.

Note: For help look at the [installation script](https://github.com/Binozo/GoRemoteScriptExecutor/blob/master/install.sh).

## GitHub Actions support
Take a look at [GitHubActionRemoteScriptExecutor](https://github.com/Binozo/GitHubActionRemoteScriptExecutor).
