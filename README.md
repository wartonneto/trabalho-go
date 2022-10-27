How to Run:

first you need to go to "configs.toml" and fill in with your configs.

after this, you need to do a "docker compose up" command so you'll get the docker image configurated in your pc.

after that you'll need the migrations, just run "docker run -v {your_entire_path_to_project}\migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "postgres://127.0.0.1/api_biblioteca?sslmode=disable&user=postgres&password=postgres" up 1"

Then you can run the project with "go run.\main.go"

the default routes is 

POST http://localhost:9000/
GET http://localhost:9000/
GET http://localhost:9000/{id}
PUT http://localhost:9000/{id}
DELETE http://localhost:9000/{id}


example of raw POST:
    
    {
        "title":"Clean Code",
        "description":"Aprenda a programar! <3",
        "value": 150
    }

# Situação de exceção
![Alt text](./exceptions/entrypoint.jpg?raw=true "Exception png")