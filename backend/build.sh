#!/bin/sh

goimports -w ../backend

go build -o OnlineJudgeBackend && ./OnlineJudgeBackend