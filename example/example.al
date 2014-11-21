(define three (+ 1 2))
(define times-three (lambda (x) (* x (three))))
(define do-it (lambda (f x) (f x)))
(define do-it-do-it (lambda (f fx) (do-it f (fx))))
(do-it-do-it times-three three)

(times-three 5)
