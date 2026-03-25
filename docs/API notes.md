#                                   QuickBite API Notes

## 1. First, what are we even building?

We are building a backend application.

A backend application receives requests like:
- "give me all restaurants"
- "create a new order"
- "give me order by id"

Then it does some work and sends back a response.

Example:
- client sends `GET /api/v1/restaurants`
- backend reads request
- backend gets restaurant data
- backend sends JSON back

So backend work is not only "write code".
It has small stages.
That is why we split the code.

---

## 2. Why we split the code

A backend request does not do only one thing.
A request usually needs:

- app startup
- route matching
- request reading
- business logic
- data access
- response writing

These are different kinds of work.
So we do **not** put everything in one file.

We split code into layers.

The goal is not "more files".
The goal is:

- easier to understand
- easier to change
- easier to debug
- easier to grow later

### Simple memory line

```md
main starts
router maps
handler reads
service works
repository stores
response writes
```

That one line explains the full idea.

---

## 3. The most important understanding first

There are **two different things** happening when we build features:

### A. Writing code
This means creating functions and structs.

Examples:
- writing a handler function
- writing a service function
- writing a repository function
- writing a model struct

### B. Connecting code
This means making one part use another part.

Examples:
- route connected to handler
- handler calling service
- service calling repository

### Very important rule

Only **routes** are truly registered.

Example:

```go
mux.HandleFunc("/api/v1/orders", handler.CreateOrder)
```

That is registration.

But this is **not** registration:

```go
order, validationError := service.CreateOrder(req)
```

This is just a normal function call.

### Remember this line

```md
routes are registered
handlers/services/repositories are written and called
```

---

## 4. What exactly is `main.go`?

### Simple answer

`main.go` is the starting point of the whole application.

When you run the Go app, Go looks for:
- `package main`
- `func main()`

That is where execution begins.

### Why do we need it?

Because the operating system needs one place to start the program.
Without `main.go`, your app does not know where to begin.

### What should `main.go` do?

Only startup work.

Usually:
- create the router
- start the HTTP server
- maybe print startup message

### What should `main.go` NOT do?

It should not:
- validate order fields
- read JSON request body
- search restaurants
- write business logic

That work belongs in other layers.

### Example

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/akhileshkasarapu3/quickbite/internal/router"
)

func main() {
    appRouter := router.RegisterRoutes()

    fmt.Println("Server is running on port 8080")

    err := http.ListenAndServe(":8080", appRouter)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}
```

### What this means in plain English

- ask router package to build the app routes
- start HTTP server
- server will now listen for requests

### Tiny memory line

```md
main.go = app starting place
```

---

## 5. What exactly is `router`?

### Simple answer

Router decides:

**which URL path should go to which handler**

Example:
- `/api/v1/orders` should go to order handler
- `/api/v1/restaurants` should go to restaurant handler

### Why do we need router?

Because many requests can come into the app.
The app must know where to send each one.

Without a router, every request would have no clear path.

### Example thinking

If request comes for:

```text
/api/v1/restaurants
```

router should say:

"Send this to `handler.GetRestaurants`"

---

## 6. What exactly is `mux`?

### Simple answer

`mux` means **multiplexer**.

In Go, a `ServeMux` is a built-in router.
It receives incoming request paths and matches them to handler functions.

### Example

```go
mux := http.NewServeMux()
```

This means:

"Create a new router object for my application."

### Example route registration

```go
mux.HandleFunc("/api/v1/health", handler.HealthHandler)
```

This means:

"If request path is `/api/v1/health`, call `handler.HealthHandler`."

### Tiny memory line

```md
request -> mux checks path -> matching handler runs
```

---

## 7. What does route registration mean?

### Simple answer

Route registration means connecting:

**URL path -> handler function**

### Example

```go
mux.HandleFunc("/api/v1/restaurants/open", handler.GetOpenRestaurants)
```

This means:
- path is `/api/v1/restaurants/open`
- handler is `handler.GetOpenRestaurants`

### Why do this?

Because when client sends request to that path, router must know which function to run.

### Very important understanding

This is the place where route and handler become connected.

Before this line exists, the handler function may be written, but no request can reach it.

### Memory line

```md
writing a handler creates it
registering a route makes it reachable
```

---

## 8. What exactly is a handler?

### Simple answer

Handler is the layer that deals with the HTTP request and HTTP response.

Handler reads request details and then calls service.

### What does handler usually read?

Handler may read:
- HTTP method
- query parameters
- request body JSON
- path values
- headers

### What does handler usually do?

- check method
- decode request body if needed
- read query params if needed
- call service
- send response using response helper

### What should handler NOT do too much of?

It should not contain heavy business rules.
That belongs in service.

### Example

```go
var req model.CreateOrderRequest
err := json.NewDecoder(r.Body).Decode(&req)
if err != nil {
    response.WriteError(w, http.StatusBadRequest, "invalid request body")
    return
}

