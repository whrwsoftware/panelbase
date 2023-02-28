#!/bin/bash

start() {
    kill -9 $(cat {{.PidFile}})
    nohup {{.Command}} {{.Option}} > {{.LogFile}} 2>&1 &
    echo $! > {{.PidFile}}
}

stop() {
    kill -9 $(cat {{.PidFile}})
}

status() {
    ps -p {{.PidFile}}
}

log() {
    tail -f -n 100 {{.LogFile}}
}

restart() {
    stop
    start
}

case "$1" in
  start)
        start
        ;;
  stop)
        stop
        ;;
  restart)
        restart
        ;;
  status)
        status
        ;;
  log)
        log
        ;;
  *)
        echo $"Usage: $0 {start|stop|restart|status|log}"
        exit 1
esac
