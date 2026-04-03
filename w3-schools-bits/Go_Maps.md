## GO MAPS
Go maps are built-in data structures used to store data as unordered, mutable collections of key-value pairs, where each key is unique. They act as references to underlying hash tables and have a default, uninitialized value of nil. The total number of elements in a map can be determined using the len() function. 

### Key Characteristics

* Unordered: Maps do not maintain the order in which elements are added.
* Unique Keys: Duplicate keys are not allowed.
* Dynamic: Elements can be added or removed efficiently. 

### Creating Maps
Go offers several methods to create maps:

* Map Literals (var or :=): Used to declare and initialize a map with predefined key-value pairs simultaneously.
* make() Function: Recommended for creating an empty map, allowing for dynamic addition of elements later. 

### Note on Order
The order in which map elements are defined is often different from the order in which they are stored and displayed. This is a deliberate design to ensure efficient data retrieval. 