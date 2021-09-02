package expection

type BizErr struct {
	Code int64
	Msg  string
}

func (biz *BizErr) Error() string {
	return biz.Msg
}
