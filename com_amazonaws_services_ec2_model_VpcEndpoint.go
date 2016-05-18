package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelVpcEndpointInterface interface {
	JavaLangObjectInterface

	// public void setVpcEndpointId(java.lang.String)
	SetVpcEndpointId(a string) 

	// public java.lang.String getVpcEndpointId()
	GetVpcEndpointId() string

	// public com.amazonaws.services.ec2.model.VpcEndpoint withVpcEndpointId(java.lang.String)
	WithVpcEndpointId(a string) *ServicesEc2ModelVpcEndpoint

	// public void setVpcId(java.lang.String)
	SetVpcId(a string) 

	// public java.lang.String getVpcId()
	GetVpcId() string

	// public com.amazonaws.services.ec2.model.VpcEndpoint withVpcId(java.lang.String)
	WithVpcId(a string) *ServicesEc2ModelVpcEndpoint

	// public void setServiceName(java.lang.String)
	SetServiceName(a string) 

	// public java.lang.String getServiceName()
	GetServiceName() string

	// public com.amazonaws.services.ec2.model.VpcEndpoint withServiceName(java.lang.String)
	WithServiceName(a string) *ServicesEc2ModelVpcEndpoint

	// public void setState(java.lang.String)
	SetState2(a string) 

	// public java.lang.String getState()
	GetState() string

	// public com.amazonaws.services.ec2.model.VpcEndpoint withState(java.lang.String)
	WithState2(a string) *ServicesEc2ModelVpcEndpoint

	// public void setState(com.amazonaws.services.ec2.model.State)
	SetState(a ServicesEc2ModelStateInterface) 

	// public com.amazonaws.services.ec2.model.VpcEndpoint withState(com.amazonaws.services.ec2.model.State)
	WithState(a ServicesEc2ModelStateInterface) *ServicesEc2ModelVpcEndpoint

	// public void setPolicyDocument(java.lang.String)
	SetPolicyDocument(a string) 

	// public java.lang.String getPolicyDocument()
	GetPolicyDocument() string

	// public com.amazonaws.services.ec2.model.VpcEndpoint withPolicyDocument(java.lang.String)
	WithPolicyDocument(a string) *ServicesEc2ModelVpcEndpoint

	// public java.util.List<java.lang.String> getRouteTableIds()
	GetRouteTableIds() []string

	// public void setRouteTableIds(java.util.Collection<java.lang.String>)
	SetRouteTableIds(a []string) 

	// public com.amazonaws.services.ec2.model.VpcEndpoint withRouteTableIds(java.lang.String...)
	WithRouteTableIds(a ...string) *ServicesEc2ModelVpcEndpoint

	// public com.amazonaws.services.ec2.model.VpcEndpoint withRouteTableIds(java.util.Collection<java.lang.String>)
	WithRouteTableIds2(a []string) *ServicesEc2ModelVpcEndpoint

	// public void setCreationTimestamp(java.util.Date)
	SetCreationTimestamp(a time.Time) 

	// public java.util.Date getCreationTimestamp()
	GetCreationTimestamp() time.Time

	// public com.amazonaws.services.ec2.model.VpcEndpoint withCreationTimestamp(java.util.Date)
	WithCreationTimestamp(a time.Time) *ServicesEc2ModelVpcEndpoint

	// public com.amazonaws.services.ec2.model.VpcEndpoint clone()
	Clone() *ServicesEc2ModelVpcEndpoint
}

