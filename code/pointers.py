a = 1
print(id(a)) #10964896
b=a
print(id(b)) #10964896
a = 2
print(b)     #1
print(id(a)) #10964928
print(id(b)) #10964896
