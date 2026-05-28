include "./user/user.thrift"
include "./interviews/interviews.thrift"

namespace go interview

service UserService extends user.UserService {
    // 可在此添加新的接口方法
}

service InterviewsService extends interviews.InterviewsService {}
