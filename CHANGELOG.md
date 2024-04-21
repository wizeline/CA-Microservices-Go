# Change Log

## [v0.1.0] - 2024-04-21
- Initial release of `CAMGO`.
- Added basic functionality.
- Clean Architecture layers: 
    - Entity 
    - Controller with handling errors unit and auto-set routes.
    - Service
    - Repository
- Handy `deployments` through make statements.
    - `Docker` Containerized environment
    - Application: build, run, stop, remove, and clean
    - Mocks generation
    - Postgresql database
    - PgAdmin database GUI
- GitHub Actions:
    - Workflows:
        - CI
        - Release
    - Dependabot
- Implemented database migrations.
- `Viper` configuration through environment variables.
- `PostgreSQL` database support.
- `ZeroLog` logger support
- `Chi` router support
- `Mockery` automate mocks generations.
