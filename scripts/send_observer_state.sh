#!/bin/bash

 curl -X POST \
     localhost:8080/set-state\
     -d @observer_state.json
