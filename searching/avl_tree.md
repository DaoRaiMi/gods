# <center> 平衡二叉搜索树
## 定义
平衡二叉树(Self-Balancing Binary Search Tree)或(Height-Balanced Binary Search Tree)
是一种二叉排序(搜索)树,其中每个结点的左子树与右子树的高度差的绝对值不超过1.

有两位俄罗斯数学家G.M.Adelson-Velskii和E.M.Landis在1962年共同发明了一种解决
平衡二叉树的算法，所以有不少资料中也称这样的平衡二叉树为`AVL树`

## 平衡因子
将二叉树上结点的左子树深度减去右子树深度的值称为：平衡因子BF(Balancing Factor)

所以BF的取值范围只能是-1,0,1.

## 最小不平衡子树
距离插入结点最近的，且以平衡因子的绝对值大于1的结点为根的子树，称为最小不平衡子树。

