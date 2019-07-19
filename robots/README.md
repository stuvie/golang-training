# GO Robots

A fun project to re-write a node.js application in Golang

## Setup Project Dependencies

go get -v github.com/gin-gonic/gin github.com/jinzhu/gorm github.com/go-sql-driver/mysql github.com/stretchr/testify/assert

## Testing

go test

## Run

go run robots.go
curl http://localhost:3000/ping

# Insert some test data

insert into todo (title, url, completed) values ('Vegan Lentil Taco Salad Bowls', 'http://thefitchen.com/vegan-lentil-taco-salad-bowls-2/', 1);
insert into todo (title, url, completed) values ('Bbq chicken shawarma', 'https://www.feastingathome.com/grilled-chicken-shawarma-recipe/', 5);
insert into todo (title, url, completed) values ('Amazing chicken tacos', 'http://www.onceuponachef.com/2011/02/chicken-tacos.html"', 5);


# Todo

connect to robots MySQL database
