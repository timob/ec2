package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelPlacementInterface interface {
	JavaLangObjectInterface

	// public void setAvailabilityZone(java.lang.String)
	SetAvailabilityZone(a string) 

	// public java.lang.String getAvailabilityZone()
	GetAvailabilityZone() string

	// public com.amazonaws.services.ec2.model.Placement withAvailabilityZone(java.lang.String)
	WithAvailabilityZone(a string) *ServicesEc2ModelPlacement

	// public void setGroupName(java.lang.String)
	SetGroupName(a string) 

	// public java.lang.String getGroupName()
	GetGroupName() string

	// public com.amazonaws.services.ec2.model.Placement withGroupName(java.lang.String)
	WithGroupName(a string) *ServicesEc2ModelPlacement

	// public void setTenancy(java.lang.String)
	SetTenancy2(a string) 

	// public java.lang.String getTenancy()
	GetTenancy() string

	// public com.amazonaws.services.ec2.model.Placement withTenancy(java.lang.String)
	WithTenancy2(a string) *ServicesEc2ModelPlacement

	// public void setTenancy(com.amazonaws.services.ec2.model.Tenancy)
	SetTenancy(a ServicesEc2ModelTenancyInterface) 

	// public com.amazonaws.services.ec2.model.Placement withTenancy(com.amazonaws.services.ec2.model.Tenancy)
	WithTenancy(a ServicesEc2ModelTenancyInterface) *ServicesEc2ModelPlacement

	// public void setHostId(java.lang.String)
	SetHostId(a string) 

	// public java.lang.String getHostId()
	GetHostId() string

	// public com.amazonaws.services.ec2.model.Placement withHostId(java.lang.String)
	WithHostId(a string) *ServicesEc2ModelPlacement

	// public void setAffinity(java.lang.String)
	SetAffinity(a string) 

	// public java.lang.String getAffinity()
	GetAffinity() string

	// public com.amazonaws.services.ec2.model.Placement withAffinity(java.lang.String)
	WithAffinity(a string) *ServicesEc2ModelPlacement

	// public com.amazonaws.services.ec2.model.Placement clone()
	Clone() *ServicesEc2ModelPlacement
}

type ServicesEc2ModelPlacement struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.Placement()
func NewServicesEc2ModelPlacement() (*ServicesEc2ModelPlacement) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/Placement")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelPlacement{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.Placement(java.lang.String)
func NewServicesEc2ModelPlacement2(a string) (*ServicesEc2ModelPlacement) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/Placement", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ServicesEc2ModelPlacement{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelPlacement) SetAvailabilityZone(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAvailabilityZone", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAvailabilityZone()
func (jbobject *ServicesEc2ModelPlacement) GetAvailabilityZone() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAvailabilityZone", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Placement withAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelPlacement) WithAvailabilityZone(a string) *ServicesEc2ModelPlacement {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailabilityZone", "com/amazonaws/services/ec2/model/Placement", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPlacement{}
	unique_x.Callable = dst
	return unique_x
}

// public void setGroupName(java.lang.String)
func (jbobject *ServicesEc2ModelPlacement) SetGroupName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGroupName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getGroupName()
func (jbobject *ServicesEc2ModelPlacement) GetGroupName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGroupName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Placement withGroupName(java.lang.String)
func (jbobject *ServicesEc2ModelPlacement) WithGroupName(a string) *ServicesEc2ModelPlacement {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroupName", "com/amazonaws/services/ec2/model/Placement", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPlacement{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTenancy(java.lang.String)
func (jbobject *ServicesEc2ModelPlacement) SetTenancy2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTenancy", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getTenancy()
func (jbobject *ServicesEc2ModelPlacement) GetTenancy() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTenancy", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Placement withTenancy(java.lang.String)
func (jbobject *ServicesEc2ModelPlacement) WithTenancy2(a string) *ServicesEc2ModelPlacement {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTenancy", "com/amazonaws/services/ec2/model/Placement", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPlacement{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTenancy(com.amazonaws.services.ec2.model.Tenancy)
func (jbobject *ServicesEc2ModelPlacement) SetTenancy(a ServicesEc2ModelTenancyInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTenancy", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tenancy"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Placement withTenancy(com.amazonaws.services.ec2.model.Tenancy)
func (jbobject *ServicesEc2ModelPlacement) WithTenancy(a ServicesEc2ModelTenancyInterface) *ServicesEc2ModelPlacement {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTenancy", "com/amazonaws/services/ec2/model/Placement", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tenancy"))
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
	unique_x := &ServicesEc2ModelPlacement{}
	unique_x.Callable = dst
	return unique_x
}

// public void setHostId(java.lang.String)
func (jbobject *ServicesEc2ModelPlacement) SetHostId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHostId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getHostId()
func (jbobject *ServicesEc2ModelPlacement) GetHostId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHostId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Placement withHostId(java.lang.String)
func (jbobject *ServicesEc2ModelPlacement) WithHostId(a string) *ServicesEc2ModelPlacement {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withHostId", "com/amazonaws/services/ec2/model/Placement", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPlacement{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAffinity(java.lang.String)
func (jbobject *ServicesEc2ModelPlacement) SetAffinity(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAffinity", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAffinity()
func (jbobject *ServicesEc2ModelPlacement) GetAffinity() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAffinity", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Placement withAffinity(java.lang.String)
func (jbobject *ServicesEc2ModelPlacement) WithAffinity(a string) *ServicesEc2ModelPlacement {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAffinity", "com/amazonaws/services/ec2/model/Placement", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPlacement{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelPlacement) ToString() string {
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
func (jbobject *ServicesEc2ModelPlacement) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelPlacement) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.Placement clone()
func (jbobject *ServicesEc2ModelPlacement) Clone() *ServicesEc2ModelPlacement {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/Placement")
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
	unique_x := &ServicesEc2ModelPlacement{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelPlacement) Clone2() (*JavaLangObject, error) {
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


