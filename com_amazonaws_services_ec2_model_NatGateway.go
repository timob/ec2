package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelNatGatewayInterface interface {
	JavaLangObjectInterface

	// public void setVpcId(java.lang.String)
	SetVpcId(a string) 

	// public java.lang.String getVpcId()
	GetVpcId() string

	// public com.amazonaws.services.ec2.model.NatGateway withVpcId(java.lang.String)
	WithVpcId(a string) *ServicesEc2ModelNatGateway

	// public void setSubnetId(java.lang.String)
	SetSubnetId(a string) 

	// public java.lang.String getSubnetId()
	GetSubnetId() string

	// public com.amazonaws.services.ec2.model.NatGateway withSubnetId(java.lang.String)
	WithSubnetId(a string) *ServicesEc2ModelNatGateway

	// public void setNatGatewayId(java.lang.String)
	SetNatGatewayId(a string) 

	// public java.lang.String getNatGatewayId()
	GetNatGatewayId() string

	// public com.amazonaws.services.ec2.model.NatGateway withNatGatewayId(java.lang.String)
	WithNatGatewayId(a string) *ServicesEc2ModelNatGateway

	// public void setCreateTime(java.util.Date)
	SetCreateTime(a time.Time) 

	// public java.util.Date getCreateTime()
	GetCreateTime() time.Time

	// public com.amazonaws.services.ec2.model.NatGateway withCreateTime(java.util.Date)
	WithCreateTime(a time.Time) *ServicesEc2ModelNatGateway

	// public void setDeleteTime(java.util.Date)
	SetDeleteTime(a time.Time) 

	// public java.util.Date getDeleteTime()
	GetDeleteTime() time.Time

	// public com.amazonaws.services.ec2.model.NatGateway withDeleteTime(java.util.Date)
	WithDeleteTime(a time.Time) *ServicesEc2ModelNatGateway

	// public java.util.List<com.amazonaws.services.ec2.model.NatGatewayAddress> getNatGatewayAddresses()
	GetNatGatewayAddresses() []*ServicesEc2ModelNatGatewayAddress

	// public void setNatGatewayAddresses(java.util.Collection<com.amazonaws.services.ec2.model.NatGatewayAddress>)
	SetNatGatewayAddresses(a []*ServicesEc2ModelNatGatewayAddress) 

	// public com.amazonaws.services.ec2.model.NatGateway withNatGatewayAddresses(com.amazonaws.services.ec2.model.NatGatewayAddress...)
	WithNatGatewayAddresses(a ...*ServicesEc2ModelNatGatewayAddress) *ServicesEc2ModelNatGateway

	// public com.amazonaws.services.ec2.model.NatGateway withNatGatewayAddresses(java.util.Collection<com.amazonaws.services.ec2.model.NatGatewayAddress>)
	WithNatGatewayAddresses2(a []*ServicesEc2ModelNatGatewayAddress) *ServicesEc2ModelNatGateway

	// public void setState(java.lang.String)
	SetState2(a string) 

	// public java.lang.String getState()
	GetState() string

	// public com.amazonaws.services.ec2.model.NatGateway withState(java.lang.String)
	WithState2(a string) *ServicesEc2ModelNatGateway

	// public void setState(com.amazonaws.services.ec2.model.NatGatewayState)
	SetState(a ServicesEc2ModelNatGatewayStateInterface) 

	// public com.amazonaws.services.ec2.model.NatGateway withState(com.amazonaws.services.ec2.model.NatGatewayState)
	WithState(a ServicesEc2ModelNatGatewayStateInterface) *ServicesEc2ModelNatGateway

	// public void setFailureCode(java.lang.String)
	SetFailureCode(a string) 

	// public java.lang.String getFailureCode()
	GetFailureCode() string

	// public com.amazonaws.services.ec2.model.NatGateway withFailureCode(java.lang.String)
	WithFailureCode(a string) *ServicesEc2ModelNatGateway

	// public void setFailureMessage(java.lang.String)
	SetFailureMessage(a string) 

	// public java.lang.String getFailureMessage()
	GetFailureMessage() string

	// public com.amazonaws.services.ec2.model.NatGateway withFailureMessage(java.lang.String)
	WithFailureMessage(a string) *ServicesEc2ModelNatGateway

	// public com.amazonaws.services.ec2.model.NatGateway clone()
	Clone() *ServicesEc2ModelNatGateway
}

