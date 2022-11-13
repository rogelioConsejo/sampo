# sampo

## Tagger
Generates a `Tag` and an automatically generated `Id` and QR code for a specific, identified `Item`. 

The QR code should be printed and added to the item to keep track of it without the need to add any other details
(making the process as streamlined as possible).

Optionally, other details can be added to the tag. This should be used through some sort of configuration to allow you 
to dynamically add extra fields to the tag and retrieve them as needed from the consumer component.

### Inventory
Track the existing `Stock` from any amount of items in an `Inventory`.

Items are identified by an `Id` string.

Amounts cannot be negative. You cannot have, add or remove a negative amount. Amounts are always int, only natural 
numbers can be used.

For smaller amounts (that are smaller than one unit), use smaller units (if you need to measure 1/2 a cm don't use the 
cm as a unit, in this case use mm, for example).

