// Demo seed data for user-service (MongoDB). Demo account password: DemoPass123!
// Run:
//   podman exec -i user-mongodb mongo userService < scripts/seed/mongo.js
// Database name matches user-service config (userService).
// Safe to rerun: missing users are inserted and seeded passwords are refreshed.

var now = new Date();
var users = [
    {
        "userId": 1001,
        "phoneNumber": "5552000001",
        "email": "demo.customer01@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "CUSTOMER",
        "firstName": "Avery",
        "lastName": "Bennett"
    },
    {
        "userId": 1002,
        "phoneNumber": "5552000002",
        "email": "demo.customer02@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "CUSTOMER",
        "firstName": "Jordan",
        "lastName": "Carter"
    },
    {
        "userId": 1003,
        "phoneNumber": "5552000003",
        "email": "demo.customer03@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "CUSTOMER",
        "firstName": "Taylor",
        "lastName": "Diaz"
    },
    {
        "userId": 1004,
        "phoneNumber": "5552000004",
        "email": "demo.customer04@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "CUSTOMER",
        "firstName": "Morgan",
        "lastName": "Ellis"
    },
    {
        "userId": 1005,
        "phoneNumber": "5552000005",
        "email": "demo.customer05@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "CUSTOMER",
        "firstName": "Riley",
        "lastName": "Foster"
    },
    {
        "userId": 1006,
        "phoneNumber": "5552000006",
        "email": "demo.customer06@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "CUSTOMER",
        "firstName": "Casey",
        "lastName": "Garcia"
    },
    {
        "userId": 1007,
        "phoneNumber": "5552000007",
        "email": "demo.customer07@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "CUSTOMER",
        "firstName": "Cameron",
        "lastName": "Hayes"
    },
    {
        "userId": 1008,
        "phoneNumber": "5552000008",
        "email": "demo.customer08@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "CUSTOMER",
        "firstName": "Drew",
        "lastName": "Irwin"
    },
    {
        "userId": 1009,
        "phoneNumber": "5552000009",
        "email": "demo.customer09@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "CUSTOMER",
        "firstName": "Quinn",
        "lastName": "Jordan"
    },
    {
        "userId": 1010,
        "phoneNumber": "5552000010",
        "email": "demo.customer10@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "CUSTOMER",
        "firstName": "Hayden",
        "lastName": "Kim"
    },
    {
        "userId": 1011,
        "phoneNumber": "5552000011",
        "email": "demo.customer11@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "CUSTOMER",
        "firstName": "Reese",
        "lastName": "Lane"
    },
    {
        "userId": 1012,
        "phoneNumber": "5552000012",
        "email": "demo.customer12@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "CUSTOMER",
        "firstName": "Parker",
        "lastName": "Morris"
    },
    {
        "userId": 1013,
        "phoneNumber": "5552000013",
        "email": "demo.customer13@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "CUSTOMER",
        "firstName": "Rowan",
        "lastName": "Nguyen"
    },
    {
        "userId": 1014,
        "phoneNumber": "5552000014",
        "email": "demo.customer14@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "CUSTOMER",
        "firstName": "Skyler",
        "lastName": "Owens"
    },
    {
        "userId": 1015,
        "phoneNumber": "5552000015",
        "email": "demo.customer15@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "CUSTOMER",
        "firstName": "Finley",
        "lastName": "Patel"
    },
    {
        "userId": 1016,
        "phoneNumber": "5552000016",
        "email": "demo.customer16@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "CUSTOMER",
        "firstName": "Emerson",
        "lastName": "Quinn"
    },
    {
        "userId": 1017,
        "phoneNumber": "5552000017",
        "email": "demo.customer17@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "CUSTOMER",
        "firstName": "Dakota",
        "lastName": "Reed"
    },
    {
        "userId": 1018,
        "phoneNumber": "5552000018",
        "email": "demo.customer18@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "CUSTOMER",
        "firstName": "Arden",
        "lastName": "Singh"
    },
    {
        "userId": 1019,
        "phoneNumber": "5552000019",
        "email": "demo.customer19@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "CUSTOMER",
        "firstName": "Sage",
        "lastName": "Turner"
    },
    {
        "userId": 1020,
        "phoneNumber": "5552000020",
        "email": "demo.customer20@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "CUSTOMER",
        "firstName": "Logan",
        "lastName": "Upton"
    },
    {
        "userId": 1021,
        "phoneNumber": "5552000021",
        "email": "demo.merchant21@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "MERCHANT",
        "firstName": "Kendall",
        "lastName": "Vargas"
    },
    {
        "userId": 1022,
        "phoneNumber": "5552000022",
        "email": "demo.merchant22@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "MERCHANT",
        "firstName": "Jamie",
        "lastName": "Walker"
    },
    {
        "userId": 1023,
        "phoneNumber": "5552000023",
        "email": "demo.merchant23@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "MERCHANT",
        "firstName": "Blake",
        "lastName": "Xu"
    },
    {
        "userId": 1024,
        "phoneNumber": "5552000024",
        "email": "demo.merchant24@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "MERCHANT",
        "firstName": "Payton",
        "lastName": "Young"
    },
    {
        "userId": 1025,
        "phoneNumber": "5552000025",
        "email": "demo.merchant25@example.com",
        "password": "$2b$10$NdHeFH3Cf8BrwvqM7.3v9u0O3GqKuf1f0aLBuFoRW3KXA2UxZglam",
        "role": "MERCHANT",
        "firstName": "Alex",
        "lastName": "Zimmer"
    }
];

users.forEach(function (user) {
    var userOnInsert = Object.assign({}, user, {
        createdDate: now,
        isActive: true,
        address: [],
    });
    delete userOnInsert.password;

    db.getCollection("users").updateOne(
        { userId: user.userId },
        {
            $set: { password: user.password },
            $setOnInsert: userOnInsert,
        },
        { upsert: true }
    );
});

// Keep the userId counter ahead of the seeded users.
db.getCollection("counter").updateOne(
    { _id: "userid" },
    { $max: { seq: 1025 } },
    { upsert: true }
);
