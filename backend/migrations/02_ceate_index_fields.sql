-- +goose Up
CREATE INDEX IF NOT EXISTS idx_whitelist_domain ON whitelist (domain);
CREATE INDEX IF NOT EXISTS idx_blacklist_domain ON blacklist (domain);
CREATE INDEX IF NOT EXISTS idx_findedlist_domain ON findedlist (domain);


-- +goose Down
DROP INDEX idx_whitelist_domain;
DROP INDEX idx_blacklist_domain;
DROP INDEX idx_findedlist_domain