#!/bin/bash

start() {
    kill -9 $(cat {{.PidFile}})
    nohup {{.Command}} {{.Option}} > {{.LogFile}} 2>&1 &
    echo $! > {{.PidFile}}
}

stop() {
    kill -9 $(cat {{.PidFile}})
}

version() {
    echo {{.Version}}
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
  version)
        version
        ;;
  *)
        echo $"Usage: $0 {start|stop|restart|version}"
        exit 1
esac
