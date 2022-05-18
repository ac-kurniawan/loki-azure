# Loki Azure

## How to run
1. Create `.env` file, you can refer to `env.example` to know the parameters needed
2. To migrate all model: `make migration`, please add `uuid` extension in your db using this query `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`
3. To start apps: `make run` (server should be run in :3222 by default)
