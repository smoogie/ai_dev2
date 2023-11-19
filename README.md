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
Task search requires data in artilces table and some data in articles collection in QDrant.

There is a go srcipt to fill in data. subdirectory: newsletter_search

You need to copy .env.example to .env and adjust it (for example put there your real open ai token)

After you fill in .env and started databases locally you can run script.
```
go get
go build
chmod +x newsletter_search
newsletter_search create-mysql-table
newsletter_search create-qdrant-collection
newsletter_search fill-data
```

##### Data for task people
Task people requires data in table people in MySQL.

There is a go srcipt to fill in data. subdirectory: people_data

You need to copy .env.example to .env and adjust it (for example put there your real open ai token)

After you fill in .env and started databases locally you can run script.
```
go get
go build
chmod +x people_data
people_data create-mysql-table
people_data fill-data
```

##### Data structure for tasks with own private api like ownapi, ownapipro
Some tasks require additional private api, those tasks use also MySQL.

There is a go srcipt to create require table. subdirectory: private_api_setup

You need to copy .env.example to .env and adjust it

After you fill in .env and started databases locally you can run script.
```
go get
go build
chmod +x private_api_setup
private_api_setup create-mysql-table
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


Part of the tasks require private API that is used by ai_dev tests servers. You can find that server in other directory, Make sure to run it before you run specific commands in main cli.

directory: go_lang_server
```
go get
go build
chmos +x ai_dev_private_api
ai_dev_private_api
```