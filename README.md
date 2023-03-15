# e-com-store

## Quick Start
```
1. open terminal
2. create a new directory
3. git clone https://github.com/shashikant231/e-com-store
4. create a new db credentials and change dsn in main.go file
5. run the server with "go run .\app\main.go" command
```


## Major urls
```
e.GET("/sync", handler.Sync) // sync categories and product data
e.GET("/shop/categories", handler.GetCategories)  
e.GET("/shop/products", handler.GetProducts)