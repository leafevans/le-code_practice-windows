/*
 * @lc app=leetcode.cn id=437 lang=cpp
 *
 * [437] 路径总和 III
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
    int pathSum(TreeNode* root, int targetSum) {
        prefix[0] = 1;
        dfs(root, 0, targetSum);
        return count;
    }

private:
    int count = 0;
    unordered_map<long, int> prefix;
    void dfs(TreeNode* node, long currSum, int targetSum) {
        if (!node) {
            return;
        }
        currSum += node->val;
        if (prefix.count(currSum - targetSum)) {
            count += prefix[currSum - targetSum];
        }
        ++prefix[currSum];
        dfs(node->left, currSum, targetSum);
        dfs(node->right, currSum, targetSum);
        --prefix[currSum];
    }
};
// @lc code=end
