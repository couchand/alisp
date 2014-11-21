(define fib-kernel (lambda (f n)
    (if
        (+ (eq? 0 n) (eq? 1 n))
        1
        (+ (f f (+ n -1)) (f f (+ n -2)))
    )
))
(define fib (lambda (n)
    (fib-kernel fib-kernel n)
))
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

(select fib (range 30))
