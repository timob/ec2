package ec2

import "github.com/timob/javabind"

type HandlersRequestHandler2Interface interface {
	JavaLangObjectInterface

	// public com.amazonaws.AmazonWebServiceRequest beforeMarshalling(com.amazonaws.AmazonWebServiceRequest)
	BeforeMarshalling(a AmazonWebServiceRequestInterface) *AmazonWebServiceRequest

	// public void beforeRequest(com.amazonaws.Request<?>)
	BeforeRequest(a RequestInterface) 

	// public com.amazonaws.http.HttpResponse beforeUnmarshalling(com.amazonaws.Request<?>, com.amazonaws.http.HttpResponse)
	BeforeUnmarshalling(a RequestInterface, b HttpHttpResponseInterface) *HttpHttpResponse

	// public void afterResponse(com.amazonaws.Request<?>, com.amazonaws.Response<?>)
	AfterResponse(a RequestInterface, b ResponseInterface) 

	// public static com.amazonaws.handlers.RequestHandler2 adapt(com.amazonaws.handlers.RequestHandler)
	Adapt(a HandlersRequestHandlerInterface) *HandlersRequestHandler2
}

type HandlersRequestHandler2 struct {
	JavaLangObject
}

// public com.amazonaws.handlers.RequestHandler2()
func NewHandlersRequestHandler2() (*HandlersRequestHandler2) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/handlers/RequestHandler2")
	if err != nil {
		panic(err)
	}
	x := &HandlersRequestHandler2{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.AmazonWebServiceRequest beforeMarshalling(com.amazonaws.AmazonWebServiceRequest)
func (jbobject *HandlersRequestHandler2) BeforeMarshalling(a AmazonWebServiceRequestInterface) *AmazonWebServiceRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "beforeMarshalling", "com/amazonaws/AmazonWebServiceRequest", conv_a.Value().Cast("com/amazonaws/AmazonWebServiceRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &AmazonWebServiceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void beforeRequest(com.amazonaws.Request<?>)
func (jbobject *HandlersRequestHandler2) BeforeRequest(a RequestInterface)  {
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

// public com.amazonaws.http.HttpResponse beforeUnmarshalling(com.amazonaws.Request<?>, com.amazonaws.http.HttpResponse)
func (jbobject *HandlersRequestHandler2) BeforeUnmarshalling(a RequestInterface, b HttpHttpResponseInterface) *HttpHttpResponse {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "beforeUnmarshalling", "com/amazonaws/http/HttpResponse", conv_a.Value().Cast("com/amazonaws/Request"), conv_b.Value().Cast("com/amazonaws/http/HttpResponse"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &HttpHttpResponse{}
	unique_x.Callable = dst
	return unique_x
}

// public void afterResponse(com.amazonaws.Request<?>, com.amazonaws.Response<?>)
func (jbobject *HandlersRequestHandler2) AfterResponse(a RequestInterface, b ResponseInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "afterResponse", javabind.Void, conv_a.Value().Cast("com/amazonaws/Request"), conv_b.Value().Cast("com/amazonaws/Response"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public static com.amazonaws.handlers.RequestHandler2 adapt(com.amazonaws.handlers.RequestHandler)
func (jbobject *HandlersRequestHandler2) Adapt(a HandlersRequestHandlerInterface) *HandlersRequestHandler2 {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/handlers/RequestHandler2", "adapt", "com/amazonaws/handlers/RequestHandler2", conv_a.Value().Cast("com/amazonaws/handlers/RequestHandler"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &HandlersRequestHandler2{}
	unique_x.Callable = dst
	return unique_x
}


