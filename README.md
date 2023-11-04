# bankApi🐿
This is a great json server that runs easily with docker compose and you can try or develop it further.🐳<br/>

## How to use

1- Make sure you installed docker and docker-compose in your machine.<br/>
2- after clone in root project run:<br/>
```bash
  sudo docker compose up -d
```
You should see this message in your terminal:<br/>
![image](https://github.com/pooulad/bankApi/assets/86445458/0f1312f1-3701-4f8a-8d66-9c79db7185c4)

After compose your database created based on your config.json in project and docker-compose.yml file(remember you can change or customize it).<br/>


3- You can check with this commands below.

```bash
  sudo docker exec -it NAME_OF_YOUR_CONTAINER psql -U postgres
```
Now you connect to postgres user in psql
Run:
```sql
  \connect YOUR_DB_NAME
```
For example:<br/>
![image](https://github.com/pooulad/bankApi/assets/86445458/e78719cb-56c2-45e8-a6fa-99add57deb50)

4- After that for making sure that project works and connected to db we should create a new user:

## Run app

```bash
  make run
```
Or 
```bash
  go run ./cmd/app/main.go
```
## Create new user

Make a post request:
![image](https://github.com/pooulad/bankApi/assets/86445458/b307edc6-b302-4ecb-a423-b466a4172a44)

```json
  {
    "firstName" : "mahdi",
    "lastName" : "dfdf",
    "password" : "1234"
  }
```
## Check database now

In psql run:

```sql
  \connect YOUR_DB_NAME
  SELECT * FROM account;
```
And you should see this response in your terminal:<br/>

![Screenshot from 2023-11-05 00-41-02](https://github.com/pooulad/bankApi/assets/86445458/499c013f-9a16-4b0e-aeef-64d4b56cdf02)


#### All endpoints

```http
  GET /account -> get all accounts
  POST /account -> create new account
  GET /account/:id -> get single account
  DELETE /account/:id -> delete account
  POST /login -> login
```

## Create account json data : 

```json
  {
    "firstName" : "mahdi",
    "lastName" : "dfdf",
    "password" : "1234"
  }
```


## Login json data : 

```json
  {
    "number" : BANK_NUMBER,
    "password" : "1234"
  }
```

## Get single account data : 

Parameter : id -> account id in endpoint : /account/[YOUR_BANK_ID]<br/>
+
Set the JWT token that you recieved in login response to header with name of "x-jwt-token"<br/>

x-jwt-token : YOUR_TOKEN

