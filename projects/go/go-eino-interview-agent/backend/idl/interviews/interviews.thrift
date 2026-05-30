namespace go interviews

// === 简历相关 ===

// 简历信息
struct ResumeInfo {
    1: required i64 id           // 简历ID
    2: required i32 user_id      // 用户ID
    3: required string file_name // 文件名
    4: required i64 file_size    // 文件大小（字节）
    5: required string file_type // 文件类型（仅支持 pdf）
    6: required i32 is_default   // 是否为默认简历（0=否，1=是）
    7: required i64 created_at   // 创建时间（毫秒时间戳）
    8: required i64 updated_at   // 更新时间（毫秒时间戳）
}

// 上传简历请求（使用 FormFile 上传 PDF 文件）
struct UploadResumeRequest {
    1: required string resume (api.form = "resume") // 简历文件（PDF 格式）
}

// 上传简历响应
struct UploadResumeResponse {
    1: required i64 resume_id // 上传后的简历ID
    2: required string message // 消息说明
}

// 获取简历详情请求
struct GetResumeRequest {
    1: required i64 resume_id (api.path = "resume_id") // 简历ID
}

// 获取简历详情响应
struct GetResumeResponse {
    1: required ResumeInfo resume
}

// 获取用户简历列表请求
struct GetUserResumesRequest {
    1: optional i32 page (api.query = "page") // 页码，默认 1
    2: optional i32 page_size (api.query = "page_size") // 每页数量，默认 10
}

// 获取用户简历列表响应
struct GetUserResumesResponse {
    1: required list<ResumeInfo> resumes // 简历列表
    2: required i64 total                // 总条数
    3: required i32 page                 // 当前页码
    4: required i32 page_size            // 每页数量
}

// 更新简历请求（仅支持更新文件名，文件内容通过重新上传）
struct UpdateResumeRequest {
    1: required i64 resume_id (api.path = "resume_id")    // 简历ID
    2: optional string file_name (api.body = "file_name") // 文件名
}

// 更新简历响应
struct UpdateResumeResponse {
    1: required string message // 消息说明
}

// 删除简历请求
struct DeleteResumeRequest {
    1: required i64 resume_id (api.path = "resume_id") // 简历ID
}

// 删除简历响应
struct DeleteResumeResponse {
    1: required string message // 消息说明
}

// === 服务定义 ===

service InterviewsService {
    // ==================== 简历相关接口 ====================

    // 上传简历
    UploadResumeResponse UploadResume(1: UploadResumeRequest request) (
        api.post = "/api/resume/upload",
        api.category = "resumes",
        api.gen_path = "interviews"
    )

    // 获取简历详情
    GetResumeResponse GetResume(1: GetResumeRequest request) (
        api.get = "/api/resume/:resume_id",
        api.category = "resumes",
        api.gen_path = "interviews"
    )

    // 获取用户简历列表
    GetUserResumesResponse GetUserResumes(1: GetUserResumesRequest request) (
        api.get = "/api/resume/list",
        api.category = "resumes",
        api.gen_path = "interviews"
    )

    // 更新简历
    UpdateResumeResponse UpdateResume(1: UpdateResumeRequest request) (
        api.put = "/api/resume/:resume_id",
        api.category = "resumes",
        api.gen_path = "interviews"
    )

    // 删除简历
    DeleteResumeResponse DeleteResume(1: DeleteResumeRequest request) (
        api.delete = "/api/resume/:resume_id",
        api.category = "resumes",
        api.gen_path = "interviews"
    )
}
