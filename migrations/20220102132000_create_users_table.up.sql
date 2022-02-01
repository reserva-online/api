CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "email" TEXT NOT NULL,
  "password" TEXT NOT NULL,
  "type" TEXT NOT NULL
);