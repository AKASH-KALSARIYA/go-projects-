# 🍽️ Restaurant Management API

Simple backend to manage restaurant entities: users, menus, foods, tables, orders, order items, and invoices. Tested to run locally with MongoDB.

- Server URL: `http://localhost:1111`
- Database: MongoDB `restaurant`

---

## 👤 Users

### 1) Create (Signup)
- Method: `POST`
- URL: `http://localhost:1111/users/signup`
- Body:
```json
{
  "first_name": "John",
  "last_name": "Doe",
  "email": "john@example.com",
  "password": "secret123",
  "phone": "+1234567890",
  "avatar": "https://example.com/a.jpg"
}
```

### 2) Login
- Method: `POST`
- URL: `http://localhost:1111/users/login`
- Body:
```json
{
  "email": "john@example.com",
  "password": "secret123"
}
```

### 3) List
- Method: `GET`
- URL: `http://localhost:1111/users`

### 4) Get by ID
- Method: `GET`
- URL: `http://localhost:1111/users/{user_id}`

---

## 🧾 Menu

### 1) Create
- Method: `POST`
- URL: `http://localhost:1111/menu`
- Body:
```json
{
  "name": "Breakfast",
  "category": "Morning",
  "start_date": "2026-03-10T06:00:00Z",
  "end_date": "2026-03-10T11:00:00Z"
}
```

### 2) Update
- Method: `PATCH`
- URL: `http://localhost:1111/menu/{menu_id}`
- Body:
```json
{
  "name": "Breakfast Specials",
  "category": "Morning",
  "start_date": "2026-03-10T07:00:00Z",
  "end_date": "2026-03-10T12:00:00Z"
}
```

### 3) List
- Method: `GET`
- URL: `http://localhost:1111/menu`

### 4) Get by ID
- Method: `GET`
- URL: `http://localhost:1111/menu/{menu_id}`

### 5) Delete by ID
- Method: `DELETE`
- URL: `http://localhost:1111/menu/{menu_id}`

### 6) Delete All
- Method: `DELETE`
- URL: `http://localhost:1111/menu`

---

## 🍲 Food

### 1) Create
- Method: `POST`
- URL: `http://localhost:1111/foods`
- Body:
```json
{
  "name": "Pancakes",
  "price": 5.99,
  "food_image": "https://example.com/p.jpg",
  "menu_id": "{menu_id}"
}
```

### 2) Update
- Method: `PATCH`
- URL: `http://localhost:1111/foods/{Food_id}`
- Body:
```json
{
  "name": "Blueberry Pancakes",
  "price": 6.49,
  "food_image": "https://example.com/b.jpg"
}
```

### 3) List (Paginated)
- Method: `GET`
- URL: `http://localhost:1111/foods?RecordPerPage=10&page=1`
- Optional: `startIndex` query param

### 4) Get by ID
- Method: `GET`
- URL: `http://localhost:1111/foods/{Food_id}`

### 5) Delete by ID
- Method: `DELETE`
- URL: `http://localhost:1111/foods/{Food_id}`

### 6) Delete All
- Method: `DELETE`
- URL: `http://localhost:1111/foods`

---

## 🪑 Table

### 1) Create
- Method: `POST`
- URL: `http://localhost:1111/Table`
- Body:
```json
{
  "number_of_guest": 4,
  "table_number": 12
}
```

### 2) Update
- Method: `PATCH`
- URL: `http://localhost:1111/Table/{Order_id}`
- Body:
```json
{
  "number_of_guest": 5,
  "table_number": 12
}
```

### 3) List
- Method: `GET`
- URL: `http://localhost:1111/Table`

### 4) Get by ID
- Method: `GET`
- URL: `http://localhost:1111/Table/{Order_id}`

### 5) Delete by ID
- Method: `DELETE`
- URL: `http://localhost:1111/Table/{Order_id}`

### 6) Delete All
- Method: `DELETE`
- URL: `http://localhost:1111/Table`

---

## 🧾 Order

### 1) Create
- Method: `POST`
- URL: `http://localhost:1111/Order`
- Body:
```json
{
  "order_date": "2026-03-10T08:30:00Z",
  "table_id": "{table_id}"
}
```

### 2) Update
- Method: `PATCH`
- URL: `http://localhost:1111/Order/{Order_id}`
- Body:
```json
{
  "order_date": "2026-03-10T09:00:00Z",
  "table_id": "{table_id}"
}
```

### 3) List
- Method: `GET`
- URL: `http://localhost:1111/Order`

### 4) Get by ID
- Method: `GET`
- URL: `http://localhost:1111/Order/{Order_id}`

### 5) Delete by ID
- Method: `DELETE`
- URL: `http://localhost:1111/Order/{Order_id}`

### 6) Delete All
- Method: `DELETE`
- URL: `http://localhost:1111/Order`

---

## 🧾 Order Items

### 1) Create
- Method: `POST`
- URL: `http://localhost:1111/OrderItem`
- Body:
```json
{
  "quantity": "M",
  "unit_price": 6.49,
  "food_id": "{Food_id}",
  "order_id": "{Order_id}"
}
```

### 2) Update
- Method: `PATCH`
- URL: `http://localhost:1111/OrderItem/{OrderItem_id}`
- Body:
```json
{
  "quantity": "L",
  "unit_price": 7.49,
  "food_id": "{Food_id}",
  "order_id": "{Order_id}"
}
```

### 3) List
- Method: `GET`
- URL: `http://localhost:1111/OrderItem`

### 4) Get by ID
- Method: `GET`
- URL: `http://localhost:1111/OrderItem/{OrderItem_id}`

### 5) Get Items by Order
- Method: `GET`
- URL: `http://localhost:1111/OrderItem-order/{OrderItem_id}`

### 6) Delete by ID
- Method: `DELETE`
- URL: `http://localhost:1111/OrderItem/{OrderItem_id}`

### 7) Delete All
- Method: `DELETE`
- URL: `http://localhost:1111/OrderItem`

---

## 💳 Invoice

### 1) Create
- Method: `POST`
- URL: `http://localhost:1111/invoice`
- Body:
```json
{
  "order_id": "{Order_id}",
  "payment_method": "CARD",
  "payment_status": "PENDING",
  "payment_due_date": "2026-03-15T00:00:00Z"
}
```

### 2) Update
- Method: `PATCH`
- URL: `http://localhost:1111/invoice/{invoice_id}`
- Body:
```json
{
  "payment_status": "PAID"
}
```

### 3) List
- Method: `GET`
- URL: `http://localhost:1111/invoice`

### 4) Get by ID
- Method: `GET`
- URL: `http://localhost:1111/invoice/{invoice_id}`

### 5) Delete by ID
- Method: `DELETE`
- URL: `http://localhost:1111/invoice/{invoice_id}`

### 6) Delete All
- Method: `DELETE`
- URL: `http://localhost:1111/invoice`
