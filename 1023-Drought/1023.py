
def sort(person_in_resident,consume):
    for i in range(0,len(person_in_resident)-1,1):
        if consume[i] / person_in_resident[i] > consume[i+1] / person_in_resident[i+1]:
            temp_con = consume[i]
            temp_per = person_in_resident[i]
            consume[i] = consume[i+1]
            person_in_resident[i] = person_in_resident[i+1]
            consume[i+1] = temp_con
            person_in_resident[i+1] = temp_per
   
    person_in_resident1 = []
    consume1 = []
   
    i = 0
    qtd = 0
    print(person_in_resident)
    while i < len(person_in_resident):
        print(i)
        qtd = i
        while (consume[qtd] // person_in_resident[qtd] == consume[qtd+1] //person_in_resident[qtd+1]) and qtd < len(consume)-2: 
            qtd = qtd+1
    #        print(qtd,consume[i] // person_in_resident[i] == consume[i+1] //person_in_resident[i+1])
            print(qtd)
  #      for j in range(i,qtd,1):

      #      print('a',j)
      #  print(person_in_resident1)
    #    consume1.append(consume[qtd+i-1])
     #   i = i + qtd
   #     qtd = 1
    #    print (consume1, person_in_resident1, i)

    #return consume1, person_in_resident1
        i = i+1
    return 0

N = int(input())
city = 1
while (N != 0):
    
    person_in_resident = []
    consume = []

    for i in range(N):
        res, con = map(int, input().split())
        person_in_resident.append(res)
        consume.append(con)

    cons_per_person = sum(consume) / sum(person_in_resident)

    print("Cidade# {}:".format(city))

    
    b = sort(person_in_resident, consume)
   # for i in range(len(consume)):
   #     print("{:d}-{:d}".format(a[i],a[i] //b[i]),end=' ')
    
    print("\nConsumo medio: {:.2f} m3.".format(cons_per_person))
    
    N = int(input())
    city = city+1