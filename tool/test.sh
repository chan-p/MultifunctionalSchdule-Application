#!/bin/sh

ps aux | grep ./main | grep -v grep | awk '{ print "kill -9", $2 }' | sh
