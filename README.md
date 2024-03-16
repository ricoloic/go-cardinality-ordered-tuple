# Calculating the cadinality of ordered n-tuple

Given a tuple of 4 elements, if each of the elements can have the same amount of different position as there is elements in the tuple. We can calculate the cardinality of the total possible count of different tuples when there elements are positionned in order.

```
aaaa	a	a	a	a
aaab	a	a	a	b
aaac	a	a	a	c
aaad	a	a	a	d
				
aabb	a	a	b	b
aabc	a	a	b	c
aabd	a	a	b	d
				
aacc	a	a	c	c
aacd	a	a	c	d
				
aadd	a	a	d	d
				
abbb	a	b	b	b
abbc	a	b	b	c
abbd	a	b	b	d
				
abcc	a	b	c	c
abcd	a	b	c	d
				
abdd	a	b	d	d
				
accc	a	c	c	c
accd	a	c	c	d
				
acdd	a	c	d	d
				
addd	a	d	d	d
				
				
bbbb	b	b	b	b
bbbc	b	b	b	c
bbbd	b	b	b	d
				
bbcc	b	b	c	c
bbcd	b	b	c	d
				
bbdd	b	b	d	d
				
bccc	b	c	c	c
bccd	b	c	c	d
				
bcdd	b	c	d	d
				
bddd	b	d	d	d
				
				
cccc	c	c	c	c
cccd	c	c	c	d
				
ccdd	c	c	d	d
				
cddd	c	d	d	d
				
				
dddd	d	d	d	d
```

In this case we can count the number of rows present and notice that with a tuple of 4 elements we have a cardinality of 4.