type ServicesEc2ModelVpcEndpoint struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.VpcEndpoint()
func NewServicesEc2ModelVpcEndpoint() (*ServicesEc2ModelVpcEndpoint) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/VpcEndpoint")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelVpcEndpoint{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpcEndpointId(java.lang.String)
func (jbobject *ServicesEc2ModelVpcEndpoint) SetVpcEndpointId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpcEndpointId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVpcEndpointId()
func (jbobject *ServicesEc2ModelVpcEndpoint) GetVpcEndpointId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcEndpointId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VpcEndpoint withVpcEndpointId(java.lang.String)
func (jbobject *ServicesEc2ModelVpcEndpoint) WithVpcEndpointId(a string) *ServicesEc2ModelVpcEndpoint {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcEndpointId", "com/amazonaws/services/ec2/model/VpcEndpoint", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpcEndpoint{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelVpcEndpoint) SetVpcId(a string)  {
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
func (jbobject *ServicesEc2ModelVpcEndpoint) GetVpcId() string {
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

// public com.amazonaws.services.ec2.model.VpcEndpoint withVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelVpcEndpoint) WithVpcId(a string) *ServicesEc2ModelVpcEndpoint {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcId", "com/amazonaws/services/ec2/model/VpcEndpoint", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpcEndpoint{}
	unique_x.Callable = dst
	return unique_x
}

// public void setServiceName(java.lang.String)
func (jbobject *ServicesEc2ModelVpcEndpoint) SetServiceName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setServiceName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getServiceName()
func (jbobject *ServicesEc2ModelVpcEndpoint) GetServiceName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getServiceName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VpcEndpoint withServiceName(java.lang.String)
func (jbobject *ServicesEc2ModelVpcEndpoint) WithServiceName(a string) *ServicesEc2ModelVpcEndpoint {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withServiceName", "com/amazonaws/services/ec2/model/VpcEndpoint", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpcEndpoint{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(java.lang.String)
func (jbobject *ServicesEc2ModelVpcEndpoint) SetState2(a string)  {
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
func (jbobject *ServicesEc2ModelVpcEndpoint) GetState() string {
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

// public com.amazonaws.services.ec2.model.VpcEndpoint withState(java.lang.String)
func (jbobject *ServicesEc2ModelVpcEndpoint) WithState2(a string) *ServicesEc2ModelVpcEndpoint {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/VpcEndpoint", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpcEndpoint{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(com.amazonaws.services.ec2.model.State)
func (jbobject *ServicesEc2ModelVpcEndpoint) SetState(a ServicesEc2ModelStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/State"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VpcEndpoint withState(com.amazonaws.services.ec2.model.State)
func (jbobject *ServicesEc2ModelVpcEndpoint) WithState(a ServicesEc2ModelStateInterface) *ServicesEc2ModelVpcEndpoint {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/VpcEndpoint", conv_a.Value().Cast("com/amazonaws/services/ec2/model/State"))
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
	unique_x := &ServicesEc2ModelVpcEndpoint{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPolicyDocument(java.lang.String)
func (jbobject *ServicesEc2ModelVpcEndpoint) SetPolicyDocument(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPolicyDocument", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPolicyDocument()
func (jbobject *ServicesEc2ModelVpcEndpoint) GetPolicyDocument() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPolicyDocument", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VpcEndpoint withPolicyDocument(java.lang.String)
func (jbobject *ServicesEc2ModelVpcEndpoint) WithPolicyDocument(a string) *ServicesEc2ModelVpcEndpoint {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPolicyDocument", "com/amazonaws/services/ec2/model/VpcEndpoint", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpcEndpoint{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getRouteTableIds()
func (jbobject *ServicesEc2ModelVpcEndpoint) GetRouteTableIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRouteTableIds", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setRouteTableIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelVpcEndpoint) SetRouteTableIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRouteTableIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VpcEndpoint withRouteTableIds(java.lang.String...)
func (jbobject *ServicesEc2ModelVpcEndpoint) WithRouteTableIds(a ...string) *ServicesEc2ModelVpcEndpoint {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRouteTableIds", "com/amazonaws/services/ec2/model/VpcEndpoint", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpcEndpoint{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.VpcEndpoint withRouteTableIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelVpcEndpoint) WithRouteTableIds2(a []string) *ServicesEc2ModelVpcEndpoint {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRouteTableIds", "com/amazonaws/services/ec2/model/VpcEndpoint", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelVpcEndpoint{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCreationTimestamp(java.util.Date)
func (jbobject *ServicesEc2ModelVpcEndpoint) SetCreationTimestamp(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCreationTimestamp", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getCreationTimestamp()
func (jbobject *ServicesEc2ModelVpcEndpoint) GetCreationTimestamp() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCreationTimestamp", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.VpcEndpoint withCreationTimestamp(java.util.Date)
func (jbobject *ServicesEc2ModelVpcEndpoint) WithCreationTimestamp(a time.Time) *ServicesEc2ModelVpcEndpoint {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCreationTimestamp", "com/amazonaws/services/ec2/model/VpcEndpoint", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelVpcEndpoint{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVpcEndpoint) ToString() string {
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
func (jbobject *ServicesEc2ModelVpcEndpoint) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelVpcEndpoint) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.VpcEndpoint clone()
func (jbobject *ServicesEc2ModelVpcEndpoint) Clone() *ServicesEc2ModelVpcEndpoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/VpcEndpoint")
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
	unique_x := &ServicesEc2ModelVpcEndpoint{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelVpcEndpoint) Clone2() (*JavaLangObject, error) {
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


