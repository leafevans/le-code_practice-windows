/*
 * @lc app=leetcode.cn id=155 lang=cpp
 *
 * [155] 最小栈
 */

// @lc code=start
class MinStack {
public:
    MinStack() {}

    void push(int val) {
        data.push(val);
        if (mins.empty() || val <= mins.top()) {
            mins.push(val);
        }
    }

    void pop() {
        if (data.empty()) {
            return;
        }
        if (mins.top() == data.top()) {
            mins.pop();
        }
        data.pop();
    }

    int top() { return data.top(); }

    int getMin() { return mins.top(); }

private:
    stack<int> data;
    stack<int> mins;
};

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack* obj = new MinStack();
 * obj->push(val);
 * obj->pop();
 * int param_3 = obj->top();
 * int param_4 = obj->getMin();
 */
// @lc code=end
