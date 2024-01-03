code1, unit1, price1 = map(float, input().split())
codE2, unit2, price2 = map(float, input().split())

total = unit1* price1 + unit2*price2

print("VALOR A PAGAR: R$ {:.2f}".format(total))
