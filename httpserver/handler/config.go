package handler

type config struct {
	withCtx         bool
	withoutResponse bool
	withResult      bool
}

type Option func(config *config)

func newConfig(options ...Option) config {
	conf := config{}
	for _, option := range options {
		option(&conf)
	}
	return conf
}

// WithContext handler func must in 3 params, and the third param must be *gin.Context
//
// example:
// func (h *HelloHandler) HelloWithGinCtx(ctx context.Context, req *proto.HelloHandlerReq, c *gin.Context) (*proto.HelloHandlerResp, error) {
// 	 resp := proto.HelloHandlerResp{}
//	 resp.Welcome = fmt.Sprintf("hello, %s! the uri is %s", req.Name, c.Request.RequestURI)
//	 return &resp, nil
// }
func WithContext() Option {
	return func(conf *config) {
		conf.withCtx = true
	}
}

// WithoutResponse handler do not deal with the response, but deal with the error.
// you need do it by yourself, and sometimes use both "no response" and "ginCtx"
//
// example:
// func HelloWithGinCtxAndNoResponse(ctx context.Context, req *proto.HelloHandlerReq, c *gin.Context) error {
//	 var err error
//	 if err != nil {
//		 return err
//	 }
//	 resp := proto.HelloHandlerResp{}
//	 resp.Welcome = fmt.Sprintf("hello, %s! the uri is %s", req.Name, c.Request.RequestURI)
//	 c.JSON(http.StatusOK, resp)
//	 return nil
// }
func WithoutResponse() Option {
	return func(conf *config) {
		conf.withoutResponse = true
	}
}

// WithResult handler func must return 1 param and the type must be *result.Result
//
// example:
// func HelloWithResult(ctx context.Context, req *proto.HelloHandlerReq) *result.Result {
//	 resp := proto.HelloHandlerResp{}
//	 resp.Welcome = fmt.Sprintf("hello, %s!", req.Name)
//	 return result.Succeed(resp)
// }
func WithResult() Option {
	return func(conf *config) {
		conf.withResult = true
	}
}
