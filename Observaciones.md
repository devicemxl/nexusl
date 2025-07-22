# To Do

1. the token, lexer and parser have a extreme complexity, review the token, lexer and parser of Go in order to decrease complexity and increase modularity. If it is possible.
    - /usr/local/go/src/go/token/
2. I see a list of super-internal Golang Implementations, it could be very beneficial to improve the nexusL implementation:
    - biteAlg:
        - Rabin-Karp algorithm for bit ops
        - /usr/local/go/src/internal/bytealg/bytealg.go
    - arena:
        - provides the ability to allocate memory for a collection of Go values and free that space manually all at once, safely. The purpose of this functionality is to improve efficiency: manually freeing memory before a garbage collection delays that cycle. Less frequent cycles means the CPU cost of the garbage collector is incurred less frequently.
        - /usr/local/go/src/arena/arena.go
        - useful to build a "rustic" memory management.
    - sync:
        - concurrency
        - /usr/local/go/src/sync/
    - semaphore:
        - sync semaphore
        - /usr/local/go/src/cmd/vendor/golang.org/x/sync/semaphore/
    - bites:
        - implements functions for the manipulation of byte slices.
        - /usr/local/go/src/bytes/
    - script:
        - implements a small, customizable, platform-agnostic scripting language.
        - /usr/local/go/src/cmd/internal/script/
    - hash:
        - hash functions for the compiler
        - /usr/local/go/src/cmd/internal/hash/
    - heap:
        - is a common way to implement a priority queue.
        - /usr/local/go/src/container/heap/
    - compress:
        - compression tools
        - /usr/local/go/src/compress/
    list:
        - doubly linked list
        - /usr/local/go/src/container/list/
    ring:
        - an element of a circular list, or ring
        - /usr/local/go/src/container/ring/
    crypto:
        - crypto
        - /usr/local/go/src/crypto/
        - /usr/local/go/src/crypto/internal/fips140/subtle/xor_arm64.s

3. implements trunKV with Fermatean Neutrosophic Cognitive Map,
    - qw