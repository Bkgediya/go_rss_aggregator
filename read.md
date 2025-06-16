DATABASE_URL=postgresql://postgres:1234@localhost:5432/rssaggregator

| Part          | Example         | Meaning                                                      |
| ------------- | --------------- | ------------------------------------------------------------ |
| Protocol      | `postgresql://` | The type of database. (PostgreSQL in this case)              |
| Username      | `postgres`      | The database user you are logging in as.                     |
| Password      | `1234`          | The user’s password.                                         |
| Host          | `localhost`     | Where the database is located (here, on your local machine). |
| Port          | `5432`          | The PostgreSQL default port number.                          |
| Database Name | `rssaggregator` | The specific database you want to connect to.                |


