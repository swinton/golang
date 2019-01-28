package ch4

// The compiler won't allow cyclical imports (A -> B -> A, or A -> B -> C -> D -> A)

// The compiler will throw *import cycle not allowed*

// Often, this is solved by introducing a package for shared structures
