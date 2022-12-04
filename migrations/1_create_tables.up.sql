CREATE TABLE "users"(
    "id" SERIAL,
    "first_name" VARCHAR(255),
    "last_name" VARCHAR(255),
    "email" VARCHAR(255),
    "password" VARCHAR(255),
    "image_url" VARCHAR(255),
    "created_at" TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(0) WITH TIME ZONE ,
    "deleted_at" TIMESTAMP(0) WITH TIME ZONE
);
ALTER TABLE
    "users" ADD PRIMARY KEY("id");
CREATE TABLE "hotels"(
    "id" SERIAL,
    "hotel_name" VARCHAR(255) NOT NULL,
    "adress" VARCHAR(255) NOT NULL,
    "image_id" INTEGER,
    "rooms_count" INTEGER NOT NULL
);
ALTER TABLE
    "hotels" ADD PRIMARY KEY("id");
CREATE TABLE "hotel_images"(
    "id" INTEGER NOT NULL,
    "hotel_id" INTEGER NOT NULL,
    "imageurl" VARCHAR(255) NOT NULL,
    "sequence_number" INTEGER NOT NULL
);
ALTER TABLE
    "hotel_images" ADD PRIMARY KEY("id");
CREATE TABLE "orders"(
    "id" SERIAL,
    "hotel_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "room_id" INTEGER NOT NULL
);
ALTER TABLE
    "orders" ADD PRIMARY KEY("id");
CREATE TABLE "rooms"(
    "id" SERIAL,
    "hotel_id" INTEGER NOT NULL,
    "room_number" INTEGER NOT NULL,
    "boked_from" TIMESTAMP(0) WITH TIME ZONE,
    "boked_to" TIMESTAMP(0) WITH TIME ZONE
);
ALTER TABLE
    "rooms" ADD PRIMARY KEY("id");
ALTER TABLE
    "orders" ADD CONSTRAINT "orders_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE
    "orders" ADD CONSTRAINT "orders_hotel_id_foreign" FOREIGN KEY("hotel_id") REFERENCES "hotels"("id");
ALTER TABLE
    "hotels" ADD CONSTRAINT "hotels_image_id_foreign" FOREIGN KEY("image_id") REFERENCES "hotel_images"("id");
ALTER TABLE
    "orders" ADD CONSTRAINT "orders_room_id_foreign" FOREIGN KEY("room_id") REFERENCES "rooms"("id");
ALTER TABLE
    "rooms" ADD CONSTRAINT "rooms_hotel_id_foreign" FOREIGN KEY("hotel_id") REFERENCES "hotels"("id");