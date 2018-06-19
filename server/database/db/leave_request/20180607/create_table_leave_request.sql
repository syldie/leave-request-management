CREATE TABLE IF NOT EXISTS leave_request 
(
  "id" SERIAL PRIMARY KEY NOT NULL,
  "employee_number" int NOT NULL REFERENCES users(employee_number),
  "type_of_leave" text NOT NULL,
  "from"  text NOT NULL,
  "to"  text NOT NULL,
  "total" int NOT NULL,
  "leave_remaining" int NOT NULL,
  "reason" text NOT NULL,
  "address" text NOT NULL,
  "contact_leave" text,
  "status" text NOT NULL,
  "reject_reason" text,
  "approved_by" text
)