run:
	export SERVER_HOST=localhost;\
	export SERVER_PORT=8000;\
	export DB_DRIVER=mysql;\
	export DB_USER=root;\
	export DB_PASS="";\
	export DB_HOST=localhost;\
	export DB_PORT=3306;\
	export DB_SCHEMA=db_user;\
	go run main.go;
