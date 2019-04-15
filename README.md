# DemoServer
 How to:

CREATE DATABASE dvdrental;

CREATE USER testuser WITH PASSWORD 'Aa123456!'

GRANT ALL PRIVILEGES ON DATABASE "dvdrental" to testuser;

pg_dump -U testuser dvdrental < db.sql
