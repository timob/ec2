package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelRequestSpotFleetRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setSpotFleetRequestConfig(com.amazonaws.services.ec2.model.SpotFleetRequestConfigData)
	SetSpotFleetRequestConfig(a ServicesEc2ModelSpotFleetRequestConfigDataInterface) 

	// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData getSpotFleetRequestConfig()
	GetSpotFleetRequestConfig() *ServicesEc2ModelSpotFleetRequestConfigData

	// public com.amazonaws.services.ec2.model.RequestSpotFleetRequest withSpotFleetRequestConfig(com.amazonaws.services.ec2.model.SpotFleetRequestConfigData)
	WithSpotFleetRequestConfig(a ServicesEc2ModelSpotFleetRequestConfigDataInterface) *ServicesEc2ModelRequestSpotFleetRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.RequestSpotFleetRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.RequestSpotFleetRequest clone()
	Clone3() *ServicesEc2ModelRequestSpotFleetRequest
}

type ServicesEc2ModelRequestSpotFleetRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.RequestSpotFleetRequest()
func NewServicesEc2ModelRequestSpotFleetRequest() (*ServicesEc2ModelRequestSpotFleetRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/RequestSpotFleetRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelRequestSpotFleetRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSpotFleetRequestConfig(com.amazonaws.services.ec2.model.SpotFleetRequestConfigData)
func (jbobject *ServicesEc2ModelRequestSpotFleetRequest) SetSpotFleetRequestConfig(a ServicesEc2ModelSpotFleetRequestConfigDataInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSpotFleetRequestConfig", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotFleetRequestConfigData"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData getSpotFleetRequestConfig()
func (jbobject *ServicesEc2ModelRequestSpotFleetRequest) GetSpotFleetRequestConfig() *ServicesEc2ModelSpotFleetRequestConfigData {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSpotFleetRequestConfig", "com/amazonaws/services/ec2/model/SpotFleetRequestConfigData")
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
	unique_x := &ServicesEc2ModelSpotFleetRequestConfigData{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.RequestSpotFleetRequest withSpotFleetRequestConfig(com.amazonaws.services.ec2.model.SpotFleetRequestConfigData)
func (jbobject *ServicesEc2ModelRequestSpotFleetRequest) WithSpotFleetRequestConfig(a ServicesEc2ModelSpotFleetRequestConfigDataInterface) *ServicesEc2ModelRequestSpotFleetRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotFleetRequestConfig", "com/amazonaws/services/ec2/model/RequestSpotFleetRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotFleetRequestConfigData"))
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
	unique_x := &ServicesEc2ModelRequestSpotFleetRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.RequestSpotFleetRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelRequestSpotFleetRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelRequestSpotFleetRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelRequestSpotFleetRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelRequestSpotFleetRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.RequestSpotFleetRequest clone()
func (jbobject *ServicesEc2ModelRequestSpotFleetRequest) Clone3() *ServicesEc2ModelRequestSpotFleetRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/RequestSpotFleetRequest")
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
	unique_x := &ServicesEc2ModelRequestSpotFleetRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelRequestSpotFleetRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelRequestSpotFleetRequest) Clone2() (*JavaLangObject, error) {
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


