# Part 2: Refactor HTTP Handlers

This part of your assignment calls for refactoring your existing product server’s HTTP handlers so that they use the products database you seeded in Part 1 above.

## Acceptance Criteria

1. Your server is initialized with a connection to the `products` database using the `database/sql` package and a suitable SQLite driver.
2. Your server verifies its connection to the database with a Ping (see [https://pkg.go.dev/database/sql@go1.19.3#DB.Ping](https://pkg.go.dev/database/sql@go1.19.3#DB.Ping)) and stops initialization cleanly (i.e. no `panic`) if the connection is invalid.
3. Your server supports the following operations
    1. `GET /products?limit=<numeric>&sort=<column>`
    2. `GET /products/<id>`
    3. `POST /products` - with a request body containing a `product` JSON payload
    4. `PUT /products/<id>` - with a request body containing a `product` JSON payload
    5. `DELETE /products/<id>`
4. You have a `getProducts` handler that retrieves a list of products from the DB to return as a JSON list to the client, limiting the number of results based on the `limit` parameter and sorting the results based on the `sort` parameter.
    1. Unknown sort fields should be ignored to prevent your SQL statements from causing an error.
    2. You must use parameterized prepared statements.
5. You have a `getProduct` handler that retrieves a single product from the database based on the passed in `id`.
    1. Return a `404 Not Found` to the client for unknown product IDs
    2. Return a `400 Bad Request` if a non-numeric ID is passed in
    3. Returns a `200 OK` and a JSON representation of the `product` back to the client.
6. You have an `addProduct` handler that inserts a new product in the DB.
    1. Return a `400 Bad Request` if any required fields are missing in the request body, including `name` and `price`.
    2. Return a `400 Bad Request` if a non-numeric `price` is passed in.
    3. Returns a `201 Created` and no response body upon success.
7. You have an `updateProduct` handler that updates an existing product in the DB.
    1. Return a `404 Not Found` to the client for unknown product IDs
    2. Return a `400 Bad Request` if a non-numeric ID is passed in
    3. Return a `400 Bad Request` if any required fields are missing in the request body, including `name` and `price`.
    4. Returns a `200 OK` and no response body upon success.
8. You have a `deleteProduct` handler that removes an existing product from the DB.
    1. Return a `404 Not Found` to the client for unknown product IDs
    2. Return a `400 Bad Request` if a non-numeric ID is passed in
    3. Returns a `200 OK` and no response body upon success.
9. Your work is done on a separate branch (e.g. `refactor-product-server-handlers`) and you provide a link to a GitHub pull request (PR) against your repository’s `main` branch with your solution for evaluation.
10. Your PR’s comment includes a copy of the output of your program for each endpoint you call via a client such as `curl`. You must show the `curl` calls and the server’s response underneath.

# Resources

1. [SQLite tutorial](https://www.youtube.com/watch?v=zLQ03DeH04c&list=PL-1QdJ8od_eyxntzYQhwCkcVZlqWVrmSf&index=1)
2. [SQLite data types](https://www.sqlite.org/datatype3.html)
3. Go [database/sql](https://pkg.go.dev/database/sql@go1.19.3) package documentation
4. SQLite driver usage [examples](https://github.com/mattn/go-sqlite3/blob/master/_example/simple/simple.go) in Go with the https://github.com/mattn/go-sqlite3 package

# Keep In Mind

1. Directory names should be lowercased, avoid MixCasedDirectoryNames.
2. Avoid directory bloat. No need for separate `handlers` , `objects` , `server` , `models` , etc folders. Keep it simple.
3. Handle your errors. Do not discard them with the blank identifier ( `_` ).
4. Do not crash your own server with `log.Fatal` within your handlers.


# Outputs
```
localhost:8080/products?limit=3&sort=price
[{"id":72,"name":"Pepper - Green Thai","description":"nisi vulputate nonummy maecenas tincidunt lacus at velit vivamus vel nulla","price":1},{"id":44,"name":"7up Diet, 355 Ml","description":"tincidunt ante vel ipsum praesent blandit lacinia erat vestibulum sed magna at nunc commodo","price":2},{"id":7,"name":"Doilies - 5, Paper","description":"ut massa volutpat convallis morbi odio odio elementum eu interdum eu tincidunt in leo maecenas","price":3}]

localhost:8080/products?limit=3&sort=value (unknown sort fields result in defaulting to sort by ID)
[{"id":1,"name":"Water - San Pellegrino","description":"curae nulla dapibus dolor vel est donec odio justo sollicitudin ut suscipit a feugiat et eros vestibulum ac est lacinia","price":80},{"id":2,"name":"Cape Capensis - Fillet","description":"rutrum neque aenean auctor gravida sem praesent id massa id nisl venenatis lacinia aenean sit amet justo morbi ut odio","price":53}]

localhost:8080/products?limit=-1&sort=price (negative or missing limiters default to listing 20 items)
[{"id":72,"name":"Pepper - Green Thai","description":"nisi vulputate nonummy maecenas tincidunt lacus at velit vivamus vel nulla","price":1},{"id":44,"name":"7up Diet, 355 Ml","description":"tincidunt ante vel ipsum praesent blandit lacinia erat vestibulum sed magna at nunc commodo","price":2}....

localhost:8080/products?limit=150&sort=price (limits over 100 items default to listing 20 items)
[{"id":72,"name":"Pepper - Green Thai","description":"nisi vulputate nonummy maecenas tincidunt lacus at velit vivamus vel nulla","price":1},{"id":44,"name":"7up Diet, 355 Ml","description":"tincidunt ante vel ipsum praesent blandit lacinia erat vestibulum sed magna at nunc commodo","price":2}....

localhost:8080/products/1000 (unknown product ID's result in a 404 status)
Status: 404 Not Found

localhost:8080/products/Bread (Non-numeric ID's result in a 400 status)
Status: 400 Bad Request

localhost:8080/products/5 (Successful arguments result in the item being sent back to the user)
{"id":5,"name":"Lettuce - Arugula","description":"vel nisl duis ac nibh fusce lacus purus aliquet at feugiat non pretium","price":11}

localhost:8080/products (Missing data in the JSON packet results in a 400 status)
POST: {"name":"Really Bad Product", "description":"I've updated this to see if it will be updated in the DB", "price":10}
Status: 400 Bad Request

localhost:8080/products (Posting to an already existing ID results in a 400 status)
POST: {"id": 1,"name":"Really Bad Product", "description":"I've updated this to see if it will be updated in the DB", "price":10}
Status: 400 Bad Request

localhost:8080/products (Meeting requirements results in a 201 status)
POST: {"id": 101,"name":"Really Bad Product", "description":"I've updated this to see if it will be updated in the DB", "price":10}
Status: 201 Created

localhost:8080/products/104 (Non-Existant ID's cannot be updated and result in a 404 status)
PUT: {"id": 104, "name":"Really GREAT Product", "description":"This product was bad but is now good", "price":10}
Status: 404 Not Found

localhost:8080/products/101 (ID's that don't cannot be updated and result in a 400 status)
PUT: {"id": 104, "name":"Really GREAT Product", "description":"This product was bad but is now good", "price":10}
Status: 400 Bad Request

localhost:8080/products/101 (Requests that are missing data (ID, Name, Price) result in a 400 status)
PUT: {"id": 104, "description":"This product was bad but is now good", "price":10}
Status: 400 Bad Request

localhost:8080/products/101 (Meeting requirements results in a 200 status)
PUT: {"id": 101, "name":"Really GREAT Product", "description":"This product was bad but is now good", "price":10}
Status: 200 OK

DELETE: localhost:8080/products/102 (Non-Existant ID's cannot be deleted and result in a 404 status)
Status: 404 Not Found

DELETE: localhost:8080/products/Bread (Non-numeric ID's result in a 400 status)
Status: 400 Bad Request

Delete: localhost:8080/products/101 (Meeting requirements results in a 200 status)
Status: 200 OK
```