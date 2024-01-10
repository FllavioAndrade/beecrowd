time = int(input())

hours = time // 3600
time = time % 3600

minutes = time // 60
time = time % 60

secounds = time % 60

print("{:d}:{:d}:{:d}".format(hours,minutes,secounds))