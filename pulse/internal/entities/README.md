# Entities

Think of this as entities from the perspective of CLEAN architecture. 
> Entities represent business logic. They also represent relationships between entities.

### Ground rules:
- Entities must never depend on anything but entities. Reason being that business logic should not depend on implementation of anything.
- They must be unit tested at all times.