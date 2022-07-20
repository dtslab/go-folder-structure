#### Go language folder structure for web/server

1. You can place the readme texts here. Like setup steps of the project, or any other information about the project that you think would be helpful for other develepors/user

2. Create ".env" file, where we will keep all environment related constants, like database username, password, secret key of jwt if any, etc

3. Create the "main.go" file in root, this will serve as the project/web-server entry-point

4. Create a "Makefile" this will contain some basic commands to run the project

5. Create these folders:
   1. Controllers: controllers same names as handler. Contains controller that accepts a request and call a particular service to process it
   2. Services: Contains service that has the logic for doing all the process like manipulating DB and giving back the desired result
   3. models: contrains the struct for storing data in DB or fetching data from DB
   4. DB: keep your DB related operations here, these can be used in services to fetch/update/insert/delete data from DB
      \*\*\* place DB Connection file as well inside this folder
   5. Routes: keep all your routes here and pass the name of controller
   6. config: keep all your configuration things here like, fetching variables from .env file or even db config can be placed here
   7. Constants: keep all constants here, can also categories them in separates files
   8. Utils: Keep commonly used function here like capitalising first letter, etc

so here we ara ready with one of the best folder structure for a web-server
so many folder structure. but i use this for small project
