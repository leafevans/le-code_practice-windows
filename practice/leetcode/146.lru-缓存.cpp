/*
 * @lc app=leetcode.cn id=146 lang=cpp
 *
 * [146] LRU 缓存
 */

// @lc code=start
class LRUCache {
public:
    LRUCache(int capacity) {
        cap = capacity;
        head = new Node(0, 0);
        tail = new Node(0, 0);
        head->next = tail;
        tail->prev = head;
    }

    int get(int key) {
        if (map.find(key) == map.end()) {
            return -1;
        }
        Node* node = map[key];
        removeNode(node);
        addToTail(node);
        return node->val;
    }

    void put(int key, int value) {
        if (map.find(key) != map.end()) {
            Node* node = map[key];
            node->val = value;
            removeNode(node);
            addToTail(node);
        } else {
            if (map.size() == cap) {
                Node* toRemove = head->next;
                removeNode(toRemove);
                map.erase(toRemove->key);
                delete toRemove;
            }
            Node* newNode = new Node(key, value);
            addToTail(newNode);
            map[key] = newNode;
        }
    }

private:
    struct Node {
        int key, val;
        Node *prev, *next;
        Node(int k, int v) : key(k), val(v), prev(nullptr), next(nullptr) {}
    };

    int cap;
    Node *head, *tail;
    unordered_map<int, Node*> map;

    void addToTail(Node* node) {
        node->prev = tail->prev;
        tail->prev->next = node;
        node->next = tail;
        tail->prev = node;
    }

    void removeNode(Node* node) {
        node->next->prev = node->prev;
        node->prev->next = node->next;
    }
};

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache* obj = new LRUCache(capacity);
 * int param_1 = obj->get(key);
 * obj->put(key,value);
 */
