package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelScheduledInstancesIamInstanceProfileInterface interface {
	JavaLangObjectInterface

	// public void setArn(java.lang.String)
	SetArn(a string) 

	// public java.lang.String getArn()
	GetArn() string

	// public com.amazonaws.services.ec2.model.ScheduledInstancesIamInstanceProfile withArn(java.lang.String)
	WithArn(a string) *ServicesEc2ModelScheduledInstancesIamInstanceProfile

	// public void setName(java.lang.String)
	SetName(a string) 

	// public java.lang.String getName()
	GetName() string

	// public com.amazonaws.services.ec2.model.ScheduledInstancesIamInstanceProfile withName(java.lang.String)
	WithName(a string) *ServicesEc2ModelScheduledInstancesIamInstanceProfile

	// public com.amazonaws.services.ec2.model.ScheduledInstancesIamInstanceProfile clone()
	Clone() *ServicesEc2ModelScheduledInstancesIamInstanceProfile
}

type ServicesEc2ModelScheduledInstancesIamInstanceProfile struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ScheduledInstancesIamInstanceProfile()
func NewServicesEc2ModelScheduledInstancesIamInstanceProfile() (*ServicesEc2ModelScheduledInstancesIamInstanceProfile) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ScheduledInstancesIamInstanceProfile")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelScheduledInstancesIamInstanceProfile{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setArn(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesIamInstanceProfile) SetArn(a string)  {
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
func (jbobject *ServicesEc2ModelScheduledInstancesIamInstanceProfile) GetArn() string {
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesIamInstanceProfile withArn(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesIamInstanceProfile) WithArn(a string) *ServicesEc2ModelScheduledInstancesIamInstanceProfile {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withArn", "com/amazonaws/services/ec2/model/ScheduledInstancesIamInstanceProfile", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelScheduledInstancesIamInstanceProfile{}
	unique_x.Callable = dst
	return unique_x
}

// public void setName(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesIamInstanceProfile) SetName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getName()
func (jbobject *ServicesEc2ModelScheduledInstancesIamInstanceProfile) GetName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesIamInstanceProfile withName(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesIamInstanceProfile) WithName(a string) *ServicesEc2ModelScheduledInstancesIamInstanceProfile {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withName", "com/amazonaws/services/ec2/model/ScheduledInstancesIamInstanceProfile", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelScheduledInstancesIamInstanceProfile{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelScheduledInstancesIamInstanceProfile) ToString() string {
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
func (jbobject *ServicesEc2ModelScheduledInstancesIamInstanceProfile) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelScheduledInstancesIamInstanceProfile) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ScheduledInstancesIamInstanceProfile clone()
func (jbobject *ServicesEc2ModelScheduledInstancesIamInstanceProfile) Clone() *ServicesEc2ModelScheduledInstancesIamInstanceProfile {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ScheduledInstancesIamInstanceProfile")
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
	unique_x := &ServicesEc2ModelScheduledInstancesIamInstanceProfile{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelScheduledInstancesIamInstanceProfile) Clone2() (*JavaLangObject, error) {
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


