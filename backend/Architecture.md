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
