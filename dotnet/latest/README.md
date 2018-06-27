# Profiles Sample - .NET Core

If you need to install dotnet, consider using or learning from [../scripts/setup.sh](../scripts/setup.sh).

To run:

1. `cp .env.tpl .env` and fill in .env with values from `az ad sp
   create-for-rbac`. Comment out any env vars you won't use.
2. `make run`
