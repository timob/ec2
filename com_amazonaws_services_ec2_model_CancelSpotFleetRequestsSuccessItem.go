package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCancelSpotFleetRequestsSuccessItemInterface interface {
	JavaLangObjectInterface

	// public void setSpotFleetRequestId(java.lang.String)
	SetSpotFleetRequestId(a string) 

	// public java.lang.String getSpotFleetRequestId()
	GetSpotFleetRequestId() string

	// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsSuccessItem withSpotFleetRequestId(java.lang.String)
	WithSpotFleetRequestId(a string) *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem

	// public void setCurrentSpotFleetRequestState(java.lang.String)
	SetCurrentSpotFleetRequestState2(a string) 

	// public java.lang.String getCurrentSpotFleetRequestState()
	GetCurrentSpotFleetRequestState() string

	// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsSuccessItem withCurrentSpotFleetRequestState(java.lang.String)
	WithCurrentSpotFleetRequestState2(a string) *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem

	// public void setCurrentSpotFleetRequestState(com.amazonaws.services.ec2.model.BatchState)
	SetCurrentSpotFleetRequestState(a ServicesEc2ModelBatchStateInterface) 

	// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsSuccessItem withCurrentSpotFleetRequestState(com.amazonaws.services.ec2.model.BatchState)
	WithCurrentSpotFleetRequestState(a ServicesEc2ModelBatchStateInterface) *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem

	// public void setPreviousSpotFleetRequestState(java.lang.String)
	SetPreviousSpotFleetRequestState2(a string) 

	// public java.lang.String getPreviousSpotFleetRequestState()
	GetPreviousSpotFleetRequestState() string

	// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsSuccessItem withPreviousSpotFleetRequestState(java.lang.String)
	WithPreviousSpotFleetRequestState2(a string) *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem

	// public void setPreviousSpotFleetRequestState(com.amazonaws.services.ec2.model.BatchState)
	SetPreviousSpotFleetRequestState(a ServicesEc2ModelBatchStateInterface) 

	// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsSuccessItem withPreviousSpotFleetRequestState(com.amazonaws.services.ec2.model.BatchState)
	WithPreviousSpotFleetRequestState(a ServicesEc2ModelBatchStateInterface) *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem

	// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsSuccessItem clone()
	Clone() *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem
}

type ServicesEc2ModelCancelSpotFleetRequestsSuccessItem struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsSuccessItem()
func NewServicesEc2ModelCancelSpotFleetRequestsSuccessItem() (*ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CancelSpotFleetRequestsSuccessItem")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCancelSpotFleetRequestsSuccessItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSpotFleetRequestId(java.lang.String)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) SetSpotFleetRequestId(a string)  {
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
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) GetSpotFleetRequestId() string {
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

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsSuccessItem withSpotFleetRequestId(java.lang.String)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) WithSpotFleetRequestId(a string) *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotFleetRequestId", "com/amazonaws/services/ec2/model/CancelSpotFleetRequestsSuccessItem", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCancelSpotFleetRequestsSuccessItem{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCurrentSpotFleetRequestState(java.lang.String)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) SetCurrentSpotFleetRequestState2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCurrentSpotFleetRequestState", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getCurrentSpotFleetRequestState()
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) GetCurrentSpotFleetRequestState() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCurrentSpotFleetRequestState", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsSuccessItem withCurrentSpotFleetRequestState(java.lang.String)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) WithCurrentSpotFleetRequestState2(a string) *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCurrentSpotFleetRequestState", "com/amazonaws/services/ec2/model/CancelSpotFleetRequestsSuccessItem", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCancelSpotFleetRequestsSuccessItem{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCurrentSpotFleetRequestState(com.amazonaws.services.ec2.model.BatchState)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) SetCurrentSpotFleetRequestState(a ServicesEc2ModelBatchStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCurrentSpotFleetRequestState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/BatchState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsSuccessItem withCurrentSpotFleetRequestState(com.amazonaws.services.ec2.model.BatchState)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) WithCurrentSpotFleetRequestState(a ServicesEc2ModelBatchStateInterface) *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCurrentSpotFleetRequestState", "com/amazonaws/services/ec2/model/CancelSpotFleetRequestsSuccessItem", conv_a.Value().Cast("com/amazonaws/services/ec2/model/BatchState"))
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
	unique_x := &ServicesEc2ModelCancelSpotFleetRequestsSuccessItem{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPreviousSpotFleetRequestState(java.lang.String)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) SetPreviousSpotFleetRequestState2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPreviousSpotFleetRequestState", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPreviousSpotFleetRequestState()
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) GetPreviousSpotFleetRequestState() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPreviousSpotFleetRequestState", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsSuccessItem withPreviousSpotFleetRequestState(java.lang.String)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) WithPreviousSpotFleetRequestState2(a string) *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPreviousSpotFleetRequestState", "com/amazonaws/services/ec2/model/CancelSpotFleetRequestsSuccessItem", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCancelSpotFleetRequestsSuccessItem{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPreviousSpotFleetRequestState(com.amazonaws.services.ec2.model.BatchState)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) SetPreviousSpotFleetRequestState(a ServicesEc2ModelBatchStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPreviousSpotFleetRequestState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/BatchState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsSuccessItem withPreviousSpotFleetRequestState(com.amazonaws.services.ec2.model.BatchState)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) WithPreviousSpotFleetRequestState(a ServicesEc2ModelBatchStateInterface) *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPreviousSpotFleetRequestState", "com/amazonaws/services/ec2/model/CancelSpotFleetRequestsSuccessItem", conv_a.Value().Cast("com/amazonaws/services/ec2/model/BatchState"))
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
	unique_x := &ServicesEc2ModelCancelSpotFleetRequestsSuccessItem{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) ToString() string {
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
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsSuccessItem clone()
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) Clone() *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CancelSpotFleetRequestsSuccessItem")
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
	unique_x := &ServicesEc2ModelCancelSpotFleetRequestsSuccessItem{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsSuccessItem) Clone2() (*JavaLangObject, error) {
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


