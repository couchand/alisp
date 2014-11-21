(define range-kernel (lambda (f n i)
    (if
        (eq? n i)
        nil
        (cons i (f f n (+ i 1)))
    )
))
(define range (lambda (n)
    (range-kernel range-kernel n 0)
))
(define select-kernel (lambda (f s l)
    (if
        (nil? l)
        nil
        (cons
            (s (car l))
            (f f s (cdr l))
        )
    )
))
(define select (lambda (s l)
    (select-kernel select-kernel s l)
))
(define sieve-finish (lambda (s)
    s
))
(define mark-primes-kernel (lambda (f s p i)
    (if
        (nil? s)
        nil
        (if
            (* (eq? 0 (% i p)) (eq? 0 (eq? i p)))
            cons(
                0
                (f f (cdr s) p (+ 1 i))
            )
            cons(
                (car s)
                (f f (cdr s) p (+ 1 i))
            )
        )
    )
))
(define mark-primes (lambda (s p)
    (mark-primes-kernel mark-primes-kernel s p (+ p 1))
))
(define next-prime (lambda (s p)
    (+ 1 (len s))
))
(define sieve-kernel (lambda (f s p)
    (if
        (eq? p (+ 1 (len s)))
        (sieve-finish s)
        (f f (mark-primes s p) (next-prime s p))
    )
))
(define sieve-initial (lambda (r)
    (select (lambda (x) 1) (range r))
))
(define sieve (lambda (r)
    (sieve-kernel sieve-kernel s l (sieve-initial r) 1)
))

(sieve 2)
(sieve-initial 10)
