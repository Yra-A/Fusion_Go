namespace go api

/* =========================== user =========================== */


struct UserInfo {
    1: i32 user_id,
    2: i32 gender,
    4: i32 enrollment_year,
    3: string mobile_phone
    5: string college,
    6: string nickname,
    7: string realname,
    8: bool has_profile,
    9: string avatar_url,
}




struct UserProfileInfo {
    1: string introduction,
    2: string qq_number,
    3: string wechat_number,
    4: list<string> honors,
    5: UserInfo user_info,
}

// 用户注册

struct UserRegisterRequest {
    1: string username
    2: string password
}

struct UserRegisterResponse {
    1: i32 status_code,
    2: string status_msg,
}

// 用户登录

struct UserLoginRequest {
    1: string username
    2: string password
}

struct UserLoginResponse {
    1: i32 status_code,
    2: string status_msg,
    3: string token,
}

// 获取用户信息

struct UserInfoRequest {
    1: i32 user_id (api.query="user_id")
    2: string authorization (api.header="Authorization")
}

struct UserInfoResponse {
    1: i32 status_code,
    2: string status_msg,
    3: UserInfo user_info,
}

// 上传用户信息
struct UserInfoUploadRequest {
    1: string authorization (api.header="Authorization")
    2: UserInfo user_info
}

struct UserInfoUploadResponse {
    1: i32 status_code,
    2: string status_msg,
}

// 获取用户档案信息

struct UserProfileInfoRequest {
    1: i32 user_id (api.path="user_id")
    2: string authorization (api.header="Authorization")
}

struct UserProfileInfoResponse {
    1: i32 status_code,
    2: string status_msg,
    3: UserProfileInfo user_profile_info,
}

// 上传用户档案信息

struct UserProfileUploadRequest {
    1: i32 user_id
    2: string authorization (api.header="Authorization")
    3: UserProfileInfo user_profile_info
}

struct UserProfileUploadResponse {
    1: i32 status_code,
    2: string status_msg,
}

/* =========================== contest =========================== */

struct TeamSize {
  1: i32 min,
  2: i32 max,
}

struct Contact {
  1: string name,
  2: string phone,
  3: string email,
}

struct ContestCoreInfo {
  1: i32 deadline,
  2: string fee,
  3: TeamSize team_size,
  4: string participant_requirements,
  5: string official_website,
  6: string additional_info,
  7: list<Contact> contact,
}

struct Contest {
  1: i32 contest_id,
  2: string title,
  3: string description,
  4: i64 created_time,
  5: string field,
  6: string format,
  7: string Image_url,
  8: ContestCoreInfo contest_core_info,
}

struct ContestBrief {
    1: i32 contest_id,
    2: string title,
    3: string description,
    4: i64 created_time,
    5: string field,
    6: string format,
}
struct ContestBriefInfo {
    ContestBrief contest_brief_info,
}

struct ContestListRequest {
  1: string keyword (api.query="keyword")
  2: list<string> fields (api.query="fields")
  3: list<string> formats (api.query="formats")
  4: i32 limit (api.query="limit")
  5: i32 offset (api.query="offset")
 }


struct ContestListResponse {
    1: i32 status_code,
    2: string status_msg,
    3: i32 total,
    4: list<ContestBriefInfo> contest_list,
}

struct ContestInfoRequest {
    1: i32 contest_id (api.path="contest_id")
}

struct ContestInfoResponse {
    1: i32 status_code,
    2: string status_msg,
    3: Contest contest,
}

/* =========================== utils =========================== */

struct ImageUploadRequest {
    1: string authorization (api.header="Authorization")
    2: binary file (api.form="file")
}

struct ImageUploadResponse {
    1: i32 status_code,
    2: string status_msg,
    3: string Image_url,
}

/* =========================== team =========================== */

struct MemberInfo {
    1: i32 user_id,
    2: string nickname,
    3: string college,
    4: string avatar_url,
    5: i32 gender,
    6: i32 enrollment_year,
    7: list<string> honors,
}

