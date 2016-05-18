package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelUserIdGroupPairInterface interface {
	JavaLangObjectInterface

	// public void setUserId(java.lang.String)
	SetUserId(a string) 

	// public java.lang.String getUserId()
	GetUserId() string

	// public com.amazonaws.services.ec2.model.UserIdGroupPair withUserId(java.lang.String)
	WithUserId(a string) *ServicesEc2ModelUserIdGroupPair

	// public void setGroupName(java.lang.String)
	SetGroupName(a string) 

	// public java.lang.String getGroupName()
	GetGroupName() string

	// public com.amazonaws.services.ec2.model.UserIdGroupPair withGroupName(java.lang.String)
	WithGroupName(a string) *ServicesEc2ModelUserIdGroupPair

	// public void setGroupId(java.lang.String)
	SetGroupId(a string) 

	// public java.lang.String getGroupId()
	GetGroupId() string

	// public com.amazonaws.services.ec2.model.UserIdGroupPair withGroupId(java.lang.String)
	WithGroupId(a string) *ServicesEc2ModelUserIdGroupPair

	// public void setVpcId(java.lang.String)
	SetVpcId(a string) 

	// public java.lang.String getVpcId()
	GetVpcId() string

	// public com.amazonaws.services.ec2.model.UserIdGroupPair withVpcId(java.lang.String)
	WithVpcId(a string) *ServicesEc2ModelUserIdGroupPair

	// public void setVpcPeeringConnectionId(java.lang.String)
	SetVpcPeeringConnectionId(a string) 

	// public java.lang.String getVpcPeeringConnectionId()
	GetVpcPeeringConnectionId() string

	// public com.amazonaws.services.ec2.model.UserIdGroupPair withVpcPeeringConnectionId(java.lang.String)
	WithVpcPeeringConnectionId(a string) *ServicesEc2ModelUserIdGroupPair

	// public void setPeeringStatus(java.lang.String)
	SetPeeringStatus(a string) 

	// public java.lang.String getPeeringStatus()
	GetPeeringStatus() string

	// public com.amazonaws.services.ec2.model.UserIdGroupPair withPeeringStatus(java.lang.String)
	WithPeeringStatus(a string) *ServicesEc2ModelUserIdGroupPair

	// public com.amazonaws.services.ec2.model.UserIdGroupPair clone()
	Clone() *ServicesEc2ModelUserIdGroupPair
}

type ServicesEc2ModelUserIdGroupPair struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.UserIdGroupPair()
func NewServicesEc2ModelUserIdGroupPair() (*ServicesEc2ModelUserIdGroupPair) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/UserIdGroupPair")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelUserIdGroupPair{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setUserId(java.lang.String)
func (jbobject *ServicesEc2ModelUserIdGroupPair) SetUserId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUserId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getUserId()
func (jbobject *ServicesEc2ModelUserIdGroupPair) GetUserId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUserId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.UserIdGroupPair withUserId(java.lang.String)
func (jbobject *ServicesEc2ModelUserIdGroupPair) WithUserId(a string) *ServicesEc2ModelUserIdGroupPair {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUserId", "com/amazonaws/services/ec2/model/UserIdGroupPair", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelUserIdGroupPair{}
	unique_x.Callable = dst
	return unique_x
}

// public void setGroupName(java.lang.String)
func (jbobject *ServicesEc2ModelUserIdGroupPair) SetGroupName(a string)  {
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
func (jbobject *ServicesEc2ModelUserIdGroupPair) GetGroupName() string {
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

// public com.amazonaws.services.ec2.model.UserIdGroupPair withGroupName(java.lang.String)
func (jbobject *ServicesEc2ModelUserIdGroupPair) WithGroupName(a string) *ServicesEc2ModelUserIdGroupPair {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroupName", "com/amazonaws/services/ec2/model/UserIdGroupPair", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelUserIdGroupPair{}
	unique_x.Callable = dst
	return unique_x
}

// public void setGroupId(java.lang.String)
func (jbobject *ServicesEc2ModelUserIdGroupPair) SetGroupId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGroupId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getGroupId()
func (jbobject *ServicesEc2ModelUserIdGroupPair) GetGroupId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGroupId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.UserIdGroupPair withGroupId(java.lang.String)
func (jbobject *ServicesEc2ModelUserIdGroupPair) WithGroupId(a string) *ServicesEc2ModelUserIdGroupPair {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroupId", "com/amazonaws/services/ec2/model/UserIdGroupPair", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelUserIdGroupPair{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelUserIdGroupPair) SetVpcId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpcId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVpcId()
func (jbobject *ServicesEc2ModelUserIdGroupPair) GetVpcId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.UserIdGroupPair withVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelUserIdGroupPair) WithVpcId(a string) *ServicesEc2ModelUserIdGroupPair {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcId", "com/amazonaws/services/ec2/model/UserIdGroupPair", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelUserIdGroupPair{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVpcPeeringConnectionId(java.lang.String)
func (jbobject *ServicesEc2ModelUserIdGroupPair) SetVpcPeeringConnectionId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpcPeeringConnectionId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVpcPeeringConnectionId()
func (jbobject *ServicesEc2ModelUserIdGroupPair) GetVpcPeeringConnectionId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcPeeringConnectionId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.UserIdGroupPair withVpcPeeringConnectionId(java.lang.String)
func (jbobject *ServicesEc2ModelUserIdGroupPair) WithVpcPeeringConnectionId(a string) *ServicesEc2ModelUserIdGroupPair {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcPeeringConnectionId", "com/amazonaws/services/ec2/model/UserIdGroupPair", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelUserIdGroupPair{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPeeringStatus(java.lang.String)
func (jbobject *ServicesEc2ModelUserIdGroupPair) SetPeeringStatus(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPeeringStatus", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPeeringStatus()
func (jbobject *ServicesEc2ModelUserIdGroupPair) GetPeeringStatus() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPeeringStatus", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.UserIdGroupPair withPeeringStatus(java.lang.String)
func (jbobject *ServicesEc2ModelUserIdGroupPair) WithPeeringStatus(a string) *ServicesEc2ModelUserIdGroupPair {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPeeringStatus", "com/amazonaws/services/ec2/model/UserIdGroupPair", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelUserIdGroupPair{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelUserIdGroupPair) ToString() string {
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
func (jbobject *ServicesEc2ModelUserIdGroupPair) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelUserIdGroupPair) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.UserIdGroupPair clone()
func (jbobject *ServicesEc2ModelUserIdGroupPair) Clone() *ServicesEc2ModelUserIdGroupPair {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/UserIdGroupPair")
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
	unique_x := &ServicesEc2ModelUserIdGroupPair{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelUserIdGroupPair) Clone2() (*JavaLangObject, error) {
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


