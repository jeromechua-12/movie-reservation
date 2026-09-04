CREATE TABLE IF NOT EXISTS reserved_seats (
    id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    reservation_id bigint NOT NULL,
    showtime_id bigint NOT NULL,
    seat_id bigint NOT NULL
);

ALTER TABLE reserved_seats
ADD CONSTRAINT reserved_seats_fk_reservations
FOREIGN KEY (reservation_id) REFERENCES reservations (id)
ON DELETE RESTRICT 
ON UPDATE RESTRICT;

ALTER TABLE reserved_seats
ADD CONSTRAINT reserved_seats_fk_showtimes
FOREIGN KEY (showtime_id) REFERENCES showtimes (id)
ON DELETE RESTRICT 
ON UPDATE RESTRICT;

ALTER TABLE reserved_seats
ADD CONSTRAINT reserved_seats_fk_seats
FOREIGN KEY (seat_id) REFERENCES seats (id)
ON DELETE RESTRICT 
ON UPDATE RESTRICT;
