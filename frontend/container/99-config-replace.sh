#!/bin/sh
contents="$(jq --arg v "$API_URL" '.api = $v' "/usr/share/nginx/html/config.json")" && echo "${contents}" > "/usr/share/nginx/html/config.json"