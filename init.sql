CREATE TABLE "contacts" (
    "id" INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "email" VARCHAR(80) UNIQUE NOT NULL,
    "phone" VARCHAR(20),
    "created_at" TIMESTAMP DEFAULT NOW(),
    "updated_at" TIMESTAMP DEFAULT NOW()
);

INSERT INTO "contacts" ("email", "phone") VALUES
('jargg@mail.com', '081234567'),
('jhon@mail.com', '08123456123');

SELECT * FROM "contacts";