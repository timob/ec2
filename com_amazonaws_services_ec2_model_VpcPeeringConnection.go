package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelVpcPeeringConnectionInterface interface {
	JavaLangObjectInterface

	// public void setAccepterVpcInfo(com.amazonaws.services.ec2.model.VpcPeeringConnectionVpcInfo)
	SetAccepterVpcInfo(a ServicesEc2ModelVpcPeeringConnectionVpcInfoInterface) 

	// public com.amazonaws.services.ec2.model.VpcPeeringConnectionVpcInfo getAccepterVpcInfo()
	GetAccepterVpcInfo() *ServicesEc2ModelVpcPeeringConnectionVpcInfo

	// public com.amazonaws.services.ec2.model.VpcPeeringConnection withAccepterVpcInfo(com.amazonaws.services.ec2.model.VpcPeeringConnectionVpcInfo)
	WithAccepterVpcInfo(a ServicesEc2ModelVpcPeeringConnectionVpcInfoInterface) *ServicesEc2ModelVpcPeeringConnection

	// public void setExpirationTime(java.util.Date)
	SetExpirationTime(a time.Time) 

	// public java.util.Date getExpirationTime()
	GetExpirationTime() time.Time

	// public com.amazonaws.services.ec2.model.VpcPeeringConnection withExpirationTime(java.util.Date)
	WithExpirationTime(a time.Time) *ServicesEc2ModelVpcPeeringConnection

	// public void setRequesterVpcInfo(com.amazonaws.services.ec2.model.VpcPeeringConnectionVpcInfo)
	SetRequesterVpcInfo(a ServicesEc2ModelVpcPeeringConnectionVpcInfoInterface) 

	// public com.amazonaws.services.ec2.model.VpcPeeringConnectionVpcInfo getRequesterVpcInfo()
	GetRequesterVpcInfo() *ServicesEc2ModelVpcPeeringConnectionVpcInfo

	// public com.amazonaws.services.ec2.model.VpcPeeringConnection withRequesterVpcInfo(com.amazonaws.services.ec2.model.VpcPeeringConnectionVpcInfo)
	WithRequesterVpcInfo(a ServicesEc2ModelVpcPeeringConnectionVpcInfoInterface) *ServicesEc2ModelVpcPeeringConnection

	// public void setStatus(com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReason)
	SetStatus(a ServicesEc2ModelVpcPeeringConnectionStateReasonInterface) 

	// public com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReason getStatus()
	GetStatus() *ServicesEc2ModelVpcPeeringConnectionStateReason

	// public com.amazonaws.services.ec2.model.VpcPeeringConnection withStatus(com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReason)
	WithStatus(a ServicesEc2ModelVpcPeeringConnectionStateReasonInterface) *ServicesEc2ModelVpcPeeringConnection

	// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
	GetTags() []*ServicesEc2ModelTag

	// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	SetTags(a []*ServicesEc2ModelTag) 

	// public com.amazonaws.services.ec2.model.VpcPeeringConnection withTags(com.amazonaws.services.ec2.model.Tag...)
	WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelVpcPeeringConnection

	// public com.amazonaws.services.ec2.model.VpcPeeringConnection withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelVpcPeeringConnection

	// public void setVpcPeeringConnectionId(java.lang.String)
	SetVpcPeeringConnectionId(a string) 

	// public java.lang.String getVpcPeeringConnectionId()
	GetVpcPeeringConnectionId() string

	// public com.amazonaws.services.ec2.model.VpcPeeringConnection withVpcPeeringConnectionId(java.lang.String)
	WithVpcPeeringConnectionId(a string) *ServicesEc2ModelVpcPeeringConnection

	// public com.amazonaws.services.ec2.model.VpcPeeringConnection clone()
	Clone() *ServicesEc2ModelVpcPeeringConnection
}

