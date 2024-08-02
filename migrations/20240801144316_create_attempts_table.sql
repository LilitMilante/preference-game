-- +goose Up
-- +goose StatementBegin
CREATE TABLE attempts (
      id UUID PRIMARY KEY,
      user_id TEXT NOT NULL,
      first_card_suit TEXT NOT NULL,
      first_card_value TEXT NOT NULL,
      second_card_suit TEXT NOT NULL,
      second_card_value TEXT NOT NULL,
      is_win BOOLEAN,
      promo_code TEXT,
      created_at TIMESTAMP NOT NULL,
      updated_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE attempts;
-- +goose StatementEnd
