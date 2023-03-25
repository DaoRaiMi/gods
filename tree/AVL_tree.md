# <center>AVL Tree

## 定义
平衡二叉树(Self-Balancing Binary Search Tree 或 Height-Balanced Binary Search Tree),
是一种二叉排序树，其中每个结点的左右子树的高度差不超过1

有两位俄罗斯数学家 **G.M.Adelson-Velskii** 和 **E.M.Landis** 在1962年共同发明了一种解决平衡二叉树的算法，
所以也可以称这样的平衡二叉树为**AVL树**

从平衡二叉树的英文名称就可以看出，它是一种高度(Height)平衡的二叉排序树。

高度平衡：
> 要么它是一棵空树，要么它的左右子树都是平衡二叉树;且左右子树的高度之差的绝对值不超过1
> 

平衡因子(Balance Factor)：
> 结点左子树的高度减去右子树的高度的值,取值范围是[-1,0,1]。只要树上有一个结点的平衡因子的绝对值大于1,
> 那么该树就不是平衡的。