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
(define sieve-initial (lambda (r)
    (select (lambda (x) 1) (range r))
))

(sieve-initial 10)
