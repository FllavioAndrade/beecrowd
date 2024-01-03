import math
a, b, c = map(float, input().split())

pi = 3.14159
triangle = (a * c) / 2.0
circle = pi * math.pow(c, 2)
trapezium = (a + b) / 2.0 * c
square = math.pow(b, 2)
rectangle = a * b

print("TRIANGULO: {:.3f}\nCIRCULO: {:.3f}\nTRAPEZIO: {:.3f}\nQUADRADO: {:.3f}\nRETANGULO: {:.3f}".format( triangle, circle, trapezium, square, rectangle))
