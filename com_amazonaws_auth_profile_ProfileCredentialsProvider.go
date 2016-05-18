package ec2

import "github.com/timob/javabind"

type AuthProfileProfileCredentialsProviderInterface interface {
	JavaLangObjectInterface

	// public com.amazonaws.auth.AWSCredentials getCredentials()
	GetCredentials() *AuthAWSCredentials

	// public void refresh()
	Refresh() 

	// public long getRefreshIntervalNanos()
	GetRefreshIntervalNanos() int64

	// public void setRefreshIntervalNanos(long)
	SetRefreshIntervalNanos(a int64) 

	// public long getRefreshForceIntervalNanos()
	GetRefreshForceIntervalNanos() int64

	// public void setRefreshForceIntervalNanos(long)
	SetRefreshForceIntervalNanos(a int64) 
}

type AuthProfileProfileCredentialsProvider struct {
	JavaLangObject
}

// public com.amazonaws.auth.profile.ProfileCredentialsProvider()
func NewAuthProfileProfileCredentialsProvider() (*AuthProfileProfileCredentialsProvider) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/auth/profile/ProfileCredentialsProvider")
	if err != nil {
		panic(err)
	}
	x := &AuthProfileProfileCredentialsProvider{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.auth.profile.ProfileCredentialsProvider(java.lang.String)
func NewAuthProfileProfileCredentialsProvider2(a string) (*AuthProfileProfileCredentialsProvider) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/auth/profile/ProfileCredentialsProvider", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AuthProfileProfileCredentialsProvider{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.auth.profile.ProfileCredentialsProvider(java.lang.String, java.lang.String)
func NewAuthProfileProfileCredentialsProvider4(a string, b string) (*AuthProfileProfileCredentialsProvider) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/auth/profile/ProfileCredentialsProvider", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AuthProfileProfileCredentialsProvider{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.auth.profile.ProfileCredentialsProvider(com.amazonaws.auth.profile.ProfilesConfigFile, java.lang.String)
func NewAuthProfileProfileCredentialsProvider3(a AuthProfileProfilesConfigFileInterface, b string) (*AuthProfileProfileCredentialsProvider) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/auth/profile/ProfileCredentialsProvider", conv_a.Value().Cast("com/amazonaws/auth/profile/ProfilesConfigFile"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AuthProfileProfileCredentialsProvider{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.auth.AWSCredentials getCredentials()
func (jbobject *AuthProfileProfileCredentialsProvider) GetCredentials() *AuthAWSCredentials {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCredentials", "com/amazonaws/auth/AWSCredentials")
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

// public void refresh()
func (jbobject *AuthProfileProfileCredentialsProvider) Refresh()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "refresh", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public long getRefreshIntervalNanos()
func (jbobject *AuthProfileProfileCredentialsProvider) GetRefreshIntervalNanos() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRefreshIntervalNanos", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public void setRefreshIntervalNanos(long)
func (jbobject *AuthProfileProfileCredentialsProvider) SetRefreshIntervalNanos(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRefreshIntervalNanos", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public long getRefreshForceIntervalNanos()
func (jbobject *AuthProfileProfileCredentialsProvider) GetRefreshForceIntervalNanos() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRefreshForceIntervalNanos", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public void setRefreshForceIntervalNanos(long)
func (jbobject *AuthProfileProfileCredentialsProvider) SetRefreshForceIntervalNanos(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRefreshForceIntervalNanos", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


