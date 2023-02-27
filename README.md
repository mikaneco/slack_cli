# slack_cli
## Usage
### Build
Building this application and moving bin file
```
go build -o slk && mv slk /usr/local/bin/
```

Command
### Setting Config
```
slk config --channel "channel name" --token "token"
```

### Sending message
```
slk post -m "message"
```

### Sending specific channel to message
```
slk post -c "channel" -m "message"
```