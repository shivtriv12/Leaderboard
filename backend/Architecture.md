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

## Tie-Breaking Mechanism

- I am choosing **1224(Standard Competition Ranking)** to rank the users.
- It gives better picture about where actually the user stands and is also used in olympics.
- 1223 Dense Rankings was another choice but may not be good for a app where users are coming to learn(quiz app - what i am keeping in mind while building).
- Also in future with more parameters such as time to complete, section wise scores, tie breaking mechanism could be enhanced.

## User Search (globally sorted by name)

- Text search happens from sql in `O(logn)` time as username is indexed using `pg_trgm`.
- Then rank is calculated from the result returned from db using redis in O(log n) time.
- Overall time would be O(logn) as pagination would be implemented.
- I am using cursor based pagination to display the users as:-
  - It would be much fast compare to offset based in scale(it would directly jump to user rather than skipping n items).
  - No duplicates and consistent scrolling.
  - Better for infinite scroll ui.

## Leaderboard

- Instead of calculating rank for every user individually, I anchor the first user's rank using ZCount and mathematically infer the rest.
- This reduces the operation to exactly **3 Redis calls (ZRevRank + ZRevRange + ZCount)** regardless of the page size, maintaining **log N** complexity.
- Pagination and Tie Ranking Mechanism remains the same.
