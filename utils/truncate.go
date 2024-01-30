package utils

import (
	"fmt"

	"gorm.io/gorm"
)

// TruncateTables truncates all tables in the database using PostgreSQL specific query
func TruncateTables(db *gorm.DB) error {
	// Raw SQL query to
	query := fmt.Sprintf("TRUNCATE TABLE %s CASCADE; ALTER SEQUENCE %s_id_seq RESTART WITH 1;", "urls", "urls")

	// Execute the raw SQL query
	if err := db.Exec(query).Error; err != nil {
		return err
	}

	return nil
}

// psql to truncate all tables
// query := `
// 	CREATE OR REPLACE FUNCTION truncate_tables(postgres IN VARCHAR) RETURNS void AS $$
// 	DECLARE
// 		statements CURSOR FOR
// 			SELECT tablename FROM pg_tables
// 			WHERE tableowner = username AND schemaname = 'public';
// 	BEGIN
// 		FOR stmt IN statements LOOP
// 			EXECUTE 'TRUNCATE TABLE ' || quote_ident(stmt.tablename) || ' CASCADE;';
// 		END LOOP;
// 	END;
// 	$$ LANGUAGE plpgsql;

// 	-- Call the function to truncate tables
// 	SELECT truncate_tables(?);
// `
