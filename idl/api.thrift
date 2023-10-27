namespace go api

/* =========================== user =========================== */

struct UserInfo {
    1: i32 user_id,
    2: i32 gender,
    3: string nickname,
    4: string realname,
    5: i32 contest_favorite_count,
    6: string avatar_url,
    7: i32 enrollment_year,
    8: string college,
}


struct UserProfileInfo {
    1: i32 user_id,
    2: string mobile_phone,
    3: string introduction,
    4: string qq_number,
    5: string wechat_number,
    6: list<string> honors,
    7: list<string> images,
    8: list<bool> is_show,
}

// 用户注册

struct UserRegisterRequest {
    1: string username
    2: string password
}

struct UserRegisterResponse {
    1: i32 status_code,
    2: string status_msg,
    3: string token,
}

// 用户登录

struct UserLoginRequest {
    1: string username
    2: string password
}

struct UserLoginResponse {
    1: i32 status_code,
    2: string status_msg,
    3: i32 user_id,
    4: string token,
}

// 获取用户信息

struct UserInfoRequest {
    1: i32 user_id (api.query="user_id")
    2: string token (api.query="token")
}

struct UserInfoResponse {
    1: i32 status_code,
    2: string status_msg,
    3: UserInfo user_info,
}

// 上传用户信息
struct UserInfoUploadRequest {
    1: i32 user_id
    2: string token
    3: UserInfo user_info
}

struct UserInfoUploadResponse {
    1: i32 status_code,
    2: string status_msg,
}

// 获取用户档案信息

struct UserProfileInfoRequest {
    1: i32 user_id (api.path="user_id")
    2: string token (api.query="token")
}

struct UserProfileInfoResponse {
    1: i32 status_code,
    2: string status_msg,
    3: UserProfileInfo user_profile_info,
    4: UserInfo user_info,
}

// 上传用户档案信息

struct UserProfileUploadRequest {
    1: i32 user_id
    2: string token
    3: UserProfileInfo user_profile_info
}

struct UserProfileUploadResponse {
    1: i32 status_code,
    2: string status_msg,
}

/* =========================== contest =========================== */

struct Contest {
    1: i32 contest_id,
    2: string title,
    3: string description,
    4: string created_time,
    5: list<string> image_list,
}

struct ContestListRequest {
    1: i32 limit (api.query="limit")
    2: i32 offset (api.query="offset")
    3: string latest_time (api.query="latest_time")
}

struct ContestListResponse {
    1: i32 status_code,
    2: string status_msg,
    3: i32 total,
    4: list<Contest> contest_list,
}

struct ContestInfoRequest {
    1: i32 contest_id (api.path="contest_id")
}

struct ContestInfoResponse {
    1: i32 status_code,
    2: string status_msg,
    3: Contest contest,
}

service ApiService {
    // 用户注册操作
    UserRegisterResponse UserRegister(1: UserRegisterRequest req) (api.post="/fusion/user/register/")
    // 用户登录操作
    UserLoginResponse UserLogin(1: UserLoginRequest req) (api.post="/fusion/user/login/")
    // 获取用户信息
    UserInfoResponse UserInfo(1: UserInfoRequest req) (api.get="/fusion/user/info/")
    // 上传用户信息
    UserInfoUploadResponse UserInfoUpload(1: UserInfoUploadRequest req) (api.post="/fusion/user/info/upload/")
    // 获取用户档案信息
    UserProfileInfoResponse UserProfileInfo(1: UserProfileInfoRequest req) (api.get="/fusion/user/profile/:user_id")
    // 上传用户档案信息
    UserProfileUploadResponse UserProfileUpload(1: UserProfileUploadRequest req) (api.post="/fusion/user/profile/upload/")

    // 获取赛事资讯列表
    ContestListResponse ContestList(1: ContestListRequest req) (api.get="/fusion/contest/list/")
    // 获取赛事资讯详情
    ContestInfoResponse ContestInfo(1: ContestInfoRequest req) (api.get="/fusion/contest/info/:contest_id")
}