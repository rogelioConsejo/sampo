# sampo

## Inventory
Track the existing `Stock` from any amount of items in an `Inventory`.

Items are identified by an `Id` string.

Amounts cannot be negative. You cannot have, add or remove a negative amount. Amounts are always int, only natural numbers can be used.

For smaller amounts (that are smaller than one unit), use smaller units (if you need to measure 1/2 a cm don't use the cm as a unit, in this case use mm, for example). 