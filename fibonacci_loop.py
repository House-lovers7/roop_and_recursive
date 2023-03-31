from typing import List

def fibo_loop(n: int):
    count: int = 3
    orders: List[int] = list(range(1,n +1))
    fibo_list: List[int] = [1, 1]

    while len(orders) > 2:
        new_fibo_num: int = fibo_list[count-3] + fibo_list[count-2]
        fibo_list.append(new_fibo_num)
        count+= 1
        orders.pop(0)
        # print(fibo_list)
    print(fibo_list[-1])

fibo_loop(2)
