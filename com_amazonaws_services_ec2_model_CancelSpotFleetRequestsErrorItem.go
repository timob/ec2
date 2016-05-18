package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCancelSpotFleetRequestsErrorItemInterface interface {
	JavaLangObjectInterface

	// public void setSpotFleetRequestId(java.lang.String)
	SetSpotFleetRequestId(a string) 

	// public java.lang.String getSpotFleetRequestId()
	GetSpotFleetRequestId() string

	// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsErrorItem withSpotFleetRequestId(java.lang.String)
	WithSpotFleetRequestId(a string) *ServicesEc2ModelCancelSpotFleetRequestsErrorItem

	// public void setError(com.amazonaws.services.ec2.model.CancelSpotFleetRequestsError)
	SetError(a ServicesEc2ModelCancelSpotFleetRequestsErrorInterface) 

	// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsError getError()
	GetError() *ServicesEc2ModelCancelSpotFleetRequestsError

	// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsErrorItem withError(com.amazonaws.services.ec2.model.CancelSpotFleetRequestsError)
	WithError(a ServicesEc2ModelCancelSpotFleetRequestsErrorInterface) *ServicesEc2ModelCancelSpotFleetRequestsErrorItem

	// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsErrorItem clone()
	Clone() *ServicesEc2ModelCancelSpotFleetRequestsErrorItem
}

type ServicesEc2ModelCancelSpotFleetRequestsErrorItem struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsErrorItem()
func NewServicesEc2ModelCancelSpotFleetRequestsErrorItem() (*ServicesEc2ModelCancelSpotFleetRequestsErrorItem) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CancelSpotFleetRequestsErrorItem")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCancelSpotFleetRequestsErrorItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSpotFleetRequestId(java.lang.String)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsErrorItem) SetSpotFleetRequestId(a string)  {
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
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsErrorItem) GetSpotFleetRequestId() string {
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

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsErrorItem withSpotFleetRequestId(java.lang.String)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsErrorItem) WithSpotFleetRequestId(a string) *ServicesEc2ModelCancelSpotFleetRequestsErrorItem {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotFleetRequestId", "com/amazonaws/services/ec2/model/CancelSpotFleetRequestsErrorItem", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCancelSpotFleetRequestsErrorItem{}
	unique_x.Callable = dst
	return unique_x
}

// public void setError(com.amazonaws.services.ec2.model.CancelSpotFleetRequestsError)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsErrorItem) SetError(a ServicesEc2ModelCancelSpotFleetRequestsErrorInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setError", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/CancelSpotFleetRequestsError"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsError getError()
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsErrorItem) GetError() *ServicesEc2ModelCancelSpotFleetRequestsError {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getError", "com/amazonaws/services/ec2/model/CancelSpotFleetRequestsError")
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
	unique_x := &ServicesEc2ModelCancelSpotFleetRequestsError{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsErrorItem withError(com.amazonaws.services.ec2.model.CancelSpotFleetRequestsError)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsErrorItem) WithError(a ServicesEc2ModelCancelSpotFleetRequestsErrorInterface) *ServicesEc2ModelCancelSpotFleetRequestsErrorItem {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withError", "com/amazonaws/services/ec2/model/CancelSpotFleetRequestsErrorItem", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CancelSpotFleetRequestsError"))
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
	unique_x := &ServicesEc2ModelCancelSpotFleetRequestsErrorItem{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsErrorItem) ToString() string {
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
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsErrorItem) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsErrorItem) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsErrorItem clone()
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsErrorItem) Clone() *ServicesEc2ModelCancelSpotFleetRequestsErrorItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CancelSpotFleetRequestsErrorItem")
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
	unique_x := &ServicesEc2ModelCancelSpotFleetRequestsErrorItem{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsErrorItem) Clone2() (*JavaLangObject, error) {
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


