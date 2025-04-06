-- Схема auth
CREATE SCHEMA IF NOT EXISTS auth;
-- Таблица пользователей
CREATE TABLE IF NOT EXISTS auth.users
(
    idUser SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL
);