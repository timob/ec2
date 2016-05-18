package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAllocateAddressResultInterface interface {
	JavaLangObjectInterface

	// public void setPublicIp(java.lang.String)
	SetPublicIp(a string) 

	// public java.lang.String getPublicIp()
	GetPublicIp() string

	// public com.amazonaws.services.ec2.model.AllocateAddressResult withPublicIp(java.lang.String)
	WithPublicIp(a string) *ServicesEc2ModelAllocateAddressResult

	// public void setDomain(java.lang.String)
	SetDomain2(a string) 

	// public java.lang.String getDomain()
	GetDomain() string

	// public com.amazonaws.services.ec2.model.AllocateAddressResult withDomain(java.lang.String)
	WithDomain2(a string) *ServicesEc2ModelAllocateAddressResult

	// public void setDomain(com.amazonaws.services.ec2.model.DomainType)
	SetDomain(a ServicesEc2ModelDomainTypeInterface) 

	// public com.amazonaws.services.ec2.model.AllocateAddressResult withDomain(com.amazonaws.services.ec2.model.DomainType)
	WithDomain(a ServicesEc2ModelDomainTypeInterface) *ServicesEc2ModelAllocateAddressResult

	// public void setAllocationId(java.lang.String)
	SetAllocationId(a string) 

	// public java.lang.String getAllocationId()
	GetAllocationId() string

	// public com.amazonaws.services.ec2.model.AllocateAddressResult withAllocationId(java.lang.String)
	WithAllocationId(a string) *ServicesEc2ModelAllocateAddressResult

	// public com.amazonaws.services.ec2.model.AllocateAddressResult clone()
	Clone() *ServicesEc2ModelAllocateAddressResult
}

type ServicesEc2ModelAllocateAddressResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.AllocateAddressResult()
func NewServicesEc2ModelAllocateAddressResult() (*ServicesEc2ModelAllocateAddressResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/AllocateAddressResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelAllocateAddressResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setPublicIp(java.lang.String)
func (jbobject *ServicesEc2ModelAllocateAddressResult) SetPublicIp(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPublicIp", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPublicIp()
func (jbobject *ServicesEc2ModelAllocateAddressResult) GetPublicIp() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPublicIp", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.AllocateAddressResult withPublicIp(java.lang.String)
func (jbobject *ServicesEc2ModelAllocateAddressResult) WithPublicIp(a string) *ServicesEc2ModelAllocateAddressResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPublicIp", "com/amazonaws/services/ec2/model/AllocateAddressResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAllocateAddressResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDomain(java.lang.String)
func (jbobject *ServicesEc2ModelAllocateAddressResult) SetDomain2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDomain", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getDomain()
func (jbobject *ServicesEc2ModelAllocateAddressResult) GetDomain() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDomain", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.AllocateAddressResult withDomain(java.lang.String)
func (jbobject *ServicesEc2ModelAllocateAddressResult) WithDomain2(a string) *ServicesEc2ModelAllocateAddressResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDomain", "com/amazonaws/services/ec2/model/AllocateAddressResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAllocateAddressResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDomain(com.amazonaws.services.ec2.model.DomainType)
func (jbobject *ServicesEc2ModelAllocateAddressResult) SetDomain(a ServicesEc2ModelDomainTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDomain", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DomainType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.AllocateAddressResult withDomain(com.amazonaws.services.ec2.model.DomainType)
func (jbobject *ServicesEc2ModelAllocateAddressResult) WithDomain(a ServicesEc2ModelDomainTypeInterface) *ServicesEc2ModelAllocateAddressResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDomain", "com/amazonaws/services/ec2/model/AllocateAddressResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DomainType"))
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
	unique_x := &ServicesEc2ModelAllocateAddressResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAllocationId(java.lang.String)
func (jbobject *ServicesEc2ModelAllocateAddressResult) SetAllocationId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAllocationId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAllocationId()
func (jbobject *ServicesEc2ModelAllocateAddressResult) GetAllocationId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAllocationId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.AllocateAddressResult withAllocationId(java.lang.String)
func (jbobject *ServicesEc2ModelAllocateAddressResult) WithAllocationId(a string) *ServicesEc2ModelAllocateAddressResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAllocationId", "com/amazonaws/services/ec2/model/AllocateAddressResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAllocateAddressResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelAllocateAddressResult) ToString() string {
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
func (jbobject *ServicesEc2ModelAllocateAddressResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelAllocateAddressResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.AllocateAddressResult clone()
func (jbobject *ServicesEc2ModelAllocateAddressResult) Clone() *ServicesEc2ModelAllocateAddressResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/AllocateAddressResult")
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
	unique_x := &ServicesEc2ModelAllocateAddressResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelAllocateAddressResult) Clone2() (*JavaLangObject, error) {
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


