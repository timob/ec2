package ec2

import "github.com/timob/javabind"

type AuthProfileProfilesConfigFileInterface interface {
	JavaLangObjectInterface

	// public com.amazonaws.auth.AWSCredentials getCredentials(java.lang.String)
	GetCredentials(a string) *AuthAWSCredentials

	// public void refresh()
	Refresh() 
}

type AuthProfileProfilesConfigFile struct {
	JavaLangObject
}

// public com.amazonaws.auth.profile.ProfilesConfigFile() throws com.amazonaws.AmazonClientException
func NewAuthProfileProfilesConfigFile() (*AuthProfileProfilesConfigFile, error) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/auth/profile/ProfilesConfigFile")
	if err != nil {
		return nil, err
	}
	x := &AuthProfileProfilesConfigFile{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public com.amazonaws.auth.profile.ProfilesConfigFile(java.lang.String)
func NewAuthProfileProfilesConfigFile2(a string) (*AuthProfileProfilesConfigFile) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/auth/profile/ProfilesConfigFile", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AuthProfileProfilesConfigFile{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.auth.AWSCredentials getCredentials(java.lang.String)
func (jbobject *AuthProfileProfilesConfigFile) GetCredentials(a string) *AuthAWSCredentials {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCredentials", "com/amazonaws/auth/AWSCredentials", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &AuthAWSCredentials{}
	unique_x.Callable = dst
	return unique_x
}

// public void refresh()
func (jbobject *AuthProfileProfilesConfigFile) Refresh()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "refresh", javabind.Void)
	if err != nil {
		panic(err)
	}

}

func (jbobject *AuthProfileProfilesConfigFile) AWS_PROFILE_ENVIRONMENT_VARIABLE() string {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/auth/profile/ProfilesConfigFile", "AWS_PROFILE_ENVIRONMENT_VARIABLE", "java/lang/String")
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

func (jbobject *AuthProfileProfilesConfigFile) SetFieldAWS_PROFILE_ENVIRONMENT_VARIABLE(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/auth/profile/ProfilesConfigFile", "AWS_PROFILE_ENVIRONMENT_VARIABLE", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *AuthProfileProfilesConfigFile) AWS_PROFILE_SYSTEM_PROPERTY() string {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/auth/profile/ProfilesConfigFile", "AWS_PROFILE_SYSTEM_PROPERTY", "java/lang/String")
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

func (jbobject *AuthProfileProfilesConfigFile) SetFieldAWS_PROFILE_SYSTEM_PROPERTY(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/auth/profile/ProfilesConfigFile", "AWS_PROFILE_SYSTEM_PROPERTY", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *AuthProfileProfilesConfigFile) DEFAULT_PROFILE_NAME() string {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/auth/profile/ProfilesConfigFile", "DEFAULT_PROFILE_NAME", "java/lang/String")
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

func (jbobject *AuthProfileProfilesConfigFile) SetFieldDEFAULT_PROFILE_NAME(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/auth/profile/ProfilesConfigFile", "DEFAULT_PROFILE_NAME", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


