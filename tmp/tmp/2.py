n = int(input())
a, b = 1, 1
total = 0.0
for i in range(1, n + 1):
    if i == 1:
        fib = 1
    else:
        a, b = b, a + b
        fib = b
    term = ((-1) ** (i + 1)) * (i / fib)
    total += term
print("{:.6f}".format(total))  