type ServicesEc2ModelNatGateway struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.NatGateway()
func NewServicesEc2ModelNatGateway() (*ServicesEc2ModelNatGateway) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/NatGateway")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelNatGateway{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelNatGateway) SetVpcId(a string)  {
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
func (jbobject *ServicesEc2ModelNatGateway) GetVpcId() string {
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

// public com.amazonaws.services.ec2.model.NatGateway withVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelNatGateway) WithVpcId(a string) *ServicesEc2ModelNatGateway {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcId", "com/amazonaws/services/ec2/model/NatGateway", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNatGateway{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelNatGateway) SetSubnetId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSubnetId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSubnetId()
func (jbobject *ServicesEc2ModelNatGateway) GetSubnetId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSubnetId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.NatGateway withSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelNatGateway) WithSubnetId(a string) *ServicesEc2ModelNatGateway {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSubnetId", "com/amazonaws/services/ec2/model/NatGateway", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNatGateway{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNatGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelNatGateway) SetNatGatewayId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNatGatewayId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getNatGatewayId()
func (jbobject *ServicesEc2ModelNatGateway) GetNatGatewayId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNatGatewayId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.NatGateway withNatGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelNatGateway) WithNatGatewayId(a string) *ServicesEc2ModelNatGateway {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNatGatewayId", "com/amazonaws/services/ec2/model/NatGateway", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNatGateway{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCreateTime(java.util.Date)
func (jbobject *ServicesEc2ModelNatGateway) SetCreateTime(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCreateTime", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getCreateTime()
func (jbobject *ServicesEc2ModelNatGateway) GetCreateTime() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCreateTime", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.NatGateway withCreateTime(java.util.Date)
func (jbobject *ServicesEc2ModelNatGateway) WithCreateTime(a time.Time) *ServicesEc2ModelNatGateway {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCreateTime", "com/amazonaws/services/ec2/model/NatGateway", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelNatGateway{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDeleteTime(java.util.Date)
func (jbobject *ServicesEc2ModelNatGateway) SetDeleteTime(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDeleteTime", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getDeleteTime()
func (jbobject *ServicesEc2ModelNatGateway) GetDeleteTime() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDeleteTime", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.NatGateway withDeleteTime(java.util.Date)
func (jbobject *ServicesEc2ModelNatGateway) WithDeleteTime(a time.Time) *ServicesEc2ModelNatGateway {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDeleteTime", "com/amazonaws/services/ec2/model/NatGateway", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelNatGateway{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.NatGatewayAddress> getNatGatewayAddresses()
func (jbobject *ServicesEc2ModelNatGateway) GetNatGatewayAddresses() []*ServicesEc2ModelNatGatewayAddress {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNatGatewayAddresses", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelNatGatewayAddress)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setNatGatewayAddresses(java.util.Collection<com.amazonaws.services.ec2.model.NatGatewayAddress>)
func (jbobject *ServicesEc2ModelNatGateway) SetNatGatewayAddresses(a []*ServicesEc2ModelNatGatewayAddress)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNatGatewayAddresses", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.NatGateway withNatGatewayAddresses(com.amazonaws.services.ec2.model.NatGatewayAddress...)
func (jbobject *ServicesEc2ModelNatGateway) WithNatGatewayAddresses(a ...*ServicesEc2ModelNatGatewayAddress) *ServicesEc2ModelNatGateway {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/NatGatewayAddress")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNatGatewayAddresses", "com/amazonaws/services/ec2/model/NatGateway", conv_a.Value().Cast("com/amazonaws/services/ec2/model/NatGatewayAddress"))
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
	unique_x := &ServicesEc2ModelNatGateway{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.NatGateway withNatGatewayAddresses(java.util.Collection<com.amazonaws.services.ec2.model.NatGatewayAddress>)
func (jbobject *ServicesEc2ModelNatGateway) WithNatGatewayAddresses2(a []*ServicesEc2ModelNatGatewayAddress) *ServicesEc2ModelNatGateway {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNatGatewayAddresses", "com/amazonaws/services/ec2/model/NatGateway", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelNatGateway{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(java.lang.String)
func (jbobject *ServicesEc2ModelNatGateway) SetState2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getState()
func (jbobject *ServicesEc2ModelNatGateway) GetState() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getState", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.NatGateway withState(java.lang.String)
func (jbobject *ServicesEc2ModelNatGateway) WithState2(a string) *ServicesEc2ModelNatGateway {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/NatGateway", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNatGateway{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(com.amazonaws.services.ec2.model.NatGatewayState)
func (jbobject *ServicesEc2ModelNatGateway) SetState(a ServicesEc2ModelNatGatewayStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/NatGatewayState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.NatGateway withState(com.amazonaws.services.ec2.model.NatGatewayState)
func (jbobject *ServicesEc2ModelNatGateway) WithState(a ServicesEc2ModelNatGatewayStateInterface) *ServicesEc2ModelNatGateway {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/NatGateway", conv_a.Value().Cast("com/amazonaws/services/ec2/model/NatGatewayState"))
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
	unique_x := &ServicesEc2ModelNatGateway{}
	unique_x.Callable = dst
	return unique_x
}

// public void setFailureCode(java.lang.String)
func (jbobject *ServicesEc2ModelNatGateway) SetFailureCode(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFailureCode", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getFailureCode()
func (jbobject *ServicesEc2ModelNatGateway) GetFailureCode() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFailureCode", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.NatGateway withFailureCode(java.lang.String)
func (jbobject *ServicesEc2ModelNatGateway) WithFailureCode(a string) *ServicesEc2ModelNatGateway {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFailureCode", "com/amazonaws/services/ec2/model/NatGateway", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNatGateway{}
	unique_x.Callable = dst
	return unique_x
}

// public void setFailureMessage(java.lang.String)
func (jbobject *ServicesEc2ModelNatGateway) SetFailureMessage(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFailureMessage", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getFailureMessage()
func (jbobject *ServicesEc2ModelNatGateway) GetFailureMessage() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFailureMessage", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.NatGateway withFailureMessage(java.lang.String)
func (jbobject *ServicesEc2ModelNatGateway) WithFailureMessage(a string) *ServicesEc2ModelNatGateway {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFailureMessage", "com/amazonaws/services/ec2/model/NatGateway", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNatGateway{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelNatGateway) ToString() string {
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
func (jbobject *ServicesEc2ModelNatGateway) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelNatGateway) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.NatGateway clone()
func (jbobject *ServicesEc2ModelNatGateway) Clone() *ServicesEc2ModelNatGateway {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/NatGateway")
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
	unique_x := &ServicesEc2ModelNatGateway{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelNatGateway) Clone2() (*JavaLangObject, error) {
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


