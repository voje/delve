#!/bin/bash

 curl -X POST \
     localhost:8080/configure \
     -d @agent_conf.json
