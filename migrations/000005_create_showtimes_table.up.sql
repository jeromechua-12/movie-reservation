CREATE TABLE IF NOT EXISTS showtimes (
    id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    movie_id bigint NOT NULL,
    hall_id bigint NOT NULL,
    seat_id bigint NOT NULL,
    start_time timestamp(0) with time zone NOT NULL
);

ALTER TABLE showtimes
ADD CONSTRAINT showtimes_fk_movies
FOREIGN KEY (movie_id) REFERENCES movies(id)
ON DELETE RESTRICT 
ON UPDATE RESTRICT;

ALTER TABLE showtimes
ADD CONSTRAINT showtimes_fk_halls
FOREIGN KEY (hall_id) REFERENCES halls (id)
ON DELETE RESTRICT 
ON UPDATE RESTRICT;

ALTER TABLE showtimes
ADD CONSTRAINT showtimes_fk_seats
FOREIGN KEY (seat_id) REFERENCES seats (id)
ON DELETE RESTRICT 
ON UPDATE RESTRICT;
