# Backend - System Booking Hotel (Server Side)

## Models :

### User Service

_User_

```json
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	Role      string    `json:"role" gorm:"default:'customer'"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
```

### Hotels Service

_Hotel_

```json
type Hotel struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:100;not null" json:"name"`
	Address     string    `gorm:"not null" json:"address"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Rooms       []Room    `gorm:"foreignKey:HotelID" json:"rooms"`
}
```

_Room_

```json
type Room struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	HotelID   uint      `json:"hotelId"`
	Number    string    `gorm:"size:10;not null" json:"number"`
	Type      string    `gorm:"size:50;not null" json:"type"`
	Price     uint   	`gorm:"not null" json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
```

### Booking Service

_Booking_

```json
type Booking struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint      `json:"userId"`
	RoomID       uint      `json:"roomId"`
	CheckInDate  time.Time `json:"check_in_date"`
	CheckOutDate time.Time `json:"check_out_date"`
	Status       string    `json:"status" gorm:"default:pending"` // pending, confirmed, cancelled
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
```

_Payment_

```json
type Payment struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	BookingID     uint      `json:"bookingId"`
	Amount        int       `json:"amount"`
	PaymentMethod string    `json:"payment_method"` // e.g., card, paypal
	Status        string    `json:"status"`         // pending, completed, failed
	PaymentDate   time.Time `json:"payment_date"`
}
```

_Refund_

```json
type Refund struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	BookingID    uint      `json:"bookingId"`
	Amount       int       `json:"amount"`
	Reason       string    `json:"reason"`
	Status       string    `json:"status"` // pending, approved, rejected
	RequestDate  time.Time `json:"request_date"`
	ApprovedDate *time.Time `json:"approved_date,omitempty"`
}
```

### Notification Service

_Notification_

```json
type Notification struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"userId"`
	Type      string    `json:"type"` // e.g., "booking", "payment", etc.
	Message   string    `json:"message"`
	IsRead    bool      `json:"is_read" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
