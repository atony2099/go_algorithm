# algorithm


##  references

[Problem Solving with Algorithms and Data Structures using Python](https://runestone.academy/runestone/books/published/pythonds/index.html)
[https://www.codesdope.com/course/algorithms-introduction/](https://www.codesdope.com/course/algorithms-introduction/)


[Coin Change](https://leetcode.com/articles/coin-change/)


![qp7LSd](https://cdn.jsdelivr.net/gh/atony2099/imgs@master/20200408/qp7LSd.jpg)
![I2okpi](https://cdn.jsdelivr.net/gh/atony2099/imgs@master/20200407/I2okpi.jpg)


## singleLinkedList

1. the operatin of  insertion and deletion don't mean to insert at specify index, not including the search operation, so the bigO is 1;


## doubleLinkedList


## LRU cahe

1. hashmap: record the key,value

2. doubleLinkedList: put the rencently reading or writing data in front of the list(first delete,then insert at the top);

   > doubleLinkedList can easily delete and insert; 





## binary search
![](https://www.geeksforgeeks.org/wp-content/uploads/Binary-Search.png)


time complexiy analyze:

1. we need to divide by 2 until with have only one element-
𝑛/(2𝑘)=1

2. we can rewrite it as -

2𝑘=𝑛

by taking log both side, we get

𝑘=log2𝑛




### quick sort

![](https://tse2-mm.cn.bing.net/th?id=OIP.i0MKaJl0e9up7IpRtWb8yAHaGo&pid=Api&rs=1)



## dynamic programing 

divide big problems into small subproblem;

vs: divide and conquer

the subproblemss are interdependent;






## recursion

> breaking a problems into smaller and smaller problems;

feature:
1. algorithm have a base case;
2. algorithm must change its state and move toward the base case.
3. call itselft recursively





### dynamic programing
1. what is?
    an optimization algorithm, optimise rescursion,store the soultion of the subproblems 

2. approach: bottom up vs top up
3. feature
   > solve the prolem withe overlapping subproblems
   

### divide and conquer
1. what is？
e.g.,quick sort, merge sort, binary search


2. featrue？
   > solve the problem indepent
   










