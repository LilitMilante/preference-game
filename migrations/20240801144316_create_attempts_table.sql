-- +goose Up
-- +goose StatementBegin
CREATE TABLE attempts (
      id SERIAL PRIMARY KEY,
      user_id VARCHAR(255) NOT NULL,
      first_card_suit VARCHAR(50) NOT NULL,
      first_card_value VARCHAR(50) NOT NULL,
      second_card_suit VARCHAR(50) NOT NULL,
      second_card_value VARCHAR(50) NOT NULL,
      is_win BOOLEAN NOT NULL,
      promocode VARCHAR(50),
      attempt_date TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE attempts;
-- +goose StatementEnd
