Short answer: use one workflow.

  Recommended (migration‑based, for real deployments)

  - Update GORM models.
  - Generate migration file:
    atlas migrate diff add_<change> --env gorm
  - Apply it to the DB (this updates schema and records the migration):
    atlas migrate apply --env gorm --url "$DATABASE_URL"

  Quick sync (no migration history)

  - Just make the DB match models:
    atlas schema apply --env gorm --url "$DATABASE_URL"

  So yes, if you want both “update schema” and “create a migration file,” the right plan is migrate diff then migrate apply. Don’t mix that
  with schema apply on the same DB.
