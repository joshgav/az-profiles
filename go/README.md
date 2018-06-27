# Profiles Sample - Go

This directory contains a package which calls two subpackages. The `hybrid` subpackage utilizes the 2017-03-09 profile to target Stack. The `latest` subpackage does not use profiles and targets Azure public. Both are called in series in this main package.

To run:

1. `cp .env.tpl .env` and fill in .env with values from `az ad sp
   create-for-rbac`. Comment out any env vars you won't use.
2. `make run`
