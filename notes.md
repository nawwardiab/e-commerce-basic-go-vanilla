1. Big picture:
   1. Packages:
      - internal/config – load DB credential
      - internal/db – open SQL connection
      - internal/model/user.go – User struct, NewUser constructor
      - internal/repository/user_repo.go – query logic
      - internal/service/user_service.go – business rules/logic
      - internal/handler/auth.go – HTTP endpoint handler
      - templates/register.tpl – frontend form

```bash
Data flow:

[config] → [db] → [repo/user_repo] ←→ [service/user_service] → [handler/auth.Register]
                                             ↓
                                         [templates/register.tpl]

```

2. Zoom-in:
   1. start with handler
   2. service
   3. repo
   4. re-zoom-out [validate]: AuthHandler only depend on UserService? config and DB setup shared (main.go, config–>db–>repos–>services–>handlers)?
   5. next slice!

---
