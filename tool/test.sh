#!/bin/sh

ps aux | grep ./app | grep -v grep | awk '{ print "kill -9", $2 }' | sh
