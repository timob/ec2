package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreatePlacementGroupRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setGroupName(java.lang.String)
	SetGroupName(a string) 

	// public java.lang.String getGroupName()
	GetGroupName() string

	// public com.amazonaws.services.ec2.model.CreatePlacementGroupRequest withGroupName(java.lang.String)
	WithGroupName(a string) *ServicesEc2ModelCreatePlacementGroupRequest

	// public void setStrategy(java.lang.String)
	SetStrategy2(a string) 

	// public java.lang.String getStrategy()
	GetStrategy() string

	// public com.amazonaws.services.ec2.model.CreatePlacementGroupRequest withStrategy(java.lang.String)
	WithStrategy2(a string) *ServicesEc2ModelCreatePlacementGroupRequest

	// public void setStrategy(com.amazonaws.services.ec2.model.PlacementStrategy)
	SetStrategy(a ServicesEc2ModelPlacementStrategyInterface) 

	// public com.amazonaws.services.ec2.model.CreatePlacementGroupRequest withStrategy(com.amazonaws.services.ec2.model.PlacementStrategy)
	WithStrategy(a ServicesEc2ModelPlacementStrategyInterface) *ServicesEc2ModelCreatePlacementGroupRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreatePlacementGroupRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.CreatePlacementGroupRequest clone()
	Clone3() *ServicesEc2ModelCreatePlacementGroupRequest
}

type ServicesEc2ModelCreatePlacementGroupRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.CreatePlacementGroupRequest()
func NewServicesEc2ModelCreatePlacementGroupRequest() (*ServicesEc2ModelCreatePlacementGroupRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreatePlacementGroupRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreatePlacementGroupRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.CreatePlacementGroupRequest(java.lang.String, java.lang.String)
func NewServicesEc2ModelCreatePlacementGroupRequest3(a string, b string) (*ServicesEc2ModelCreatePlacementGroupRequest) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreatePlacementGroupRequest", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &ServicesEc2ModelCreatePlacementGroupRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.CreatePlacementGroupRequest(java.lang.String, com.amazonaws.services.ec2.model.PlacementStrategy)
func NewServicesEc2ModelCreatePlacementGroupRequest2(a string, b ServicesEc2ModelPlacementStrategyInterface) (*ServicesEc2ModelCreatePlacementGroupRequest) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreatePlacementGroupRequest", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("com/amazonaws/services/ec2/model/PlacementStrategy"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &ServicesEc2ModelCreatePlacementGroupRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setGroupName(java.lang.String)
func (jbobject *ServicesEc2ModelCreatePlacementGroupRequest) SetGroupName(a string)  {
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
func (jbobject *ServicesEc2ModelCreatePlacementGroupRequest) GetGroupName() string {
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

// public com.amazonaws.services.ec2.model.CreatePlacementGroupRequest withGroupName(java.lang.String)
func (jbobject *ServicesEc2ModelCreatePlacementGroupRequest) WithGroupName(a string) *ServicesEc2ModelCreatePlacementGroupRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroupName", "com/amazonaws/services/ec2/model/CreatePlacementGroupRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreatePlacementGroupRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStrategy(java.lang.String)
func (jbobject *ServicesEc2ModelCreatePlacementGroupRequest) SetStrategy2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStrategy", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getStrategy()
func (jbobject *ServicesEc2ModelCreatePlacementGroupRequest) GetStrategy() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStrategy", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreatePlacementGroupRequest withStrategy(java.lang.String)
func (jbobject *ServicesEc2ModelCreatePlacementGroupRequest) WithStrategy2(a string) *ServicesEc2ModelCreatePlacementGroupRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStrategy", "com/amazonaws/services/ec2/model/CreatePlacementGroupRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreatePlacementGroupRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStrategy(com.amazonaws.services.ec2.model.PlacementStrategy)
func (jbobject *ServicesEc2ModelCreatePlacementGroupRequest) SetStrategy(a ServicesEc2ModelPlacementStrategyInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStrategy", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/PlacementStrategy"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreatePlacementGroupRequest withStrategy(com.amazonaws.services.ec2.model.PlacementStrategy)
func (jbobject *ServicesEc2ModelCreatePlacementGroupRequest) WithStrategy(a ServicesEc2ModelPlacementStrategyInterface) *ServicesEc2ModelCreatePlacementGroupRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStrategy", "com/amazonaws/services/ec2/model/CreatePlacementGroupRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/PlacementStrategy"))
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
	unique_x := &ServicesEc2ModelCreatePlacementGroupRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreatePlacementGroupRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelCreatePlacementGroupRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelCreatePlacementGroupRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCreatePlacementGroupRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreatePlacementGroupRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreatePlacementGroupRequest clone()
func (jbobject *ServicesEc2ModelCreatePlacementGroupRequest) Clone3() *ServicesEc2ModelCreatePlacementGroupRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreatePlacementGroupRequest")
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
	unique_x := &ServicesEc2ModelCreatePlacementGroupRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelCreatePlacementGroupRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelCreatePlacementGroupRequest) Clone2() (*JavaLangObject, error) {
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


