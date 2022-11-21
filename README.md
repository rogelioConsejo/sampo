# SAMPO
# Asset
We have two types of asset: the `item` and the `sku`.

The `item` is a unique non-fungible asset.

The `sku` is a cataloged fungible asset.
## Item
### Tag
A tag allows to append information to an item, like a Price or some extra information.
#### Tagger
Generates a `Tag` and an automatically generated `Id` and QR code for a specific, identified `Item`. 

The QR code should be printed and added to the item to keep track of it without the need to add any other details
(making the process as streamlined as possible).

Optionally, other details can be added to the tag. This should be used through some sort of configuration to allow you 
to dynamically add extra fields to the tag and retrieve them as needed from the consumer component.

## SKU
### Inventory
Track the existing `Stock` from any amount of items in an `Inventory`.

Items are identified by an `Id` string.

Amounts cannot be negative. You cannot have, add or remove a negative amount. Amounts are always int, only natural 
numbers can be used.

For smaller amounts (that are smaller than one unit), use smaller units (if you need to measure 1/2 a cm don't use the 
cm as a unit, in this case use mm, for example).

## Container
A container contains items and SKUs. An instance of a container is separate from other instances of containers.

A container can also contain other containers.