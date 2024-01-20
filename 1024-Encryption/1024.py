n = int(input())

for i in range(n):
    message = input()
    
    message[0] = (message[0])
    encrypt = ""
    print(message)
    for j in range(len(message),0, -1):
        encrypt += message[j]

    print(encrypt) 