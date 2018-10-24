# adminarea-cli

Administrative Area CLI

## Features

* Can create sql table

## Installation

To install the library and command line program, use the following:

```bash
go get -v github.com/ramadani/adminarea-cli/...
```

## Usage

Create `adminarea_config.yml` and fill it with following this values:

```yaml
db:
  driver: mysql
  dsn: user:password@tcp(127.0.0.1:3306)/dbname?parseTime=true
```

### Available Commands

```cmd
setup    Create administrative areas table
country  Save a country
```

### Create SQL table

To create `administrative_areas` table to your database, run this command:

```cmd
adminarea-cli setup
```

### Save the country

```cmd
// adminarea-cli country [COUNTRY_CODE_ID]
adminarea-cli country id
```

## TODO

- [x] Can create db table
- [ ] Config file can pass by argument
- [x] Save a country
- [ ] Save provinces by country
- [ ] Save regencies by country
- [ ] Save districts by country
- [ ] Save villages by country

## License

This library is distributed under the [MIT](LICENSE) license.
