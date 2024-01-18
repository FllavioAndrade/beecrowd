qtd = int(input())

def calc (n1,d1,n2,d2,ope):
    if ope == "+":
        numerator = n1*d2 + n2*d1
        denominator = d1*d2

    elif ope == "-":
        numerator  = n1*d2 - n2*d1
        denominator  = d1*d2

    elif ope == "*":
        numerator  = n1*n2
        denominator  = d1*d2

    elif "/":
        numerator  = n1*d2
        denominator  = n2*d1

    numerator1 = numerator
    denominator1 = denominator
    maior = max(denominator,numerator)
    
    for j in range(2,maior,1):
        while numerator1%j == 0 and denominator1 % j == 0:
            numerator1 = int(numerator1/j)
            denominator1 = int(denominator1/j)

    return numerator, denominator, numerator1, denominator1     


for i in range(qtd):
    n1,div,d1,ope,n2,div,d2 = map(str, input().split()) 
    n1 = int(n1)
    d1 = int(d1)
    n2 = int(n2)
    d2 = int(d2)
    
    valor = calc(n1,d1,n2,d2,ope)
    print("{}{}{} = {}{}{}".format((valor[0]),div,(valor[1]),(valor[2]),div,(valor[3])))

