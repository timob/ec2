package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelIamInstanceProfileInterface interface {
	JavaLangObjectInterface

	// public void setArn(java.lang.String)
	SetArn(a string) 

	// public java.lang.String getArn()
	GetArn() string

	// public com.amazonaws.services.ec2.model.IamInstanceProfile withArn(java.lang.String)
	WithArn(a string) *ServicesEc2ModelIamInstanceProfile

	// public void setId(java.lang.String)
	SetId(a string) 

	// public java.lang.String getId()
	GetId() string

	// public com.amazonaws.services.ec2.model.IamInstanceProfile withId(java.lang.String)
	WithId(a string) *ServicesEc2ModelIamInstanceProfile

	// public com.amazonaws.services.ec2.model.IamInstanceProfile clone()
	Clone() *ServicesEc2ModelIamInstanceProfile
}

type ServicesEc2ModelIamInstanceProfile struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.IamInstanceProfile()
func NewServicesEc2ModelIamInstanceProfile() (*ServicesEc2ModelIamInstanceProfile) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/IamInstanceProfile")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelIamInstanceProfile{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setArn(java.lang.String)
func (jbobject *ServicesEc2ModelIamInstanceProfile) SetArn(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setArn", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getArn()
func (jbobject *ServicesEc2ModelIamInstanceProfile) GetArn() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getArn", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.IamInstanceProfile withArn(java.lang.String)
func (jbobject *ServicesEc2ModelIamInstanceProfile) WithArn(a string) *ServicesEc2ModelIamInstanceProfile {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withArn", "com/amazonaws/services/ec2/model/IamInstanceProfile", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelIamInstanceProfile{}
	unique_x.Callable = dst
	return unique_x
}

// public void setId(java.lang.String)
func (jbobject *ServicesEc2ModelIamInstanceProfile) SetId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getId()
func (jbobject *ServicesEc2ModelIamInstanceProfile) GetId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.IamInstanceProfile withId(java.lang.String)
func (jbobject *ServicesEc2ModelIamInstanceProfile) WithId(a string) *ServicesEc2ModelIamInstanceProfile {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withId", "com/amazonaws/services/ec2/model/IamInstanceProfile", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelIamInstanceProfile{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelIamInstanceProfile) ToString() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toString", "java/lang/String")
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

// public boolean equals(java.lang.Object)
func (jbobject *ServicesEc2ModelIamInstanceProfile) Equals(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "equals", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public int hashCode()
func (jbobject *ServicesEc2ModelIamInstanceProfile) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.IamInstanceProfile clone()
func (jbobject *ServicesEc2ModelIamInstanceProfile) Clone() *ServicesEc2ModelIamInstanceProfile {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/IamInstanceProfile")
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
	unique_x := &ServicesEc2ModelIamInstanceProfile{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelIamInstanceProfile) Clone2() (*JavaLangObject, error) {
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


