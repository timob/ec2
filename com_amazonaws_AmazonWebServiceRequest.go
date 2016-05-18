package ec2

import "github.com/timob/javabind"

type AmazonWebServiceRequestInterface interface {
	JavaLangObjectInterface

	// public void setRequestCredentials(com.amazonaws.auth.AWSCredentials)
	SetRequestCredentials(a AuthAWSCredentialsInterface) 

	// public com.amazonaws.auth.AWSCredentials getRequestCredentials()
	GetRequestCredentials() *AuthAWSCredentials

	// public com.amazonaws.metrics.RequestMetricCollector getRequestMetricCollector()
	GetRequestMetricCollector() *MetricsRequestMetricCollector

	// public void setRequestMetricCollector(com.amazonaws.metrics.RequestMetricCollector)
	SetRequestMetricCollector(a MetricsRequestMetricCollectorInterface) 

	// public void setGeneralProgressListener(com.amazonaws.event.ProgressListener)
	SetGeneralProgressListener(a EventProgressListenerInterface) 

	// public com.amazonaws.event.ProgressListener getGeneralProgressListener()
	GetGeneralProgressListener() *EventProgressListener

	// public java.util.Map<java.lang.String, java.lang.String> getCustomRequestHeaders()
	GetCustomRequestHeaders() map[string]string

	// public java.lang.String putCustomRequestHeader(java.lang.String, java.lang.String)
	PutCustomRequestHeader(a string, b string) string

	// public java.util.Map<java.lang.String, java.util.List<java.lang.String>> getCustomQueryParameters()
	GetCustomQueryParameters() map[string]*[]string

	// public void putCustomQueryParameter(java.lang.String, java.lang.String)
	PutCustomQueryParameter(a string, b string) 

	// public final int getReadLimit()
	GetReadLimit() int

	// public com.amazonaws.AmazonWebServiceRequest getCloneSource()
	GetCloneSource() *AmazonWebServiceRequest

	// public com.amazonaws.AmazonWebServiceRequest getCloneRoot()
	GetCloneRoot() *AmazonWebServiceRequest

	// public java.lang.Integer getSdkRequestTimeout()
	GetSdkRequestTimeout() int

	// public void setSdkRequestTimeout(int)
	SetSdkRequestTimeout(a int) 

	// public java.lang.Integer getSdkClientExecutionTimeout()
	GetSdkClientExecutionTimeout() int

	// public void setSdkClientExecutionTimeout(int)
	SetSdkClientExecutionTimeout(a int) 

	// public com.amazonaws.AmazonWebServiceRequest clone()
	Clone() *AmazonWebServiceRequest
}

type AmazonWebServiceRequest struct {
	JavaLangObject
}

