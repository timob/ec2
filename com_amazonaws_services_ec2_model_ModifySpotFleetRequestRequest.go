package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelModifySpotFleetRequestRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setSpotFleetRequestId(java.lang.String)
	SetSpotFleetRequestId(a string) 

	// public java.lang.String getSpotFleetRequestId()
	GetSpotFleetRequestId() string

	// public com.amazonaws.services.ec2.model.ModifySpotFleetRequestRequest withSpotFleetRequestId(java.lang.String)
	WithSpotFleetRequestId(a string) *ServicesEc2ModelModifySpotFleetRequestRequest

	// public void setTargetCapacity(java.lang.Integer)
	SetTargetCapacity(a int) 

	// public java.lang.Integer getTargetCapacity()
	GetTargetCapacity() int

	// public com.amazonaws.services.ec2.model.ModifySpotFleetRequestRequest withTargetCapacity(java.lang.Integer)
	WithTargetCapacity(a int) *ServicesEc2ModelModifySpotFleetRequestRequest

	// public void setExcessCapacityTerminationPolicy(java.lang.String)
	SetExcessCapacityTerminationPolicy2(a string) 

	// public java.lang.String getExcessCapacityTerminationPolicy()
	GetExcessCapacityTerminationPolicy() string

	// public com.amazonaws.services.ec2.model.ModifySpotFleetRequestRequest withExcessCapacityTerminationPolicy(java.lang.String)
	WithExcessCapacityTerminationPolicy2(a string) *ServicesEc2ModelModifySpotFleetRequestRequest

	// public void setExcessCapacityTerminationPolicy(com.amazonaws.services.ec2.model.ExcessCapacityTerminationPolicy)
	SetExcessCapacityTerminationPolicy(a ServicesEc2ModelExcessCapacityTerminationPolicyInterface) 

	// public com.amazonaws.services.ec2.model.ModifySpotFleetRequestRequest withExcessCapacityTerminationPolicy(com.amazonaws.services.ec2.model.ExcessCapacityTerminationPolicy)
	WithExcessCapacityTerminationPolicy(a ServicesEc2ModelExcessCapacityTerminationPolicyInterface) *ServicesEc2ModelModifySpotFleetRequestRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifySpotFleetRequestRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.ModifySpotFleetRequestRequest clone()
	Clone3() *ServicesEc2ModelModifySpotFleetRequestRequest
}

type ServicesEc2ModelModifySpotFleetRequestRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.ModifySpotFleetRequestRequest()
func NewServicesEc2ModelModifySpotFleetRequestRequest() (*ServicesEc2ModelModifySpotFleetRequestRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ModifySpotFleetRequestRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelModifySpotFleetRequestRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSpotFleetRequestId(java.lang.String)
func (jbobject *ServicesEc2ModelModifySpotFleetRequestRequest) SetSpotFleetRequestId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSpotFleetRequestId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSpotFleetRequestId()
func (jbobject *ServicesEc2ModelModifySpotFleetRequestRequest) GetSpotFleetRequestId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSpotFleetRequestId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ModifySpotFleetRequestRequest withSpotFleetRequestId(java.lang.String)
func (jbobject *ServicesEc2ModelModifySpotFleetRequestRequest) WithSpotFleetRequestId(a string) *ServicesEc2ModelModifySpotFleetRequestRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotFleetRequestId", "com/amazonaws/services/ec2/model/ModifySpotFleetRequestRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifySpotFleetRequestRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTargetCapacity(java.lang.Integer)
func (jbobject *ServicesEc2ModelModifySpotFleetRequestRequest) SetTargetCapacity(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTargetCapacity", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getTargetCapacity()
func (jbobject *ServicesEc2ModelModifySpotFleetRequestRequest) GetTargetCapacity() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTargetCapacity", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.ModifySpotFleetRequestRequest withTargetCapacity(java.lang.Integer)
func (jbobject *ServicesEc2ModelModifySpotFleetRequestRequest) WithTargetCapacity(a int) *ServicesEc2ModelModifySpotFleetRequestRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTargetCapacity", "com/amazonaws/services/ec2/model/ModifySpotFleetRequestRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelModifySpotFleetRequestRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setExcessCapacityTerminationPolicy(java.lang.String)
func (jbobject *ServicesEc2ModelModifySpotFleetRequestRequest) SetExcessCapacityTerminationPolicy2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExcessCapacityTerminationPolicy", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getExcessCapacityTerminationPolicy()
func (jbobject *ServicesEc2ModelModifySpotFleetRequestRequest) GetExcessCapacityTerminationPolicy() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getExcessCapacityTerminationPolicy", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ModifySpotFleetRequestRequest withExcessCapacityTerminationPolicy(java.lang.String)
func (jbobject *ServicesEc2ModelModifySpotFleetRequestRequest) WithExcessCapacityTerminationPolicy2(a string) *ServicesEc2ModelModifySpotFleetRequestRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withExcessCapacityTerminationPolicy", "com/amazonaws/services/ec2/model/ModifySpotFleetRequestRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifySpotFleetRequestRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setExcessCapacityTerminationPolicy(com.amazonaws.services.ec2.model.ExcessCapacityTerminationPolicy)
func (jbobject *ServicesEc2ModelModifySpotFleetRequestRequest) SetExcessCapacityTerminationPolicy(a ServicesEc2ModelExcessCapacityTerminationPolicyInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExcessCapacityTerminationPolicy", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ExcessCapacityTerminationPolicy"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ModifySpotFleetRequestRequest withExcessCapacityTerminationPolicy(com.amazonaws.services.ec2.model.ExcessCapacityTerminationPolicy)
func (jbobject *ServicesEc2ModelModifySpotFleetRequestRequest) WithExcessCapacityTerminationPolicy(a ServicesEc2ModelExcessCapacityTerminationPolicyInterface) *ServicesEc2ModelModifySpotFleetRequestRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withExcessCapacityTerminationPolicy", "com/amazonaws/services/ec2/model/ModifySpotFleetRequestRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ExcessCapacityTerminationPolicy"))
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
	unique_x := &ServicesEc2ModelModifySpotFleetRequestRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifySpotFleetRequestRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelModifySpotFleetRequestRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelModifySpotFleetRequestRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelModifySpotFleetRequestRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelModifySpotFleetRequestRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ModifySpotFleetRequestRequest clone()
func (jbobject *ServicesEc2ModelModifySpotFleetRequestRequest) Clone3() *ServicesEc2ModelModifySpotFleetRequestRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ModifySpotFleetRequestRequest")
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
	unique_x := &ServicesEc2ModelModifySpotFleetRequestRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelModifySpotFleetRequestRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelModifySpotFleetRequestRequest) Clone2() (*JavaLangObject, error) {
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


