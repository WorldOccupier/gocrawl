CREATE TABLE IF NOT EXISTS t_web_page_details (
    url VARCHAR(2048) PRIMARY KEY,
    crawled_at TIMESTAMP NOT NULL,
    content TEXT NOT NULL,
    processed BOOLEAN NOT NULL DEFAULT FALSE
);
