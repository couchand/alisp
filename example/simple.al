(define one 1)
(define two 2)
(define add (lambda (x y) (+ x y)))
(define do-it (lambda (f a b) (f a b)))
(do-it add one two)
