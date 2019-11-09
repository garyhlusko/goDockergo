# goDockergo

This builds and creates a docker image. 

Put this in the same directory as your docker-compose.yml file and run 

`go run main.go --dbPwd dbPassword --dbName databaseName --dbUser dbUsername --network_name NetworkName`

This will write to an .env file that your docker-compose.yml file will use.
