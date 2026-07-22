#!/bin/bash

SERVER_PORT=8081
ADDRESS="localhost:${SERVER_PORT}"
TEMP_FILE="tempfile"
SERVER_BINARY=cmd/server/server.exe
AGENT_BINARY=cmd/agent/agent.exe
LOG_DIR=logs

if [ ! -e $LOG_DIR ] ; then
    mkdir ${LOG_DIR}
fi

if [ ! -d $LOG_DIR ] ; then
    echo "${LOG_DIR} is not a directory"
    exit 1
fi

echo -n "Increment 1	...	"
if ./metricstest.exe -test.v -test.run=^TestIteration1$ \
            -binary-path=${SERVER_BINARY} &> ${LOG_DIR}/iter1.log
then
	echo "[ OK ]"
else
	echo "[FAIL]"
	exit 1
fi

echo -n "Increment 2	...	"			
if ./metricstest.exe -test.v -test.run=^TestIteration2[AB]*$ \
            -source-path=. \
            -agent-binary-path=${AGENT_BINARY} &> ${LOG_DIR}/iter2.log
then
	echo "[ OK ]"
else
	echo "[FAIL]"
	exit 1
fi

echo -n "Increment 3	...	"
if ./metricstest.exe -test.v -test.run=^TestIteration3[AB]*$ \
            -source-path=. \
            -agent-binary-path=${AGENT_BINARY} \
            -binary-path=${SERVER_BINARY} &> ${LOG_DIR}/iter3.log
then
	echo "[ OK ]"
else
	echo "[FAIL]"
	exit 1
fi

echo -n "Increment 4	...	"
if ./metricstest.exe -test.v -test.run=^TestIteration4$ \
            -agent-binary-path=${AGENT_BINARY} \
            -binary-path=${SERVER_BINARY} \
            -server-port=${SERVER_PORT} \
            -source-path=. &> ${LOG_DIR}/iter4.log
then
	echo "[ OK ]"
else
	echo "[FAIL]"
	exit 1
fi

echo -n "Increment 5	...	"
if ./metricstest.exe -test.v -test.run=^TestIteration5$ \
            -agent-binary-path=${AGENT_BINARY} \
            -binary-path=${SERVER_BINARY} \
            -server-port=${SERVER_PORT} \
            -source-path=. &> ${LOG_DIR}/iter5.log
then
	echo "[ OK ]"
else
	echo "[FAIL]"
	exit 1
fi

echo -n "Increment 6	...	"
if ./metricstest.exe -test.v -test.run=^TestIteration6$ \
            -agent-binary-path=${AGENT_BINARY} \
            -binary-path=${SERVER_BINARY} \
            -server-port=${SERVER_PORT} \
            -source-path=. &> ${LOG_DIR}/iter6.log
then
	echo "[ OK ]"
else
	echo "[FAIL]"
	exit 1
fi

echo -n "Increment 7	...	"
if ./metricstest.exe -test.v -test.run=^TestIteration7$ \
            -agent-binary-path=${AGENT_BINARY} \
            -binary-path=${SERVER_BINARY} \
            -server-port=${SERVER_PORT} \
            -source-path=. &> ${LOG_DIR}/iter7.log
then
	echo "[ OK ]"
else
	echo "[FAIL]"
	exit 1
fi

echo -n "Increment 8	...	"
if ./metricstest.exe -test.v -test.run=^TestIteration8$ \
            -agent-binary-path=${AGENT_BINARY} \
            -binary-path=${SERVER_BINARY} \
            -server-port=${SERVER_PORT} \
            -source-path=. &> ${LOG_DIR}/iter8.log
then
	echo "[ OK ]"
else
	echo "[FAIL]"
	exit 1
fi

echo -n "Increment 9	...	"
if ./metricstest -test.v -test.run=^TestIteration9$ \
            -agent-binary-path=${AGENT_BINARY} \
            -binary-path=${SERVER_BINARY} \
            -file-storage-path=${TEMP_FILE} \
            -server-port=${SERVER_PORT} \
            -source-path=. &> ${LOG_DIR}/iter9.log
then
	echo "[ OK ]"
else
	echo "[FAIL]"
	exit 1
fi
