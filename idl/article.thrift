namespace go article

struct ArticleBrief {
    1: i32 article_id,
    2: string title,
    3: i32 author_id,
    4: string author,
    5: i64 created_time,
    6: string link,
}

struct ArticleBriefInfo {
    1: ArticleBrief article_brief_info,
}

struct ArticleListRequest {
    1: i32 contest_id,
    2: i32 limit,
    3: i32 offset,
}

struct ArticleListResponse {
    1: i32 status_code,
    2: string status_msg,
    3: i32 total,
    4: list<ArticleBriefInfo> article_list,
}

struct ArticleCreateRequest {
    1: i32 article_id,
    2: string title,
    3: i32 author_id,
    4: string author,
    5: string link,
    6: i32 contest_id,
}

struct ArticleCreateResponse {
    1: i32 status_code,
    2: string status_msg,
    3: i32 article_id,
}

service ArticleService {
    // 获取赛事资讯文章列表
    ArticleListResponse ArticleList(1: ArticleListRequest req)
    // 创建赛事资讯文章
    ArticleCreateResponse ArticleCreate(1: ArticleCreateRequest req)
}