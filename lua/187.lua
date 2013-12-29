function factors (n)
    prime_fact = {}
    d = 2
    while d*d <= n do
        while (n % d) == 0 do
            table.insert(prime_fact, d) -- supposing you want multiple factors repeated
            n = n / d
        end
        d = d + 1
    end
    if n > 1 then
        table.insert(prime_fact, n)
    end
    return prime_fact
end

print("enter the upper limit:")
a = io.read("*number") -- upper limit
i = 4 -- first composite
j = 0 -- counter
while i < a do
    b = factors(i)
    if #b == 2 then
        j = j + 1
    end
    if i % 1000 == 0 then
        print(i)
    end
    i = i + 1
end

print("answer:",j)
