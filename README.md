# Vote service

- Written in Go
- Stores data in PostgreSQL with gorm

# Quickstart
```shell
git clone https://github.com/cenkayla/poll.git
cd poll
go run main.go
```
# Usage

To create a poll with multiple choices
```shell
curl --request POST \
	--url "localhost:8080/api/createpoll" \
	--header "Content-Type: application/json" \
	--data "{\"id\":1,\"name\":\"What's your favorite game?\",\"choice\":[{\"name\":\"League of Legends\"},{\"name\":\"Dota2\"}]}"
```
Response
```shell
"Succesfully created."
```

To vote for a specific choice
```shell
curl --request POST \
	--url "localhost:8080/api/poll" \
	--header "Content-Type: application/json" \
	--data "{\"id\":1,\"name\":\"Dota2\"}"
```
Response
```shell
"Successfully voted to Dota2"
```

To get the result for a specific vote
```shell
curl --request POST \
	--url "localhost:8080/api/getresult" \
	--header "Content-Type: application/json" \
	--data "{\"id\":1}"
```
Response
```shell
{
  "ID": 1,
  "name": "What's your favorite game?",
  "choice": [
    {
      "id": 1,
      "name": "League of Legends",
      "votes": 0,
      "CreatedAt": "2021-07-15T02:22:45.96279+03:00",
      "UpdatedAt": "2021-07-15T02:22:45.96279+03:00",
      "DeletedAt": null
    },
    {
      "id": 1,
      "name": "Dota2",
      "votes": 1,
      "CreatedAt": "2021-07-15T02:22:45.96279+03:00",
      "UpdatedAt": "2021-07-15T02:27:41.6796+03:00",
      "DeletedAt": null
    }
  ],
  "CreatedAt": "2021-07-15T02:22:45.961113+03:00",
  "UpdatedAt": "2021-07-15T02:22:45.961113+03:00",
  "DeletedAt": null
}
```
