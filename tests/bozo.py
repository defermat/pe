import random

arrays = [[1,2,4,3],[1,4,2,3],[1,4,3,2],[1,3,4,2],[1,3,2,4],[2,1,3,4],[2,1,4,3],[2,3,1,4],[2,3,4,1],[2,4,1,3],[2,4,3,1],[3,1,2,4],[3,1,4,2],[3,2,1,4],[3,2,4,1],[3,4,1,2],[3,4,2,1],[4,1,2,3],[4,1,3,2],[4,2,1,3],[4,2,3,1],[4,3,1,2],[4,3,2,1]]

average = 0
b = 0
while b<23:
	a = arrays[b]
	isSorted = False
	count = 0;
	print a
	while (not isSorted):
		isSorted = True
		random.seed()
		index1 = random.randint(0,3)
		index2 = random.randint(0,3)
		temp = a[index1]
		a[index1] = a[index2]
		a[index2] = temp
		for i in range(len(a) - 1):
			if(a[i] > a[i+1]):
				isSorted = False
		count+=1
	average += count
	print count
	b += 1

print average,"/24=",average/24.0
