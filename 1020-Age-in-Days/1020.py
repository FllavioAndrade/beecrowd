days = int(input())

period = days // 365
print("{:d} ano(s)".format(period))
days = days % 365
period = days // 30
print("{:d} mes(es)".format(period))
days = days % 30
print("{:d} dia(s)".format(days))