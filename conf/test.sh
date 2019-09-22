#!/usr/bin/env bash

envoydir=$1
basedir=$2

if [ -z "$basedir" ]; then
  basedir=.
fi

if [ -z "$envoydir" ]; then
  envoydir=.
fi

if [ ! -r $basedir/bin/personserver ]; then
	echo "$basedir/bin/personserver not accessible"
	exit
fi

if [ ! -r $envoydir/envoy ]; then
	echo "$envoydir/envoy not accessible"
	exit
fi

echo "-------- Starting GRPC server -----------"
$basedir/bin/personserver &
echo "-------- Starting Envoy -----------"
$envoydir/envoy -c $basedir/conf/envoy.yaml &

sleep 5

echo "-------- Running tests -----------"
curl -X POST -d '{"name": "Sean", "age": 21}' -H "Content-Type: application/json" http://0.0.0.0:7778/create
curl -X POST -d '{"name": "Jeff", "age": 42}' -H "Content-Type: application/json" http://0.0.0.0:7778/create
curl -X POST -d '{"name": "Jeff", "age": 42}' -H "Content-Type: application/json" http://0.0.0.0:7778/create
curl -H "Content-Type: application/json" http://0.0.0.0:7778/lookup?name=Sean\&age=21
curl -H "Content-Type: application/json" http://0.0.0.0:7778/lookup?name=Jeff\&age=42
curl -H "Content-Type: application/json" http://0.0.0.0:7778/lookup?name=Neil\&age=42

kill %2
kill %1
