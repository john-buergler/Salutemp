DROP TABLE IF EXISTS medication_constraint;
DROP TABLE IF EXISTS alert;
DROP TABLE IF EXISTS status_report;
DROP TYPE IF EXISTS condition_type;
DROP TABLE IF EXISTS stored_medication;
DROP TABLE IF EXISTS medication;
DROP TABLE IF EXISTS user_device;
DROP TABLE IF EXISTS "user";

CREATE TYPE condition_type AS ENUM ('TEMPERATURE', 'HUMIDITY', 'LIGHT_EXPOSURE');

CREATE TABLE IF NOT EXISTS "user" (
    user_id varchar NOT NULL UNIQUE,
    first_name varchar NOT NULL,
    last_name varchar NOT NULL,
    email varchar NOT NULL,
    push_notification_enabled BOOLEAN DEFAULT FALSE,
    PRIMARY KEY (user_id)
);

CREATE TABLE IF NOT EXISTS expo_notification_token (
    expo_notification_token_id SERIAL PRIMARY KEY,
    user_id varchar UNIQUE NOT NULL,  -- Add UNIQUE constraint here
    device_token VARCHAR NOT NULL,
    FOREIGN KEY (user_id) REFERENCES "user" (user_id)
);

CREATE TABLE IF NOT EXISTS user_device (
    user_device_id SERIAL PRIMARY KEY,
    user_id varchar NOT NULL,
    device_id VARCHAR NOT NULL,
    FOREIGN KEY (user_id) REFERENCES "user" (user_id)
);

CREATE TABLE IF NOT EXISTS medication (
    medication_id integer NOT NULL UNIQUE,
    medication_name varchar NOT NULL,
    PRIMARY KEY (medication_id)
);

CREATE TABLE IF NOT EXISTS stored_medication ( /*medication_instance?*/
    stored_medication_id integer NOT NULL,
    medication_id integer NOT NULL,
    user_id varchar NOT NULL,
    current_temperature float,
    current_humidity float,
    current_light float,
    PRIMARY KEY (stored_medication_id),
    FOREIGN KEY (medication_id) REFERENCES medication (medication_id),
    FOREIGN KEY (user_id) REFERENCES "user" (user_id)
    );

CREATE TABLE IF NOT EXISTS alert (
    warning_id integer NOT NULL UNIQUE,
    stored_medication_id integer NOT NULL,
    warning_timestamp timestamp NOT NULL,
    warning_description varchar NOT NULL,
    condition_type condition_type NOT NULL,
    PRIMARY KEY (warning_id),
    FOREIGN KEY (stored_medication_id) REFERENCES stored_medication (stored_medication_id)
);

CREATE TABLE IF NOT EXISTS status_report (
    event_time timestamp NOT NULL,
    stored_medication_id integer NOT NULL,
    temperature float,
    humidity float,
    light float,
    PRIMARY KEY (event_time, stored_medication_id),
    FOREIGN KEY (stored_medication_id) REFERENCES stored_medication (stored_medication_id)
);

CREATE TABLE IF NOT EXISTS medication_constraint (
    stored_medication_id integer NOT NULL,
    condition_type condition_type NOT NULL,
    max_threshold float,
    min_threshold float,
    duration varchar, /*Not sure if we should store this as a time object.*/
    PRIMARY KEY (stored_medication_id, condition_type),
    FOREIGN KEY (stored_medication_id) REFERENCES stored_medication (stored_medication_id)
);

-- populate
-- Add users
-- Insert sample data into "user" table
INSERT INTO "user" (user_id, first_name, last_name, email, push_notification_enabled)
VALUES
  ('1', 'John', 'Doe', 'john.doe@example.com', true),
  ('2', 'Jane', 'Smith', 'jane.smith@example.com', false),
  ('3', 'Bob', 'Johnson', 'bob.johnson@example.com', true),
  ('4', 'Alice', 'Williams', 'alice.williams@example.com', false),
  ('5', 'Charlie', 'Brown', 'charlie.brown@example.com', true);

-- Insert sample data into "user_device" table
INSERT INTO user_device (user_id, device_id)
VALUES
  (1, 'device_1'),
  (2, 'device_2'),
  (3, 'device_3'),
  (4, 'device_4'),
  (5, 'device_5');

-- Insert sample data into "medication" table
INSERT INTO medication (medication_id, medication_name)
VALUES
  (1, 'Medication A'),
  (2, 'Medication B'),
  (3, 'Medication C'),
  (4, 'Medication D'),
  (5, 'Medication E'),
  (6, 'Medication F'),
  (7, 'Medication G'),
  (8, 'Medication H'),
  (9, 'Medication I'),
  (10, 'Medication J'),
  (11, 'Medication K'),
  (12, 'Medication L'),
  (13, 'Medication M'),
  (14, 'Medication N'),
  (15, 'Medication O'),
  (16, 'Medication P'),
  (17, 'Medication Q'),
  (18, 'Medication R'),
  (19, 'Medication S'),
  (20, 'Medication T');

