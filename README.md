structures
==========

Let us learn some golang via data structures

Questions/notes
---------------

1. Why is the golang linked list unencapsulated, i.e. why are **Element**s
programatically accessible? _Ideas_: Collapses both datum and iterator into
the same struct, to help avoid repeated traversals when accessing the list
in a sequence of operations.

2. Next: a fully encapsulated, singly-linked list with:
    * O(1) Append
    * O(1) Concat (in place, obviously)
    * O(m+n) CopyConcat
    * O(lg n) Search (binary underneath)
    * O(n) Remove (nil if failed)