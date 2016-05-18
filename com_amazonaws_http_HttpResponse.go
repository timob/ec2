package ec2

import "github.com/timob/javabind"

type HttpHttpResponseInterface interface {
	JavaLangObjectInterface

	// public com.amazonaws.Request<?> getRequest()
	GetRequest() *Request

	// public java.util.Map<java.lang.String, java.lang.String> getHeaders()
	GetHeaders() map[string]string

	// public void addHeader(java.lang.String, java.lang.String)
	AddHeader(a string, b string) 

	// public void setStatusText(java.lang.String)
	SetStatusText(a string) 

	// public java.lang.String getStatusText()
	GetStatusText() string

	// public void setStatusCode(int)
	SetStatusCode(a int) 

	// public int getStatusCode()
	GetStatusCode() int
}

type HttpHttpResponse struct {
	JavaLangObject
}

// public com.amazonaws.Request<?> getRequest()
func (jbobject *HttpHttpResponse) GetRequest() *Request {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRequest", "com/amazonaws/Request")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &Request{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.Map<java.lang.String, java.lang.String> getHeaders()
func (jbobject *HttpHttpResponse) GetHeaders() map[string]string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHeaders", "java/util/Map")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoMap(javabind.NewJavaToGoString(), javabind.NewJavaToGoString())
	dst := new(map[string]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void addHeader(java.lang.String, java.lang.String)
func (jbobject *HttpHttpResponse) AddHeader(a string, b string)  {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addHeader", javabind.Void, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void setStatusText(java.lang.String)
func (jbobject *HttpHttpResponse) SetStatusText(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatusText", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getStatusText()
func (jbobject *HttpHttpResponse) GetStatusText() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStatusText", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setStatusCode(int)
func (jbobject *HttpHttpResponse) SetStatusCode(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatusCode", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public int getStatusCode()
func (jbobject *HttpHttpResponse) GetStatusCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStatusCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}


