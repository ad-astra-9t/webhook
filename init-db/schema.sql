CREATE TABLE webhooks (
    id SERIAL PRIMARY KEY,
    callback VARCHAR(255) NOT NULL
);

CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);


CREATE TABLE event_webhooks (
    event_id INTEGER REFERENCES events(id),
    webhook_id INTEGER REFERENCES webhooks(id)
);
