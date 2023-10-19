namespace go user

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
    1: string username,
    2: string password,
}

struct UserRegisterResponse {
    1: i32 status_code,
    2: string status_msg,
    3: string token,
}

struct UserLoginRequest {
    1: string username,
    2: string password,
}

struct UserLoginResponse {
    1: i32 status_code,
    2: string status_msg,
    3: i32 user_id,
    4: string token,
}

struct UserInfoRequest {
    1: i32 user_id,
    2: string token,
}

struct UserInfoResponse {
    1: i32 status_code,
    2: string status_msg,
    3: UserInfo user_info,
}

struct UserProfileInfoRequest {
    1: i32 user_id,
    2: string token,
}

struct UserProfileInfoResponse {
    1: i32 status_code,
    2: string status_msg,
    3: UserProfileInfo user_profile_info,
}

struct UserProfileUploadRequest {
    1: string token,
    2: i32 user_id,
    3: string mobile_phone,
    4: string introduction,
    5: string qq_number,
    6: string wechat_number,
    7: list<string> honors,
    8: list<string> images,
}

struct UserProfileUploadResponse {
    1: i32 status_code,
    2: string status_msg,
}

service UserService {
    // 用户注册操作
    UserRegisterResponse UserRegister(1: UserRegisterRequest req)
    // 用户登录操作
    UserLoginResponse UserLogin(1: UserLoginRequest req)
    // 获取用户信息
    UserInfoResponse UserInfo(1: UserInfoRequest req)
    // 获取用户档案信息
    UserProfileInfoResponse UserProfileInfo(1: UserProfileInfoRequest req)
    // 上传用户档案信息
    UserProfileUploadResponse UserProfileUpload(1: UserProfileUploadRequest req)
}