# Details

Date : 2024-04-11 14:43:19

Directory g:\\goProj\\goalboard-collab

Total : 78 files,  6204 codes, 218 comments, 555 blanks, all 6977 lines

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)

## Files
| filename | language | code | comment | blank | total |
| :--- | :--- | ---: | ---: | ---: | ---: |
| [README.md](/README.md) | Markdown | 42 | 0 | 29 | 71 |
| [backend/Dockerfile](/backend/Dockerfile) | Docker | 17 | 0 | 14 | 31 |
| [backend/UnitTesting/go.sum](/backend/UnitTesting/go.sum) | Go Checksum File | 3 | 0 | 1 | 4 |
| [backend/UnitTesting/unitTest_test.go](/backend/UnitTesting/unitTest_test.go) | Go | 115 | 26 | 34 | 175 |
| [backend/auth/auth.go](/backend/auth/auth.go) | Go | 55 | 4 | 17 | 76 |
| [backend/auth/mod.go](/backend/auth/mod.go) | Go | 1 | 0 | 1 | 2 |
| [backend/buildtool/main.go](/backend/buildtool/main.go) | Go | 1 | 19 | 3 | 23 |
| [backend/buildtool/runbuild/build.go](/backend/buildtool/runbuild/build.go) | Go | 0 | 55 | 20 | 75 |
| [backend/database/migrations/000001_create_user_table.down.sql](/backend/database/migrations/000001_create_user_table.down.sql) | SQL | 3 | 0 | 2 | 5 |
| [backend/database/migrations/000001_create_user_table.up.sql](/backend/database/migrations/000001_create_user_table.up.sql) | SQL | 7 | 0 | 3 | 10 |
| [backend/database/migrations/000002_create_goals_table.down.sql](/backend/database/migrations/000002_create_goals_table.down.sql) | SQL | 3 | 0 | 3 | 6 |
| [backend/database/migrations/000002_create_goals_table.up.sql](/backend/database/migrations/000002_create_goals_table.up.sql) | SQL | 12 | 0 | 3 | 15 |
| [backend/database/migrations/000003_create_activity_table.down.sql](/backend/database/migrations/000003_create_activity_table.down.sql) | SQL | 3 | 0 | 3 | 6 |
| [backend/database/migrations/000003_create_activity_table.up.sql](/backend/database/migrations/000003_create_activity_table.up.sql) | SQL | 10 | 0 | 4 | 14 |
| [backend/database/mod.go](/backend/database/mod.go) | Go | 1 | 0 | 0 | 1 |
| [backend/database/psql.go](/backend/database/psql.go) | Go | 32 | 1 | 17 | 50 |
| [backend/env/mod.go](/backend/env/mod.go) | Go | 1 | 0 | 1 | 2 |
| [backend/env/readfile.go](/backend/env/readfile.go) | Go | 29 | 2 | 7 | 38 |
| [backend/go.mod](/backend/go.mod) | Go Module File | 14 | 0 | 4 | 18 |
| [backend/go.sum](/backend/go.sum) | Go Checksum File | 18 | 0 | 1 | 19 |
| [backend/main.go](/backend/main.go) | Go | 31 | 3 | 13 | 47 |
| [backend/models/dbQueryModels.go](/backend/models/dbQueryModels.go) | Go | 14 | 5 | 6 | 25 |
| [backend/models/goal/goal_model.go](/backend/models/goal/goal_model.go) | Go | 48 | 0 | 19 | 67 |
| [backend/models/goal/mod.go](/backend/models/goal/mod.go) | Go | 1 | 0 | 1 | 2 |
| [backend/models/mod.go](/backend/models/mod.go) | Go | 1 | 0 | 1 | 2 |
| [backend/models/user/mod.go](/backend/models/user/mod.go) | Go | 1 | 0 | 1 | 2 |
| [backend/models/user/user.go](/backend/models/user/user.go) | Go | 37 | 0 | 16 | 53 |
| [backend/pw/check.go](/backend/pw/check.go) | Go | 12 | 1 | 3 | 16 |
| [backend/pw/mod.go](/backend/pw/mod.go) | Go | 1 | 0 | 1 | 2 |
| [backend/router/main.go](/backend/router/main.go) | Go | 93 | 26 | 27 | 146 |
| [backend/router/mod.go](/backend/router/mod.go) | Go | 1 | 0 | 1 | 2 |
| [backend/routes/activities.go](/backend/routes/activities.go) | Go | 42 | 0 | 14 | 56 |
| [backend/routes/goals.go](/backend/routes/goals.go) | Go | 156 | 9 | 42 | 207 |
| [backend/routes/healthcheck.go](/backend/routes/healthcheck.go) | Go | 15 | 0 | 6 | 21 |
| [backend/routes/login.go](/backend/routes/login.go) | Go | 70 | 6 | 18 | 94 |
| [backend/routes/logout.go](/backend/routes/logout.go) | Go | 17 | 1 | 4 | 22 |
| [backend/routes/mod.go](/backend/routes/mod.go) | Go | 1 | 0 | 1 | 2 |
| [backend/routes/register.go](/backend/routes/register.go) | Go | 80 | 3 | 26 | 109 |
| [backend/utils/cors.go](/backend/utils/cors.go) | Go | 5 | 0 | 3 | 8 |
| [backend/utils/ctx.go](/backend/utils/ctx.go) | Go | 2 | 0 | 2 | 4 |
| [backend/utils/invariant.go](/backend/utils/invariant.go) | Go | 11 | 0 | 3 | 14 |
| [backend/utils/jwt.go](/backend/utils/jwt.go) | Go | 14 | 0 | 7 | 21 |
| [backend/utils/logger.go](/backend/utils/logger.go) | Go | 16 | 5 | 8 | 29 |
| [backend/utils/mod.go](/backend/utils/mod.go) | Go | 1 | 0 | 1 | 2 |
| [compose.yaml](/compose.yaml) | YAML | 62 | 0 | 7 | 69 |
| [frontend/.eslintrc.cjs](/frontend/.eslintrc.cjs) | JavaScript | 18 | 0 | 1 | 19 |
| [frontend/Dockerfile](/frontend/Dockerfile) | Docker | 16 | 1 | 10 | 27 |
| [frontend/README.md](/frontend/README.md) | Markdown | 22 | 0 | 9 | 31 |
| [frontend/index.html](/frontend/index.html) | HTML | 13 | 0 | 1 | 14 |
| [frontend/package-lock.json](/frontend/package-lock.json) | JSON | 4,261 | 0 | 1 | 4,262 |
| [frontend/package.json](/frontend/package.json) | JSON | 33 | 0 | 1 | 34 |
| [frontend/postcss.config.js](/frontend/postcss.config.js) | JavaScript | 6 | 0 | 1 | 7 |
| [frontend/public/vite.svg](/frontend/public/vite.svg) | XML | 1 | 0 | 0 | 1 |
| [frontend/src/Home.tsx](/frontend/src/Home.tsx) | TypeScript JSX | 41 | 0 | 3 | 44 |
| [frontend/src/Login.tsx](/frontend/src/Login.tsx) | TypeScript JSX | 69 | 22 | 15 | 106 |
| [frontend/src/Register.tsx](/frontend/src/Register.tsx) | TypeScript JSX | 97 | 15 | 20 | 132 |
| [frontend/src/UserActivitiesNew.tsx](/frontend/src/UserActivitiesNew.tsx) | TypeScript JSX | 94 | 1 | 13 | 108 |
| [frontend/src/UserGoals.tsx](/frontend/src/UserGoals.tsx) | TypeScript JSX | 36 | 8 | 5 | 49 |
| [frontend/src/UserGoalsNew.tsx](/frontend/src/UserGoalsNew.tsx) | TypeScript JSX | 97 | 0 | 14 | 111 |
| [frontend/src/assets/react.svg](/frontend/src/assets/react.svg) | XML | 1 | 0 | 0 | 1 |
| [frontend/src/components/Button.tsx](/frontend/src/components/Button.tsx) | TypeScript JSX | 21 | 0 | 5 | 26 |
| [frontend/src/components/Heading.tsx](/frontend/src/components/Heading.tsx) | TypeScript JSX | 27 | 0 | 5 | 32 |
| [frontend/src/components/Input.tsx](/frontend/src/components/Input.tsx) | TypeScript JSX | 24 | 0 | 5 | 29 |
| [frontend/src/components/Nav.tsx](/frontend/src/components/Nav.tsx) | TypeScript JSX | 12 | 0 | 2 | 14 |
| [frontend/src/components/Select.tsx](/frontend/src/components/Select.tsx) | TypeScript JSX | 24 | 0 | 5 | 29 |
| [frontend/src/constants.ts](/frontend/src/constants.ts) | TypeScript | 1 | 0 | 1 | 2 |
| [frontend/src/index.css](/frontend/src/index.css) | CSS | 3 | 0 | 1 | 4 |
| [frontend/src/main.tsx](/frontend/src/main.tsx) | TypeScript JSX | 46 | 0 | 6 | 52 |
| [frontend/src/utils/activity.ts](/frontend/src/utils/activity.ts) | TypeScript | 19 | 0 | 4 | 23 |
| [frontend/src/utils/goal.ts](/frontend/src/utils/goal.ts) | TypeScript | 21 | 0 | 7 | 28 |
| [frontend/src/utils/loader-data.ts](/frontend/src/utils/loader-data.ts) | TypeScript | 6 | 0 | 2 | 8 |
| [frontend/src/utils/user.ts](/frontend/src/utils/user.ts) | TypeScript | 43 | 0 | 11 | 54 |
| [frontend/src/vite-env.d.ts](/frontend/src/vite-env.d.ts) | TypeScript | 0 | 1 | 1 | 2 |
| [frontend/tailwind.config.js](/frontend/tailwind.config.js) | JavaScript | 11 | 1 | 1 | 13 |
| [frontend/tsconfig.json](/frontend/tsconfig.json) | JSON with Comments | 21 | 2 | 3 | 26 |
| [frontend/tsconfig.node.json](/frontend/tsconfig.node.json) | JSON | 10 | 0 | 1 | 11 |
| [frontend/vite.config.ts](/frontend/vite.config.ts) | TypeScript | 5 | 1 | 2 | 8 |
| [package-lock.json](/package-lock.json) | JSON | 6 | 0 | 1 | 7 |

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)