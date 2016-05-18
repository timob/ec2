package ec2

import "github.com/timob/javabind"

type AmazonWebServiceClientInterface interface {
	JavaLangObjectInterface

	// public void setEndpoint(java.lang.String, java.lang.String, java.lang.String)
	SetEndpoint2(a string, b string, c string) 

	// public final void configureRegion(com.amazonaws.regions.Regions)
	ConfigureRegion(a RegionsRegionsInterface) 

	// public void shutdown()
	Shutdown() 

	// public void addRequestHandler(com.amazonaws.handlers.RequestHandler)
	AddRequestHandler(a HandlersRequestHandlerInterface) 

	// public void addRequestHandler(com.amazonaws.handlers.RequestHandler2)
	AddRequestHandler2(a HandlersRequestHandler2Interface) 

	// public void removeRequestHandler(com.amazonaws.handlers.RequestHandler)
	RemoveRequestHandler(a HandlersRequestHandlerInterface) 

	// public void removeRequestHandler(com.amazonaws.handlers.RequestHandler2)
	RemoveRequestHandler2(a HandlersRequestHandler2Interface) 

	// public void setTimeOffset(int)
	SetTimeOffset(a int) 

	// public com.amazonaws.AmazonWebServiceClient withTimeOffset(int)
	WithTimeOffset(a int) *AmazonWebServiceClient

	// public int getTimeOffset()
	GetTimeOffset() int

	// public com.amazonaws.metrics.RequestMetricCollector getRequestMetricsCollector()
	GetRequestMetricsCollector() *MetricsRequestMetricCollector

	// public java.lang.String getServiceName()
	GetServiceName() string

	// public final void setServiceNameIntern(java.lang.String)
	SetServiceNameIntern(a string) 

	// public final java.lang.String getSignerRegionOverride()
	GetSignerRegionOverride() string

	// public final void setSignerRegionOverride(java.lang.String)
	SetSignerRegionOverride(a string) 
}

type AmazonWebServiceClient struct {
	JavaLangObject
}

// public com.amazonaws.AmazonWebServiceClient(com.amazonaws.ClientConfiguration)
func NewAmazonWebServiceClient(a ClientConfigurationInterface) (*AmazonWebServiceClient) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/AmazonWebServiceClient", conv_a.Value().Cast("com/amazonaws/ClientConfiguration"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AmazonWebServiceClient{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.AmazonWebServiceClient(com.amazonaws.ClientConfiguration, com.amazonaws.metrics.RequestMetricCollector)
func NewAmazonWebServiceClient2(a ClientConfigurationInterface, b MetricsRequestMetricCollectorInterface) (*AmazonWebServiceClient) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/AmazonWebServiceClient", conv_a.Value().Cast("com/amazonaws/ClientConfiguration"), conv_b.Value().Cast("com/amazonaws/metrics/RequestMetricCollector"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AmazonWebServiceClient{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setEndpoint(java.lang.String) throws java.lang.IllegalArgumentException
func (jbobject *AmazonWebServiceClient) SetEndpoint(a string) error {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEndpoint", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		return err
	}
	conv_a.CleanUp()
	return nil
}

// public void setEndpoint(java.lang.String, java.lang.String, java.lang.String)
func (jbobject *AmazonWebServiceClient) SetEndpoint2(a string, b string, c string)  {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEndpoint", javabind.Void, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()

}

// public void setRegion(com.amazonaws.regions.Region) throws java.lang.IllegalArgumentException
func (jbobject *AmazonWebServiceClient) SetRegion(a RegionsRegionInterface) error {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRegion", javabind.Void, conv_a.Value().Cast("com/amazonaws/regions/Region"))
	if err != nil {
		return err
	}
	conv_a.CleanUp()
	return nil
}

// public final void configureRegion(com.amazonaws.regions.Regions)
func (jbobject *AmazonWebServiceClient) ConfigureRegion(a RegionsRegionsInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "configureRegion", javabind.Void, conv_a.Value().Cast("com/amazonaws/regions/Regions"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void shutdown()
func (jbobject *AmazonWebServiceClient) Shutdown()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "shutdown", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void addRequestHandler(com.amazonaws.handlers.RequestHandler)
func (jbobject *AmazonWebServiceClient) AddRequestHandler(a HandlersRequestHandlerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addRequestHandler", javabind.Void, conv_a.Value().Cast("com/amazonaws/handlers/RequestHandler"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addRequestHandler(com.amazonaws.handlers.RequestHandler2)
func (jbobject *AmazonWebServiceClient) AddRequestHandler2(a HandlersRequestHandler2Interface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addRequestHandler", javabind.Void, conv_a.Value().Cast("com/amazonaws/handlers/RequestHandler2"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeRequestHandler(com.amazonaws.handlers.RequestHandler)
func (jbobject *AmazonWebServiceClient) RemoveRequestHandler(a HandlersRequestHandlerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeRequestHandler", javabind.Void, conv_a.Value().Cast("com/amazonaws/handlers/RequestHandler"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeRequestHandler(com.amazonaws.handlers.RequestHandler2)
func (jbobject *AmazonWebServiceClient) RemoveRequestHandler2(a HandlersRequestHandler2Interface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeRequestHandler", javabind.Void, conv_a.Value().Cast("com/amazonaws/handlers/RequestHandler2"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setTimeOffset(int)
func (jbobject *AmazonWebServiceClient) SetTimeOffset(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTimeOffset", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public com.amazonaws.AmazonWebServiceClient withTimeOffset(int)
func (jbobject *AmazonWebServiceClient) WithTimeOffset(a int) *AmazonWebServiceClient {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTimeOffset", "com/amazonaws/AmazonWebServiceClient", a)
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
	unique_x := &AmazonWebServiceClient{}
	unique_x.Callable = dst
	return unique_x
}

// public int getTimeOffset()
func (jbobject *AmazonWebServiceClient) GetTimeOffset() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTimeOffset", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.metrics.RequestMetricCollector getRequestMetricsCollector()
func (jbobject *AmazonWebServiceClient) GetRequestMetricsCollector() *MetricsRequestMetricCollector {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRequestMetricsCollector", "com/amazonaws/metrics/RequestMetricCollector")
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

// public java.lang.String getServiceName()
func (jbobject *AmazonWebServiceClient) GetServiceName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getServiceName", "java/lang/String")
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

// public final void setServiceNameIntern(java.lang.String)
func (jbobject *AmazonWebServiceClient) SetServiceNameIntern(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setServiceNameIntern", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public final java.lang.String getSignerRegionOverride()
func (jbobject *AmazonWebServiceClient) GetSignerRegionOverride() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSignerRegionOverride", "java/lang/String")
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

// public final void setSignerRegionOverride(java.lang.String)
func (jbobject *AmazonWebServiceClient) SetSignerRegionOverride(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSignerRegionOverride", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

func (jbobject *AmazonWebServiceClient) LOGGING_AWS_REQUEST_METRIC() bool {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/AmazonWebServiceClient", "LOGGING_AWS_REQUEST_METRIC", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *AmazonWebServiceClient) SetFieldLOGGING_AWS_REQUEST_METRIC(val bool) {
	err := javabind.GetEnv().SetStaticField("com/amazonaws/AmazonWebServiceClient", "LOGGING_AWS_REQUEST_METRIC", val)
	if err != nil {
		panic(err)
	}

}


