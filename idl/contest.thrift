namespace go contest

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
  7: string image_url,
  8: ContestCoreInfo contest_core_info,
  9: bool is_favorite,
}

struct ContestBrief {
    1: i32 contest_id,
    2: string title,
    3: string description,
    4: i64 created_time,
    5: string field,
    6: string format,
}

struct ContestListRequest {
  1: string keyword
  2: list<string> fields
  3: list<string> formats
  4: i32 limit
  5: i32 offset
}
struct ContestBriefInfo {
    ContestBrief contest_brief_info,
}

struct ContestListResponse {
    1: i32 status_code,
    2: string status_msg,
    3: i32 total,
    4: list<ContestBriefInfo> contest_list,
}

struct ContestInfoRequest {
    1: i32 contest_id
    2: i32 user_id
}

struct ContestInfoResponse {
    1: i32 status_code,
    2: string status_msg,
    3: Contest contest,
}

struct ContestCreateRequest {
    1: Contest contest
}

struct ContestCreateResponse {
    1: i32 status_code,
    2: string status_msg,
    3: i32 contest_id,
}


//The following interface is specifically designed for the 'favorite' module to retrieve favorite contest list
struct GetContestsByFavoritesRequest {
    1: list<i32> contest_ids
}

struct GetContestsByFavoritesResponse {
    1: list<ContestBriefInfo> contest_list
}


service ContestService {
    // 获取赛事资讯列表
    ContestListResponse ContestList(1: ContestListRequest req)
    // 获取赛事资讯详情
    ContestInfoResponse ContestInfo(1: ContestInfoRequest req)
    // 创建赛事资讯
    ContestCreateResponse ContestCreate(1: ContestCreateRequest req)

    //The following interface is specifically designed for the 'favorite' module to retrieve contest information
    GetContestsByFavoritesResponse GetContestsByFavorites(1: GetContestsByFavoritesRequest req)
}

