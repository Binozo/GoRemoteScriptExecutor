# GoRemoteScriptExecutor
Securely execute scripts on remote Servers using simple HTTP GET requests.

[![Compilation check](https://github.com/Binozo/GoRemoteScriptExecutor/actions/workflows/compile-check.yaml/badge.svg)](https://github.com/Binozo/GoRemoteScriptExecutor/actions/workflows/compile-check.yaml)
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

## HTTP Endpoints

### Note: The HTTP Server runs on port 5123
Note: you will need to pass the Authorization header with your password.
```Authorization: <password>```
### `/`
Returns some basic info like application name and version.
Example:
```json
{ 
  "name" : "GoRemoteScriptExecutor",
  "version" : "0.1"
}
```

### `/scripts`
Returns a list of all scripts found in the scripts directory.
Example:
```json
{
  "scripts" : [
    "script1.sh",
    "script2.sh"
  ],
  "status" : "ok"
}
```

### `/runScript/<scriptname>`
Executes the script with the given name.
Possible GET parameters:
- `blocking` (optional) - If set to `true` the script will be executed in blocking mode. Default is `false`.
- `responseOutput` (optional) - If set to `true` the script output will be sent back to you. Default is `false`.
Example:
```json
{
  "status" : "ok"
}
```

### `/update`
Checks if an update is available and updates automatically.
Example:
```json
{
  "status" : "ok",
  "message" : "Update available. Update will be downloaded and installed automatically."
}
```