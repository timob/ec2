package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelRunScheduledInstancesRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setClientToken(java.lang.String)
	SetClientToken(a string) 

	// public java.lang.String getClientToken()
	GetClientToken() string

	// public com.amazonaws.services.ec2.model.RunScheduledInstancesRequest withClientToken(java.lang.String)
	WithClientToken(a string) *ServicesEc2ModelRunScheduledInstancesRequest

	// public void setInstanceCount(java.lang.Integer)
	SetInstanceCount(a int) 

	// public java.lang.Integer getInstanceCount()
	GetInstanceCount() int

	// public com.amazonaws.services.ec2.model.RunScheduledInstancesRequest withInstanceCount(java.lang.Integer)
	WithInstanceCount(a int) *ServicesEc2ModelRunScheduledInstancesRequest

	// public void setScheduledInstanceId(java.lang.String)
	SetScheduledInstanceId(a string) 

	// public java.lang.String getScheduledInstanceId()
	GetScheduledInstanceId() string

	// public com.amazonaws.services.ec2.model.RunScheduledInstancesRequest withScheduledInstanceId(java.lang.String)
	WithScheduledInstanceId(a string) *ServicesEc2ModelRunScheduledInstancesRequest

	// public void setLaunchSpecification(com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification)
	SetLaunchSpecification(a ServicesEc2ModelScheduledInstancesLaunchSpecificationInterface) 

	// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification getLaunchSpecification()
	GetLaunchSpecification() *ServicesEc2ModelScheduledInstancesLaunchSpecification

	// public com.amazonaws.services.ec2.model.RunScheduledInstancesRequest withLaunchSpecification(com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification)
	WithLaunchSpecification(a ServicesEc2ModelScheduledInstancesLaunchSpecificationInterface) *ServicesEc2ModelRunScheduledInstancesRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.RunScheduledInstancesRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.RunScheduledInstancesRequest clone()
	Clone3() *ServicesEc2ModelRunScheduledInstancesRequest
}

type ServicesEc2ModelRunScheduledInstancesRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.RunScheduledInstancesRequest()
func NewServicesEc2ModelRunScheduledInstancesRequest() (*ServicesEc2ModelRunScheduledInstancesRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/RunScheduledInstancesRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelRunScheduledInstancesRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelRunScheduledInstancesRequest) SetClientToken(a string)  {
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
func (jbobject *ServicesEc2ModelRunScheduledInstancesRequest) GetClientToken() string {
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

// public com.amazonaws.services.ec2.model.RunScheduledInstancesRequest withClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelRunScheduledInstancesRequest) WithClientToken(a string) *ServicesEc2ModelRunScheduledInstancesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClientToken", "com/amazonaws/services/ec2/model/RunScheduledInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRunScheduledInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelRunScheduledInstancesRequest) SetInstanceCount(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceCount", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getInstanceCount()
func (jbobject *ServicesEc2ModelRunScheduledInstancesRequest) GetInstanceCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceCount", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.RunScheduledInstancesRequest withInstanceCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelRunScheduledInstancesRequest) WithInstanceCount(a int) *ServicesEc2ModelRunScheduledInstancesRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceCount", "com/amazonaws/services/ec2/model/RunScheduledInstancesRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelRunScheduledInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setScheduledInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelRunScheduledInstancesRequest) SetScheduledInstanceId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setScheduledInstanceId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getScheduledInstanceId()
func (jbobject *ServicesEc2ModelRunScheduledInstancesRequest) GetScheduledInstanceId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getScheduledInstanceId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.RunScheduledInstancesRequest withScheduledInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelRunScheduledInstancesRequest) WithScheduledInstanceId(a string) *ServicesEc2ModelRunScheduledInstancesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withScheduledInstanceId", "com/amazonaws/services/ec2/model/RunScheduledInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRunScheduledInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setLaunchSpecification(com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification)
func (jbobject *ServicesEc2ModelRunScheduledInstancesRequest) SetLaunchSpecification(a ServicesEc2ModelScheduledInstancesLaunchSpecificationInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLaunchSpecification", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ScheduledInstancesLaunchSpecification"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification getLaunchSpecification()
func (jbobject *ServicesEc2ModelRunScheduledInstancesRequest) GetLaunchSpecification() *ServicesEc2ModelScheduledInstancesLaunchSpecification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLaunchSpecification", "com/amazonaws/services/ec2/model/ScheduledInstancesLaunchSpecification")
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
	unique_x := &ServicesEc2ModelScheduledInstancesLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.RunScheduledInstancesRequest withLaunchSpecification(com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification)
func (jbobject *ServicesEc2ModelRunScheduledInstancesRequest) WithLaunchSpecification(a ServicesEc2ModelScheduledInstancesLaunchSpecificationInterface) *ServicesEc2ModelRunScheduledInstancesRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withLaunchSpecification", "com/amazonaws/services/ec2/model/RunScheduledInstancesRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ScheduledInstancesLaunchSpecification"))
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
	unique_x := &ServicesEc2ModelRunScheduledInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.RunScheduledInstancesRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelRunScheduledInstancesRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelRunScheduledInstancesRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelRunScheduledInstancesRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelRunScheduledInstancesRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.RunScheduledInstancesRequest clone()
func (jbobject *ServicesEc2ModelRunScheduledInstancesRequest) Clone3() *ServicesEc2ModelRunScheduledInstancesRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/RunScheduledInstancesRequest")
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
	unique_x := &ServicesEc2ModelRunScheduledInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelRunScheduledInstancesRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelRunScheduledInstancesRequest) Clone2() (*JavaLangObject, error) {
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


