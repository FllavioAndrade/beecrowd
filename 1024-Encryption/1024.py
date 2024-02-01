import math

n = int(input())

for i in range (n):
    text = input()
    text = list(text)
    encryption = []
    h = len(text)
    qtd = math.floor(h/2)

    for j in range(h):
        if  (ord(text[j]) >= 65 and ord(text[j]) <= 176):
            text[j] = chr(ord(text[j]) + 3)

    if qtd%2 == 0:
        qtd = qtd-1
        for j in range(h-1,-1,-1):
            encryption.append(text[j])
            if j <= qtd+1:
                encryption[qtd] = (chr(ord(encryption[qtd]) -1))
                qtd = qtd + 1
        
    else:
        for j in range(h-1,-1,-1):
            encryption.append(text[j])
            print(qtd,j)
            if j <= qtd:
                encryption[qtd] = (chr(ord(encryption[qtd]) -1))
                
                qtd = qtd + 1
    
    print(len(encryption))
    text = ''.join(map(str, encryption))
    print(text,end='\n')