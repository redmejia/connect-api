#!/bin/bash

# .env enviroment variables file
DOT_ENV=internal/env/.env

# clean and stop service
if [[ $1 == "sp" ]]; then
	make clean
	make stop 2> /dev/null
fi

if [[ $2 == "st" ]];then
	# run enviroment varabiables and build and start server
	if [ -f "$DOT_ENV" ]; then
		# echo "$DOT_ENV found setting and running server"
		# \:*[a-zA-Z0-9_]* vim or neovim  :w :53 on save error
		rm \:*[a-zA-Z0-9_]* 2> /dev/null || rm [^a-zA-Z0-9_] 2> /dev/null  || source internal/env/.env && make start
		
	else
		# crate .env file and edit with the enviroment variables
		echo "$DOT_ENV was not found"
		echo "Creating"
		touch internal/env/.env

	fi
fi



