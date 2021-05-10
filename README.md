# boink

send a slack message to a webhook when you see a line in a log file. make sure to put usernames in brackets.

usage:
```
export SLACK_WEBHOOK_URL=https://...
boink logfile 'pattern' 'message'
```

example:
```
export SLACK_WEBHOOK_URL=https://...
boink /var/log/cassandra/system.log 'sessions completed' '<@patcable> node is done'
```


