namespace go api

/* =========================== user =========================== */

struct UserInfo {
    1: i32 user_id,
    2: string nickname,
    3: string realname,
    4: i32 contest_favorite_count,
    5: string avatar_url,
}

struct UserProfileInfo {
    1: i32 user_id,
    2: string mobile_phone,
    3: string introduction,
    4: string qq_number,
    5: string wechat_number,
    6: list<string> honors,
    7: list<string> images,
}

struct UserRegisterRequest {
    1: string username (api.raw_body="username")
    2: string password (api.raw_body="password")
}

struct UserRegisterResponse {
    1: i32 status_code,
    2: string status_msg,
    3: string token,
}

struct UserLoginRequest {
    1: string username (api.raw_body="username")
    2: string password (api.raw_body="password")
}

struct UserLoginResponse {
    1: i32 status_code,
    2: string status_msg,
    3: i32 user_id,
    4: string token,
}

struct UserInfoRequest {
    1: i32 user_id (api.query="user_id")
    2: string token (api.query="token")
}

struct UserInfoResponse {
    1: i32 status_code,
    2: string status_msg,
    3: UserInfo user_info,
}

struct UserProfileInfoRequest {
    1: i32 user_id (api.path="user_id")
    2: string token (api.query="token")
}

struct UserProfileInfoResponse {
    1: i32 status_code,
    2: string status_msg,
    3: UserProfileInfo user_profile_info,
}

struct UserProfileUploadRequest {
    1: string token (api.query="token")
    2: i32 user_id (api.raw_body="user_id")
    3: string mobile_phone (api.raw_body="mobile_phone")
    4: string introduction (api.raw_body="introduction")
    5: string qq_number (api.raw_body="qq_number")
    6: string wechat_number (api.raw_body="wechat_number")
    7: list<string> honors (api.raw_body="honors")
    8: list<string> images (api.raw_body="images")
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
    // 获取用户档案信息
    UserProfileInfoResponse UserProfileInfo(1: UserProfileInfoRequest req) (api.get="/fusion/user/profile/{user_id}")
    // 上传用户档案信息
    UserProfileUploadResponse UserProfileUpload(1: UserProfileUploadRequest req) (api.post="/fusion/user/profile/upload/")

    // 获取赛事资讯列表
    ContestListResponse ContestList(1: ContestListRequest req) (api.get="/fusion/contest/list/")
    // 获取赛事资讯详情
    ContestInfoResponse ContestInfo(1: ContestInfoRequest req) (api.get="/fusion/contest/info/{contest_id}")
}