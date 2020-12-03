import fileinput 

slopes = [(7,1)] 

G = [] 

for line in fileinput.input():
    G.append(list(line.strip())) 


ans = 1
print(len(G))
for (dc, dr) in slopes:
     r = 0
     c = 0 


     score = 0 
     n = 0

     
     while r+1 < len(G):
         c += dc  
         r += dr  


         if G[r][c % len(G[r])] == '#':
             score += 1 
             print(score, r,c % len(G[r]))

         n = n +1  

     print("n count ",n)   
     ans *= score 

print(ans)               