```

&nbsp;

## List of available endpoints:

### Orchestrator - Cek Connection

- `GET /` http://localhost:8080/
- `GET /health` http://localhost:8080/health

### User Service

- `POST /users/register` http://localhost:8080/users/register
- `POST /users/login` http://localhost:8080/users/login
- `GET /users/` http://localhost:8080/users/
- `GET /users/:id` http://localhost:8080/users/:id

### Hotel Service

- `POST /hotels` http://localhost:8080/hotels
- `POST /rooms` http://localhost:8082/rooms
- `GET /hotels/` http://localhost:8080/hotels/
- `GET /hotels/:id` http://localhost:8080/hotels/:id

### Booking Service

- `POST /bookings` http://localhost:8080/bookings
- `GET /bookings/` http://localhost:8080/bookings/
- `PACTH /bookings/:id` http://localhost:8080/bookings/:id
- `DELETE /bookings/:id` http://localhost:8080/bookings/:id
- `POST /bookings/payments` http://localhost:8080/bookings/payments/:bookingId
- `GET /bookings/paymens/` http://localhost:8080/bookings/payments/

- `POST /bookings/refunds` http://localhost:8080/bookings/refunds
- `GET /bookings/refunds/` http://localhost:8080/bookings/refunds/
- `PATCH /bookings/refunds/:refundId` http://localhost:8080/bookings/refunds/:refundId

&nbsp;

## Testing In Postman :

Cara setting di postman :

1. Pilih endpoint yang ingin dijalankan
2. Jika method GET langsung jalankan / klik start
3. Jika method POST / PATCH / PUT memerlukan Request dengan klik Body pilih raw dan JSON

### 1.POST /users/register

Request:

- body: Body - raw - JSON

```json
{
  "name": "user1",
  "email": "user1@mail.com",
  "password": "password123"
}
```

atau test membuat akun dengan role admin

- body: Body - raw - JSON

```json
{
  "name": "user2",
  "email": "user2@mail.com",
  "password": "password123",
  "role": "admin"
}
```

_Response (201 - OK)_

Respon tanpa spesifik role

```json
{
  "id": 1,
  "name": "user1",
  "email": "user1@mail.com",
  "password": "$2a$10$Kh5w9h.Fy0fJow3.JcrHg.xUvU/TDIFn5qqeeLHqIFAKH4Fw9daPS",
  "role": "customer",
  "created_at": "2024-12-09T07:45:57.571481782Z",
  "updated_at": "2024-12-09T07:45:57.571481782Z"
}
```

Respon spesifik role admin

```json
{
  "id": 2,
  "name": "user2",
  "email": "user2@mail.com",
  "password": "$2a$11$Kh5w9h.Fy0fJow3.JcrHg.xUvU/TDIFn5qqeeLHqIFAKH4Fw9dsaj",
  "role": "admin",
  "created_at": "2024-12-09T07:45:57.571481782Z",
  "updated_at": "2024-12-09T07:45:57.571481782Z"
}
```

_Response (400 - Bad Request)_

```json
{
  "error": "ERROR: duplicate key value violates unique constraint \"uni_users_email\" (SQLSTATE 23505)"
}
```

### 2. POST /users/login

Request:

- body: Body - raw - JSON

```json
{
  "email": "user1@mail.com",
  "password": "password123"
}
```

_Response (200 - OK)_

```json
{
  "message": "Login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzM4MTcyMTksIm5hbWUiOiJ1c2VyMSIsInJvbGUiOiJhZG1pbiIsInVzZXJJZCI6MX0.NqquECyjQwOe8kJQoEZTe_5DlSQL9yYfn8BmfGk4mx0"
}
```

_Response (404 - Not Found)_

```json
{
  "error": "User not found"
}
```

_Response (401 - Unauthorized)_

```json
{
  "error": "Invalid credentials"
}
```

### 3. POST GET /users/:id

Description: Get all users from database

- headers:

```json
{
  "Authorization": "Bearer <access_token>"
}
```

_Response (200 - OK)_

```json
[
  {
    "id": 1,
    "name": "user1",
    "email": "user1@mail.com",
    "password": "$2a$10$Kh5w9h.Fy0fJow3.JcrHg.xUvU/TDIFn5qqeeLHqIFAKH4Fw9daPS",
    "role": "customer",
    "created_at": "2024-12-09T07:45:57.571481782Z",
    "updated_at": "2024-12-09T07:45:57.571481782Z"
  },
  {
    "id": 2,
    "name": "user2",
    "email": "user2@mail.com",
    "password": "$2a$10$Kh5w9h.Fy0fJow3.JcrHg.xUvU/TDIFn5qqeeLHqIFAKH4Fw9daPS",
    "role": "customer",
    "created_at": "2024-12-09T07:45:57.571481782Z",
    "updated_at": "2024-12-09T07:45:57.571481782Z"
  }
]
```

_Response (401 - Unauthorized)_

```json
{
  "message": "Unauthorized"
}
OR
{
  "message": "Invalid Token"
}
```

### 4. POST GET /users/:id

Description: Get by id users from database

- headers:

```json
{
  "Authorization": "Bearer <access_token>"
}
```

_Response (200 - OK)_

```json
[
  {
    "id": 1,
    "name": "user1",
    "email": "user1@mail.com",
    "password": "$2a$10$Kh5w9h.Fy0fJow3.JcrHg.xUvU/TDIFn5qqeeLHqIFAKH4Fw9daPS",
    "role": "customer",
    "created_at": "2024-12-09T07:45:57.571481782Z",
    "updated_at": "2024-12-09T07:45:57.571481782Z"
  }
]
```

_Response (401 - Unauthorized)_

```json
{
  "message": "Unauthorized"
}
OR
{
  "message": "Invalid Token"
}
```

### 5.POST /hotels

Request:

- body: Body - raw - JSON

```json
{
  "name": "Hotel Liliput",
  "address": "Indonesia, BSD",
  "description": "Tempat yang menyegarkan"
}
```

_Response (201 - Created)_

```json
{
  "id": 4,
  "name": "Hotel Liliput",
  "address": "Indonesia, BSD",
  "description": "Tempat yang menyegarkan",
  "created_at": "2024-12-09T08:06:39.874473812Z",
  "updated_at": "2024-12-09T08:06:39.874473812Z",
  "rooms": null
}
```

_Response (400 - Bad Request)_

```json
{
  "message": "Invalid input"
}
```

### 6.POST /rooms

Description: Create room berdasarkan id hotel

Request:

- body: Body - raw - JSON

```json
{
  "hotelId": 4,
  "number": "A 102",
  "type": "Double",
  "price": 800000
}
```

_Response (201 - Created)_

```json
{
  "id": 6,
  "hotelId": 4,
  "number": "A 102",
  "type": "Double",
  "price": 800000,
  "created_at": "2024-12-09T08:10:46.338696147Z",
  "updated_at": "2024-12-09T08:10:46.338696147Z"
}
```

_Response (400 - Bad Request)_

```json
{
  "message": "Invalid input"
}
```

### 7. GET /hotels/

Description: Get All hotels from database

_Response (200 - OK)_

```json
[
  {
    "id": 1,
    "name": "Hotel Liliput",
    "address": "Indonesia, BSD",
    "description": "Tempat yang menyegarkan",
    "created_at": "2024-12-08T10:09:39.939054Z",
    "updated_at": "2024-12-08T10:09:39.939054Z",
    "rooms": [
      {
        "id": 2,
        "hotelId": 1,
        "number": "A 102",
        "type": "Double",
        "price": 800000,
        "created_at": "2024-12-08T10:10:03.17437Z",
        "updated_at": "2024-12-08T10:10:03.17437Z"
      },
      {
        "id": 3,
        "hotelId": 1,
        "number": "A 102",
        "type": "Double",
        "price": 800000,
        "created_at": "2024-12-08T10:10:05.337639Z",
        "updated_at": "2024-12-08T10:10:05.337639Z"
      }
    ]
  },
  {
    "id": 2,
    "name": "Hotel Liliput",
    "address": "Indonesia, BSD",
    "description": "Tempat yang menyegarkan",
    "created_at": "2024-12-08T15:06:02.284132Z",
    "updated_at": "2024-12-08T15:06:02.284132Z",
    "rooms": []
  },
  {
    "id": 3,
    "name": "Hotel Liliput",
    "address": "Indonesia, BSD",
    "description": "Tempat yang menyegarkan",
    "created_at": "2024-12-09T06:49:20.151517Z",
    "updated_at": "2024-12-09T06:49:20.151517Z",
    "rooms": []
  },
  {
    "id": 4,
    "name": "Hotel Liliput",
    "address": "Indonesia, BSD",
    "description": "Tempat yang menyegarkan",
    "created_at": "2024-12-09T08:06:39.874473Z",
    "updated_at": "2024-12-09T08:06:39.874473Z",
    "rooms": [
      {
        "id": 5,
        "hotelId": 4,
        "number": "A 102",
        "type": "Double",
        "price": 800000,
        "created_at": "2024-12-09T08:08:43.511474Z",
        "updated_at": "2024-12-09T08:08:43.511474Z"
      },
      {
        "id": 6,
        "hotelId": 4,
        "number": "A 102",
        "type": "Double",
        "price": 800000,
        "created_at": "2024-12-09T08:10:46.338696Z",
        "updated_at": "2024-12-09T08:10:46.338696Z"
      }
    ]
  },
  {
    "id": 5,
    "name": "",
    "address": "Indonesia, BSD",
    "description": "Tempat yang menyegarkan",
    "created_at": "2024-12-09T08:07:42.091101Z",
    "updated_at": "2024-12-09T08:07:42.091101Z",
    "rooms": []
  },
  {
    "id": 6,
    "name": "",
    "address": "",
    "description": "",
    "created_at": "2024-12-09T08:07:54.834478Z",
    "updated_at": "2024-12-09T08:07:54.834478Z",
    "rooms": []
  }
]
```

### 8. GET /hotels/:hotelId

Description: Get By id hotels from database

_Response (200 - OK)_

```json
{
  "id": 1,
  "name": "Hotel Liliput",
  "address": "Indonesia, BSD",
  "description": "Tempat yang menyegarkan",
  "created_at": "2024-12-08T10:09:39.939054Z",
  "updated_at": "2024-12-08T10:09:39.939054Z",
  "rooms": [
    {
      "id": 2,
      "hotelId": 1,
      "number": "A 102",
      "type": "Double",
      "price": 800000,
      "created_at": "2024-12-08T10:10:03.17437Z",
      "updated_at": "2024-12-08T10:10:03.17437Z"
    },
    {
      "id": 3,
      "hotelId": 1,
      "number": "A 102",
      "type": "Double",
      "price": 800000,
      "created_at": "2024-12-08T10:10:05.337639Z",
      "updated_at": "2024-12-08T10:10:05.337639Z"
    }
  ]
}
```

### 9.POST /bookings

Description: Create booking

Request:

- body: Body - raw - JSON

```json
{
  "userId": 1,
  "roomId": 3,
  "check_in_date": "2024-12-10",
  "check_out_date": "2024-12-15"
}
```

_Response (201 - Created)_

```json
{
  "id": 5,
  "userId": 1,
  "roomId": 3,
  "check_in_date": "2024-12-10T00:00:00Z",
  "check_out_date": "2024-12-15T00:00:00Z",
  "status": "pending",
  "created_at": "2024-12-09T08:17:25.225084673Z",
  "updated_at": "2024-12-09T08:17:25.225084843Z"
}
```

_Response (400 - Bad Request)_

```json
{
    "error": "Missing required fields"
}
or
{
    "error": "Invalid input format"
}
```

### 10. GET /booking/

Description: Get all bookings from database

_Response (200 - OK)_

```json
[
  {
    "id": 2,
    "userId": 1,
    "roomId": 3,
    "check_in_date": "2024-12-10T00:00:00Z",
    "check_out_date": "2024-12-15T00:00:00Z",
    "status": "pending",
    "created_at": "2024-12-08T10:28:40.666953Z",
    "updated_at": "2024-12-08T10:28:40.666953Z"
  },
  {
    "id": 3,
    "userId": 1,
    "roomId": 2,
    "check_in_date": "2024-12-10T00:00:00Z",
    "check_out_date": "2024-12-15T00:00:00Z",
    "status": "pending",
    "created_at": "2024-12-08T11:05:35.026627Z",
    "updated_at": "2024-12-08T11:05:35.026628Z"
  },
  {
    "id": 4,
    "userId": 1,
    "roomId": 3,
    "check_in_date": "2024-12-10T00:00:00Z",
    "check_out_date": "2024-12-15T00:00:00Z",
    "status": "pending",
    "created_at": "2024-12-08T11:42:01.806126Z",
    "updated_at": "2024-12-08T11:42:01.806126Z"
  },
  {
    "id": 5,
    "userId": 1,
    "roomId": 3,
    "check_in_date": "2024-12-10T00:00:00Z",
    "check_out_date": "2024-12-15T00:00:00Z",
    "status": "pending",
    "created_at": "2024-12-09T08:17:25.225084Z",
    "updated_at": "2024-12-09T08:17:25.225084Z"
  }
]
```

### 11.PATCH /bookings/:bookingId

Description: Create booking

Request:

- body: Body - raw - JSON

```json
{
  "status": "confirm"
}
```

_Response (201 - Created)_

```json
{
  "booking": {
    "id": 3,
    "userId": 1,
    "roomId": 2,
    "check_in_date": "2024-12-10T00:00:00Z",
    "check_out_date": "2024-12-15T00:00:00Z",
    "status": "confirm",
    "created_at": "2024-12-08T11:05:35.026627Z",
    "updated_at": "2024-12-09T08:23:18.814176504Z"
  },
  "message": "Booking updated successfully"
}
```

_Response (404 - Not Found)_

```json
{
  "error": "Booking not found"
}
```

_Response (400 - Bad Request)_

```json
{
  "error": "Invalid input"
}
```

### 12.DELETE /bookings/:bookingId

Description: Create booking

Request:

- body: Body - raw - JSON

```json
{
  "status": "confirm"
}
```

_Response (200 - Created)_

```json
{
  "message": "Booking deleted successfully"
}
```

_Response (404 - Not Found)_

```json
{
  "error": "Booking not found"
}
```

### 13.POST /bookings/payments/:bookingId

Description: Create payment by id booking

Request:

- body: Body - raw - JSON

```json
{
  "amount": 800000,
  "method": "card"
}
```

_Response (201 - Created)_

```json
{
  "client_secret": "pi_3QU2LbAEeIeJxDdv0TkI6n2a_secret_qmzq7GBtYVT6YqE9eEiO5npRT",
  "message": "Payment initiated"
}
```

_Response (400 - Bad Request)_

```json
{
    "error": "Missing required fields"
}
or
{
    "error": "Invalid input format"
}
```

### 14. GET /bookings/payments/

Description: Get all bookings from database

_Response (200 - OK)_

```json
[
  {
    "id": 1,
    "bookingId": 2,
    "amount": 800000,
    "payment_method": "card",
    "status": "pending",
    "payment_date": "2024-12-08T11:07:33.120104Z"
  },
  {
    "id": 2,
    "bookingId": 3,
    "amount": 800000,
    "payment_method": "card",
    "status": "pending",
    "payment_date": "2024-12-08T11:50:08.025181Z"
  },
  {
    "id": 3,
    "bookingId": 3,
    "amount": 800000,
    "payment_method": "card",
    "status": "pending",
    "payment_date": "2024-12-09T08:29:47.967727Z"
  },
  {
    "id": 4,
    "bookingId": 3,
    "amount": 800000,
    "payment_method": "card",
    "status": "pending",
    "payment_date": "2024-12-09T08:30:17.14246Z"
  }
]
```

### 15.POST /bookings/refunds

Description: Create refund

Request:

- body: Body - raw - JSON

```json
{
  "bookingId": 4,
  "reason": "Customer requested cancellation"
}
```

_Response (201 - Created)_

```json
{
  "message": "Refund requested",
  "refund": {
    "id": 2,
    "bookingId": 4,
    "amount": 100,
    "reason": "Customer requested cancellation",
    "status": "pending",
    "request_date": "2024-12-09T08:31:57.680550929Z"
  }
}
```

_Response (404 - Not Found)_

```json
{
  "error": "Booking not found"
}
```

_Response (400 - Bad Request)_

```json
{
  "error": "Invalid input"
}
```

### 16. GET /bookings/refunds

Description: Get all refund from database

_Response (200 - OK)_

```json
[
  {
    "id": 1,
    "bookingId": 4,
    "amount": 100,
    "reason": "Customer requested cancellation",
    "status": "approved",
    "request_date": "2024-12-08T11:42:12.720786Z",
    "approved_date": "2024-12-08T11:46:14.903125Z"
  },
  {
    "id": 2,
    "bookingId": 4,
    "amount": 100,
    "reason": "Customer requested cancellation",
    "status": "pending",
    "request_date": "2024-12-09T08:31:57.68055Z"
  }
]
```

### 17.PATCH /bookings/refunds/:refundId

Description: Update status refund

Request:

- body: Body - raw - JSON

```json
{
  "status": "approved"
}
```

_Response (200 - Created)_

```json
{
  "message": "Refund status updated",
  "refund": {
    "id": 1,
    "bookingId": 4,
    "amount": 100,
    "reason": "Customer requested cancellation",
    "status": "approved",
    "request_date": "2024-12-08T11:42:12.720786Z",
    "approved_date": "2024-12-09T08:35:54.784986728Z"
  }
}
```

_Response (404 - Not Found)_

```json
{
  "error": "Refunds not found"
}
```

_Response (400 - Bad Request)_

```json
{
  "error": "Invalid input"
}
```
