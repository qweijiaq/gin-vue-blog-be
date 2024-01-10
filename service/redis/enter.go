package redis

const (
	articleLookPrefix         = "article_look"          // 文章浏览量
	articleCommentCountPrefix = "article_comment_count" // 文章评论数
	articleDiggPrefix         = "article_digg"          // 文章点赞数
	commentDiggPrefix         = "comment_digg"          // 评论点赞数
)

func NewDigg() CountDB {
	return CountDB{
		Index: articleDiggPrefix,
	}
}
func NewArticleLook() CountDB {
	return CountDB{
		Index: articleLookPrefix,
	}
}
func NewCommentCount() CountDB {
	return CountDB{
		Index: articleCommentCountPrefix,
	}
}
func NewCommentDigg() CountDB {
	return CountDB{
		Index: commentDiggPrefix,
	}
}