struct TeamBriefInfo {
    1: i32 team_id,
    2: string title,
    3: string goal,
    4: i32 cur_people_num,
    5: i64 created_time,
    6: MemberInfo leader_info,
    7: i32 contest_id,
}

struct TeamInfo {
    1: TeamBriefInfo team_brief_info,
    2: string description,
    3: list<MemberInfo> members,
}

struct TeamApplication {
    1: i32 team_id,
    2: string reason,
    3: i64 created_time,
    4: i32 application_type,
    5: MemberInfo member_info,
    6: i32 application_id,
}

struct TeamCreateRequest {
    1: string authorization (api.header="Authorization")
    2: i32 user_id
    3: i32 team_id
    4: string title
    5: string goal
    6: string description
    7: i32 contest_id
}

struct TeamCreateResponse {
    1: i32 status_code,
    2: string status_msg,
}

struct TeamListRequest {
    1: string authorization (api.header="Authorization")
    2: i32 contest_id (api.path="contest_id")
    3: i32 limit (api.query="limit")
    4: i32 offset (api.query="offset")
}

struct TeamListResponse {
    1: i32 status_code,
    2: string status_msg,
    3: i32 total,
    4: list<TeamBriefInfo> team_list,
}

struct TeamInfoRequest {
    1: string authorization (api.header="Authorization")
    2: i32 contest_id (api.path="contest_id")
    3: i32 team_id (api.path="team_id")
}

struct TeamInfoResponse {
    1: i32 status_code,
    2: string status_msg,
    3: TeamInfo team_info,
}

struct TeamApplicationSubmitRequest {
    1: string authorization (api.header="Authorization")
    2: i32 team_id,
    3: string reason,
    4: i64 created_time,
    5: i32 application_type,
    6: MemberInfo member_info,
}

struct TeamApplicationSubmitResponse {
    1: i32 status_code,
    2: string status_msg,
}

struct TeamManageListRequest {
    1: string authorization (api.header="Authorization")
    2: i32 user_id (api.query="user_id")
    3: i32 team_id (api.query="team_id")
}

struct TeamManageListResponse {
    1: i32 status_code,
    2: string status_msg,
    3: list<TeamApplication> application_list,
}

struct TeamManageActionRequest {
    1: string authorization (api.header="Authorization")
    2: i32 user_id (api.path="user_id")
    3: i32 application_id (api.path="application_id")
    4: i32 action_type (api.path="action_type")
}

struct TeamManageActionResponse {
    1: i32 status_code,
    2: string status_msg,
}

service ApiService {
    /* user */
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

    /* contest */
    // 获取赛事资讯列表
    ContestListResponse ContestList(1: ContestListRequest req) (api.get="/fusion/contest/list/")
    // 获取赛事资讯详情
    ContestInfoResponse ContestInfo(1: ContestInfoRequest req) (api.get="/fusion/contest/info/:contest_id")


    /* utils */

    // 上传图片
    ImageUploadResponse ImageUpload(1: ImageUploadRequest req) (api.post="/fusion/utils/upload/img")

    /* team */
    // 创建队伍
    TeamCreateResponse TeamCreate(1: TeamCreateRequest req) (api.post="/fusion/team/create/")
    // 获取队伍列表
    TeamListResponse TeamList(1: TeamListRequest req) (api.get="/fusion/contest/:contest_id/team/list")
    // 获取队伍详情
    TeamInfoResponse TeamInfo(1: TeamInfoRequest req) (api.get="/fusion/contest/:contest_id/team/info/:team_id")
    // 提交队伍申请
    TeamApplicationSubmitResponse TeamApplicationSubmit(1: TeamApplicationSubmitRequest req) (api.post="/fusion/team/application/submit/")
    // 获取队伍申请列表
    TeamManageListResponse TeamManageList(1: TeamManageListRequest req) (api.get="/fusion/team/manage/list/")
    // 队伍申请操作
    TeamManageActionResponse TeamManageAction(1: TeamManageActionRequest req) (api.post="/fusion/team/manage/action/")
}