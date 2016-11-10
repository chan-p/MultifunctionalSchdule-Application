#!/bin/sh

ps aux | grep ./application | grep -v grep | awk '{ print "kill -9", $2 }' | sh
