package main

//type withStack   struct{
//	error
//	*stack
//}

func WithStack(err error)  error {
	if  err == nil {
		return  nil
	}
	//return &withStack{
    //     err,
    //     runtime.callers(3),
	//}
	return  err

}

