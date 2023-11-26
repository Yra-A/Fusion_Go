namespace go team

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
    1: i32 user_id
    2: i32 team_id
    3: string title
    4: string goal
    5: string description
    6: i32 contest_id
}

struct TeamCreateResponse {
    1: i32 status_code,
    2: string status_msg,
    3: i32 team_id
}

struct TeamListRequest {
    1: i32 contest_id,
    2: i32 limit,
    3: i32 offset,
}

struct TeamListResponse {
    1: i32 status_code,
    2: string status_msg,
    3: i32 total,
    4: list<TeamBriefInfo> team_list,
}

struct TeamInfoRequest {
    1: i32 contest_id,
    2: i32 team_id,
}

struct TeamInfoResponse {
    1: i32 status_code,
    2: string status_msg,
    3: TeamInfo team_info,
}

struct TeamApplicationSubmitRequest {
    1: i32 team_id,
    2: string reason,
    3: i64 created_time,
    4: i32 application_type,
    5: MemberInfo member_info,
}

struct TeamApplicationSubmitResponse {
    1: i32 status_code,
    2: string status_msg,
}

struct TeamManageListRequest {
    1: i32 user_id,
    2: i32 team_id,
}

struct TeamManageListResponse {
    1: i32 status_code,
    2: string status_msg,
    3: list<TeamApplication> application_list,
}

struct TeamManageActionRequest {
    1: i32 user_id,
    2: i32 application_id,
    3: i32 action_type,
}

struct TeamManageActionResponse {
    1: i32 status_code,
    2: string status_msg,
}

service TeamService {
    /* team */
    // 创建队伍
    TeamCreateResponse TeamCreate(1: TeamCreateRequest req)
    // 获取队伍列表
    TeamListResponse TeamList(1: TeamListRequest req)
    // 获取队伍详情
    TeamInfoResponse TeamInfo(1: TeamInfoRequest req)
    // 提交队伍申请
    TeamApplicationSubmitResponse TeamApplicationSubmit(1: TeamApplicationSubmitRequest req)
    // 获取队伍申请列表
    TeamManageListResponse TeamManageList(1: TeamManageListRequest req)
    // 队伍申请操作
    TeamManageActionResponse TeamManageAction(1: TeamManageActionRequest req)
}