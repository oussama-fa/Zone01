#! /bin/bash

curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq --arg id "$HERO_ID" '.[] | select(.id == ($id | tonumber)) | .connections.relatives' | tr -d "\""