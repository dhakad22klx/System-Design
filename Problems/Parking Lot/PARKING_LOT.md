# Parking Lot

Initial Go models and constructors are implemented:

- **Vehicle:** vehicle number and size (`SMALL`, `MEDIUM`, `LARGE`).
- **Lot:** parking spot ID, size, and lock state; starts unlocked.
- **Level:** level ID and spots grouped by size.
- **ParkingSystem:** collection of levels.

Interfaces define charging, spot access and locking, level access, parking,
unparking, and availability lookup. Their methods are not implemented yet.

Run with `go run .`; currently prints `Parking Lot Problems`.
