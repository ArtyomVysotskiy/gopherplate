-- Схема fileProcessing
CREATE SCHEMA IF NOT EXISTS fileProcessing;
-- Таблица файлов
CREATE TABLE IF NOT EXISTS fileProcessing.files
(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    file_path TEXT NOT NULL,
    size BIGINT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                             CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES auth.users(idUser) ON DELETE CASCADE
    );
-- Таблица результатов поиска
CREATE TABLE IF NOT EXISTS fileProcessing.search_results
(
    id SERIAL PRIMARY KEY,
    file_id INT NOT NULL,
    query_text VARCHAR(255) NOT NULL,
    result_text TEXT NOT NULL,
    position INT NOT NULL,
    searched_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                              CONSTRAINT fk_file FOREIGN KEY (file_id) REFERENCES fileProcessing.files(id) ON DELETE CASCADE
    );
