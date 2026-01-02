#include <algorithm>
#include <iostream>
#include <memory>
#include <string>
#include <string_view>
#include <vector>

namespace modern_cpp {

enum class Priority { Low, Medium, High };

struct Task {
    int id;
    std::string description;
    Priority priority;
    bool completed = false;
};

class TaskManager {
public:
    explicit TaskManager(std::string_view name) : owner_name_(name) {}

    void add_task(int id, std::string_view desc, Priority prio) { tasks_.push_back({id, std::string(desc), prio}); }

    [[nodiscard]] std::vector<Task> get_pending_tasks() const {
        std::vector<Task> pending;
        std::copy_if(tasks_.begin(), tasks_.end(), std::back_inserter(pending),
                     [](const Task& t) { return !t.completed; });
        return pending;
    }

    void complete_task(int id) {
        auto it = std::find_if(tasks_.begin(), tasks_.end(), [id](const Task& t) { return t.id == id; });
        if (it != tasks_.end()) {
            it->completed = true;
            std::cout << "Task " << id << " completed.\n";
        }
    }

    void print_status() const {
        std::cout << "Task Manager for: " << owner_name_ << "\n";
        for (const auto& task : tasks_) {
            std::cout << " - [" << (task.completed ? "X" : " ") << "] " << task.description
                      << " (Priority: " << static_cast<int>(task.priority) << ")\n";
        }
    }

private:
    std::string owner_name_;
    std::vector<Task> tasks_;
};

} // namespace modern_cpp

int main() {
    auto manager = std::make_unique<modern_cpp::TaskManager>("Developer");

    manager->add_task(1, "Fix critical bugs", modern_cpp::Priority::High);
    manager->add_task(2, "Write documentation", modern_cpp::Priority::Medium);
    manager->add_task(3, "Coffee break", modern_cpp::Priority::Low);

    manager->complete_task(1);

    std::cout << "\n--- Current Status ---\n";
    manager->print_status();

    return 0;
}
