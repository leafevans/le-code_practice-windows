/*
 * @lc app=leetcode.cn id=23 lang=cpp
 *
 * [23] 合并 K 个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        if (lists.empty()) {
            return nullptr;
        }
        auto cmp = [](ListNode* a, ListNode* b) { return a->val > b->val; };
        priority_queue<ListNode*, vector<ListNode*>, decltype(cmp)> pq(cmp);
        for (auto list : lists) {
            if (list) {
                pq.push(list);
            }
        }
        ListNode* dummy = new ListNode();
        ListNode* curr = dummy;
        while (!pq.empty()) {
            ListNode* p = pq.top();
            pq.pop();
            curr->next = p;
            curr = curr->next;
            if (p->next) {
                p = p->next;
                pq.push(p);
            }
        }
        return dummy->next;
    }
};
// @lc code=end
