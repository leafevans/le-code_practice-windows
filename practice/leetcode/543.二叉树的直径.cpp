/*
 * @lc app=leetcode.cn id=543 lang=cpp
 *
 * [543] 二叉树的直径
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    int diameterOfBinaryTree(TreeNode* root) {
        int maxDiameter = 0;
        depth(root, maxDiameter);
        return maxDiameter;
    }

private:
    int depth(TreeNode* node, int& maxDiameter) {
        if (!node) {
            return 0;
        }
        int left = depth(node->left, maxDiameter);
        int right = depth(node->right, maxDiameter);
        maxDiameter = max(maxDiameter, left + right);
        return max(left, right) + 1;
    }
};
// @lc code=end
