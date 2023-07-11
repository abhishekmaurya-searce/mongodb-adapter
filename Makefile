emulator_run:
	podman-compose -f podman-compose.yml up
spanner_cli:
	podman-compose -f podman-compose.yml exec spanner-cli spanner-cli -p spanner-project -i spanner-instance -d spanner-database
gorun:
	#mongodb+srv://abhishekmaurya:abhishek@clusterblog.r9oxx61.mongodb.net 
	#library_management 
	go run cmd/library/*.go mongodb+srv://abhishekmaurya:abhishek@clusterblog.r9oxx61.mongodb.net sample_geospatial projects/spanner-project/instances/spanner-instance/databases/spanner-database

.PHONY: emulator_run spanner_cli gorun