order, validationError := service.CreateOrder(req)
if validationError != "" {
    response.WriteError(w, http.StatusBadRequest, validationError)
    return
}

response.WriteSuccess(w, http.StatusCreated, order)
```

### Plain English meaning

- read JSON body from request
- convert it into Go struct
- ask service to do real work
- send result back

### Tiny memory line

```md
handler = HTTP layer
```

---

## 9. What exactly is service?

### Simple answer

Service contains the **business logic** of the application.

### What does business logic mean?

It means the app rules.

Examples:
- customer name is required
- restaurant id is required
- item list cannot be empty
- only open restaurants should be returned
- cuisine filtering should be case-insensitive

These are application rules.

### Why should handler not do this work?

Because handler should stay focused on HTTP work.

If handler starts doing:
- validation
- filtering
- sorting
- calculations
- core domain decisions

then handler becomes too large and messy.

Service keeps the business rules separate.

### Example

```go
if customerName == "" {
    return model.Order{}, "customer name is required"
}

if len(req.Items) == 0 {
    return model.Order{}, "at least one item is required"
}
```

This belongs in service because it is an app rule.

### Another example

```go
if restaurant.IsOpen {
    openRestaurants = append(openRestaurants, restaurant)
}
```

Filtering only open restaurants is business behavior.
So it belongs in service.

### Tiny memory line

```md
service = app brain / business rules
```

---

## 10. What exactly is repository?

### Simple answer

Repository handles **data storage and data retrieval**.

### What does that mean?

Repository does work like:
- save order
- get all orders
- get order by id
- get all restaurants
- get restaurant by id

### Why is repository different from service?

Because service decides **what should happen**.
Repository decides **where and how data is stored or fetched**.

### Example difference

#### Service thinking
"I must create an order if the request is valid."

#### Repository thinking
"I will store this order in memory / DB / somewhere else."

### Storage logic means
Any logic related to:
- saving data
- reading data
- searching stored data
- returning stored records

### Example

```go
orders = append(orders, order)
```

This is storage logic.
So it belongs in repository.

Another example:

```go
for _, order := range orders {
    if order.ID == id {
        return order, true
    }
}
```

This is storage lookup logic.
So it belongs in repository.

### Tiny memory line

```md
repository = storage layer
```

---

## 11. What exactly is model?

### Simple answer

Model holds shared data structures used across multiple layers.

### Why do we need model?

Because types like:
- `Order`
- `Restaurant`
- `CreateOrderRequest`

are not only for one layer.
They are used by handler, service, and repository.

So they should live in one neutral place.

### Example

```go
type Order struct {
    ID              string   `json:"id"`
    CustomerName    string   `json:"customer_name"`
    RestaurantID    string   `json:"restaurant_id"`
    Items           []string `json:"items"`
    DeliveryAddress string   `json:"delivery_address"`
    Status          string   `json:"status"`
}
```

### Why not keep `Order` inside service?

Because repository also needs it.
Handler also sends it in JSON response.
So it should not belong only to service.

### Tiny memory line

```md
model = shared data shape
```

---

## 12. What exactly is response?

### Simple answer

Response package helps us send JSON back in a clean and reusable way.

### Why do we need it?

Because otherwise every handler would repeat:
- set content type
- set status code
- encode JSON
- send error format

That causes duplication.

### Example

```go
response.WriteSuccess(w, http.StatusOK, data)
response.WriteError(w, http.StatusBadRequest, "invalid request")
```

### Why keep JSON writing there?

So all responses look consistent.
This makes handlers smaller and cleaner.

### Tiny memory line

```md
response = JSON output helper
```

---

## 13. What gets registered and what gets called?

This is one of the most important confusing points.
So let’s make it very clear.

### Registered
Only routes are registered.

Example:

```go
mux.HandleFunc("/api/v1/orders", handler.CreateOrder)
```

This is route registration.

### Called
Handlers, services, and repositories are not registered here.
They are just normal functions that are called.

Example handler calling service:

```go
order, validationError := service.CreateOrder(req)
```

Example service calling repository:

```go
savedOrder := repository.SaveOrder(order)
```

### Final memory line

```md
route is registered to handler
handler calls service
service calls repository
```

---

## 14. In what order should you think when creating a new feature?

Use this exact thinking order.

```md
feature idea
-> model
-> repository
-> service
-> handler
-> route registration
-> test
```

Now let’s explain each one slowly.

### Step 1: Feature idea
Ask:

**What endpoint do I want?**

Example:
- `GET /api/v1/restaurants/open`
- `POST /api/v1/orders`

### Step 2: Model
Ask:

**What data shape do I need?**

Example:
- `Restaurant`
- `Order`
- `CreateOrderRequest`

### Step 3: Repository
Ask:

**What data access is needed?**

Example:
- get all restaurants
- save order
- get order by id

### Step 4: Service
Ask:

**What business logic is needed?**

Example:
- filter only open restaurants
- validate order request
- build order object

### Step 5: Handler
Ask:

**How do I read the HTTP request and send the HTTP response?**

Example:
- read query param
- decode JSON body
- call service
- return success/error

### Step 6: Route registration
Ask:

**Which URL should call which handler?**

Example:

```go
mux.HandleFunc("/api/v1/restaurants/open", handler.GetOpenRestaurants)
```

### Step 7: Test
Use `curl` or browser and check if the endpoint works.

---

## 15. Full request flow: `POST /api/v1/orders`

Now let’s explain one complete request in a spoon-feeding way.

### Client request

Client sends:

```http
POST /api/v1/orders
```

with JSON body like:

```json
{
  "customer_name": "Akhilesh",
  "restaurant_id": "res_101",
  "items": ["Mutton Biryani"],
  "delivery_address": "1417 Sage Way, Aubrey, TX"
}
```

---

### Step 1: Router receives the path

Router sees path:

```text
/api/v1/orders
```

and route registration says:

```go
mux.HandleFunc("/api/v1/orders", handler.CreateOrder)
```

So router sends request to:

```go
handler.CreateOrder
```

---

### Step 2: Handler starts

Handler receives:
- `w` = response writer
- `r` = request

It checks method:

```go
if r.Method != http.MethodPost {
    response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
    return
}
```

Why?
Because this endpoint should only allow POST.

---

### Step 3: Handler reads request body

Handler does:

```go
var req model.CreateOrderRequest
err := json.NewDecoder(r.Body).Decode(&req)
```

This means:
- read JSON from request body
- convert it into Go struct

Now this JSON:

```json
{
  "customer_name": "Akhilesh"
}
```

becomes Go value inside `req`.

---

### Step 4: Handler calls service

Handler does:

```go
order, validationError := service.CreateOrder(req)
```

Meaning:
"Service, please do the real order creation work."

Handler does not create the order fully by itself.
That is service’s job.

---

### Step 5: Service validates request

Service checks business rules like:

```go
if customerName == "" {
    return model.Order{}, "customer name is required"
}