// public com.amazonaws.AmazonWebServiceRequest()
func NewAmazonWebServiceRequest() (*AmazonWebServiceRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/AmazonWebServiceRequest")
	if err != nil {
		panic(err)
	}
	x := &AmazonWebServiceRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setRequestCredentials(com.amazonaws.auth.AWSCredentials)
func (jbobject *AmazonWebServiceRequest) SetRequestCredentials(a AuthAWSCredentialsInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRequestCredentials", javabind.Void, conv_a.Value().Cast("com/amazonaws/auth/AWSCredentials"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.auth.AWSCredentials getRequestCredentials()
func (jbobject *AmazonWebServiceRequest) GetRequestCredentials() *AuthAWSCredentials {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRequestCredentials", "com/amazonaws/auth/AWSCredentials")
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
	unique_x := &AuthAWSCredentials{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.metrics.RequestMetricCollector getRequestMetricCollector()
func (jbobject *AmazonWebServiceRequest) GetRequestMetricCollector() *MetricsRequestMetricCollector {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRequestMetricCollector", "com/amazonaws/metrics/RequestMetricCollector")
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
	unique_x := &MetricsRequestMetricCollector{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRequestMetricCollector(com.amazonaws.metrics.RequestMetricCollector)
func (jbobject *AmazonWebServiceRequest) SetRequestMetricCollector(a MetricsRequestMetricCollectorInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRequestMetricCollector", javabind.Void, conv_a.Value().Cast("com/amazonaws/metrics/RequestMetricCollector"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setGeneralProgressListener(com.amazonaws.event.ProgressListener)
func (jbobject *AmazonWebServiceRequest) SetGeneralProgressListener(a EventProgressListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGeneralProgressListener", javabind.Void, conv_a.Value().Cast("com/amazonaws/event/ProgressListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.event.ProgressListener getGeneralProgressListener()
func (jbobject *AmazonWebServiceRequest) GetGeneralProgressListener() *EventProgressListener {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGeneralProgressListener", "com/amazonaws/event/ProgressListener")
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
	unique_x := &EventProgressListener{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.Map<java.lang.String, java.lang.String> getCustomRequestHeaders()
func (jbobject *AmazonWebServiceRequest) GetCustomRequestHeaders() map[string]string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCustomRequestHeaders", "java/util/Map")
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

// public java.lang.String putCustomRequestHeader(java.lang.String, java.lang.String)
func (jbobject *AmazonWebServiceRequest) PutCustomRequestHeader(a string, b string) string {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "putCustomRequestHeader", "java/lang/String", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.util.Map<java.lang.String, java.util.List<java.lang.String>> getCustomQueryParameters()
func (jbobject *AmazonWebServiceRequest) GetCustomQueryParameters() map[string]*[]string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCustomQueryParameters", "java/util/Map")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoMap(javabind.NewJavaToGoString(), javabind.NewJavaToGoList(javabind.NewJavaToGoString()))
	dst := new(map[string]*[]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void putCustomQueryParameter(java.lang.String, java.lang.String)
func (jbobject *AmazonWebServiceRequest) PutCustomQueryParameter(a string, b string)  {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "putCustomQueryParameter", javabind.Void, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public final int getReadLimit()
func (jbobject *AmazonWebServiceRequest) GetReadLimit() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReadLimit", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.AmazonWebServiceRequest getCloneSource()
func (jbobject *AmazonWebServiceRequest) GetCloneSource() *AmazonWebServiceRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCloneSource", "com/amazonaws/AmazonWebServiceRequest")
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
	unique_x := &AmazonWebServiceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest getCloneRoot()
func (jbobject *AmazonWebServiceRequest) GetCloneRoot() *AmazonWebServiceRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCloneRoot", "com/amazonaws/AmazonWebServiceRequest")
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
	unique_x := &AmazonWebServiceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Integer getSdkRequestTimeout()
func (jbobject *AmazonWebServiceRequest) GetSdkRequestTimeout() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSdkRequestTimeout", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setSdkRequestTimeout(int)
func (jbobject *AmazonWebServiceRequest) SetSdkRequestTimeout(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSdkRequestTimeout", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public java.lang.Integer getSdkClientExecutionTimeout()
func (jbobject *AmazonWebServiceRequest) GetSdkClientExecutionTimeout() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSdkClientExecutionTimeout", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setSdkClientExecutionTimeout(int)
func (jbobject *AmazonWebServiceRequest) SetSdkClientExecutionTimeout(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSdkClientExecutionTimeout", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *AmazonWebServiceRequest) Clone() *AmazonWebServiceRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/AmazonWebServiceRequest")
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
	unique_x := &AmazonWebServiceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *AmazonWebServiceRequest) Clone2() (*JavaLangObject, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "java/lang/Object")
	if err != nil {
		var zero *JavaLangObject
		return zero, err
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x, nil
}

func (jbobject *AmazonWebServiceRequest) NOOP() *AmazonWebServiceRequest {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/AmazonWebServiceRequest", "NOOP", "com/amazonaws/AmazonWebServiceRequest")
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
	unique_x := &AmazonWebServiceRequest{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *AmazonWebServiceRequest) SetFieldNOOP(val AmazonWebServiceRequestInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/AmazonWebServiceRequest", "NOOP", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


