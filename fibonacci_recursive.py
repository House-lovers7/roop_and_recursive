def fibo(n:int) -> int:
    if n ==1:
        return 1
    elif n == 2:
        return 1
    else:
        answer: int = fibo(n-1) + fibo(n-2)
        #print(f'n:{n}')
        #print(f'answer:{answer}')
        return answer
print(fibo(100))
