
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
   
    for i in range(len(consume)):
      __doc__  




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
    for i in range(len(consume)):
        print("{:d}-{:d}".format(person_in_resident[i],consume[i] //person_in_resident[i]),end=' ')
    
    print("\nConsumo medio: {:.2f} m3.".format(cons_per_person))
    
    N = int(input())
    city = city+1