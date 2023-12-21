-- name: insert:company
INSERT OR IGNORE INTO
  B_COMPANY (company_name)
  values (?);

-- name: select:company
SELECT c.company_id
  FROM B_COMPANY c
  WHERE c.company_name = ? COLLATE NOCASE;
