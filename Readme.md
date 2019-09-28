
# Developing

## Prerequisites

  * [Golang development environment](https://golang.org/dl/)
  * Postgres 10 server
 
For data model changes you'll need:
  * [bindata](https://github.com/kevinburke/go-bindata) installed and in your path
  * [golang-migrate cli tool](https://github.com/golang-migrate/migrate/releases) installed and in your path
  * [sqlboiler](https://github.com/volatiletech/sqlboiler#download) installed
  
Also recommended is [fresh](https://github.com/gravityblast/fresh) for automatic recompiling.

## Downloading 

## Config file setup

  1. Copy `sqlboiler.sample.toml.sample` and `config.sample.json` to new files - removing the `.sample`
  2. You'll need to at least change the Database Name (`<DATABASE>`), DB User (`<DBUSER>`) and DB Password (`<DBPASS>`) in those files to a suitable config for your system.
     * It's a good idea to use a unique database and user and a random password to avoid issues with your dev environment
  3. You may need to change the host and port entries to reflect your local setup. The example fields in `config.sample.json` are the defaults. 
  4. `config.sample.json` is used by the live application, while `sqlboiler.sample.toml` is only used to generate new datamodels from database data.
  5. the root "development" flag is only used to perform migrations from files.  
  
## Database Setup

  1. You'll need to have a postgres install to connect to, and the ability to log in as the  `postgres` user.
  2. Create the login role from the settings you did above: 
     * `CREATE ROLE "<DBUSER>" WITH PASSWORD '<DBPASS>' LOGIN` (the difference in quoting is intentional)
  3. Create the database via:
     * `CREATE DATABASE "<DBNAME>" OWNER "<DBUSER>"`
  4. The application should create any appropriate database fields on bootup!

## Changing data models

  1. Ensure you have the development flag turned on to test your migration scripts
  2. Create a migration via 
      1. `create_migration.sh` (on linux or macOS)
         1. this takes a single parameter - the name of the migration. a 
      2. running `create_migration.ps1` in powershell 
         1. this will ask you the name of the migration when you run it.
  3. Specify the migration syntax in the new files created in the `datasource/migrations/` directory
  4. Test your migration via either:
      1. compiling and running the server with the development flag turned on
      2. manually running the migration via `migrate -source file://datasource/migrations -database postgres://<DBUSER>:<DBPASS>@localhost:15432/<DBNAME>?sslmode=disable up`
  5. Once your migration has the correct SQL, regenerate the migration bindata
     1. Ensure you have [bindata](https://github.com/kevinburke/go-bindata) installed and in your path
     2. Run the following from the project root: `go-bindata -o datasource\migrationData\main.go -pkg migrationData -prefix datasource\migrations\ datasource\migrations\`
  6. Then generate the models:
     1. Ensure you have a configured `sqlboiler.sample.toml` file
     2. Run `sqlboiler psql -o datamodels_raw -p datamodels_raw` from the root of the repo