package ec2

import "github.com/timob/javabind"

type ClientConfigurationInterface interface {
	JavaLangObjectInterface

	// public com.amazonaws.Protocol getProtocol()
	GetProtocol() *Protocol

	// public void setProtocol(com.amazonaws.Protocol)
	SetProtocol(a ProtocolInterface) 

	// public com.amazonaws.ClientConfiguration withProtocol(com.amazonaws.Protocol)
	WithProtocol(a ProtocolInterface) *ClientConfiguration

	// public int getMaxConnections()
	GetMaxConnections() int

	// public void setMaxConnections(int)
	SetMaxConnections(a int) 

	// public com.amazonaws.ClientConfiguration withMaxConnections(int)
	WithMaxConnections(a int) *ClientConfiguration

	// public java.lang.String getUserAgent()
	GetUserAgent() string

	// public void setUserAgent(java.lang.String)
	SetUserAgent(a string) 

	// public com.amazonaws.ClientConfiguration withUserAgent(java.lang.String)
	WithUserAgent(a string) *ClientConfiguration

	// public java.lang.String getProxyHost()
	GetProxyHost() string

	// public void setProxyHost(java.lang.String)
	SetProxyHost(a string) 

	// public com.amazonaws.ClientConfiguration withProxyHost(java.lang.String)
	WithProxyHost(a string) *ClientConfiguration

	// public int getProxyPort()
	GetProxyPort() int

	// public void setProxyPort(int)
	SetProxyPort(a int) 

	// public com.amazonaws.ClientConfiguration withProxyPort(int)
	WithProxyPort(a int) *ClientConfiguration

	// public java.lang.String getProxyUsername()
	GetProxyUsername() string

	// public void setProxyUsername(java.lang.String)
	SetProxyUsername(a string) 

	// public com.amazonaws.ClientConfiguration withProxyUsername(java.lang.String)
	WithProxyUsername(a string) *ClientConfiguration

	// public java.lang.String getProxyPassword()
	GetProxyPassword() string

	// public void setProxyPassword(java.lang.String)
	SetProxyPassword(a string) 

	// public com.amazonaws.ClientConfiguration withProxyPassword(java.lang.String)
	WithProxyPassword(a string) *ClientConfiguration

	// public java.lang.String getProxyDomain()
	GetProxyDomain() string

	// public void setProxyDomain(java.lang.String)
	SetProxyDomain(a string) 

	// public com.amazonaws.ClientConfiguration withProxyDomain(java.lang.String)
	WithProxyDomain(a string) *ClientConfiguration

	// public java.lang.String getProxyWorkstation()
	GetProxyWorkstation() string

	// public void setProxyWorkstation(java.lang.String)
	SetProxyWorkstation(a string) 

	// public com.amazonaws.ClientConfiguration withProxyWorkstation(java.lang.String)
	WithProxyWorkstation(a string) *ClientConfiguration

	// public int getMaxErrorRetry()
	GetMaxErrorRetry() int

	// public void setMaxErrorRetry(int)
	SetMaxErrorRetry(a int) 

	// public com.amazonaws.ClientConfiguration withMaxErrorRetry(int)
	WithMaxErrorRetry(a int) *ClientConfiguration

	// public int getSocketTimeout()
	GetSocketTimeout() int

	// public void setSocketTimeout(int)
	SetSocketTimeout(a int) 

	// public com.amazonaws.ClientConfiguration withSocketTimeout(int)
	WithSocketTimeout(a int) *ClientConfiguration

	// public int getConnectionTimeout()
	GetConnectionTimeout() int

	// public void setConnectionTimeout(int)
	SetConnectionTimeout(a int) 

	// public com.amazonaws.ClientConfiguration withConnectionTimeout(int)
	WithConnectionTimeout(a int) *ClientConfiguration

	// public int getRequestTimeout()
	GetRequestTimeout() int

	// public void setRequestTimeout(int)
	SetRequestTimeout(a int) 

	// public com.amazonaws.ClientConfiguration withRequestTimeout(int)
	WithRequestTimeout(a int) *ClientConfiguration

	// public int getClientExecutionTimeout()
	GetClientExecutionTimeout() int

	// public void setClientExecutionTimeout(int)
	SetClientExecutionTimeout(a int) 

	// public com.amazonaws.ClientConfiguration withClientExecutionTimeout(int)
	WithClientExecutionTimeout(a int) *ClientConfiguration

	// public boolean useReaper()
	UseReaper() bool

	// public void setUseReaper(boolean)
	SetUseReaper(a bool) 

	// public com.amazonaws.ClientConfiguration withReaper(boolean)
	WithReaper(a bool) *ClientConfiguration

	// public boolean useGzip()
	UseGzip() bool

	// public void setUseGzip(boolean)
	SetUseGzip(a bool) 

	// public com.amazonaws.ClientConfiguration withGzip(boolean)
	WithGzip(a bool) *ClientConfiguration

	// public int[] getSocketBufferSizeHints()
	GetSocketBufferSizeHints() []int

	// public void setSocketBufferSizeHints(int, int)
	SetSocketBufferSizeHints(a int, b int) 

	// public com.amazonaws.ClientConfiguration withSocketBufferSizeHints(int, int)
	WithSocketBufferSizeHints(a int, b int) *ClientConfiguration

	// public java.lang.String getSignerOverride()
	GetSignerOverride() string

	// public void setSignerOverride(java.lang.String)
	SetSignerOverride(a string) 

	// public com.amazonaws.ClientConfiguration withSignerOverride(java.lang.String)
	WithSignerOverride(a string) *ClientConfiguration

	// public boolean isPreemptiveBasicProxyAuth()
	IsPreemptiveBasicProxyAuth() bool

	// public void setPreemptiveBasicProxyAuth(java.lang.Boolean)
	SetPreemptiveBasicProxyAuth(a bool) 

	// public com.amazonaws.ClientConfiguration withPreemptiveBasicProxyAuth(boolean)
	WithPreemptiveBasicProxyAuth(a bool) *ClientConfiguration

	// public long getConnectionTTL()
	GetConnectionTTL() int64

	// public void setConnectionTTL(long)
	SetConnectionTTL(a int64) 

	// public com.amazonaws.ClientConfiguration withConnectionTTL(long)
	WithConnectionTTL(a int64) *ClientConfiguration

	// public long getConnectionMaxIdleMillis()
	GetConnectionMaxIdleMillis() int64

	// public void setConnectionMaxIdleMillis(long)
	SetConnectionMaxIdleMillis(a int64) 

	// public com.amazonaws.ClientConfiguration withConnectionMaxIdleMillis(long)
	WithConnectionMaxIdleMillis(a int64) *ClientConfiguration

	// public boolean useTcpKeepAlive()
	UseTcpKeepAlive() bool

	// public void setUseTcpKeepAlive(boolean)
	SetUseTcpKeepAlive(a bool) 

	// public com.amazonaws.ClientConfiguration withTcpKeepAlive(boolean)
	WithTcpKeepAlive(a bool) *ClientConfiguration

	// public com.amazonaws.DnsResolver getDnsResolver()
	GetDnsResolver() *DnsResolver

	// public void setDnsResolver(com.amazonaws.DnsResolver)
	SetDnsResolver(a DnsResolverInterface) 

	// public com.amazonaws.ClientConfiguration withDnsResolver(com.amazonaws.DnsResolver)
	WithDnsResolver(a DnsResolverInterface) *ClientConfiguration

	// public int getResponseMetadataCacheSize()
	GetResponseMetadataCacheSize() int

	// public void setResponseMetadataCacheSize(int)
	SetResponseMetadataCacheSize(a int) 

	// public com.amazonaws.ClientConfiguration withResponseMetadataCacheSize(int)
	WithResponseMetadataCacheSize(a int) *ClientConfiguration

	// public com.amazonaws.ApacheHttpClientConfig getApacheHttpClientConfig()
	GetApacheHttpClientConfig() *ApacheHttpClientConfig

	// public boolean isUseExpectContinue()
	IsUseExpectContinue() bool

	// public void setUseExpectContinue(boolean)
	SetUseExpectContinue(a bool) 

	// public com.amazonaws.ClientConfiguration withUseExpectContinue(boolean)
	WithUseExpectContinue(a bool) *ClientConfiguration
}

