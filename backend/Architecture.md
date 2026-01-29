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

1. Batch update in DB, followed by a Redis pipeline update.
2. Since DB is the source of truth and Redis is eventually consistent, Redis can always be rebuilt from DB.
3. Some future improvements to make the system more scalable:
   - Exponential backoff for DB updates.
   - Asynchronous Redis updates using a queue.
   - Reconciliation logic for fixing inconsistencies.
