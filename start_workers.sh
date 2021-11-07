#!/usr/bin/env bash

if [[ "$1" == "kill" ]]; then
  ps aux | grep engine | awk '{print $2}' | xargs kill
else
  for i in {1..4} ; do
      nohup engine/engine --master=:50050 --addr=:5005$i &
  done
fi
