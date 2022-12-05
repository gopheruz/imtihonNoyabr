CREATE TABLE "users"(
    "id" SERIAl,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "image_url" VARCHAR(255) NOT NULL,
    "user_type"  VARCHAR(255) CHECK ("user_type" IN('superadmin', 'user')),
    "created_at" TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(0) WITH TIME ZONE,
    "deleted_at" TIMESTAMP(0) WITH TIME ZONE
);
ALTER TABLE
    "users" ADD PRIMARY KEY("id");
CREATE TABLE "hotels"(
    "id" SERIAL,
    "hotel_name" VARCHAR(255) NOT NULL,
    "adress" VARCHAR(255) NOT NULL,
    "desctiption" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "hotels" ADD PRIMARY KEY("id");
CREATE TABLE "hotel_images"(
    "id" SERIAL,
    "hotel_id" INTEGER NOT NULL,
    "imageurl" VARCHAR(255) NOT NULL,
    "sequence_number" INTEGER NOT NULL
);
ALTER TABLE
    "hotel_images" ADD PRIMARY KEY("id");
CREATE TABLE "bookings"(
    "id" SERIAL,
    "hotel_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "room_id" INTEGER NOT NULL,
    "from" DATE NOT NULL,
    "to" DATE NOT NULL,
    "total_price" DECIMAL(8, 2) NOT NULL
);
ALTER TABLE
    "bookings" ADD PRIMARY KEY("id");
CREATE TABLE "rooms"(
    "id" SERIAL,
    "hotel_id" INTEGER NOT NULL,
    "room_number" INTEGER NOT NULL
);
ALTER TABLE
    "rooms" ADD PRIMARY KEY("id");
ALTER TABLE
    "bookings" ADD CONSTRAINT "bookings_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE
    "hotel_images" ADD CONSTRAINT "hotel_images_hotel_id_foreign" FOREIGN KEY("hotel_id") REFERENCES "hotels"("id");
ALTER TABLE
    "bookings" ADD CONSTRAINT "bookings_hotel_id_foreign" FOREIGN KEY("hotel_id") REFERENCES "hotels"("id");
ALTER TABLE
    "bookings" ADD CONSTRAINT "bookings_room_id_foreign" FOREIGN KEY("room_id") REFERENCES "rooms"("id");
ALTER TABLE
    "rooms" ADD CONSTRAINT "rooms_hotel_id_foreign" FOREIGN KEY("hotel_id") REFERENCES "hotels"("id");