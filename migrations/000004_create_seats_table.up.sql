CREATE TABLE IF NOT EXISTS seats (
    id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    hall_id bigint NOT NULL,
    row_label text NOT NULL,
    seat_number int NOT NULL
);

ALTER TABLE seats
ADD CONSTRAINT seats_fk_halls
FOREIGN KEY (hall_id) REFERENCES halls (id)
ON DELETE RESTRICT 
ON UPDATE RESTRICT;