-- Insert sample data into "stored_medication" table
INSERT INTO stored_medication (stored_medication_id, medication_id, user_id, current_temperature, current_humidity, current_light)
VALUES
  (1, 1, 1, 25.5, 50.0, 300),
  (2, 2, 2, 22.0, 40.0, 200),
  (3, 3, 3, 26.5, 60.0, 400),
  (4, 4, 4, 23.0, 45.0, 250),
  (5, 5, 5, 24.5, 55.0, 350),
  (6, 6, 1, 25.0, 50.0, 300),
  (7, 7, 1, 22.0, 40.0, 200),
  (8, 8, 1, 26.5, 60.0, 400),
  (9, 9, 1, 23.0, 45.0, 250),
  (10, 10, 1, 24.5, 55.0, 350),
  (11, 11, 1, 25.2, 48.0, 290),
  (12, 12, 1, 22.8, 42.0, 210),
  (13, 13, 1, 26.0, 58.0, 380),
  (14, 14, 1, 23.5, 46.0, 260),
  (15, 15, 1, 24.8, 53.0, 330);

-- Insert sample data into "alert" table
INSERT INTO alert (warning_id, stored_medication_id, warning_timestamp, warning_description, condition_type)
VALUES
  (1, 1, '2023-01-01 08:00:00', 'High Temperature', 'TEMPERATURE'),
  (2, 2, '2023-01-02 10:30:00', 'Low Humidity', 'HUMIDITY'),
  (3, 3, '2023-01-03 12:45:00', 'High Light Exposure', 'LIGHT_EXPOSURE'),
  (4, 4, '2023-01-04 14:15:00', 'Low Temperature', 'TEMPERATURE'),
  (5, 5, '2023-01-05 16:30:00', 'High Humidity', 'HUMIDITY'),
  (6, 3, '2023-01-03 12:45:00', 'High Humidity', 'HUMIDITY');

-- Insert sample data into "medication_constraint" table
-- Insert additional sample data into "medication_constraint" table for all medications
INSERT INTO medication_constraint (stored_medication_id, condition_type, max_threshold, min_threshold, duration)
VALUES
  -- Medication A
  (1, 'TEMPERATURE', 30.0, 20.0, '1 day'),
  (1, 'HUMIDITY', 55.0, 40.0, '2 days'),
  (1, 'LIGHT_EXPOSURE', 400.0, 250.0, '3 days'),

  -- Medication B
  (2, 'TEMPERATURE', 28.0, 18.0, '1 day'),
  (2, 'HUMIDITY', 50.0, 35.0, '2 days'),
  (2, 'LIGHT_EXPOSURE', 420.0, 20.0, '3 days'),

  -- Medication C
  (3, 'TEMPERATURE', 26.0, 16.0, '1 day'),
  (3, 'HUMIDITY', 58.0, 42.0, '2 days'),
  (3, 'LIGHT_EXPOSURE', 450.0, 280.0, '3 days'),

  -- Medication D
  (4, 'TEMPERATURE', 27.0, 17.0, '1 day'),
  (4, 'HUMIDITY', 45.0, 30.0, '2 days'),
  (4, 'LIGHT_EXPOSURE', 430.0, 20.0, '3 days'),

  -- Medication E
  (5, 'TEMPERATURE', 29.0, 19.0, '1 day'),
  (5, 'HUMIDITY', 53.0, 38.0, '2 days'),
  (5, 'LIGHT_EXPOSURE', 410.0, 2.0, '3 days'),

  
  (6, 'TEMPERATURE', 26.5, 16.5, '1 day'),
  (6, 'HUMIDITY', 49.0, 34.0, '2 days'),
  (6, 'LIGHT_EXPOSURE', 410.0, 220.0, '3 days'),

  (7, 'TEMPERATURE', 27.5, 17.5, '1 day'),
  (7, 'HUMIDITY', 50.0, 35.0, '2 days'),
  (7, 'LIGHT_EXPOSURE', 400.0, 250.0, '3 days'),

  (8, 'TEMPERATURE', 29.0, 19.0, '1 day'),
  (8, 'HUMIDITY', 51.0, 36.0, '2 days'),
  (8, 'LIGHT_EXPOSURE', 420.0, 260.0, '3 days'),

  (9, 'TEMPERATURE', 28.0, 18.0, '1 day'),
  (9, 'HUMIDITY', 48.0, 33.0, '2 days'),
  (9, 'LIGHT_EXPOSURE', 430.0, 240.0, '3 days'),

  (10, 'TEMPERATURE', 26.0, 16.0, '1 day'),
  (10, 'HUMIDITY', 47.0, 32.0, '2 days'),
  (10, 'LIGHT_EXPOSURE', 450.0, 280.0, '3 days'),

  (11, 'TEMPERATURE', 27.0, 17.0, '1 day'),
  (11, 'HUMIDITY', 45.0, 30.0, '2 days'),
  (11, 'LIGHT_EXPOSURE', 440.0, 270.0, '3 days'),

  (12, 'TEMPERATURE', 29.0, 19.0, '1 day'),
  (12, 'HUMIDITY', 53.0, 38.0, '2 days'),
  (12, 'LIGHT_EXPOSURE', 410.0, 255.0, '3 days'),

  (13, 'HUMIDITY', 46.0, 31.0, '2 days'),
  (13, 'LIGHT_EXPOSURE', 420.0, 260.0, '3 days'),

  (14, 'HUMIDITY', 49.0, 34.0, '2 days'),
  (14, 'LIGHT_EXPOSURE', 400.0, 250.0, '3 days'),

  (15, 'HUMIDITY', 50.0, 35.0, '2 days'),
  (15, 'LIGHT_EXPOSURE', 410.0, 260.0, '3 days');



  -- Insert sample data into "expo_notification_token" table
INSERT INTO expo_notification_token (user_id, device_token)
VALUES
  (2, 'expo_device_token_2'),
  (3, 'expo_device_token_3'),
  (4, 'expo_device_token_4'),
  (5, 'expo_device_token_5');


