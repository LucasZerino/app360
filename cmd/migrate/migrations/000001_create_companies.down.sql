DROP TRIGGER IF EXISTS set_updated_at_on_companies ON companies;
DROP FUNCTION IF EXISTS update_companies_updated_at();
DROP TABLE IF EXISTS companies;
DROP EXTENSION IF EXISTS citext; 