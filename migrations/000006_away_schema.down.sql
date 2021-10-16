DROP INDEX users_email_idx;
DROP INDEX room_restrictions_start_date_idx;
DROP INDEX room_restrictions_end_date_idx;
DROP INDEX room_restrictions_room_id_idx;
DROP INDEX room_restrictions_reservation_id_idx;
DROP INDEX reservations_email_idx;
DROP INDEX reservations_last_name_idx;

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS room_restrictions;
DROP TABLE IF EXISTS restrictions;
DROP TABLE IF EXISTS reservations;
DROP TABLE IF EXISTS rooms;

