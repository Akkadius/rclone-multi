# rclone-multi

A simple wrapper for rclone for multi-remote backup operations

Built for personal backup needs, available for whoever finds useful. 

```
 INFO  [rclone-multi] A simple wrapper for rclone for multi-remote backup operations
 INFO  
 INFO  > upload [source-file] [destination-path]
 INFO  > trim [duration] Delete files older than this in seconds or ms|s|m|h|d|w|M|y Ex: 10d or 10s
 INFO  > exist [duration] [destination-path] Check for existence of files newer than this in seconds or alert. ms|s|m|h|d|w|M|y Ex: 10d or 10s
 INFO  
```

### Upload

Upload to as many remotes configured locally 

```
go run main.go upload ~/Downloads/Menuetto.ttf
cmiles@cmiles-msi:~/code/rclone-multi$ go run main.go upload ~/Downloads/Menuetto.ttf 
 INFO  Uploading file [/home/cmiles/Downloads/Menuetto.ttf] via remotes to path []
 INFO  -- Remote [local] Uploading [/home/cmiles/Downloads/Menuetto.ttf] to [local:]
 SUCCESS  -- Remote [local] Uploaded [/home/cmiles/Downloads/Menuetto.ttf] to [local:]

```

### Trim

Trim back files older than [duration]

```
go run main.go trim 30d
cmiles@cmiles-msi:~/code/rclone-multi$ go run main.go trim 30d
 INFO  Deleting files older than [30d]
 INFO  Remote [local]
 INFO  Deleting [Menuetto.ttf] via [local]
 INFO  rclone deletefile local:Menuetto.ttf
```

### Existence

Check for existence of files in a path within a specified time frame or send an alert. This is useful as a watchdog to let you know if your backups are failing you.

``` 
go run main.go exist 1s .
 INFO  Checking for existence of files duration [1s] via remotes to path [.]
 INFO  rclone lsl local:. --max-age=1s
 INFO  Remote [local]
```

**Alert Example**

The following example gets sent via notifiers

``` 
[ALERT] [hostname] No files or backups found for remote [local] path [.] duration [1s]
```

### Notifications

There are currently only two implemented and supported notifiers. They both implement an "Info" channel and an "Alert" channel.

**Slack**

``` 
NOTIFY_INFO_SLACK_WEBHOOK
NOTIFY_ALERT_SLACK_WEBHOOK
``` 

**Discord**

```
NOTIFY_INFO_DISCORD_WEBHOOK
NOTIFY_ALERT_DISCORD_WEBHOOK
```

#### Notification Source Labels

By default the notifications source from the hostname of the machine running the binary. If you would like to override this you need to set the following environment variable 

``` 
NOTIFY_SOURCE_LABEL
```
