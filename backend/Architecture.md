# My Architecture Decisions and Reasons

## Redis: AOF vs RDB vs Build From DB (Persistence)

I have gone with building from DB.

---

### Pros

- No slower performance (every write = disk write).
- No large AOF files over time.
- Cannot lose recent data.

### Cons

- Have to write code for building.
- Will take more time building.

---

## Update Score

I am creating a simulation worker which will run every 10 seconds, pick 1k random users, and randomly change ratings such that  
`0 < ratings <= 10k`.

### Workflow

1. Currently Redis(pipeline)->db(batch), to make it more realtime.
2. Source of truth is db only.
3. In Production, redis first->async queue write to db.
4. Reconciliation logic for fixing inconsistencies.