if len(req.Items) == 0 {
    return model.Order{}, "at least one item is required"
}
```

Why here?
Because these are application rules.
Not HTTP rules.

---

### Step 6: Service builds order object

Service creates something like:

```go
order := model.Order{
    CustomerName:    customerName,
    RestaurantID:    restaurantID,
    Items:           req.Items,
    DeliveryAddress: deliveryAddress,
    Status:          "placed",
}
```

Meaning:
service creates a clean business object.

---

### Step 7: Service calls repository

Service does:

```go
savedOrder := repository.SaveOrder(order)
```

Meaning:
"Repository, please store this order."

---

### Step 8: Repository stores order

Repository may do:

```go
orders = append(orders, order)
```

and maybe generate ID too.

This is storage work.
So it belongs in repository.

---

### Step 9: Repository returns stored order

Repository returns saved order back to service.

Service returns it back to handler.

---

### Step 10: Handler sends response

Handler does:

```go
response.WriteSuccess(w, http.StatusCreated, order)
```

Meaning:
- status code = 201 Created
- response body = created order as JSON

---

### Final memory line for POST order

```md
router matches path
handler reads JSON
service validates and creates
repository stores
handler sends response
```

---

## 16. Full request flow: `GET /api/v1/restaurant?id=res_101`

Now let’s do a read request in the same slow way.

### Client request

```http
GET /api/v1/restaurant?id=res_101
```

---

### Step 1: Router matches path

Router sees:

```text
/api/v1/restaurant
```

and route says:

```go
mux.HandleFunc("/api/v1/restaurant", handler.GetRestaurantByID)
```

So router calls:

```go
handler.GetRestaurantByID
```

---

### Step 2: Handler starts

Handler checks method:

```go
if r.Method != http.MethodGet {
    response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
    return
}
```

Why?
Because this endpoint should only allow GET.

---

### Step 3: Handler reads query parameter

Handler does:

```go
restaurantID := r.URL.Query().Get("id")
```

If request is:

```text
/api/v1/restaurant?id=res_101
```

then `restaurantID` becomes:

```text
res_101
```

---

### Step 4: Handler validates input presence

Handler checks:

```go
if restaurantID == "" {
    response.WriteError(w, http.StatusBadRequest, "restaurant id is required")
    return
}
```

Why in handler?
Because reading query param is HTTP-level work.

---

### Step 5: Handler calls service

Handler does:

```go
restaurant, found := service.GetRestaurantByID(restaurantID)
```

Meaning:
"Service, please give me restaurant for this ID."

---

### Step 6: Service calls repository

Service may do:

```go
return repository.GetRestaurantByID(id)
```

Right now maybe no heavy business logic is needed.
That is okay.

Service still acts as the boundary.

---

### Step 7: Repository searches data

Repository may do:

```go
for _, restaurant := range restaurants {
    if restaurant.ID == id {
        return restaurant, true
    }
}
```

This is storage search logic.
So it belongs in repository.

---

### Step 8: Result comes back

Repository returns result to service.
Service returns it to handler.

---

### Step 9: Handler sends JSON response

If found:

```go
response.WriteSuccess(w, http.StatusOK, restaurant)
```

If not found:

```go
response.WriteError(w, http.StatusNotFound, "restaurant not found")
```

---

### Final memory line for GET restaurant

```md
router matches path
handler reads query param
service asks repository
repository finds data
handler sends response
```

---

## 17. `GetOrders` from `internal/handler/order.go` line by line

Now let’s explain this slowly because this is a very good handler example.

### Example handler

```go
func GetOrders(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
        return
    }

    orders := service.GetOrders()
    response.WriteSuccess(w, http.StatusOK, orders)
}
```

---

### Line 1

```go
func GetOrders(w http.ResponseWriter, r *http.Request) {
```

This defines the handler function.

- `w` is used to send response back
- `r` contains request details

This is in handler because it deals with HTTP request and response.

---

### Line 2 to 5

```go
if r.Method != http.MethodGet {
    response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
    return
}
```

This checks if request method is GET.

Why here?
Because checking HTTP method is handler work.
That is part of HTTP layer.

Why not service?
Because service should not care whether request came as GET or POST.
That is HTTP concern.

---

### Line 7

```go
orders := service.GetOrders()
```

This asks service layer to give all orders.

Why call service?
Because handler should not directly read storage or repository in this architecture.
Handler should go through service.

Why not directly repository?
Because service is the business boundary.
Even if today it only forwards the call, tomorrow service may add:
- filtering
- sorting
- permission checks
- transformations

So handler should stay dependent on service.

---

### Line 8

```go
response.WriteSuccess(w, http.StatusOK, orders)
```

This sends success response as JSON.

Why use response helper?
Because we want response format to stay consistent.
And we do not want every handler to manually repeat JSON writing code.

---

### Final meaning of `GetOrders`

In plain English:

- accept only GET
- ask service for order list
- return order list as JSON

### Memory line

```md
GetOrders handler does HTTP work, not storage work
```

---

## 18. Why service exists even when it looks small

Sometimes service looks like only this:

```go
func GetOrders() []model.Order {
    return repository.GetOrders()
}
```

You may think:
"Why do we need service if it just forwards the call?"

Good question.

### Answer

Because service is the business boundary.
Today it may be small.
Tomorrow it may grow.

Later service may add:
- filtering
- sorting
- validation
- permission checks
- combining multiple repository calls

So keeping the service layer now makes the structure consistent and future-friendly.

---

## 19. Why repository exists even for in-memory data

You may think:
"If data is just in a slice, why do I need repository?"

Because repository is not only for databases.
Repository is for **data access responsibility**.

Today:
- repository uses slice

Tomorrow:
- repository may use PostgreSQL
- MongoDB
- Redis

If storage changes, repository is the layer that should change the most.
That is the point.

---

## 20. Mutex and RWMutex in very simple language

When we store data in memory like this:

```go
var orders = []model.Order{}
var nextOrderNumber = 1001
```

many requests may try to read/write this data at the same time.

Because Go HTTP handlers can run concurrently.

That can create bugs.

### What is `Mutex`?

Think of one bathroom key.
Only one person can use the bathroom at a time.
Others must wait.

That is how `Mutex` works.

Example idea:
- one request enters critical section
- other request waits
- first request finishes
- lock opens
- next request enters

### What is `RWMutex`?

Think of a library.

- many people can sit and read books at the same time
- but if someone is rearranging or rewriting the books, others should wait

That is `RWMutex`.

### Meaning

- `RLock` = read lock, many readers allowed
- `Lock` = write lock, only one writer allowed

### Why useful?

Because apps usually have:
- many reads
- fewer writes

So `RWMutex` allows better concurrency for reads.

### Very short memory line

```md
Mutex = one at a time
RWMutex = many readers, one writer
```

---

## 21. Why we copied slices in repository

Example:

```go
result := make([]model.Order, len(orders))
copy(result, orders)
return result
```

### Why do this?

Because we do not want outside code to directly hold repository-owned slice memory.

Returning a copy is safer.

### Simple idea

Original slice belongs to repository.
Other layers should receive a safer copy.

---

## 22. Simple way to decide where code belongs

Use these simple rules.

### Put code in router if it is about:
- path registration
- route mapping

Example:

```go
mux.HandleFunc("/api/v1/orders", handler.CreateOrder)
```

### Put code in handler if it is about:
- HTTP method
- query params
- request body
- response status

Examples:

```go
if r.Method != http.MethodPost
```

```go
orderID := r.URL.Query().Get("id")
```

```go
json.NewDecoder(r.Body).Decode(&req)
```

### Put code in service if it is about:
- validation
- filtering
- sorting
- business rules
- building domain objects

Examples:

```go
if len(req.Items) == 0
```

```go
if restaurant.IsOpen
```

### Put code in repository if it is about:
- saving
- reading
- searching stored data
- storage mutation

Example:

```go
orders = append(orders, order)
```

### Put code in model if it is about:
- shared struct definitions

Examples:

```go
type Order struct { ... }
type Restaurant struct { ... }
```

### Put code in response if it is about:
- JSON formatting
- reusable success/error output

Examples:

```go
response.WriteSuccess(...)
response.WriteError(...)
```

---

## 23. One reusable feature template

Whenever you build a new feature, think like this:

```md
1. What is the endpoint?
2. What data shape do I need?
3. What repository function do I need?
4. What business logic should service do?
5. What should handler read from request?
6. What route should point to handler?
7. How will I test it?
```

### Example for open restaurants

```md
endpoint -> /api/v1/restaurants/open
model -> Restaurant
repository -> GetRestaurants()
service -> GetOpenRestaurants()
handler -> GetOpenRestaurants(w, r)
route -> mux.HandleFunc("/api/v1/restaurants/open", handler.GetOpenRestaurants)
```

---

## 24. Short revision summary

```md
main.go:
starts the app

router:
maps URL paths to handlers

handler:
reads HTTP request and calls service

service:
contains business logic

repository:
handles storage and retrieval

model:
defines shared data structures

response:
writes standard JSON responses
```

---

## 25. Final memory lines

```md
main starts
router maps
handler reads
service works
repository stores
response writes
```

```md
routes are registered
handlers/services/repositories are written and called
```

```md
repository gets data
service applies business meaning
handler exposes it over HTTP
```

```md
feature idea
-> model
-> repository
-> service
-> handler
-> route registration
-> test
```

---

# QuickBite Beginner Cheat Sheet

## Build order for a new feature

```md
What is the endpoint?
What data shape do I need?
What storage access do I need?
What business logic do I need?
What handler should expose it?
What route should point to it?
How will I test it?
```

## Layer cheat sheet

```md
Model = shape
Repository = storage
Service = business logic
Handler = HTTP
Router = path mapping
Response = JSON formatting
```

## Registration cheat sheet

```md
Only routes are registered.
Handlers, services, and repositories are normal functions that are called.
```

## Concurrency cheat sheet

```md
Shared writable memory needs locking.
Mutex = one goroutine at a time.
RWMutex = many readers, one writer.
```

## If you are confused, ask these questions first

```md
1. Am I starting the app? -> main.go
2. Am I connecting a path to a handler? -> router
3. Am I reading HTTP request details? -> handler
4. Am I applying app rules or validation? -> service
5. Am I saving or fetching stored data? -> repository
6. Am I defining shared data shape? -> model
7. Am I formatting JSON response? -> response
```

## Most important architecture rule

```md
Do not mix HTTP logic, business logic, and storage logic in one place.
Keep each layer focused on one responsibility.
```