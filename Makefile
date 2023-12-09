migratefilesup: # nc-0
	migrate create -ext sql -dir db/migration -seq init_schema

postgres: # nc_1
	docker run --name postgresTechBlog -p 5432:5432 -e POSTGRES_USER=mosleh -e POSTGRES_PASSWORD=1234 -d postgres:latest

postgresstop:
	docker stop postgresTechBlog

postgresstart:
	docker start postgresTechBlog

postgresdown:
	docker rm postgresTechBlog

createdb: # nc_2
	docker exec -it postgresTechBlog createdb --username=mosleh --owner mosleh tech_blog

dropdb:
	docker exec -it postgresTechBlog dropdb --username=mosleh tech_blog

execdb: # access to database psql command line
	docker exec -it postgresTechBlog psql -U mosleh -n tech_blog

migrateup: # nc_3
	migrate -path db/migration -database "postgres://mosleh:1234@localhost:5432/tech_blog?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://mosleh:1234@localhost:5432/tech_blog?sslmode=disable" -verbose down

sqlc: # nc_4
	sqlc generate

run: # Run program
	go run main.go

.PHONY: migratefilesup postgres postgresstop postgresstart postgresdown createdb dropdb execdb migrateup migratedown sqlc run