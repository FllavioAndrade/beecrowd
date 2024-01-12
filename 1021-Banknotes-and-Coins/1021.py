banknotes = [100,50,20,10,5,2]
bankcoins = [100,50,25,10,5,1]

money = float(input())
print("NOTAS:")
for i in range(len(banknotes)):
    number = int(money) // banknotes[i]
    print("{:d} nota(s) de R$ {:.2f}".format(number, banknotes[i]))
    money = money % banknotes[i]
print("MOEDAS:")
money *= 100
for i in range(len(bankcoins)):
    number = int(money) // bankcoins[i]
    print("{:d} moeda(s) de R$ {:.2f}".format(number, bankcoins[i]/100))
    money = money % bankcoins[i]