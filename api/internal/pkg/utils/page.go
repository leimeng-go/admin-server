package utils

func Page(page int64,pageSize int64) (offset int64,limit int64){
	if page<1{
		page=1
	}
	if pageSize<1{
		pageSize=10
	}
	offset=(page-1)*pageSize
	limit=pageSize
	return offset,limit
}
