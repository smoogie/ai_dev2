# ai_dev
Labs from AI DEV 2 course. Each directory is the solution for labs in the selected language.

### test server
directory: test_server

Simple implementation of server that simulates server used during the course by lecturers to verify tasks.

It does not represent the real logic provided by lecturers. It is only a simulator/mock for tests.
```
docker compose run web npm install
docker compose up -d

docker compose stop

docker compose start
```

test server would be available on http://localhost:8080


### additional materials required for tasks
directory: additional_task_resources

##### Databases
Some of the tasks require access to local DBs.
Subfolder databases has docker configuration for the MySQL, Qdrant.
```
docker compose up -d

docker compose stop

docker compose start
```
#### Mysql
Mysql is available on default port.

- DB = ai_dev
- user/pass = root/root

##### Data for task search
Task search require tables in artilces table and some data in QDrant.

First make sure that Qdrant and MySQL are running. 
You also need to make sure that article table exists in MySQL DB.
If table does not exists, you can create it with sql script available in file additional_task_resources/newsletter_search/init_table.sql

There is a go srcipt to fill in data. subdirectory: newsletter_search

You need to copy .env.example to .env and adjust it (for example put there your real open ai token)

After you fill in .env and started databases locally you can run script.
```
go get
go build
chmod +x newsletter_search
newsletter_search create-qdrant-collection
newsletter_search fill-data
```


### Go lang
directory: go_lang
```
go get
go build
chmod +x ai_dev
ai_dev cXXlXX
```
cXXlXX replace with lab number, for example c01l01
