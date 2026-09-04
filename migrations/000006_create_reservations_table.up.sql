CREATE TABLE IF NOT EXISTS reservations (
    id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id bigint NOT NULL,
    showtime_id bigint NOT NULL
);

ALTER TABLE reservations
ADD CONSTRAINT reservations_fk_users
FOREIGN KEY (user_id) REFERENCES users (id)
ON DELETE RESTRICT 
ON UPDATE RESTRICT;

ALTER TABLE reservations
ADD CONSTRAINT reservations_fk_showtimes
FOREIGN KEY (showtime_id) REFERENCES showtimes (id)
ON DELETE RESTRICT 
ON UPDATE RESTRICT;
