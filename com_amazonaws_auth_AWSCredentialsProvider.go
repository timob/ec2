package ec2

import "github.com/timob/javabind"

type AuthAWSCredentialsProviderInterface interface {

	// public abstract com.amazonaws.auth.AWSCredentials getCredentials()
	GetCredentials() *AuthAWSCredentials

	// public abstract void refresh()
	Refresh() 
}

type AuthAWSCredentialsProvider struct {
	JavaLangObject
}

// public abstract com.amazonaws.auth.AWSCredentials getCredentials()
func (jbobject *AuthAWSCredentialsProvider) GetCredentials() *AuthAWSCredentials {
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

// public abstract void refresh()
func (jbobject *AuthAWSCredentialsProvider) Refresh()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "refresh", javabind.Void)
	if err != nil {
		panic(err)
	}

}


