package ec2

import "github.com/timob/javabind"

type AuthAWSCredentialsInterface interface {

	// public abstract java.lang.String getAWSAccessKeyId()
	GetAWSAccessKeyId() string

	// public abstract java.lang.String getAWSSecretKey()
	GetAWSSecretKey() string
}

type AuthAWSCredentials struct {
	JavaLangObject
}

// public abstract java.lang.String getAWSAccessKeyId()
func (jbobject *AuthAWSCredentials) GetAWSAccessKeyId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAWSAccessKeyId", "java/lang/String")
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

// public abstract java.lang.String getAWSSecretKey()
func (jbobject *AuthAWSCredentials) GetAWSSecretKey() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAWSSecretKey", "java/lang/String")
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


