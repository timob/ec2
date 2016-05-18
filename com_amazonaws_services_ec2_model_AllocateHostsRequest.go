package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAllocateHostsRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setAutoPlacement(java.lang.String)
	SetAutoPlacement2(a string) 

	// public java.lang.String getAutoPlacement()
	GetAutoPlacement() string

	// public com.amazonaws.services.ec2.model.AllocateHostsRequest withAutoPlacement(java.lang.String)
	WithAutoPlacement2(a string) *ServicesEc2ModelAllocateHostsRequest

	// public void setAutoPlacement(com.amazonaws.services.ec2.model.AutoPlacement)
	SetAutoPlacement(a ServicesEc2ModelAutoPlacementInterface) 

	// public com.amazonaws.services.ec2.model.AllocateHostsRequest withAutoPlacement(com.amazonaws.services.ec2.model.AutoPlacement)
	WithAutoPlacement(a ServicesEc2ModelAutoPlacementInterface) *ServicesEc2ModelAllocateHostsRequest

	// public void setClientToken(java.lang.String)
	SetClientToken(a string) 

	// public java.lang.String getClientToken()
	GetClientToken() string

	// public com.amazonaws.services.ec2.model.AllocateHostsRequest withClientToken(java.lang.String)
	WithClientToken(a string) *ServicesEc2ModelAllocateHostsRequest

	// public void setInstanceType(java.lang.String)
	SetInstanceType(a string) 

	// public java.lang.String getInstanceType()
	GetInstanceType() string

	// public com.amazonaws.services.ec2.model.AllocateHostsRequest withInstanceType(java.lang.String)
	WithInstanceType(a string) *ServicesEc2ModelAllocateHostsRequest

	// public void setQuantity(java.lang.Integer)
	SetQuantity(a int) 

	// public java.lang.Integer getQuantity()
	GetQuantity() int

	// public com.amazonaws.services.ec2.model.AllocateHostsRequest withQuantity(java.lang.Integer)
	WithQuantity(a int) *ServicesEc2ModelAllocateHostsRequest

	// public void setAvailabilityZone(java.lang.String)
	SetAvailabilityZone(a string) 

	// public java.lang.String getAvailabilityZone()
	GetAvailabilityZone() string

	// public com.amazonaws.services.ec2.model.AllocateHostsRequest withAvailabilityZone(java.lang.String)
	WithAvailabilityZone(a string) *ServicesEc2ModelAllocateHostsRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.AllocateHostsRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.AllocateHostsRequest clone()
	Clone3() *ServicesEc2ModelAllocateHostsRequest
}

type ServicesEc2ModelAllocateHostsRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.AllocateHostsRequest()
func NewServicesEc2ModelAllocateHostsRequest() (*ServicesEc2ModelAllocateHostsRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/AllocateHostsRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelAllocateHostsRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setAutoPlacement(java.lang.String)
func (jbobject *ServicesEc2ModelAllocateHostsRequest) SetAutoPlacement2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAutoPlacement", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAutoPlacement()
func (jbobject *ServicesEc2ModelAllocateHostsRequest) GetAutoPlacement() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAutoPlacement", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.AllocateHostsRequest withAutoPlacement(java.lang.String)
func (jbobject *ServicesEc2ModelAllocateHostsRequest) WithAutoPlacement2(a string) *ServicesEc2ModelAllocateHostsRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAutoPlacement", "com/amazonaws/services/ec2/model/AllocateHostsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAllocateHostsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAutoPlacement(com.amazonaws.services.ec2.model.AutoPlacement)
func (jbobject *ServicesEc2ModelAllocateHostsRequest) SetAutoPlacement(a ServicesEc2ModelAutoPlacementInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAutoPlacement", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/AutoPlacement"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.AllocateHostsRequest withAutoPlacement(com.amazonaws.services.ec2.model.AutoPlacement)
func (jbobject *ServicesEc2ModelAllocateHostsRequest) WithAutoPlacement(a ServicesEc2ModelAutoPlacementInterface) *ServicesEc2ModelAllocateHostsRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAutoPlacement", "com/amazonaws/services/ec2/model/AllocateHostsRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AutoPlacement"))
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
	unique_x := &ServicesEc2ModelAllocateHostsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelAllocateHostsRequest) SetClientToken(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setClientToken", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getClientToken()
func (jbobject *ServicesEc2ModelAllocateHostsRequest) GetClientToken() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getClientToken", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.AllocateHostsRequest withClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelAllocateHostsRequest) WithClientToken(a string) *ServicesEc2ModelAllocateHostsRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClientToken", "com/amazonaws/services/ec2/model/AllocateHostsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAllocateHostsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelAllocateHostsRequest) SetInstanceType(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceType()
func (jbobject *ServicesEc2ModelAllocateHostsRequest) GetInstanceType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.AllocateHostsRequest withInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelAllocateHostsRequest) WithInstanceType(a string) *ServicesEc2ModelAllocateHostsRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/AllocateHostsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAllocateHostsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setQuantity(java.lang.Integer)
func (jbobject *ServicesEc2ModelAllocateHostsRequest) SetQuantity(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setQuantity", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getQuantity()
func (jbobject *ServicesEc2ModelAllocateHostsRequest) GetQuantity() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getQuantity", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.AllocateHostsRequest withQuantity(java.lang.Integer)
func (jbobject *ServicesEc2ModelAllocateHostsRequest) WithQuantity(a int) *ServicesEc2ModelAllocateHostsRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withQuantity", "com/amazonaws/services/ec2/model/AllocateHostsRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelAllocateHostsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelAllocateHostsRequest) SetAvailabilityZone(a string)  {
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
func (jbobject *ServicesEc2ModelAllocateHostsRequest) GetAvailabilityZone() string {
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

// public com.amazonaws.services.ec2.model.AllocateHostsRequest withAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelAllocateHostsRequest) WithAvailabilityZone(a string) *ServicesEc2ModelAllocateHostsRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailabilityZone", "com/amazonaws/services/ec2/model/AllocateHostsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAllocateHostsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.AllocateHostsRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelAllocateHostsRequest) GetDryRunRequest() *Request {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDryRunRequest", "com/amazonaws/Request")
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
	unique_x := &Request{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelAllocateHostsRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelAllocateHostsRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelAllocateHostsRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.AllocateHostsRequest clone()
func (jbobject *ServicesEc2ModelAllocateHostsRequest) Clone3() *ServicesEc2ModelAllocateHostsRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/AllocateHostsRequest")
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
	unique_x := &ServicesEc2ModelAllocateHostsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelAllocateHostsRequest) Clone() *AmazonWebServiceRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/AmazonWebServiceRequest")
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
	unique_x := &AmazonWebServiceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelAllocateHostsRequest) Clone2() (*JavaLangObject, error) {
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


