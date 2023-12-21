-- name: insert:title
INSERT OR IGNORE INTO
  B_TITLE (company_id, title, 'level')
  values (?, ?, ?);
