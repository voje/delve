run-agent:
	air

run-server:
	go run cmd/server/main.go

# When server's not running (serving configuration), we can update agent manually via curl 
update-agent:
	cd scripts && ./update_agent.sh
