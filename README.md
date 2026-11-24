# LRU Cache

> Striver SDE Sheet 083 · Day 14 Stack & Queue II — hash map for lookup + doubly linked list for recency, sentinels killing every nil check; all O(1). Go, stdlib only, tested.

## Problem

A fixed-capacity key-value cache where Get and Put run in O(1) and the **l**east **r**ecently **u**sed entry is evicted when space runs out. The single most-asked design problem in interviews.

## Approach

Two structures covering each other's weakness: a hash map finds a node by key instantly but knows no order; a doubly linked list keeps perfect recency order (front = most recent, back = eviction candidate) but can't be searched. Every operation is "look up in the map, splice in the list": Get unlinks and re-fronts the node; Put updates-and-fronts, or inserts after evicting the tail's neighbor. Two sentinel nodes (permanent head and tail) mean splicing never checks for nil ends — the detail that makes the implementation short enough to write under pressure. The map stores *node pointers*, and nodes carry their key so eviction can delete the map entry.

**Complexity:** Get/Put O(1); O(capacity) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
