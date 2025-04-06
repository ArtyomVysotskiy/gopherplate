-- Схема fileManagement
CREATE SCHEMA IF NOT EXISTS fileManagement;
-- Таблица файлов
CREATE TABLE IF NOT EXISTS fileManagement.files
(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    file_path TEXT NOT NULL,
    size BIGINT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                             CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES auth.users(idUser) ON DELETE CASCADE
    );
-- Таблица истории переименований
CREATE TABLE IF NOT EXISTS fileManagement.file_rename_history
(
    id SERIAL PRIMARY KEY,
    file_id INT NOT NULL,
    old_name VARCHAR(255) NOT NULL,
    new_name VARCHAR(255) NOT NULL,
    renamed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                             user_id INT NOT NULL,
                             CONSTRAINT fk_file FOREIGN KEY (file_id) REFERENCES fileManagement.files(id) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES auth.users(idUser) ON DELETE CASCADE
    );
-- Таблица разделения файлов
CREATE TABLE IF NOT EXISTS fileManagement.split_files
(
    id SERIAL PRIMARY KEY,
    original_file_id INT NOT NULL,
    part_file_id INT NOT NULL,
    user_id INT NOT NULL,
    split_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_original_file FOREIGN KEY (original_file_id) REFERENCES fileManagement.files(id) ON DELETE CASCADE,
    CONSTRAINT fk_part_file FOREIGN KEY (part_file_id) REFERENCES fileManagement.files(id) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES auth.users(idUser) ON DELETE CASCADE
);