type ClientConfiguration struct {
	JavaLangObject
}

// public com.amazonaws.ClientConfiguration()
func NewClientConfiguration() (*ClientConfiguration) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/ClientConfiguration")
	if err != nil {
		panic(err)
	}
	x := &ClientConfiguration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.ClientConfiguration(com.amazonaws.ClientConfiguration)
func NewClientConfiguration2(a ClientConfigurationInterface) (*ClientConfiguration) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/ClientConfiguration", conv_a.Value().Cast("com/amazonaws/ClientConfiguration"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ClientConfiguration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.Protocol getProtocol()
func (jbobject *ClientConfiguration) GetProtocol() *Protocol {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProtocol", "com/amazonaws/Protocol")
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
	unique_x := &Protocol{}
	unique_x.Callable = dst
	return unique_x
}

// public void setProtocol(com.amazonaws.Protocol)
func (jbobject *ClientConfiguration) SetProtocol(a ProtocolInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProtocol", javabind.Void, conv_a.Value().Cast("com/amazonaws/Protocol"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.ClientConfiguration withProtocol(com.amazonaws.Protocol)
func (jbobject *ClientConfiguration) WithProtocol(a ProtocolInterface) *ClientConfiguration {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProtocol", "com/amazonaws/ClientConfiguration", conv_a.Value().Cast("com/amazonaws/Protocol"))
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public int getMaxConnections()
func (jbobject *ClientConfiguration) GetMaxConnections() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMaxConnections", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void setMaxConnections(int)
func (jbobject *ClientConfiguration) SetMaxConnections(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMaxConnections", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public com.amazonaws.ClientConfiguration withMaxConnections(int)
func (jbobject *ClientConfiguration) WithMaxConnections(a int) *ClientConfiguration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMaxConnections", "com/amazonaws/ClientConfiguration", a)
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String getUserAgent()
func (jbobject *ClientConfiguration) GetUserAgent() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUserAgent", "java/lang/String")
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

// public void setUserAgent(java.lang.String)
func (jbobject *ClientConfiguration) SetUserAgent(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUserAgent", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.ClientConfiguration withUserAgent(java.lang.String)
func (jbobject *ClientConfiguration) WithUserAgent(a string) *ClientConfiguration {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUserAgent", "com/amazonaws/ClientConfiguration", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String getProxyHost()
func (jbobject *ClientConfiguration) GetProxyHost() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProxyHost", "java/lang/String")
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

// public void setProxyHost(java.lang.String)
func (jbobject *ClientConfiguration) SetProxyHost(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProxyHost", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.ClientConfiguration withProxyHost(java.lang.String)
func (jbobject *ClientConfiguration) WithProxyHost(a string) *ClientConfiguration {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProxyHost", "com/amazonaws/ClientConfiguration", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public int getProxyPort()
func (jbobject *ClientConfiguration) GetProxyPort() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProxyPort", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void setProxyPort(int)
func (jbobject *ClientConfiguration) SetProxyPort(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProxyPort", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public com.amazonaws.ClientConfiguration withProxyPort(int)
func (jbobject *ClientConfiguration) WithProxyPort(a int) *ClientConfiguration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProxyPort", "com/amazonaws/ClientConfiguration", a)
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String getProxyUsername()
func (jbobject *ClientConfiguration) GetProxyUsername() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProxyUsername", "java/lang/String")
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

// public void setProxyUsername(java.lang.String)
func (jbobject *ClientConfiguration) SetProxyUsername(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProxyUsername", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.ClientConfiguration withProxyUsername(java.lang.String)
func (jbobject *ClientConfiguration) WithProxyUsername(a string) *ClientConfiguration {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProxyUsername", "com/amazonaws/ClientConfiguration", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String getProxyPassword()
func (jbobject *ClientConfiguration) GetProxyPassword() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProxyPassword", "java/lang/String")
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

// public void setProxyPassword(java.lang.String)
func (jbobject *ClientConfiguration) SetProxyPassword(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProxyPassword", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.ClientConfiguration withProxyPassword(java.lang.String)
func (jbobject *ClientConfiguration) WithProxyPassword(a string) *ClientConfiguration {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProxyPassword", "com/amazonaws/ClientConfiguration", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String getProxyDomain()
func (jbobject *ClientConfiguration) GetProxyDomain() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProxyDomain", "java/lang/String")
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

// public void setProxyDomain(java.lang.String)
func (jbobject *ClientConfiguration) SetProxyDomain(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProxyDomain", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.ClientConfiguration withProxyDomain(java.lang.String)
func (jbobject *ClientConfiguration) WithProxyDomain(a string) *ClientConfiguration {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProxyDomain", "com/amazonaws/ClientConfiguration", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String getProxyWorkstation()
func (jbobject *ClientConfiguration) GetProxyWorkstation() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProxyWorkstation", "java/lang/String")
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

// public void setProxyWorkstation(java.lang.String)
func (jbobject *ClientConfiguration) SetProxyWorkstation(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProxyWorkstation", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.ClientConfiguration withProxyWorkstation(java.lang.String)
func (jbobject *ClientConfiguration) WithProxyWorkstation(a string) *ClientConfiguration {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProxyWorkstation", "com/amazonaws/ClientConfiguration", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public int getMaxErrorRetry()
func (jbobject *ClientConfiguration) GetMaxErrorRetry() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMaxErrorRetry", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void setMaxErrorRetry(int)
func (jbobject *ClientConfiguration) SetMaxErrorRetry(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMaxErrorRetry", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public com.amazonaws.ClientConfiguration withMaxErrorRetry(int)
func (jbobject *ClientConfiguration) WithMaxErrorRetry(a int) *ClientConfiguration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMaxErrorRetry", "com/amazonaws/ClientConfiguration", a)
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public int getSocketTimeout()
func (jbobject *ClientConfiguration) GetSocketTimeout() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSocketTimeout", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void setSocketTimeout(int)
func (jbobject *ClientConfiguration) SetSocketTimeout(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSocketTimeout", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public com.amazonaws.ClientConfiguration withSocketTimeout(int)
func (jbobject *ClientConfiguration) WithSocketTimeout(a int) *ClientConfiguration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSocketTimeout", "com/amazonaws/ClientConfiguration", a)
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public int getConnectionTimeout()
func (jbobject *ClientConfiguration) GetConnectionTimeout() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getConnectionTimeout", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void setConnectionTimeout(int)
func (jbobject *ClientConfiguration) SetConnectionTimeout(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setConnectionTimeout", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public com.amazonaws.ClientConfiguration withConnectionTimeout(int)
func (jbobject *ClientConfiguration) WithConnectionTimeout(a int) *ClientConfiguration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withConnectionTimeout", "com/amazonaws/ClientConfiguration", a)
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public int getRequestTimeout()
func (jbobject *ClientConfiguration) GetRequestTimeout() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRequestTimeout", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void setRequestTimeout(int)
func (jbobject *ClientConfiguration) SetRequestTimeout(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRequestTimeout", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public com.amazonaws.ClientConfiguration withRequestTimeout(int)
func (jbobject *ClientConfiguration) WithRequestTimeout(a int) *ClientConfiguration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRequestTimeout", "com/amazonaws/ClientConfiguration", a)
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public int getClientExecutionTimeout()
func (jbobject *ClientConfiguration) GetClientExecutionTimeout() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getClientExecutionTimeout", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void setClientExecutionTimeout(int)
func (jbobject *ClientConfiguration) SetClientExecutionTimeout(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setClientExecutionTimeout", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public com.amazonaws.ClientConfiguration withClientExecutionTimeout(int)
func (jbobject *ClientConfiguration) WithClientExecutionTimeout(a int) *ClientConfiguration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClientExecutionTimeout", "com/amazonaws/ClientConfiguration", a)
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean useReaper()
func (jbobject *ClientConfiguration) UseReaper() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "useReaper", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setUseReaper(boolean)
func (jbobject *ClientConfiguration) SetUseReaper(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUseReaper", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public com.amazonaws.ClientConfiguration withReaper(boolean)
func (jbobject *ClientConfiguration) WithReaper(a bool) *ClientConfiguration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReaper", "com/amazonaws/ClientConfiguration", a)
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean useGzip()
func (jbobject *ClientConfiguration) UseGzip() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "useGzip", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setUseGzip(boolean)
func (jbobject *ClientConfiguration) SetUseGzip(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUseGzip", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public com.amazonaws.ClientConfiguration withGzip(boolean)
func (jbobject *ClientConfiguration) WithGzip(a bool) *ClientConfiguration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGzip", "com/amazonaws/ClientConfiguration", a)
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public int[] getSocketBufferSizeHints()
func (jbobject *ClientConfiguration) GetSocketBufferSizeHints() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSocketBufferSizeHints", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

// public void setSocketBufferSizeHints(int, int)
func (jbobject *ClientConfiguration) SetSocketBufferSizeHints(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSocketBufferSizeHints", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public com.amazonaws.ClientConfiguration withSocketBufferSizeHints(int, int)
func (jbobject *ClientConfiguration) WithSocketBufferSizeHints(a int, b int) *ClientConfiguration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSocketBufferSizeHints", "com/amazonaws/ClientConfiguration", a, b)
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String getSignerOverride()
func (jbobject *ClientConfiguration) GetSignerOverride() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSignerOverride", "java/lang/String")
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

// public void setSignerOverride(java.lang.String)
func (jbobject *ClientConfiguration) SetSignerOverride(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSignerOverride", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.ClientConfiguration withSignerOverride(java.lang.String)
func (jbobject *ClientConfiguration) WithSignerOverride(a string) *ClientConfiguration {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSignerOverride", "com/amazonaws/ClientConfiguration", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean isPreemptiveBasicProxyAuth()
func (jbobject *ClientConfiguration) IsPreemptiveBasicProxyAuth() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isPreemptiveBasicProxyAuth", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setPreemptiveBasicProxyAuth(java.lang.Boolean)
func (jbobject *ClientConfiguration) SetPreemptiveBasicProxyAuth(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPreemptiveBasicProxyAuth", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.ClientConfiguration withPreemptiveBasicProxyAuth(boolean)
func (jbobject *ClientConfiguration) WithPreemptiveBasicProxyAuth(a bool) *ClientConfiguration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPreemptiveBasicProxyAuth", "com/amazonaws/ClientConfiguration", a)
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public long getConnectionTTL()
func (jbobject *ClientConfiguration) GetConnectionTTL() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getConnectionTTL", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public void setConnectionTTL(long)
func (jbobject *ClientConfiguration) SetConnectionTTL(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setConnectionTTL", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public com.amazonaws.ClientConfiguration withConnectionTTL(long)
func (jbobject *ClientConfiguration) WithConnectionTTL(a int64) *ClientConfiguration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withConnectionTTL", "com/amazonaws/ClientConfiguration", a)
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public long getConnectionMaxIdleMillis()
func (jbobject *ClientConfiguration) GetConnectionMaxIdleMillis() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getConnectionMaxIdleMillis", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public void setConnectionMaxIdleMillis(long)
func (jbobject *ClientConfiguration) SetConnectionMaxIdleMillis(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setConnectionMaxIdleMillis", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public com.amazonaws.ClientConfiguration withConnectionMaxIdleMillis(long)
func (jbobject *ClientConfiguration) WithConnectionMaxIdleMillis(a int64) *ClientConfiguration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withConnectionMaxIdleMillis", "com/amazonaws/ClientConfiguration", a)
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean useTcpKeepAlive()
func (jbobject *ClientConfiguration) UseTcpKeepAlive() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "useTcpKeepAlive", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setUseTcpKeepAlive(boolean)
func (jbobject *ClientConfiguration) SetUseTcpKeepAlive(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUseTcpKeepAlive", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public com.amazonaws.ClientConfiguration withTcpKeepAlive(boolean)
func (jbobject *ClientConfiguration) WithTcpKeepAlive(a bool) *ClientConfiguration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTcpKeepAlive", "com/amazonaws/ClientConfiguration", a)
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.DnsResolver getDnsResolver()
func (jbobject *ClientConfiguration) GetDnsResolver() *DnsResolver {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDnsResolver", "com/amazonaws/DnsResolver")
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
	unique_x := &DnsResolver{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDnsResolver(com.amazonaws.DnsResolver)
func (jbobject *ClientConfiguration) SetDnsResolver(a DnsResolverInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDnsResolver", javabind.Void, conv_a.Value().Cast("com/amazonaws/DnsResolver"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.ClientConfiguration withDnsResolver(com.amazonaws.DnsResolver)
func (jbobject *ClientConfiguration) WithDnsResolver(a DnsResolverInterface) *ClientConfiguration {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDnsResolver", "com/amazonaws/ClientConfiguration", conv_a.Value().Cast("com/amazonaws/DnsResolver"))
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public int getResponseMetadataCacheSize()
func (jbobject *ClientConfiguration) GetResponseMetadataCacheSize() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getResponseMetadataCacheSize", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void setResponseMetadataCacheSize(int)
func (jbobject *ClientConfiguration) SetResponseMetadataCacheSize(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setResponseMetadataCacheSize", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public com.amazonaws.ClientConfiguration withResponseMetadataCacheSize(int)
func (jbobject *ClientConfiguration) WithResponseMetadataCacheSize(a int) *ClientConfiguration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withResponseMetadataCacheSize", "com/amazonaws/ClientConfiguration", a)
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.ApacheHttpClientConfig getApacheHttpClientConfig()
func (jbobject *ClientConfiguration) GetApacheHttpClientConfig() *ApacheHttpClientConfig {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getApacheHttpClientConfig", "com/amazonaws/ApacheHttpClientConfig")
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
	unique_x := &ApacheHttpClientConfig{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean isUseExpectContinue()
func (jbobject *ClientConfiguration) IsUseExpectContinue() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isUseExpectContinue", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setUseExpectContinue(boolean)
func (jbobject *ClientConfiguration) SetUseExpectContinue(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUseExpectContinue", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public com.amazonaws.ClientConfiguration withUseExpectContinue(boolean)
func (jbobject *ClientConfiguration) WithUseExpectContinue(a bool) *ClientConfiguration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUseExpectContinue", "com/amazonaws/ClientConfiguration", a)
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
	unique_x := &ClientConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ClientConfiguration) DEFAULT_CONNECTION_TIMEOUT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_CONNECTION_TIMEOUT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ClientConfiguration) SetFieldDEFAULT_CONNECTION_TIMEOUT(val int) {
	err := javabind.GetEnv().SetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_CONNECTION_TIMEOUT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ClientConfiguration) DEFAULT_SOCKET_TIMEOUT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_SOCKET_TIMEOUT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ClientConfiguration) SetFieldDEFAULT_SOCKET_TIMEOUT(val int) {
	err := javabind.GetEnv().SetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_SOCKET_TIMEOUT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ClientConfiguration) DEFAULT_REQUEST_TIMEOUT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_REQUEST_TIMEOUT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ClientConfiguration) SetFieldDEFAULT_REQUEST_TIMEOUT(val int) {
	err := javabind.GetEnv().SetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_REQUEST_TIMEOUT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ClientConfiguration) DEFAULT_CLIENT_EXECUTION_TIMEOUT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_CLIENT_EXECUTION_TIMEOUT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ClientConfiguration) SetFieldDEFAULT_CLIENT_EXECUTION_TIMEOUT(val int) {
	err := javabind.GetEnv().SetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_CLIENT_EXECUTION_TIMEOUT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ClientConfiguration) DEFAULT_MAX_CONNECTIONS() int {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_MAX_CONNECTIONS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ClientConfiguration) SetFieldDEFAULT_MAX_CONNECTIONS(val int) {
	err := javabind.GetEnv().SetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_MAX_CONNECTIONS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ClientConfiguration) DEFAULT_USER_AGENT() string {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_USER_AGENT", "java/lang/String")
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

func (jbobject *ClientConfiguration) SetFieldDEFAULT_USER_AGENT(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_USER_AGENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ClientConfiguration) DEFAULT_USE_REAPER() bool {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_USE_REAPER", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *ClientConfiguration) SetFieldDEFAULT_USE_REAPER(val bool) {
	err := javabind.GetEnv().SetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_USE_REAPER", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ClientConfiguration) DEFAULT_USE_GZIP() bool {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_USE_GZIP", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *ClientConfiguration) SetFieldDEFAULT_USE_GZIP(val bool) {
	err := javabind.GetEnv().SetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_USE_GZIP", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ClientConfiguration) DEFAULT_CONNECTION_TTL() int64 {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_CONNECTION_TTL", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *ClientConfiguration) SetFieldDEFAULT_CONNECTION_TTL(val int64) {
	err := javabind.GetEnv().SetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_CONNECTION_TTL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ClientConfiguration) DEFAULT_CONNECTION_MAX_IDLE_MILLIS() int64 {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_CONNECTION_MAX_IDLE_MILLIS", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *ClientConfiguration) SetFieldDEFAULT_CONNECTION_MAX_IDLE_MILLIS(val int64) {
	err := javabind.GetEnv().SetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_CONNECTION_MAX_IDLE_MILLIS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ClientConfiguration) DEFAULT_TCP_KEEP_ALIVE() bool {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_TCP_KEEP_ALIVE", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *ClientConfiguration) SetFieldDEFAULT_TCP_KEEP_ALIVE(val bool) {
	err := javabind.GetEnv().SetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_TCP_KEEP_ALIVE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ClientConfiguration) DEFAULT_RESPONSE_METADATA_CACHE_SIZE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_RESPONSE_METADATA_CACHE_SIZE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ClientConfiguration) SetFieldDEFAULT_RESPONSE_METADATA_CACHE_SIZE(val int) {
	err := javabind.GetEnv().SetStaticField("com/amazonaws/ClientConfiguration", "DEFAULT_RESPONSE_METADATA_CACHE_SIZE", val)
	if err != nil {
		panic(err)
	}

}


