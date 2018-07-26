-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "public"."borrows" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "book_id" uuid,
  "user_id" uuid,
  "from" timestamptz DEFAULT now(),
  "to" timestamptz,
  CONSTRAINT "borrow_book_id_fkey" FOREIGN KEY (book_id) REFERENCES books(id) NOT DEFERRABLE,
  CONSTRAINT "borrow_user_id_fkey" FOREIGN KEY (user_id) REFERENCES users(id) NOT DEFERRABLE,
  CONSTRAINT "borrow_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP "public"."borrows"