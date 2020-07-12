# Url_shortener in Golang
=============================

## Steps to run

- Run 
```
sql.connection.go
``` 
This includes connection to database and tables.

- Then run
```
shortner.go
```
This file includes functionalities to get 
the short url.

- Finally, run
```
Data/data.go
data.go
```
Here, we can store our Long url and short url
and retrieve them accordingly.

- Check db
```
select * from url_shortener;
+----+-----------------------------+--------------------------------------------------------------------------------------+
| id | slug                        | url                                                                                  |
+----+-----------------------------+--------------------------------------------------------------------------------------+
|  1 | http://tinyurl.com/ybms2dme | https://www.geeksforgeeks.org/golang-tutorial-learn-go-programming-language/?ref=lbp |
+----+-----------------------------+--------------------------------------------------------------------------------------+
```

