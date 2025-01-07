DROP TRIGGER IF EXISTS set_updated_at_on_user_companies ON user_companies;
DROP FUNCTION IF EXISTS update_user_companies_updated_at();
DROP TABLE IF EXISTS user_companies; 