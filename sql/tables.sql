-- Create album table in default postgres database, add some sample data
--
-- psql -d postgres -U postgres -h 0.0.0.0 -a -f ${SCRIPT_DIR}/tables.sql
--
INSERT INTO albums
   ( title, artist, price)
VALUES
   ('Blue Train', 'John Coltrane', 56.99),
   ('Giant Steps', 'John Coltrane', 63.99),
   ('Jeru', 'Gerry, Mulligan', 17.99),
   ('Sarah Vaughan', 'Sarah Vaughan', 34.99);