namespace go favorite

/* =========================== favorite =========================== */
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

struct ContestFavoriteActionRequest {
    1: i32 user_id
    2: i32 contest_id
    3: i32 action_type
}

struct ContestFavoriteActionResponse {
    1: i32 status_code,
    2: string status_msg,
}

struct ContestFavoriteListRequest {
    1: i32 user_id
    2: i32 limit
    3: i32 offset
}

struct ContestFavoriteListResponse {
    1: i32 status_code,
    2: string status_msg,
    3: list<ContestBriefInfo> contest_list,
    4: i32 total
}

//The following interface is specifically designed for the 'contest' module to retrieve favorite status
struct QueryFavoriteStatusByUserIdRequest {
    1: i32 user_id
    2: i32 contest_id
}

struct QueryFavoriteStatusByUserIdResponse {
    1: bool is_favorite
}

service FavoriteService {
    // 赛事收藏操作
    ContestFavoriteActionResponse ContestFavoriteAction(1: ContestFavoriteActionRequest req)
    // 获取赛事收藏列表
    ContestFavoriteListResponse ContestFavoriteList(1: ContestFavoriteListRequest req)
    // 获取用户对某个赛事的收藏状态
    QueryFavoriteStatusByUserIdResponse QueryFavoriteStatusByUserId(1: QueryFavoriteStatusByUserIdRequest req)
}
