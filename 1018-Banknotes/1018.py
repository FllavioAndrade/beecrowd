money = int(input())
banknotes = [100,50,20,10,5,2,1]
print(money)
for i in range(len(banknotes)):
        notes = money // banknotes[i]
        money = (money % banknotes[i])
        print( "{} nota(s) de R$ {},00".format(notes,banknotes[i]))
