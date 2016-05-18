package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCancelledSpotInstanceRequestInterface interface {
	JavaLangObjectInterface

	// public void setSpotInstanceRequestId(java.lang.String)
	SetSpotInstanceRequestId(a string) 

	// public java.lang.String getSpotInstanceRequestId()
	GetSpotInstanceRequestId() string

	// public com.amazonaws.services.ec2.model.CancelledSpotInstanceRequest withSpotInstanceRequestId(java.lang.String)
	WithSpotInstanceRequestId(a string) *ServicesEc2ModelCancelledSpotInstanceRequest

	// public void setState(java.lang.String)
	SetState2(a string) 

	// public java.lang.String getState()
	GetState() string

	// public com.amazonaws.services.ec2.model.CancelledSpotInstanceRequest withState(java.lang.String)
	WithState2(a string) *ServicesEc2ModelCancelledSpotInstanceRequest

	// public void setState(com.amazonaws.services.ec2.model.CancelSpotInstanceRequestState)
	SetState(a ServicesEc2ModelCancelSpotInstanceRequestStateInterface) 

	// public com.amazonaws.services.ec2.model.CancelledSpotInstanceRequest withState(com.amazonaws.services.ec2.model.CancelSpotInstanceRequestState)
	WithState(a ServicesEc2ModelCancelSpotInstanceRequestStateInterface) *ServicesEc2ModelCancelledSpotInstanceRequest

	// public com.amazonaws.services.ec2.model.CancelledSpotInstanceRequest clone()
	Clone() *ServicesEc2ModelCancelledSpotInstanceRequest
}

type ServicesEc2ModelCancelledSpotInstanceRequest struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CancelledSpotInstanceRequest()
func NewServicesEc2ModelCancelledSpotInstanceRequest() (*ServicesEc2ModelCancelledSpotInstanceRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CancelledSpotInstanceRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCancelledSpotInstanceRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSpotInstanceRequestId(java.lang.String)
func (jbobject *ServicesEc2ModelCancelledSpotInstanceRequest) SetSpotInstanceRequestId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSpotInstanceRequestId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSpotInstanceRequestId()
func (jbobject *ServicesEc2ModelCancelledSpotInstanceRequest) GetSpotInstanceRequestId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSpotInstanceRequestId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CancelledSpotInstanceRequest withSpotInstanceRequestId(java.lang.String)
func (jbobject *ServicesEc2ModelCancelledSpotInstanceRequest) WithSpotInstanceRequestId(a string) *ServicesEc2ModelCancelledSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotInstanceRequestId", "com/amazonaws/services/ec2/model/CancelledSpotInstanceRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCancelledSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(java.lang.String)
func (jbobject *ServicesEc2ModelCancelledSpotInstanceRequest) SetState2(a string)  {
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
func (jbobject *ServicesEc2ModelCancelledSpotInstanceRequest) GetState() string {
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

// public com.amazonaws.services.ec2.model.CancelledSpotInstanceRequest withState(java.lang.String)
func (jbobject *ServicesEc2ModelCancelledSpotInstanceRequest) WithState2(a string) *ServicesEc2ModelCancelledSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/CancelledSpotInstanceRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCancelledSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(com.amazonaws.services.ec2.model.CancelSpotInstanceRequestState)
func (jbobject *ServicesEc2ModelCancelledSpotInstanceRequest) SetState(a ServicesEc2ModelCancelSpotInstanceRequestStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CancelledSpotInstanceRequest withState(com.amazonaws.services.ec2.model.CancelSpotInstanceRequestState)
func (jbobject *ServicesEc2ModelCancelledSpotInstanceRequest) WithState(a ServicesEc2ModelCancelSpotInstanceRequestStateInterface) *ServicesEc2ModelCancelledSpotInstanceRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/CancelledSpotInstanceRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState"))
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
	unique_x := &ServicesEc2ModelCancelledSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCancelledSpotInstanceRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCancelledSpotInstanceRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCancelledSpotInstanceRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CancelledSpotInstanceRequest clone()
func (jbobject *ServicesEc2ModelCancelledSpotInstanceRequest) Clone() *ServicesEc2ModelCancelledSpotInstanceRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CancelledSpotInstanceRequest")
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
	unique_x := &ServicesEc2ModelCancelledSpotInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCancelledSpotInstanceRequest) Clone2() (*JavaLangObject, error) {
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