type ServicesEc2ModelVpcPeeringConnection struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.VpcPeeringConnection()
func NewServicesEc2ModelVpcPeeringConnection() (*ServicesEc2ModelVpcPeeringConnection) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/VpcPeeringConnection")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelVpcPeeringConnection{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setAccepterVpcInfo(com.amazonaws.services.ec2.model.VpcPeeringConnectionVpcInfo)
func (jbobject *ServicesEc2ModelVpcPeeringConnection) SetAccepterVpcInfo(a ServicesEc2ModelVpcPeeringConnectionVpcInfoInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAccepterVpcInfo", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpcPeeringConnectionVpcInfo"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VpcPeeringConnectionVpcInfo getAccepterVpcInfo()
func (jbobject *ServicesEc2ModelVpcPeeringConnection) GetAccepterVpcInfo() *ServicesEc2ModelVpcPeeringConnectionVpcInfo {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAccepterVpcInfo", "com/amazonaws/services/ec2/model/VpcPeeringConnectionVpcInfo")
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
	unique_x := &ServicesEc2ModelVpcPeeringConnectionVpcInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.VpcPeeringConnection withAccepterVpcInfo(com.amazonaws.services.ec2.model.VpcPeeringConnectionVpcInfo)
func (jbobject *ServicesEc2ModelVpcPeeringConnection) WithAccepterVpcInfo(a ServicesEc2ModelVpcPeeringConnectionVpcInfoInterface) *ServicesEc2ModelVpcPeeringConnection {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAccepterVpcInfo", "com/amazonaws/services/ec2/model/VpcPeeringConnection", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpcPeeringConnectionVpcInfo"))
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
	unique_x := &ServicesEc2ModelVpcPeeringConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public void setExpirationTime(java.util.Date)
func (jbobject *ServicesEc2ModelVpcPeeringConnection) SetExpirationTime(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExpirationTime", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getExpirationTime()
func (jbobject *ServicesEc2ModelVpcPeeringConnection) GetExpirationTime() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getExpirationTime", "java/util/Date")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDate()
	dst := new(time.Time)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.VpcPeeringConnection withExpirationTime(java.util.Date)
func (jbobject *ServicesEc2ModelVpcPeeringConnection) WithExpirationTime(a time.Time) *ServicesEc2ModelVpcPeeringConnection {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withExpirationTime", "com/amazonaws/services/ec2/model/VpcPeeringConnection", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelVpcPeeringConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRequesterVpcInfo(com.amazonaws.services.ec2.model.VpcPeeringConnectionVpcInfo)
func (jbobject *ServicesEc2ModelVpcPeeringConnection) SetRequesterVpcInfo(a ServicesEc2ModelVpcPeeringConnectionVpcInfoInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRequesterVpcInfo", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpcPeeringConnectionVpcInfo"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VpcPeeringConnectionVpcInfo getRequesterVpcInfo()
func (jbobject *ServicesEc2ModelVpcPeeringConnection) GetRequesterVpcInfo() *ServicesEc2ModelVpcPeeringConnectionVpcInfo {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRequesterVpcInfo", "com/amazonaws/services/ec2/model/VpcPeeringConnectionVpcInfo")
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
	unique_x := &ServicesEc2ModelVpcPeeringConnectionVpcInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.VpcPeeringConnection withRequesterVpcInfo(com.amazonaws.services.ec2.model.VpcPeeringConnectionVpcInfo)
func (jbobject *ServicesEc2ModelVpcPeeringConnection) WithRequesterVpcInfo(a ServicesEc2ModelVpcPeeringConnectionVpcInfoInterface) *ServicesEc2ModelVpcPeeringConnection {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRequesterVpcInfo", "com/amazonaws/services/ec2/model/VpcPeeringConnection", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpcPeeringConnectionVpcInfo"))
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
	unique_x := &ServicesEc2ModelVpcPeeringConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReason)
func (jbobject *ServicesEc2ModelVpcPeeringConnection) SetStatus(a ServicesEc2ModelVpcPeeringConnectionStateReasonInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatus", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReason"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReason getStatus()
func (jbobject *ServicesEc2ModelVpcPeeringConnection) GetStatus() *ServicesEc2ModelVpcPeeringConnectionStateReason {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStatus", "com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReason")
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
	unique_x := &ServicesEc2ModelVpcPeeringConnectionStateReason{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.VpcPeeringConnection withStatus(com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReason)
func (jbobject *ServicesEc2ModelVpcPeeringConnection) WithStatus(a ServicesEc2ModelVpcPeeringConnectionStateReasonInterface) *ServicesEc2ModelVpcPeeringConnection {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/VpcPeeringConnection", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReason"))
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
	unique_x := &ServicesEc2ModelVpcPeeringConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
func (jbobject *ServicesEc2ModelVpcPeeringConnection) GetTags() []*ServicesEc2ModelTag {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTags", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelTag)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelVpcPeeringConnection) SetTags(a []*ServicesEc2ModelTag)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTags", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VpcPeeringConnection withTags(com.amazonaws.services.ec2.model.Tag...)
func (jbobject *ServicesEc2ModelVpcPeeringConnection) WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelVpcPeeringConnection {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Tag")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/VpcPeeringConnection", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tag"))
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
	unique_x := &ServicesEc2ModelVpcPeeringConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.VpcPeeringConnection withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelVpcPeeringConnection) WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelVpcPeeringConnection {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/VpcPeeringConnection", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelVpcPeeringConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVpcPeeringConnectionId(java.lang.String)
func (jbobject *ServicesEc2ModelVpcPeeringConnection) SetVpcPeeringConnectionId(a string)  {
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
func (jbobject *ServicesEc2ModelVpcPeeringConnection) GetVpcPeeringConnectionId() string {
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

// public com.amazonaws.services.ec2.model.VpcPeeringConnection withVpcPeeringConnectionId(java.lang.String)
func (jbobject *ServicesEc2ModelVpcPeeringConnection) WithVpcPeeringConnectionId(a string) *ServicesEc2ModelVpcPeeringConnection {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcPeeringConnectionId", "com/amazonaws/services/ec2/model/VpcPeeringConnection", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpcPeeringConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVpcPeeringConnection) ToString() string {
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
func (jbobject *ServicesEc2ModelVpcPeeringConnection) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelVpcPeeringConnection) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.VpcPeeringConnection clone()
func (jbobject *ServicesEc2ModelVpcPeeringConnection) Clone() *ServicesEc2ModelVpcPeeringConnection {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/VpcPeeringConnection")
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
	unique_x := &ServicesEc2ModelVpcPeeringConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelVpcPeeringConnection) Clone2() (*JavaLangObject, error) {
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


