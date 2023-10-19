namespace go contest

struct Contest {
    1: i32 contest_id,
    2: string title,
    3: string description,
    4: string created_time,
    5: list<string> image_list,
}

struct ContestListRequest {
    1: i32 limit
    2: i32 offset
    3: string latest_time
}

struct ContestListResponse {
    1: i32 status_code,
    2: string status_msg,
    3: i32 total,
    4: list<Contest> contest_list,
}

struct ContestInfoRequest {
    1: i32 contest_id
}

struct ContestInfoResponse {
    1: i32 status_code,
    2: string status_msg,
    3: Contest contest,
}

service ContestService {
    // 获取赛事资讯列表
    ContestListResponse ContestList(1: ContestListRequest req)
    // 获取赛事资讯详情
    ContestInfoResponse ContestInfo(1: ContestInfoRequest req)
}