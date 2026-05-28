namespace go user

struct RegisterRequest {
    1: required string username    (api.form = "username")
    2: required string email       (api.form = "email")
    3: required string password    (api.form = "password")
}

// 登录请求
struct LoginRequest {
    1: required string email        (api.form="email")
    2: required string password     (api.form="password")
}

// 用户资料
struct UserProfile {
    1: required i64 id
    2: required string username
    3: required string email
    4: required string role
    5: optional string wechat_open_id
    6: optional string wechat_union_id
    7: optional string nickname
    8: optional string avatar
    9: optional i64 created_at
    10: optional i64 updated_at
}

struct LoginResponse {
    1: required string token
    2: required UserProfile user
}

service UserService {
    LoginResponse Register(1: RegisterRequest request) (
        api.post     = "/api/user/register",
        api.category = "user",
        api.gen_path = "user",
    )

    // 用户登录
    LoginResponse Login(1: LoginRequest request) (
        api.post     = "/api/user/login",
        api.category = "user",
        api.gen_path = "user",
    )
}
