# slack_cli
## Usage
Build
```
go build slk
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