package ec2

import "github.com/timob/javabind"

type HandlersRequestHandlerInterface interface {

	// public abstract void beforeRequest(com.amazonaws.Request<?>)
	BeforeRequest(a RequestInterface) 

	// public abstract void afterResponse(com.amazonaws.Request<?>, java.lang.Object, com.amazonaws.util.TimingInfo)
	AfterResponse(a RequestInterface, b interface{}, c UtilTimingInfoInterface) 
}

type HandlersRequestHandler struct {
	JavaLangObject
}

// public abstract void beforeRequest(com.amazonaws.Request<?>)
func (jbobject *HandlersRequestHandler) BeforeRequest(a RequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "beforeRequest", javabind.Void, conv_a.Value().Cast("com/amazonaws/Request"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public abstract void afterResponse(com.amazonaws.Request<?>, java.lang.Object, com.amazonaws.util.TimingInfo)
func (jbobject *HandlersRequestHandler) AfterResponse(a RequestInterface, b interface{}, c UtilTimingInfoInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "afterResponse", javabind.Void, conv_a.Value().Cast("com/amazonaws/Request"), conv_b.Value().Cast("java/lang/Object"), conv_c.Value().Cast("com/amazonaws/util/TimingInfo"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()

}


