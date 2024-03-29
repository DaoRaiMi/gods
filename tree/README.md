# <center> 树
## 定义
树是(N, N>=0)个结点的有限集合。当N=0时为**空树**。在任意一个非空树中有：
1. 有且仅有一个特定的称为根(Root)的结点;
2. 当N>1时，剩余的结点可以分为(M, M>0)个互不相交的集合，其中第一个集合又是一棵树，并且称之为根树的子树(SubTree)

## 结点
在树中结点包含一个数据元素以及若干个指向其子树的分支。

### 结点的度
结点拥有的子树数(分支数)称为结点的**度(Degree)**

度为0的结点称为叶子(Leaf)结点或终端结点，度不为0的结点称为分支结点或非终端结点。

除了根结点以外，分支结点也称为内部结点。

树的度是树内各结点的度的最大值。

### 结点的层次
结点的层次(Level)从根开始，根为第一层，根的孩子结点为第二层，依此类推。

树中结点的最大层次称为树的深度(Depth)或高度(Height)。

若将树中结点的各子树看成从左到右是有次序的，不能互换的，则称该树为有序树，否则称为无序树。


## 树的存储结构
* 双亲表示法
  > 使用一组连续的空间来存储树的结点，同时在每个结点中，附加一个指针，该指针指向其双亲结点
  在数组中的下标。
  > 
  > 双亲表示法还可以基于业务需求来进行扩展。比如，可以在结点中加一个兄弟指针。
* 孩子表示法
  > 把每个结点的直接子结点按从左到右的顺序排序，再用单链表作为存储结构，则N个结点就会有N个孩子链表，
  > 如果是叶子结点则此单链表为空。然后这N个头指针，其实也就是这N个结点，又组一个线性表。
  >
  > 对于这种结构，如果还想要知道结点的双亲，则只需要在结点中再添加一个双亲域即可。此种方式
  > 又叫作“双亲孩子表示法”
* 孩子兄弟表示法
  > 对于任意一棵树，它的结点的第一个孩子如果存在就是唯一的，这个结点的右兄弟结点
  > 如果存在那么也是唯一的。因此，我们设置两个指针，分别指向该结点的第一个孩子和此结点的右兄弟。
  > 
  > 这种表示方法的最大好处就是把一个复杂的树变成了一棵二叉树，然后就可以利用二叉树的特性和算法来处理了。 

## 常见树
* [二叉树](./binary_tree.md)

* [二叉排序树](./binary_sort_tree.md)

* [AVL树](./AVL_tree.md)

* [字典树](./trie_tree.md)